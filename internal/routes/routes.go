package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/topics"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/users"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/posts"
)

func GetRoutes(conn *pgx.Conn) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, err := users.HandleList(conn, w, req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		r.Get("/topics", func(w http.ResponseWriter, req *http.Request) {
			response, err := topics.HandleList(conn, w, req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		r.Get("/posts", func(w http.ResponseWriter, req *http.Request) {
			response, err := posts.HandleList(conn, w, req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
