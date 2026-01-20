package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListUser(conn *pgxpool.Pool) ([]models.User, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT userid, name, password_hash FROM users",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func FetchUser(conn *pgxpool.Pool, username string) (UserAuth, error) {
	var user UserAuth
	err := conn.QueryRow(
		context.Background(),
		`SELECT userid, password_hash FROM users WHERE name = $1`, 
		username,
		).Scan(&user.UserID, &user.PasswordHash)
	if err != nil {
		return UserAuth{}, err
	}
	return user, err
}

func InsertUser(conn *pgxpool.Pool, username string, password string) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO users (name, password_hash)
		VALUES ($1, $2)`, username, password)
	return err
}
