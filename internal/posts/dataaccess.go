package posts

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

func ListPostByTopic(conn *pgxpool.Pool, topicID int) ([]QueryPostResponse, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT post.post_id, post.title, post.content, post.user_id, topic.title
		FROM post 
		INNER JOIN topic
		ON post.topic_id = topic.topic_id
		WHERE post.topic_id = $1`,
		topicID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	post := []QueryPostResponse{}

	for rows.Next() {
		var p QueryPostResponse
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID, &p.TopicTitle)
		if err != nil {
			return nil, err
		}
		post = append(post, p)
	}
	return post, nil
}

func ListAllPost(conn *pgxpool.Pool) ([]models.Post, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT post_id, title, content, user_id, topic_id
		FROM post`)
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

// SQL query that deletes the post with the same post id & user id from the database
// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func DeletePost(conn *pgxpool.Pool, delPostObj DeletePostInput) ([]models.Post, error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM post WHERE post_id = $1 AND user_id = $2`, delPostObj.PostID, delPostObj.UserID)
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
func InsertPost(conn *pgxpool.Pool, newPostObj CreatePostInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO post (title, content, user_id, topic_id)
		VALUES ($1, $2, $3, $4)`, newPostObj.Title, newPostObj.Content, newPostObj.UserID, newPostObj.TopicID)
	return nil, err
}

func UpdatePost(conn *pgxpool.Pool, PostObj UpdatePostInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`UPDATE post SET content = $1 WHERE post_id = $2 AND user_id = $3`,
		PostObj.Content, PostObj.PostID, PostObj.UserID)
	return nil, err
}

//TODO add better error catching
