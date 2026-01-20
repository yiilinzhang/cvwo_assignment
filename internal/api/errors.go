package api

import "net/http"

type HTTPError struct {
	Status int
	Err    error
}

func (e HTTPError) Error() string { return e.Err.Error() }
func BadRequest(err error) error   { return HTTPError{Status: http.StatusBadRequest, Err: err} }
func Unauthorized(err error) error { return HTTPError{Status: http.StatusUnauthorized, Err: err} }
//TODO not used yet
func Forbidden(err error) error    { return HTTPError{Status: http.StatusForbidden, Err: err} }

func NotFound(err error) error     { return HTTPError{Status: http.StatusNotFound, Err: err} }
