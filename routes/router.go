// routes/router.go (update the existing file)
package routes

import (
	"net/http"
	"urbverde-api/routes/address"
	"urbverde-api/routes/city" // Add this import
	"urbverde-api/routes/tracker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration (your existing code)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	apiV1 := r.Group("/api/v1")
	{
		address.SetupAddressRoutes(apiV1)
		city.SetupCityRoutes(apiV1) // Add this line

		// Existing endpoints route
		apiV1.GET("/endpoints", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"available_endpoints": tracker.AvailableEndpoints,
			})
		})
	}

	return r
}
