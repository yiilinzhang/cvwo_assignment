package dataaccess

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func InsertUser(conn *pgxpool.Pool, username string, password string) (error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO users (name, password_hash)
		VALUES ($1, $2)`, username, password)
	return err
}
