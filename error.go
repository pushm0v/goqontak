package goqontak

import "errors"

var (
	ErrStatusMethodNotAllowed    = errors.New("Method not Allowd")
	ErrStatusNotFound            = errors.New("Not Found")
	ErrStatusUnauthorized        = errors.New("Unauthorized")
	ErrStatusInternalServerError = errors.New("Internal Server Error")
	ErrStatusBadGateway          = errors.New("Bad Gateway")
	ErrStatusBadRequest          = errors.New("Bad Request")
	ErrStatusUnprocessableEntity = errors.New("Unprocessable Entity")
)
