package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func InsertTopic(conn *pgxpool.Pool, newTopicObj api.CreateTopicInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO topic (title, user_id)
		VALUES ($1, $2)`, newTopicObj.Title, newTopicObj.UserId)
	return nil, err
}
