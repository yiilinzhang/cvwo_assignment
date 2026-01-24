package api

import "net/http"

type HTTPError struct {
	Status int
	Err    error
}

func (e HTTPError) Error() string { return e.Err.Error() }
func BadRequest(err error) error   { return HTTPError{Status: http.StatusBadRequest, Err: err} }
func Unauthorized(err error) error { return HTTPError{Status: http.StatusUnauthorized, Err: err} }
func NotFound(err error) error     { return HTTPError{Status: http.StatusNotFound, Err: err} }
