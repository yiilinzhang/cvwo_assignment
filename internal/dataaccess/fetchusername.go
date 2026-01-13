package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//TODO depcreciated remove
func FetchUserName(conn *pgxpool.Pool, userid int) (string, error) {
	var userName string
	err := conn.QueryRow(context.Background(),
		`SELECT name FROM users
		WHERE userid = $1`,userid).Scan(&userName)
	if err != nil {
		return "", err
	}
	return userName, err
}
