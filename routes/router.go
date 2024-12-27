package routes

import (
	"net/http"
	_ "urbverde-api/docs"
	"urbverde-api/routes/address"
	"urbverde-api/routes/tracker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// API routes under /v1
	v1 := r.Group("/v1")
	{
		address.SetupAddressRoutes(v1)
		v1.GET("/endpoints", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"available_endpoints": tracker.AvailableEndpoints,
			})
		})
	}

	// Swagger UI route: Register this after all other routes
	url := ginSwagger.URL("/swagger/doc.json") // Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Optional: Redirect root to Swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	return r
}
