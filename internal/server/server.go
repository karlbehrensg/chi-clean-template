package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/karlbehrensg/chi-clean-template/internal/apps/users"
	"github.com/karlbehrensg/chi-clean-template/utils"
	_ "github.com/lib/pq"
)

type Server struct {
	name    string
	version string
	env     string
	port    int
	db      *sql.DB
}

func NewServer() *http.Server {
	name := os.Getenv("SERVER_NAME")
	if name == "" {
		slog.Error("SERVER_NAME is not set")
		os.Exit(1)
	}

	version := os.Getenv("SERVER_VERSION")
	if version == "" {
		version = "v1.0"
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	env := os.Getenv("ENV")
	if env != "dev" && env != "" {
		env = "prod"
	}

	dbConnString := os.Getenv("DB_CONN_STRING")
	if dbConnString == "" {
		slog.Error("DB_CONN_STRING is not set")
		os.Exit(1)
	}
	db, err := utils.NewPostgresClient(dbConnString)
	if err != nil {
		slog.Error("Error connecting to Postgres: ", err)
		os.Exit(1)
	}

	newServer := &Server{
		name:    name,
		version: version,
		env:     env,
		port:    port,
		db:      db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.registerRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) registerRoutes() http.Handler {
	// Logger
	var logLevel slog.Level
	if s.env == "dev" {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}

	logger := httplog.NewLogger(s.name, httplog.Options{
		JSON:             s.env != "dev",
		LogLevel:         logLevel,
		Concise:          false,
		RequestHeaders:   true,
		MessageFieldName: "message",
		Tags: map[string]string{
			"version": s.version,
			"env":     s.env,
		},
		QuietDownRoutes: []string{
			"/ping",
		},
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)

	r.Mount("/users", users.RegisterRoutes(s.db))

	return r
}
