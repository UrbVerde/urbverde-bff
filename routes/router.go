package routes

import (
	"net/http"
	_ "urbverde-api/docs"
	"urbverde-api/routes/address"
	"urbverde-api/routes/tracker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title API Documentation
// @version 1.0
// @description API Swagger para a aplicação Go

// @host localhost:8080
// @BasePath /api/v1
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configurações CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// Define o grupo de rotas com prefixo "/api/v1"
	apiV1 := r.Group("/api/v1")
	{
		address.SetupAddressRoutes(apiV1) // Carrega rotas do módulo "address"

		// Endpoint para listar todas as rotas disponíveis
		apiV1.GET("/endpoints", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"available_endpoints": tracker.AvailableEndpoints,
			})
		})
	}

	r.GET("/swagger/*any", swagger.WrapHandler(files.Handler))

	return r
}
