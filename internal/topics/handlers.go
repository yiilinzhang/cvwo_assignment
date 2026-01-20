package topics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers"
)

type Handler struct{ Conn *pgxpool.Pool }

const (
	ListTopics                    = "topics.HandleList"
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListTopics))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListTopicsMessage},
	}, nil
}

// Delete a specific post from database, requires postid to be passed in via url
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicID, err := strconv.Atoi(chi.URLParam(r, "topicID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("invalid topicID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input DeleteTopicInput
	input.TopicID = topicID
	input.UserID = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = DeleteTopic(h.Conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDeleteTopics, "topics.Delete"))
	}
	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulDeleteTopicsMessage},
	}, nil
}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input CreateTopicInput
	input.UserID = userID
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newTopic, err := InsertTopic(h.Conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateTopics, "topics.Create"))
	}

	//TODO check if more efficient to not unmardshell
	data, err := json.Marshal(newTopic)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListTopics))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulCreateTopicsMessage},
	}, nil
}
