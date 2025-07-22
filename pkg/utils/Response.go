package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Jika berhasil
func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Status: code,
		Data:   data,
	})
}

// Jika error
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status:  code,
		Message: message,
	})
}
