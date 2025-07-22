package routes

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/handlers"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService(&repositories.UserRepository{DB: db})
	authService := services.NewAuthService(userService)
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", func(c *gin.Context) {
			authHandler.Login(c)
		})
		authRoutes.POST("/register", func(c *gin.Context) {
			authHandler.Register(c)
		})
	}
}
