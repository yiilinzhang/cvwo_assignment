package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/posts"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/topics"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers/users"
)

type ListHandler func(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error)

func GetRoutes(conn *pgxpool.Pool) func(r chi.Router) {
	//TODO see if there is a way to modularise this code
	return func(r chi.Router) {
		r.Get("/users", Routing(conn, users.HandleList))
		r.Get("/topics", Routing(conn, topics.HandleList))
		r.Get("/posts", Routing(conn, posts.HandleList))
	}
}

func Routing(conn *pgxpool.Pool, HandleList ListHandler) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			response, err := HandleList(conn, w, req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
}