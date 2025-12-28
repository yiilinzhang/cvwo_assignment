package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func FetchUser(conn *pgxpool.Pool, username string) (int, string, error) {
	var hash string
	var userID int
	err := conn.QueryRow(context.Background(),
		`SELECT userid, password_hash FROM users
		WHERE name = $1`,username).Scan(&userID, &hash)
	if err != nil {
		return 0, "", err
	}
	return userID, hash, err
}
