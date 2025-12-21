package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

const (
	ListPosts = "posts.HandleList"

	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
	ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleListByTopic(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	topicId := chi.URLParam(r, "topicId")
	postList, err := dataaccess.ListPostByTopic(conn, topicId)
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
