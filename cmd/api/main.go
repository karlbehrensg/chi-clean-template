package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/karlbehrensg/chi-clean-template/internal/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	server := server.NewServer()

	slog.Info(fmt.Sprintf("Server is running on port %v", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
