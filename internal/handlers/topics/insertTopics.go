package topics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleInsertTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var topicJSON api.CreateTopicInput
	json.NewDecoder(r.Body).Decode(&topicJSON)
	_, claims, err := jwtauth.FromContext(r.Context())
	topicJSON.UserId = int(claims["user_id"].(float64))
	if err != nil {
		return nil, errors.New("Title later")
	}
	//Validate input

	//TODO remove this log in the future
	log.Println(topicJSON.UserId)
	log.Println(topicJSON.Title)
	log.Println("got here 3")

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newTopic, err := dataaccess.InsertTopic(conn, topicJSON)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveTopics, ListTopics))
	}

	log.Println("got here 1")
	//TODO check if more efficient to not unmardshell
	data, err := json.Marshal(newTopic)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListTopics))
	}
	log.Println("got here 2")

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListTopicsMessage},
	}, nil
}
