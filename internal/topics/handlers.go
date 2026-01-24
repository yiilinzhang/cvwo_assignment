package topics

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers"
)

type Handler struct {
	Conn *pgxpool.Pool
}

const (
	SuccessfulCreateTopicsMessage = "Successfully listed topic"
	SuccessfulListTopicsMessage   = "Successfully listed topics"
	SuccessfulUpdateTopicsMessage = "Successfully updated topic"
	SuccessfulDeleteTopicsMessage = "Successfully deleted topic"

	ErrCreateTopics   = "Failed to create topic in %s"
	ErrRetrieveTopics = "Failed to retrieve topics in %s"
	ErrUpdateTopics   = "Failed to update topic in %s"
	ErrDeleteTopics   = "Failed to delete topic in %s"

	ErrEncodeView = "Failed to encode topics in %s"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicList, err := ListTopic(h.Conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, "topics.List"))
	}

	data, err := json.Marshal(topicList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "topics.List"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListTopicsMessage},
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input CreateTopicInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserID = userID

	if err = InsertTopic(h.Conn, input); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateTopics, "topics.Create"))
	}

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulCreateTopicsMessage},
	}, nil
}
