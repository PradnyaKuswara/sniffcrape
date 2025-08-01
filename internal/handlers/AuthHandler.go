package handlers

import (
	"net/http"
	"strings"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	customerrors "github.com/PradnyaKuswara/sniffcrape/pkg/errors"
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

	response, err := h.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		return
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

func (h *AuthHandler) ValidateUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		status, message := utils.MapErrorToStatusCode(customerrors.ErrUnauthenticated)
		utils.RespondWithError(c, status, message)
		c.Abort()
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	authUser, err := h.AuthService.ValidateUser(tokenStr)

	if err != nil {
		status, message := utils.MapErrorToStatusCode(err)
		utils.RespondWithError(c, status, message)
		c.Abort()
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, authUser)
}
