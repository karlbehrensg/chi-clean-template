package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/karlbehrensg/chi-clean-template/utils"
)

type service struct {
	cases ICases
}

func RegisterRoutes(db *sql.DB) http.Handler {
	repo := newRepository(db)
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

	utils.InfoTrace(ctx, "Checking request form")
	if err := r.ParseForm(); err != nil {
		utils.ErrorTrace(ctx, "Error parsing form: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"message": "Invalid form",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	utils.InfoTrace(ctx, "Mounting request values to user struct")
	user := &User{}
	errs, err := user.MountRequestValues(r.PostForm)
	if err != nil {
		utils.ErrorTrace(ctx, "Error mounting request values: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		var errsMsg []string
		for _, e := range errs {
			errsMsg = append(errsMsg, e.Error())
		}

		response := map[string]interface{}{
			"message": err.Error(),
			"errors":  errsMsg,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	if err := s.cases.CreateUser(ctx, user); err != nil {
		utils.ErrorTrace(ctx, "Error creating user", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{
			"message": "Error creating user",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
}
