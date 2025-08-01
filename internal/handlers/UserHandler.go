package handlers

import (
	"net/http"

	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUserByID(id string, c *gin.Context) {
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, user)
}
