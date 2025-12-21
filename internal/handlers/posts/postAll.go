package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)


func HandleListAllPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	postList, err := dataaccess.ListAllPost(conn)
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
