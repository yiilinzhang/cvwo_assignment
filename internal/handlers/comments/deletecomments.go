package comments

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	_, claims, err := jwtauth.FromContext(r.Context())
	userID := int(claims["user_id"].(float64))
	//TODO add err catching
	//TODO retitlr these JSON things to paylaod
	//Extract JSON from HTTP req
	var commentJSON api.DeleteCommentInput 
	commentJSON.CommentId = commentId
	commentJSON.UserId = userID


	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = dataaccess.DeleteComment(conn, commentJSON)
	if err != nil {
		return nil, err
	}


	return nil, nil
}
