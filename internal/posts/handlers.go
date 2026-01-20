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

type Handler struct {
	Conn *pgxpool.Pool
}

const (
	SuccessfulCreatePostsMessage = "Successfully created post"
	SuccessfulListPostsMessage   = "Successfully listed posts"
	SuccessfulUpdatePostsMessage = "Successfully updated post"
	SuccessfulDeletePostsMessage = "Successfully deleted post"

	ErrCreatePosts   = "Failed to create post in %s"
	ErrRetrievePosts = "Failed to retrieve posts in %s"
	ErrUpdatePosts   = "Failed to update post in %s"
	ErrDeletePosts   = "Failed to delete post in %s"

	ErrEncodeView = "Failed to encode posts in %s"
)

func (h *Handler) ListByTopic(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	topicID, err := strconv.Atoi(chi.URLParam(r, "topicID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("invalid topicID"))
	}

	postList, err := ListPostByTopic(h.Conn, topicID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, "posts.ListByTopics"))
	}

	data, err := json.Marshal(postList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "posts.ListByTopics"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("invalid postID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input DeletePostInput
	input.PostID = postID
	input.UserID = userID

	_, err = DeletePost(h.Conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDeletePosts, "posts.Delete"))
	}

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulDeletePostsMessage},
	}, nil
}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("invalid postID"))
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

	_, err = UpdatePost(h.Conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrUpdatePosts, "posts.Update"))
	}

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulUpdatePostsMessage},
	}, nil

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input CreatePostInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserID = userID

	newPost, err := InsertPost(h.Conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreatePosts, "posts.Create"))
	}

	data, err := json.Marshal(newPost)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "posts.Create"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulCreatePostsMessage},
	}, nil
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postList, err := ListAllPost(h.Conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, "posts.List"))
	}

	data, err := json.Marshal(postList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "posts.List"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}
