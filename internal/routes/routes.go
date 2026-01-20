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
	"github.com/yiilinzhang/cvwo_assignment/internal/comments"
	"github.com/yiilinzhang/cvwo_assignment/internal/posts"
	"github.com/yiilinzhang/cvwo_assignment/internal/topics"
	"github.com/yiilinzhang/cvwo_assignment/internal/users"
)

type ListHandler func(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error)

func PrivateRoutes(conn *pgxpool.Pool) func(r chi.Router) {
	return func(r chi.Router) {
		//Private routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.TokenAuth))
			r.Use(jwtauth.Authenticator(auth.TokenAuth))

			r.Get("/me", Routing(conn, users.HandleMe))
			//TODO combine with queryparams
			r.Get("/users", Routing(conn, users.HandleListUsers))
			r.Post("/logout", Routing(conn, users.HandleLogout))

			r.Post("/posts", Routing(conn, posts.HandleInsertPosts))
			r.Delete("/posts/{postID}", Routing(conn, posts.HandleDeletePosts))
			r.Patch("/posts/{postID}", Routing(conn, posts.HandleEditPosts))
			r.Post("/posts/{postID}/comments", Routing(conn, comments.HandleInsertComments))

			r.Post("/topics", Routing(conn, topics.HandleInsertTopics))
			r.Delete("/topics/{topicID}", Routing(conn, topics.HandleDeleteTopics))

			r.Get("/comments/{commentID}", Routing(conn, comments.HandleCommentByID))
			r.Delete("/comments/{commentID}", Routing(conn, comments.HandleDeleteComments))
			r.Patch("/comments/{commentID}", Routing(conn, comments.HandleEditComments))

		})
	}
}

func PublicRoutes(conn *pgxpool.Pool) func(r chi.Router) {
	return func(r chi.Router) {
		//Public routes routes
		r.Group(func(r chi.Router) {
			//
			r.Post("/login", Routing(conn, users.HandleLoginAuth))
			r.Post("/users", Routing(conn, users.HandleAddUsers))
			r.Get("/posts/{topicID}", Routing(conn, posts.HandleListByTopic))
			r.Get("/posts/{postID}/comments", Routing(conn, comments.HandleListByPosts))
			r.Get("/posts", Routing(conn, posts.HandleListAllPosts))
			r.Get("/topics", Routing(conn, topics.HandleListTopics))
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
