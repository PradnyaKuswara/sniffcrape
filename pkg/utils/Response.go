package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithSuccess(c *gin.Context, code int, data interface{}, message ...string) {
	msg := "Success"

	if len(message) > 0 {
		msg = strings.Join(message, " ")
	}

	c.JSON(code, Response{
		Message: msg,
		Status:  code,
		Data:    data,
	})
}

func RespondWithError(c *gin.Context, code int, message ...string) {
	msg := "Something went wrong"
	if len(message) > 0 {
		msg = strings.Join(message, " ")
	}

	c.JSON(code, Response{
		Status:  code,
		Message: msg,
	})
}
