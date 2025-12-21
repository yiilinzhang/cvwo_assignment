package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func InsertPost(conn *pgxpool.Pool, newTitle string, newContent string, newUserId int) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO post (title, content, user_id)
		VALUES ($1, $2, $3)`, newTitle, newContent, newUserId)
	return nil, err
}
