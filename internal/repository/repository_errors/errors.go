package repository_errors

import "errors"

var (
	ErrNotFound         = errors.New("Not found")
	ErrAlreadyExists    = errors.New("Already exists")
	ErrNotEnoughBalance = errors.New("Not enough balance")
)
