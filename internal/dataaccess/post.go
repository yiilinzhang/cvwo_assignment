package dataaccess

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

//TODO check if i need defer rows.Close()
func ListPost(conn *pgxpool.Pool, topicId string) ([]models.Post, error) {
	topicInt, err := strconv.Atoi(topicId)
    if err != nil {
        return nil, err
    }
	rows, err := conn.Query(context.Background(),
        `SELECT post_id, title, content, user_id, topic_id 
		FROM post 
		WHERE topic_id = $1`,
		topicInt)
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
