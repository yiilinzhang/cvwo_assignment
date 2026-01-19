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
	ListComments                  = "comments.HandleListByPosts"
	SuccessfulListCommentsMessage = "Successfully listed comments"
	ErrRetrieveComments           = "Failed to retrieve comments in %s"
)

func HandleListByPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid post id"))
	}

	postList, err := dataaccess.ListCommentByPost(conn, postId)

	data, err := json.Marshal(postList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}


func HandleCommentById(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid comment id"))
	}

	comment, err := dataaccess.ListCommentById(conn, commentId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListComments))
	}

	data, err := json.Marshal(comment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}


// Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid comment id"))
	}

	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input api.DeleteCommentInput
	input.CommentId = commentId
	input.UserId = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = dataaccess.DeleteComment(conn, input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}


//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleEditComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid comment id"))
	}

	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input api.UpdateCommentInput
	if err := decodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.CommentId = commentId
	input.UserId = userID
	_, err = dataaccess.UpdateComment(conn, input)

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{},
	}, nil

}


//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleInsertComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	PostId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid post id"))
	}

	userID, err := userIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input api.CreateCommentInput
	if err := decodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserId = userID
	input.PostId = PostId

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newComment, err := dataaccess.InsertComment(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListComments))
	}

	//TODO check if more efficient to not unmardshell
	data, err := json.Marshal(newComment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}
