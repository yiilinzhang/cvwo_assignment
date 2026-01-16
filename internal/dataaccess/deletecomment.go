

package dataaccess

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/models"
)

//SQL query that deletes the post with the same post id & user id from the database
//TODO after auth pass in user id too
//TODO change the []model.post return type to just return error or smth more accurate
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
