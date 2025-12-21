package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListUser(conn *pgxpool.Pool) ([]models.User, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT userid, name FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
