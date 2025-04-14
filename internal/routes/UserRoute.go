package routes

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/handlers"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := &repositories.UserRepository{DB: db}
  userService := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			userHandler.GetUserByID(id, c)
		})
	}
}
