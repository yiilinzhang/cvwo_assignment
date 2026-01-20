package topics

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
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
		if err := rows.Scan(&t.UserID, &t.ID, &t.Title); err != nil {
			return nil, err
		}
		topic = append(topic, t)
	}
	return topic, nil
}

func DeleteTopic(conn *pgxpool.Pool, input DeleteTopicInput) error {
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM topic WHERE topic_id = $1 AND user_id = $2`,
		input.TopicID,
		input.UserID,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return api.NotFound(errors.New("topic doesnt exist"))
	}
	return nil
}

func InsertTopic(conn *pgxpool.Pool, input CreateTopicInput) error {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO topic (title, user_id) VALUES ($1, $2)`,
		input.Title,
		input.UserID,
	)
	//TODO check if i need this or no error is nil
	if err != nil {
		return err
	}
	return nil
}
