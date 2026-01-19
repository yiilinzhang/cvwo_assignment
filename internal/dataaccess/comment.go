package dataaccess

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

// SQL query that deletes the post with the same post id & user id from the database
// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func DeleteComment(conn *pgxpool.Pool, delCommentObj api.DeleteCommentInput) ([]models.Post, error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM comment WHERE comment_id = $1 AND user_id = $2`, delCommentObj.CommentId, delCommentObj.UserId)
	if err != nil {
		log.Fatalf("Exec error: %v\n", err)
	}
	if commandTag.RowsAffected() == 0 {
		log.Printf("No rows were affected\n")
	}
	return nil, err
}

//What i want returned, user.username, post.content, post.title, comment.content
func ListCommentByPost(conn *pgxpool.Pool, postId int) ([]models.QueryCommentResponse, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT comment.comment_id, comment.content, comment.user_id, users.name
		FROM comment 
		INNER JOIN users
		ON comment.user_id = users.userid
		WHERE comment.post_id = $1`,
		postId)
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

// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func InsertComment(conn *pgxpool.Pool, newCommentObj api.CreateCommentInput) ([]models.Comment, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO comment (content, user_id, post_id)
		VALUES ($1, $2, $3)`, newCommentObj.Content, newCommentObj.UserId, newCommentObj.PostId)
	return nil, err
}

//What i want returned, user.username, post.content, post.title, comment.content

func ListCommentById(conn *pgxpool.Pool, commentId int) ([]models.QueryCommentResponse, error) {
	//TODO change to conn.exec cus one row only and update storage models
	rows, err := conn.Query(context.Background(),
		`SELECT content, user_id
		FROM comment 
		WHERE comment_id = $1`,
		commentId)
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

func UpdateComment(conn *pgxpool.Pool, CommentObj api.UpdateCommentInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`UPDATE comment SET content = $1 WHERE comment_id = $2 AND user_id = $3`,
		CommentObj.Content, CommentObj.CommentId, CommentObj.UserId)
	return nil, err
}
