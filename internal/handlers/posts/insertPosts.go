package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleInsertPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var postJSON api.CreatePostInput
	json.NewDecoder(r.Body).Decode(&postJSON)

	//Validate input

	//TODO remove this log in the future
	log.Println(postJSON)

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newPost, err := dataaccess.InsertPost(conn, postJSON)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data, err := json.Marshal(newPost)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListPosts))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}
