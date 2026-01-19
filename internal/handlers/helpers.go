package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
)

func userIDFromContext(r *http.Request) (int, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0, api.Unauthorized(errors.New("Invalid token"))
	}

	userId, ok := claims["user_id"]
	if !ok {
		return 0, api.Unauthorized(errors.New("Missing user_id"))
	}

	f, ok := userId.(float64)
	if !ok {
		return 0, api.Unauthorized(errors.New("invalid user_id claim type"))
	}

	return int(f), nil
}