package comments

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
)

func DeleteComment(conn *pgxpool.Pool, input DeleteCommentInput) (error) {
	commandTag, err := conn.Exec(context.Background(),
		`DELETE FROM comment WHERE comment_id = $1 AND user_id = $2`, 
		input.CommentID, 
		input.UserID,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		api.NotFound(errors.New("No rows were affected\n"))
	}
	return nil
}

func ListCommentByPost(conn *pgxpool.Pool, postID int) ([]QueryCommentResponse, error) {
	rows, err := conn.Query(
		context.Background(),
		`SELECT 
			c.comment_id, 
			c.content, 
			c.user_id, 
			u.name,
			c.parent_comment_id
		FROM comment c 
		INNER JOIN users u
			ON c.user_id = u.userid
		WHERE c.post_id = $1
		ORDER BY c.comment_id DESC`,
		postID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comment := []QueryCommentResponse{}
	for rows.Next() {
		var c QueryCommentResponse
		err := rows.Scan(&c.ID, &c.Content, &c.UserID, &c.UserName, &c.ParentID)
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
		`INSERT INTO comment (content, user_id, post_id, parent_comment_id)
		VALUES ($1, $2, $3, $4)`, 
		input.Content, 
		input.UserID, 
		input.PostID,
		input.ParentID,
	)
	return err
}

func ListCommentByID(conn *pgxpool.Pool, commentID int) (QueryCommentResponse, error) {
	var c QueryCommentResponse
	err := conn.QueryRow(
		context.Background(),
		`SELECT content, user_id
		FROM comment 
		WHERE comment_id = $1`,
		commentID,
	).Scan(&c.Content, &c.UserID)
	if err != nil {
		return QueryCommentResponse{}, err
	}
	return c, nil
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

func GetCommentPostID(conn *pgxpool.Pool, commentID int) (int, error) {
    var postID int
    err := conn.QueryRow(
        context.Background(),
        `SELECT post_id FROM comment WHERE comment_id = $1`,
        commentID,
    ).Scan(&postID)
    return postID, err
}
