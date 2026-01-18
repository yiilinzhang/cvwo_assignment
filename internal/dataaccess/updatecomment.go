package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)
func UpdateComment(conn *pgxpool.Pool, CommentObj api.UpdateCommentInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
	`UPDATE comment SET content = $1 WHERE comment_id = $2 AND user_id = $3`, 
CommentObj.Content, CommentObj.CommentId, CommentObj.UserId)
return nil, err
}
//TODO add better error catching