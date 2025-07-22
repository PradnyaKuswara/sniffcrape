package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	customerrors "github.com/PradnyaKuswara/sniffcrape/pkg/errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func MapErrorToStatusCode(err error) (int, string) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound), errors.Is(err, customerrors.ErrDataNotFound):
		return http.StatusNotFound, customerrors.ErrDataNotFound.Error()

	case strings.Contains(strings.ToLower(err.Error()), "invalid id"):
		return http.StatusBadRequest, "Invalid ID format"

	case errors.Is(err, io.EOF):
		return http.StatusBadRequest, "Request body is empty"

	case isValidationError(err):
		return http.StatusBadRequest, formatValidationError(err)

	case errors.Is(err, customerrors.ErrUnauthorized):
		return http.StatusUnauthorized, customerrors.ErrUnauthorized.Error()

	case errors.Is(err, customerrors.ErrForbidden):
		return http.StatusForbidden, customerrors.ErrForbidden.Error()

	case errors.Is(err, customerrors.ErrInvalidCredentials):
		return http.StatusUnauthorized, customerrors.ErrInvalidCredentials.Error()

	default:
		return http.StatusInternalServerError, err.Error()
	}
}

func isValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func formatValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var sb strings.Builder
		for _, fieldErr := range validationErrors {
			sb.WriteString(fmt.Sprintf("Field '%s' failed on the '%s' rule; ", fieldErr.Field(), fieldErr.Tag()))
		}
		return strings.TrimSuffix(sb.String(), "; ")
	}
	return "Validation failed"
}
