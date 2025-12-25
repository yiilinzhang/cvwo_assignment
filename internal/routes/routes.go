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
	return func(r chi.Router) {
		//Private routes
		r.Group(func(r chi.Router){
		r.Get("/users", Routing(conn, users.HandleList))
		r.Get("/topics", Routing(conn, topics.HandleList))

		//TODO combine with queryparams
		r.Get("/posts", Routing(conn, posts.HandleListAllPosts))
		r.Get("/posts/{topicId}", Routing(conn, posts.HandleListByTopic))

		r.Post("/posts", Routing(conn, posts.HandleInsertPosts))})
	}
}

// TODO double check this code again
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

	// return func(r chi.Router) {
	// 	r.Group(func(r chi.Router) {
	// 	// Seek, verify and validate JWT tokens
	// 	r.Use(jwtauth.Verifier(auth.TokenAuth))

	// 	// Handle valid / invalid tokens. In this example, we use
	// 	// the provided authenticator middleware, but you can write your
	// 	// own very easily, look at the Authenticator method in jwtauth.go
	// 	// and tweak it, its not scary.
	// 	r.Use(jwtauth.Authenticator(auth.TokenAuth))

	// 	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
	// 		_, claims, _ := jwtauth.FromContext(r.Context())
	// 		w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
	// 	})

	// 	r.Get("/users", Routing(conn, users.HandleList))
	// 	r.Get("/topics", Routing(conn, topics.HandleList))

	// 	//TODO combine with queryparams
	// 	r.Get("/posts", Routing(conn, posts.HandleListAllPosts))
	// 	r.Get("/posts/{topicId}", Routing(conn, posts.HandleListByTopic))

	// 	r.Post("/posts", Routing(conn, posts.HandleInsertPosts))

	// })

	
	// }