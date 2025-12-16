package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/yiilinzhang/cvwo_assignment/internal/routes"
)

func Setup(conn *pgx.Conn) chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r, conn)
	return r
}

func setUpRoutes(r chi.Router, conn *pgx.Conn) {
	r.Group(routes.GetRoutes(conn))
}
