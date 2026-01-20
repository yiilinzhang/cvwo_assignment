package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/auth"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers"
	"golang.org/x/crypto/bcrypt"
)

const (
	ListUsers                  = "HandleList"
	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func HandleListUsers(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userList, err := ListUser(conn)
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
	var input CreateUserInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}

	//TODO add empty user pw validation
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = ""
	s := string(hashBytes)

	err = InsertUser(conn, input.Username, s)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}
	return nil, err

}

// return nil if password matches
func HandleLoginAuth(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var loginCred CreateUserInput
	if err := handlers.DecodeJSON(r, &loginCred); err != nil {
		return nil, api.BadRequest(errors.New("Invalid login cred"))
	}

	userID, savedPass, err := FetchUser(conn, loginCred.Username)
	//TODO change this err message
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	err = bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(loginCred.Password))
	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	_, jwtString, _ := auth.TokenAuth.Encode(map[string]interface{}{
		"user_id": userID})

	data, err := json.Marshal(jwtString)
	if err != nil {
		return nil, errors.New("Failed to convert JWT to JSON")
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwtString,
		Path:     "/",
		MaxAge:   604800,
		HttpOnly: true,
		//TODO Need to change this in prod to true
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages:  []string{"login successful"},
		ErrorCode: 0,
	}, nil

}

func HandleLogout(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		//TODO Need to change this in prod to true
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	
	return &api.Response{
		Payload: api.Payload{
			Data: nil,
		},
		Messages:  []string{"logout successful"},
		ErrorCode: 0,
	}, nil
}

//Used to check if the user is loggedin based on cookie

func HandleMe(conn *pgxpool.Pool, w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(userID)
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages:  []string{"login successful"},
		ErrorCode: 0,
	}, nil
}
