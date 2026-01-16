package dataaccess

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)



//What i want returned, user.username, post.content, post.title, comment.content


func ListCommentByPost(conn *pgxpool.Pool, postId string) ([]models.QueryCommentResponse, error) {
	postInt, err := strconv.Atoi(postId)
	if err != nil {
		return nil, err
	}


	rows, err := conn.Query(context.Background(),
		`SELECT comment.comment_id, comment.content, comment.user_id, users.name
		FROM comment 
		INNER JOIN users
		ON comment.user_id = users.userid
		WHERE comment.post_id = $1`,
		postInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comment := []models.QueryCommentResponse{}
	for rows.Next() {
		var c models.QueryCommentResponse
		err := rows.Scan(&c.ID, &c.Content, &c.UserId, &c.UserName)
		if err != nil {
			return nil, err
		}
		comment = append(comment, c)
	}
	return comment, nil
}
