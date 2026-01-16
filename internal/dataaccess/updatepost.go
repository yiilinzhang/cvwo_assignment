package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)
func UpdatePost(conn *pgxpool.Pool, PostObj api.UpdatePostInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
	`UPDATE post SET content = $1 WHERE post_id = $2 AND user_id = $3`, 
PostObj.Content, PostObj.PostId, PostObj.UserId)
return nil, err
}
//TODO add better error catching