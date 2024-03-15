package users

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/karlbehrensg/chi-clean-template/utils"
)

type service struct {
	cases ICases
}

func RegisterRoutes() http.Handler {
	repo := newRepository()
	cases := newCases(repo)
	service := newService(cases)

	r := chi.NewRouter()

	r.Post("/", service.CreateUser)

	return r
}

func newService(cases ICases) IService {
	return &service{
		cases,
	}
}

func (s *service) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := utils.NewTrace(r.Context())

	resp := make(map[string]string)
	resp["message"] = "User created successfully"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Error marshalling response: ", err)
	}

	utils.InfoTrace(ctx, "Building user object")
	user := &User{}

	s.cases.CreateUser(ctx, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}
