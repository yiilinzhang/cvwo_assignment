package dataaccess

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListTopic(conn *pgxpool.Pool) ([]models.Topic, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT user_id, topic_id, title FROM topic",
	)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	topic := []models.Topic{}
	
	for rows.Next() {
		var t models.Topic
		err := rows.Scan(&t.UserId, &t.ID, &t.Title)
		if err != nil {
			return nil, err
		}
		topic = append(topic, t)
	}
	return topic, nil
}

// SQL query that deletes the post with the same post id & user id from the database
// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func DeleteTopic(conn *pgxpool.Pool, delTopicObj api.DeleteTopicInput) ([]models.Post, error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM topic WHERE topic_id = $1 AND user_id = $2`, delTopicObj.TopicId, delTopicObj.UserId)
	if err != nil {
		log.Fatalf("Exec error: %v\n", err)
	}
	if commandTag.RowsAffected() == 0 {
		log.Printf("No rows were affected\n")
	}
	return nil, err
}

// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func InsertTopic(conn *pgxpool.Pool, newTopicObj api.CreateTopicInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO topic (title, user_id)
		VALUES ($1, $2)`, newTopicObj.Title, newTopicObj.UserId)
	return nil, err
}
