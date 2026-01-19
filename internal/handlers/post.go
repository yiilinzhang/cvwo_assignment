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
	ListPosts                  = "posts.HandleList"
	SuccessfulListPostsMessage = "Successfully listed posts"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
)

func HandleListByTopic(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {

	topicId, err := strconv.Atoi(chi.URLParam(r, "topicId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid post id"))
	}

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

// Delete a specific post from database, requires postid to be passed in via url
func HandleDeletePosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid post id"))
	}

	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}
	//TODO add err catching

	//Extract JSON from HTTP req
	var input api.DeletePostInput
	input.PostId = postId
	input.UserId = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = dataaccess.DeletePost(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	return nil, nil
}

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleEditPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid post id"))
	}

	var input api.UpdatePostInput
	if err := decodeJSON(r, &input); err != nil {
		return nil, err
	}
	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	input.PostId = postId
	input.UserId = userID
	_, err = dataaccess.UpdatePost(conn, input)
	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil

}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleInsertPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input api.CreatePostInput
	if err := decodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserId = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newPost, err := dataaccess.InsertPost(conn, input)
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
