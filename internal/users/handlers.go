package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
	"github.com/yiilinzhang/cvwo_assignment/internal/auth"
	"github.com/yiilinzhang/cvwo_assignment/internal/handlers"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Conn *pgxpool.Pool
}

const (
	SuccessfulListUsersMessage = "Successfully listed users"

	ErrRetrieveUsers = "Failed to retrieve users in %s"
	ErrCreateUsers   = "Failed to create user in %s"

	ErrEncodeView = "Failed to encode users in %s"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userList, err := ListUser(h.Conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, "users.List"))
	}

	data, err := json.Marshal(userList)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, "users.List"))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var input CreateUserInput
	if err := handlers.DecodeJSON(r, &input); err != nil {
		return nil, err
	}
	if input.Username == "" || input.Password == "" {
		return nil, api.BadRequest(errors.New("username and password are required"))
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateUsers, "users.Create"))
	}

	input.Password = ""

	s := string(hashBytes)
	err = InsertUser(h.Conn, input.Username, s)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateUsers, "users.Create"))
	}
	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{"signup successful"},
	}, nil

}

// return nil if password matches
func (h *Handler) LoginAuth(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var loginCred CreateUserInput
	if err := handlers.DecodeJSON(r, &loginCred); err != nil {
		return nil, api.BadRequest(errors.New("invalid login cred"))
	}

	user, err := FetchUser(h.Conn, loginCred.Username)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, "users.LoginAuth"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginCred.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	_, jwtString, _ := auth.TokenAuth.Encode(map[string]interface{}{
		"user_id": user.UserID})

	isProd := os.Getenv("ENV") == "production"
	sameSite := http.SameSiteLaxMode
	if isProd {
		sameSite = http.SameSiteNoneMode
	}
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwtString,
		Path:     "/",
		MaxAge:   604800,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: sameSite,
	}

	http.SetCookie(w, &cookie)

	return &api.Response{
		Payload:  api.Payload{},
		Messages: []string{"login successful"},
	}, nil

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	isProd := os.Getenv("ENV") == "production"
	sameSite := http.SameSiteLaxMode
	if isProd {
		sameSite = http.SameSiteNoneMode
	}
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: sameSite,
	}

	http.SetCookie(w, &cookie)

	return &api.Response{
		Payload: api.Payload{
			Data: nil,
		},
		Messages: []string{"logout successful"},
	}, nil
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	userID, err := handlers.UserIDFromContext(r)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(userID)
	if err != nil {
		return nil, err
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{"me okay"},
	}, nil
}
