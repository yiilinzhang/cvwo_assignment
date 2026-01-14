package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
func InsertComment(conn *pgxpool.Pool, newCommentObj api.CreateCommentInput) ([]models.Comment, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO comment (content, user_id, post_id)
		VALUES ($1, $2, $3)`,newCommentObj.Content, newCommentObj.UserId, newCommentObj.PostId)
	return nil, err
}
