package app

import "errors"

var ErrNotFound = errors.New("not found")

var ErrUnauthorized = errors.New("unauthorized")

var ErrDuplicateEntry = errors.New("duplicate entry")
