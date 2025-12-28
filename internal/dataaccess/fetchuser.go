package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func FetchUser(conn *pgxpool.Pool, username string) (string, error) {
	var hash string
	err := conn.QueryRow(context.Background(),
		`SELECT password_hash FROM users
		WHERE name = $1`,username).Scan(&hash)
	if err != nil {
		return "", err
	}
	return hash, err
}
