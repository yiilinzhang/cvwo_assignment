package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

const (
	ListPosts = "posts.HandleList"

	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrievePosts          = "Failed to retrieve posts in %s"
	ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleList(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postList, err := dataaccess.ListPost(conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data, err := json.Marshal(postList)
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
