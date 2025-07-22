package errors

import "errors"

var ErrDataNotFound = errors.New("Data not found")
var ErrInvalidCredentials = errors.New("Invalid credentials")
var ErrInternalServer = errors.New("Internal server error")
var ErrUnauthorized = errors.New("Unauthorized access")
var ErrForbidden = errors.New("Forbidden access")
var ErrBadRequest = errors.New("Bad request")
var ErrConflict = errors.New("Conflict error")
var ErrValidation = errors.New("Validation error")
var ErrDatabase = errors.New("Database error")
