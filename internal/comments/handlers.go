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

type Handler struct{ Conn *pgxpool.Pool }

const (
	SuccessfulCreateCommentsMessage = "Successfully listed comments"
	SuccessfulListCommentsMessage   = "Successfully listed comments"
	SuccessfulUpdateCommentsMessage = "Successfully updated comments"
	SuccessfulDeleteCommentsMessage = "Successfully deleted comments"

	ErrCreateComments   = "Failed to create comment in %s"
	ErrRetrieveComments = "Failed to retrieve comments in %s"
	ErrUpdateComments   = "Failed to update comment in %s"
	ErrDeleteComments   = "Failed to delete comment in %s"
	ErrEncodeComments       = "Failed to encode comment in %s"
)

func (h *Handler) ListByPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid postID"))
	}

	postList, err := ListCommentByPost(h.Conn, postID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, "comments.ListByPosts"))
	}

	data, err := json.Marshal(postList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeComments, "comments.ListByPosts"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}

func (h *Handler) ListByID(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid commentID"))
	}

	comment, err := ListCommentByID(h.Conn, commentID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, "comments.ListByID"))
	}

	data, err := json.Marshal(comment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeComments, "comments.ListByID"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}

// Delete a specific post from database, requires postid to be passed in via url
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "commentID"))
	if err != nil {
		return nil, api.BadRequest(errors.New("Invalid commentID"))
	}

	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDeleteComments, "comments.Delete"))
	}

	var input DeleteCommentInput
	input.CommentID = commentID
	input.UserID = userID

	err = DeleteComment(h.Conn, input)
	if err != nil {
		return nil, err
	}
	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulDeleteCommentsMessage},
	}, nil
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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

	if err = UpdateComment(h.Conn, input); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrUpdateComments, "comments.Update"))
	}

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{SuccessfulUpdateCommentsMessage},
	}, nil

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
	if err = InsertComment(h.Conn, input); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateComments, "comments.Create"))
	}

	return &api.Response{
		Payload: api.Payload{},
		Messages: []string{SuccessfulCreateCommentsMessage},
	}, nil
}
