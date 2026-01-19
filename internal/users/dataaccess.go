package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListUser(conn *pgxpool.Pool) ([]models.User, error) {
	rows, err := conn.Query(context.Background(),
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

// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func FetchUser(conn *pgxpool.Pool, username string) (int, string, error) {
	var hash string
	var userID int
	err := conn.QueryRow(context.Background(),
		`SELECT userid, password_hash FROM users
		WHERE name = $1`, username).Scan(&userID, &hash)
	if err != nil {
		return 0, "", err
	}
	return userID, hash, err
}

// TODO depcreciated remove
func FetchUserName(conn *pgxpool.Pool, userid int) (string, error) {
	var userName string
	err := conn.QueryRow(context.Background(),
		`SELECT name FROM users
		WHERE userid = $1`, userid).Scan(&userName)
	if err != nil {
		return "", err
	}
	return userName, err
}

// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func InsertUser(conn *pgxpool.Pool, username string, password string) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO users (name, password_hash)
		VALUES ($1, $2)`, username, password)
	return err
}
