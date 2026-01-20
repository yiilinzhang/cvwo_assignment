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

const (
	ListTopics                  = "topics.HandleList"
	SuccessfulListTopicsMessage = "Successfully listed topics"
	ErrRetrieveTopics           = "Failed to retrieve topics in %s"
	ErrEncodeView               = "Failed to retrieve topics in %s"
)

func HandleListTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicList, err := ListTopic(conn)
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

// Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicID, err := strconv.Atoi(chi.URLParam(r, "topicID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid topicID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input DeleteTopicInput
	input.TopicID = topicID
	input.UserID = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = DeleteTopic(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, ListTopics))
	}
	return nil, nil
}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleInsertTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
	newTopic, err := InsertTopic(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, ListTopics))
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
		Messages: []string{SuccessfulListTopicsMessage},
	}, nil
}
