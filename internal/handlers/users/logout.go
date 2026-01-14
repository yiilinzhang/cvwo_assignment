package users

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yiilinzhang/cvwo_assignment/internal/api" 
)

//Used to check if the user is loggedin based on cookie

func HandleLogout (conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error){
	 cookie := http.Cookie{
		Name: "jwt",
		Value: "",
		Path: "/",
		MaxAge: -1,
		HttpOnly: true,
		//TODO Need to change this in prod to true
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return &api.Response{
		Payload: api.Payload{
			Data: nil,
		},
		Messages: []string{"login successful"},
		ErrorCode: 0,
	}, nil
}