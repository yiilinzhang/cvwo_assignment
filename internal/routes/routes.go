package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/auth"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers"
)

type ListHandler func(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error)

func PrivateRoutes(conn *pgxpool.Pool) func(r chi.Router) {
	return func(r chi.Router) {
		//Private routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.TokenAuth))
			r.Use(jwtauth.Authenticator(auth.TokenAuth))

			r.Get("/me", Routing(conn, handlers.HandleMe))
			//TODO combine with queryparams
			r.Get("/users", Routing(conn, handlers.HandleListUsers))
			r.Post("/logout", Routing(conn, handlers.HandleLogout))

			r.Post("/posts", Routing(conn, handlers.HandleInsertPosts))
			r.Delete("/posts/{postId}", Routing(conn, handlers.HandleDeletePosts))
			r.Patch("/posts/{postId}", Routing(conn, handlers.HandleEditPosts))
			r.Post("/posts/{postId}/comments", Routing(conn, handlers.HandleInsertComments))

			r.Post("/topics", Routing(conn, handlers.HandleInsertTopics))
			r.Delete("/topics/{topicId}", Routing(conn, handlers.HandleDeleteTopics))

			r.Get("/comments/{commentId}", Routing(conn, handlers.HandleCommentById))
			r.Delete("/comments/{commentId}", Routing(conn, handlers.HandleDeleteComments))
			r.Patch("/comments/{commentId}", Routing(conn, handlers.HandleEditComments))

		})
	}
}

func PublicRoutes(conn *pgxpool.Pool) func(r chi.Router) {
	return func(r chi.Router) {
		//Public routes routes
		r.Group(func(r chi.Router) {
			//
			r.Post("/login", Routing(conn, handlers.HandleLoginAuth))
			r.Post("/users", Routing(conn, handlers.HandleAddUsers))
			r.Get("/posts/{topicId}", Routing(conn, handlers.HandleListByTopic))
			r.Get("/posts/{postId}/comments", Routing(conn, handlers.HandleListByPosts))
			r.Get("/posts", Routing(conn, handlers.HandleListAllPosts))
			r.Get("/topics", Routing(conn, handlers.HandleListTopics))
		})
	}
}

// TODO double check this code again
func Routing(conn *pgxpool.Pool, HandleList ListHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		response, err := HandleList(conn, w, req)
		if err != nil {
			var httpErr api.HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Error(), httpErr.Status)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
