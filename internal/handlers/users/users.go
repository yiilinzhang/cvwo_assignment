package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/auth"
	"github.com/yiilinzhang/cvwo_assignment/internal/dataaccess"
	"golang.org/x/crypto/bcrypt"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)



func HandleList(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userList, err := dataaccess.ListUser(conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	data, err := json.Marshal(userList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

func HandleAddUsers(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	//Extract JSON from HTTP req
	var userJSON api.CreateUserInput
	json.NewDecoder(r.Body).Decode(&userJSON)

	//TODO rmv prnt statm
	//TODO check if i should shift API here
	//TODO add empty user pw validation
	log.Println(userJSON)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(userJSON.Password), bcrypt.DefaultCost)
	userJSON.Password=""
	s := string(hashBytes)
	err = dataaccess.InsertUser(conn, userJSON.Username, s)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}
	return nil, err

}

//return nil if password matches
func HandleLoginAuth (conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error){
	//TODO might need to rename this api model
	//JSON->GO conversion of http body
	var loginCred api.CreateUserInput
	json.NewDecoder(r.Body).Decode(&loginCred)
	userID, savedPass, err := dataaccess.FetchUser(conn, loginCred.Username)
	//TODO change this err message
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}
	//Compare passwords returns nil if matches, err otherwise
	err = bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(loginCred.Password))
	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}
	

	//create a JWT if auth suceeds
	_, jwtString, _ := auth.TokenAuth.Encode(map[string]interface{}{
    "user_id": userID,})


	data, err := json.Marshal(jwtString)
	if err != nil {
		return nil, errors.New("Failed to convert JWT to JSON")
	}
	cookie := http.Cookie{
		Name: "auth_token",
		Value: jwtString,
		Path: "/",
		MaxAge: 604800,
		HttpOnly: true,
		//TODO Need to change this in prod to true
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"login successful"},
		ErrorCode: 0,
	}, nil
	
}
