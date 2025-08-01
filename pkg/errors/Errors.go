package errors

import "errors"

var ErrDataNotFound = errors.New("Data not found")
var ErrInvalidCredentials = errors.New("Invalid credentials")
var ErrInternalServer = errors.New("Internal server error")
var ErrUnauthorized = errors.New("Unauthorized access")
var ErrTokenExpired = errors.New("Token expired")
var ErrTokenInvalid = errors.New("Token is invalid")
var ErrTokenMalformed = errors.New("Token is malformed")
var ErrTokenSignatureInvalid = errors.New("Token signature is invalid")
var ErrUnauthenticated = errors.New("Unauthenticated access")
var ErrForbidden = errors.New("Forbidden access")
var ErrBadRequest = errors.New("Bad request")
var ErrConflict = errors.New("Conflict error")
var ErrValidation = errors.New("Validation error")
var ErrDatabase = errors.New("Database error")
