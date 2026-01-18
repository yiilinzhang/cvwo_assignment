package dataaccess

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)



//What i want returned, user.username, post.content, post.title, comment.content


func ListCommentById(conn *pgxpool.Pool, commentId string) ([]models.QueryCommentResponse, error) {
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		return nil, err
	}
	//TODO change to conn.exec cus one row only and update storage models

	rows, err := conn.Query(context.Background(),
		`SELECT content, user_id
		FROM comment 
		WHERE comment_id = $1`,
		commentIdInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comment := []models.QueryCommentResponse{}
	for rows.Next() {
		var c models.QueryCommentResponse
		err := rows.Scan(&c.Content, &c.UserId)
		if err != nil {
			return nil, err
		}
		comment = append(comment, c)
	}
	return comment, nil
}
