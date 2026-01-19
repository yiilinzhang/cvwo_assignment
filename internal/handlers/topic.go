package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

const (
	ListTopics = "topics.HandleList"

	SuccessfulListTopicsMessage = "Successfully listed topics"
	ErrRetrieveTopics           = "Failed to retrieve topics in %s"
)

func HandleListTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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

// Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicId, err := strconv.Atoi(chi.URLParam(r, "topicId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid topic id"))
	}

	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	//Extract JSON from HTTP req
	var topicJSON api.DeleteTopicInput
	topicJSON.TopicId = topicId
	topicJSON.UserId = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = dataaccess.DeleteTopic(conn, topicJSON)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, ListTopics))
	}

	return nil, nil
}

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleInsertTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var topicJSON api.CreateTopicInput
	json.NewDecoder(r.Body).Decode(&topicJSON)
	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}
	topicJSON.UserId = userID

	//Validate input

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newTopic, err := dataaccess.InsertTopic(conn, topicJSON)
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
