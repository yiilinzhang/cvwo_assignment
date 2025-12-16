package topics

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
	ListTopics = "topics.HandleList"

	SuccessfulListTopicsMessage = "Successfully listed topics"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveTopics          = "Failed to retrieve topics in %s"
	ErrEncodeView              = "Failed to retrieve topics in %s"
)

func HandleList(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicList, err := dataaccess.ListTopic(conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, ListTopics))
	}

	data, err := json.Marshal(topicList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListTopics))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListTopicsMessage},
	}, nil
}
