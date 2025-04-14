package setup

import (
	"fmt"
	"os"

	"github.com/PradnyaKuswara/sniffcrape/pkg/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	mode := os.Getenv("GIN_MODE")
	fmt.Println("GIN_MODE:", mode)
    if mode == "" {
        mode = gin.DebugMode // default fallback
    }
	gin.SetMode(mode)

	r := gin.Default()

	// trust proxy
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Set up routes
	routes.RegisterUserRoutes(r, db)

	return r
}