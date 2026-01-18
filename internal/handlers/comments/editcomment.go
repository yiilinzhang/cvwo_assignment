package comments

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleEditComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	var commentJSON api.UpdateCommentInput
	json.NewDecoder(r.Body).Decode(&commentJSON); 
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}
	commentJSON.CommentId = commentId;
	commentJSON.UserId = int(claims["user_id"].(float64))
	_, err = dataaccess.UpdateComment(conn, commentJSON);
	return &api.Response{
		Payload: api.Payload{
		},
		Messages: []string{},
	}, nil

	
}
