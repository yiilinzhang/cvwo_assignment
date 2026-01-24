package posts

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListPostByTopic(conn *pgxpool.Pool, topicID int) ([]QueryPostResponse, error) {
	query := `SELECT 
			post.post_id, 
			post.title, 
			post.content, 
			post.user_id, 
			topic.title
		FROM post 
		INNER JOIN topic
			ON post.topic_id = topic.topic_id
		WHERE post.topic_id = $1
		ORDER BY post.post_id DESC`
	rows, err := conn.Query(context.Background(), query, topicID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	post := []QueryPostResponse{}
	for rows.Next() {
		var p QueryPostResponse
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID, &p.TopicTitle); err != nil {
			return nil, err
		}
		post = append(post, p)
	}
	return post, nil
}

func ListPostByID(conn *pgxpool.Pool, postID int) (models.Post, error) {
	var p models.Post
	query := `SELECT post_id, title, content, user_id, topic_id FROM post WHERE post_id = $1`
	err := conn.QueryRow(context.Background(), query, postID).
	Scan(&p.ID, &p.Title, &p.Content, &p.UserID, &p.TopicID)
	if err != nil {
		return models.Post{}, err
	}

	return p, nil
}

func ListAllPost(conn *pgxpool.Pool) ([]models.Post, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT post_id, title, content, user_id, topic_id
		FROM post
		ORDER BY post_id DESC`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	post := []models.Post{}
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID, &p.TopicID)
		if err != nil {
			return nil, err
		}
		post = append(post, p)
	}
	return post, nil
}

func DeletePost(conn *pgxpool.Pool, input DeletePostInput) error {
	commandTag, err := conn.Exec(
		context.Background(),
		`DELETE FROM post WHERE post_id = $1 AND user_id = $2`,
		input.PostID,
		input.UserID,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return api.NotFound(errors.New("No rows were affected\n"))
	}
	return nil
}

func InsertPost(conn *pgxpool.Pool, input CreatePostInput) (error) {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO post (title, content, user_id, topic_id)
		VALUES ($1, $2, $3, $4)`,
		input.Title,
		input.Content,
		input.UserID,
		input.TopicID,
	)
	return err
}

func UpdatePost(conn *pgxpool.Pool, input UpdatePostInput) (error) {
	_, err := conn.Exec(
		context.Background(),
		`UPDATE post SET content = $1 WHERE post_id = $2 AND user_id = $3`,
		input.Content, 
		input.PostID, 
		input.UserID,
	)
	return err
}
