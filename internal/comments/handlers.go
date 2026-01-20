package comments

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
	ListComments                  = "comments.HandleListByPosts"
	SuccessfulListCommentsMessage = "Successfully listed comments"
	ErrRetrieveComments           = "Failed to retrieve comments in %s"
	ErrEncodeView                 = "Failed to retrieve comment in %s"
)

func HandleListByPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid postID"))
	}

	postList, err := ListCommentByPost(conn, postID)

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

func HandleCommentByID(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid commentID"))
	}

	comment, err := ListCommentByID(conn, commentID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, ListComments))
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
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid commentID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input DeleteCommentInput
	input.CommentID = commentID
	input.UserID = userID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = DeleteComment(conn, input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleEditComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid commentID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input UpdateCommentInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.CommentID = commentID
	input.UserID = userID
	_, err = UpdateComment(conn, input)

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{},
	}, nil

}

// Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?
func HandleInsertComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	PostID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid postID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}

	var input CreateCommentInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	input.UserID = userID
	input.PostID = PostID

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newComment, err := InsertComment(conn, input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, ListComments))
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
