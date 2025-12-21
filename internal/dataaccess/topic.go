package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListTopic(conn *pgxpool.Pool) ([]models.Topic, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT topic_id, title FROM topic",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	topic := []models.Topic{}
	for rows.Next() {
		var t models.Topic
		err := rows.Scan(&t.ID, &t.Title)
		if err != nil {
			return nil, err
		}
		topic = append(topic, t)
	}
	return topic, nil
}
