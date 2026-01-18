package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

const (
	ListComments = "comments.HandleListByPosts"

	SuccessfulListCommentsMessage = "Successfully listed comments"
	ErrRetrieveComments           = "Failed to retrieve comments in %s"
)

func HandleListByPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postId := chi.URLParam(r, "postId")
	postList, err := dataaccess.ListCommentByPost(conn, postId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListComments))
	}

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
	commentId := chi.URLParam(r, "commentId")
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

//Delete a specific post from database, requires postid to be passed in via url
func HandleDeleteComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	_, claims, err := jwtauth.FromContext(r.Context())
	userID := int(claims["user_id"].(float64))
	//TODO add err catching
	//TODO retitlr these JSON things to paylaod
	//Extract JSON from HTTP req
	var commentJSON api.DeleteCommentInput 
	commentJSON.CommentId = commentId
	commentJSON.UserId = userID


	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	_, err = dataaccess.DeleteComment(conn, commentJSON)
	if err != nil {
		return nil, err
	}


	return nil, nil
}


//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleEditComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	var commentJSON api.UpdateCommentInput
	json.NewDecoder(r.Body).Decode(&commentJSON); 
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}
	commentJSON.CommentId = commentId;
	commentJSON.UserId = int(claims["user_id"].(float64))
	_, err = dataaccess.UpdateComment(conn, commentJSON);
	return &api.Response{
		Payload: api.Payload{
		},
		Messages: []string{},
	}, nil

	
}

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleInsertComments(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var commentJSON api.CreateCommentInput
	json.NewDecoder(r.Body).Decode(&commentJSON)

commentJSON.PostId = chi.URLParam(r, "postId")
	_, claims, err := jwtauth.FromContext(r.Context())
	commentJSON.UserId = int(claims["user_id"].(float64))
	if err != nil {
		return nil, errors.New("Title later")
	}
	//Validate input

	//TODO adjust this after i decide if i wna tot return anything to fornt end chekc if okay to leave just return err
	newComment, err := dataaccess.InsertComment(conn, commentJSON)
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
