package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api" 
)

//Used to check if the user is loggedin based on cookie

func HandleMe (conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error){
	 _, claims, err := jwtauth.FromContext(r.Context())
	userID := int(claims["user_id"].(float64))
	if err != nil {
		return nil, errors.New("Title later")
	}
	//TODO depreciated
	// userid, err := dataaccess.FetchUserName(conn, userID)
	// if err != nil {
	// 	return nil, errors.New("Title later")
	// }
	data, err := json.Marshal(userID)
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"login successful"},
		ErrorCode: 0,
	}, nil
}