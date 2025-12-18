package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	"github.com/yiilinzhang/cvwo_assignment/internal/routes"
)

func Setup(conn *pgx.Conn) chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders:   []string{"Accept","Authorization","Content-Type","X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	setUpRoutes(r, conn)
	return r
}

func setUpRoutes(r chi.Router, conn *pgx.Conn) {
	r.Group(routes.GetRoutes(conn))
}
