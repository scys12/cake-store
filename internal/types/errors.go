package types

import "errors"

var (
	ErrFailedInsertCake = errors.New("failed to insert cake")
	ErrFailedUpdateCake = errors.New("failed to update cake")
	ErrFailedDeleteCake = errors.New("failed to delete cake")
	ErrGetCakeDetail    = errors.New("failed to get cake detail")

	ErrFailedGetLastIDDB  = errors.New("failed to get last inserted id")
	ErrNoRowsAffected     = errors.New("there is not any rows affected")
	ErrFailedRowsAffected = errors.New("failed to get rows affected")

	ErrBadRequest = errors.New("error bad request")
)
