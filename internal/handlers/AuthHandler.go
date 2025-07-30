package handlers

import (
	"net/http"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest models.AuthRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if ok, err := utils.ValidateStruct(loginRequest); !ok {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
	}

	user, token, err := h.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
	}

	response := &models.AuthResponse{
		Token: token,
		User: models.AuthUser{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
			Email:     user.Email,
		},
	}

	utils.RespondWithSuccess(c, http.StatusOK, response)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest models.AuthRegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if ok, err := utils.ValidateStruct(registerRequest); !ok {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
	}

	err := h.AuthService.Register(registerRequest)
	if err != nil {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User registered successfully")
}
