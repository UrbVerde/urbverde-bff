// urbverde-bff/routes/router.go
package routes

import (
	"net/http"
	_ "urbverde-api/docs"
	"urbverde-api/routes/address"
	"urbverde-api/routes/cards"
	"urbverde-api/routes/categories"

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
		cards.SetupCardsRoutes(v1)
		address.SetupAddressRoutes(v1)
		categories.SetupCategoriesRoutes(v1)
	}

	// Swagger UI route: Register this after all other routes
	config := ginSwagger.Config{
		URL:                      "/swagger/doc.json",
		DeepLinking:              true,
		DefaultModelsExpandDepth: -1,
		DocExpansion:             "list",
	}

	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&config, swaggerFiles.Handler))

	// Optional: Redirect root to Swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	return r
}
