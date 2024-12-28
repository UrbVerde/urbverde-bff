package routes

import (
	"net/http"
	"urbverde-api/routes/address"
	"urbverde-api/routes/cards"
	"urbverde-api/routes/tracker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
		cards.SetupCardsRoutes(apiV1)

		// Endpoint para listar todas as rotas disponíveis
		apiV1.GET("/endpoints", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"available_endpoints": tracker.AvailableEndpoints,
			})
		})
	}

	return r
}
