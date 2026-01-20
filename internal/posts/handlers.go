package posts

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
	ListPosts                  = "posts.HandleList"
	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
	ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleListByTopic(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	topicID, err := strconv.Atoi(chi.URLParam(r, "topicID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid topicID"))
	}

	postList, err := ListPostByTopic(conn, topicID)
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

// Delete a specific post from database, requires postid to be passed in via url
func HandleDeletePosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid postID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input DeletePostInput
	input.PostID = postID
	input.UserID = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = DeletePost(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	return nil, nil
}

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleEditPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid postID"))
	}

	var input UpdatePostInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	input.PostID = postID
	input.UserID = userID
	_, err = UpdatePost(conn, input)
	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil

}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleInsertPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input CreatePostInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserID = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newPost, err := InsertPost(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	//TODO check if more efficient to not unmardshell
	data, err := json.Marshal(newPost)
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

func HandleListAllPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postList, err := ListAllPost(conn)
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
