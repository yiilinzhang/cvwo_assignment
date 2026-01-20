package comments

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteComment(conn *pgxpool.Pool, input DeleteCommentInput) (error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM comment WHERE comment_id = $1 AND user_id = $2`, 
		input.CommentID, 
		input.UserID,
	)
	if err != nil {
		return nil
	}

	if commandTag.RowsAffected() == 0 {
		log.Printf("No rows were affected\n")
	}
	return nil
}

func ListCommentByPost(conn *pgxpool.Pool, postID int) ([]QueryCommentResponse, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT 
			comment.comment_id, 
			comment.content, 
			comment.user_id, users.name
		FROM comment 
		INNER JOIN users
			ON comment.user_id = users.userid
		WHERE comment.post_id = $1`,
		postID,
	)
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

func InsertComment(conn *pgxpool.Pool, input CreateCommentInput) (error) {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO comment (content, user_id, post_id)
		VALUES ($1, $2, $3)`, 
		input.Content, 
		input.UserID, 
		input.PostID,
	)
	return err
}

func ListCommentByID(conn *pgxpool.Pool, commentID int) ([]QueryCommentResponse, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT content, user_id
		FROM comment 
		WHERE comment_id = $1`,
		commentID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comment := []QueryCommentResponse{}
	for rows.Next() {
		var c QueryCommentResponse
		if err := rows.Scan(&c.Content, &c.UserID); err != nil {
			return nil, err
		}
		comment = append(comment, c)
	}
	return comment, nil
}

func UpdateComment(conn *pgxpool.Pool, input UpdateCommentInput) (error) {
	_, err := conn.Exec(
		context.Background(),
		`UPDATE comment SET content = $1 WHERE comment_id = $2 AND user_id = $3`,
		input.Content, 
		input.CommentID, 
		input.UserID)
	return err
}
