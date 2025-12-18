package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListPost(conn *pgxpool.Pool) ([]models.Post, error) {
	rows, err := conn.Query(context.Background(),
        "SELECT post_id, title, content, user_id, topic_id FROM post",
    )
	if err != nil {
    return nil, err
	}
	defer rows.Close()
	post := []models.Post{}
	for rows.Next() {
		var p models.Post
		err :=rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserId, &p.TopicId)
		if err != nil {
    	return nil, err
		}
		post = append(post, p)
	}
	return post, nil
}
