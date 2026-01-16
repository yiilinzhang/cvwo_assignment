package comments

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleInsertComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var commentJSON api.CreateCommentInput
	json.NewDecoder(r.Body).Decode(&commentJSON)

commentJSON.PostId = chi.URLParam(r, "postId")
	_, claims, err := jwtauth.FromContext(r.Context())
	commentJSON.UserId = int(claims["user_id"].(float64))
	if err != nil {
		return nil, errors.New("Title later")
	}
	//Validate input
	log.Println(commentJSON)

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newComment, err := dataaccess.InsertComment(conn, commentJSON)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListComments))
	}

	log.Println("got here 1")
	//TODO check if more efficient to not unmardshell
	data, err := json.Marshal(newComment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
	}
	log.Println("got here 2")

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}
