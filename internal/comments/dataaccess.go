package comments

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

// SQL query that deletes the post with the same post id & user id from the database
// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func DeleteComment(conn *pgxpool.Pool, delCommentObj DeleteCommentInput) ([]models.Post, error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM comment WHERE comment_id = $1 AND user_id = $2`, delCommentObj.CommentID, delCommentObj.UserID)
	if err != nil {
		log.Fatalf("Exec error: %v\n", err)
	}
	if commandTag.RowsAffected() == 0 {
		log.Printf("No rows were affected\n")
	}
	return nil, err
}

// What i want returned, user.username, post.content, post.title, comment.content
func ListCommentByPost(conn *pgxpool.Pool, postID int) ([]QueryCommentResponse, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT comment.comment_id, comment.content, comment.user_id, users.name
		FROM comment 
		INNER JOIN users
		ON comment.user_id = users.userid
		WHERE comment.post_id = $1`,
		postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	comment := []QueryCommentResponse{}

	for rows.Next() {
		var c QueryCommentResponse
		err := rows.Scan(&c.ID, &c.Content, &c.UserID, &c.UserName)
		if err != nil {
			return nil, err
		}
		comment = append(comment, c)
	}
	return comment, nil
}

// TODO after auth pass in user id too
// TODO change the []model.post return type to just return error or smth more accurate
func InsertComment(conn *pgxpool.Pool, newCommentObj CreateCommentInput) ([]models.Comment, error) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO comment (content, user_id, post_id)
		VALUES ($1, $2, $3)`, newCommentObj.Content, newCommentObj.UserID, newCommentObj.PostID)
	return nil, err
}

//What i want returned, user.username, post.content, post.title, comment.content

func ListCommentByID(conn *pgxpool.Pool, commentID int) ([]QueryCommentResponse, error) {
	//TODO change to conn.exec cus one row only and update storage models
	rows, err := conn.Query(context.Background(),
		`SELECT content, user_id
		FROM comment 
		WHERE comment_id = $1`,
		commentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	comment := []QueryCommentResponse{}

	for rows.Next() {
		var c QueryCommentResponse
		err := rows.Scan(&c.Content, &c.UserID)
		if err != nil {
			return nil, err
		}
		comment = append(comment, c)
	}
	return comment, nil
}

func UpdateComment(conn *pgxpool.Pool, CommentObj UpdateCommentInput) ([]models.Post, error) {
	_, err := conn.Exec(context.Background(),
		`UPDATE comment SET content = $1 WHERE comment_id = $2 AND user_id = $3`,
		CommentObj.Content, CommentObj.CommentID, CommentObj.UserID)
	return nil, err
}
