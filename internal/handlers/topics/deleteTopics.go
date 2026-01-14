package topics

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteTopics(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicId := chi.URLParam(r, "topicId")
	_, claims, err := jwtauth.FromContext(r.Context())
	userID := int(claims["user_id"].(float64))
	//TODO add err catching

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
