package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/auth"
	"github.com/yiilinzhang/cvwo_assignment/internal/comments"
	"github.com/yiilinzhang/cvwo_assignment/internal/posts"
	"github.com/yiilinzhang/cvwo_assignment/internal/topics"
	"github.com/yiilinzhang/cvwo_assignment/internal/users"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (*api.Response, error)

func PrivateRoutes(
	postsHandler *posts.Handler,
	usersHandler *users.Handler,
	topicsHandler *topics.Handler,
	commentsHandler *comments.Handler,

) func(r chi.Router) {
	return func(r chi.Router) {
		//Private routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.TokenAuth))
			r.Use(jwtauth.Authenticator(auth.TokenAuth))

			r.Get("/me", Routing(usersHandler.Me))
			//TODO combine with queryparams
			r.Get("/users", Routing(usersHandler.List))

			r.Post("/posts", Routing(postsHandler.Create))
			r.Delete("/posts/{postID}", Routing(postsHandler.Delete))
			r.Patch("/posts/{postID}", Routing(postsHandler.Update))
			r.Post("/posts/{postID}/comments", Routing(commentsHandler.Create))

			r.Post("/topics", Routing(topicsHandler.Create))

			r.Get("/comments/{commentID}", Routing(commentsHandler.ListByID))
			r.Delete("/comments/{commentID}", Routing(commentsHandler.Delete))
			r.Patch("/comments/{commentID}", Routing(commentsHandler.Update))

		})
	}
}

func PublicRoutes(postsHandler *posts.Handler,
	usersHandler *users.Handler,
	topicsHandler *topics.Handler,
	commentsHandler *comments.Handler,
) func(r chi.Router) {
	return func(r chi.Router) {
		//Public routes routes
		r.Post("/login", Routing(usersHandler.LoginAuth))
		r.Post("/logout", Routing(usersHandler.Logout))
		r.Post("/users", Routing(usersHandler.Create))

		r.Get("/posts/{topicID}", Routing(postsHandler.ListByTopic))
		r.Get("/posts", Routing(postsHandler.List))

		r.Get("/topics", Routing(topicsHandler.List))

		r.Get("/posts/{postID}/comments", Routing(commentsHandler.ListByPosts))
	}
}

// Routing Wraps handlers to fmt JSON reponses and ... smth about handle error change when completed todo
// TODO error messgaes giges too mcuh into i think
func Routing(HandleList HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response, err := HandleList(w, req)
		if err != nil {
			status := http.StatusInternalServerError
			var httpErr api.HTTPError
			if errors.As(err, &httpErr) {
				status = httpErr.Status
			}

			w.WriteHeader(status)
			json.NewEncoder(w).Encode(api.Response{
				Payload:  api.Payload{},
				Messages: []string{err.Error()},
			})
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}
