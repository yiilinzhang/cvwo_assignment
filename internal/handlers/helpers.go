package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/yiilinzhang/cvwo_assignment/internal/api"
)

// Helper func used to retreive userID from contect, return id and error
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

//Abtract out error catching in json body decoding. Used to replace json.NewDecoder.Decode + err logic
func decodeJSON(r *http.Request, destination any) error {
	if err := json.NewDecoder(r.Body).Decode(destination); err != nil {
		return api.BadRequest(errors.New("JSON body is invalid"))
	}
	return nil
}