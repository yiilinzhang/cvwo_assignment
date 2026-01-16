package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
)

//Check if there is a way to not double declare this in insert post too. Maybe split into 3 and parse here?

func HandleEditPosts(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	var postJSON api.UpdatePostInput
	json.NewDecoder(r.Body).Decode(&postJSON); 
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}
	postJSON.PostId = postId;
	postJSON.UserId = int(claims["user_id"].(float64))
	_, err = dataaccess.UpdatePost(conn, postJSON);
	return &api.Response{
		Payload: api.Payload{
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil

	
}
