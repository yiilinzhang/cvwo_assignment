package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func InsertPost(conn *pgxpool.Pool, newPostObj api.CreatePostInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO post (title, content, user_id, topic_id)
		VALUES ($1, $2, 1, $3)`, newPostObj.Title, newPostObj.Content, newPostObj.TopicId)
	return nil, err
}
