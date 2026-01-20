package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/routes"

	"github.com/yiilinzhang/cvwo_assignment/internal/comments"
	"github.com/yiilinzhang/cvwo_assignment/internal/posts"
	"github.com/yiilinzhang/cvwo_assignment/internal/topics"
	"github.com/yiilinzhang/cvwo_assignment/internal/users"
)

func Setup(conn *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	setUpRoutes(r, conn)
	return r
}

func setUpRoutes(r chi.Router, conn *pgxpool.Pool) {
	postsHandler := &posts.Handler{Conn: conn}
	usersHandler := &users.Handler{Conn: conn}
	topicsHandler := &topics.Handler{Conn: conn}
	commentsHandler := &comments.Handler{Conn: conn}
	//TODO use grpoup later when i add login auth
	//Private toutes
	r.Group(routes.PrivateRoutes(postsHandler, usersHandler, topicsHandler, commentsHandler))
	r.Group(routes.PublicRoutes(postsHandler, usersHandler, topicsHandler, commentsHandler))
}
