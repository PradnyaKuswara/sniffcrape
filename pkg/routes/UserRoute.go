package routes

import (
	"github.com/PradnyaKuswara/sniffcrape/pkg/handlers"
	"github.com/PradnyaKuswara/sniffcrape/pkg/repositories"
	"github.com/PradnyaKuswara/sniffcrape/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{UserRepository: repo}
	userHandler := &handlers.UserHandler{UserService: userService}

	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			userHandler.GetUserByID(id, c)
		})
	}
}