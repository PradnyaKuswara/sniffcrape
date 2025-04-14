package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Jika berhasil
func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Status: true,
		Data:   data,
	})
}

// Jika error
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Status:  false,
		Message: message,
	})
}
