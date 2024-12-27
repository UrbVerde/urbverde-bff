package city

import (
	"urbverde-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCityRoutes(rg *gin.RouterGroup) {
	cityController := controllers.NewCityController()

	// City data endpoint
	rg.GET("/city/data", cityController.GetCityData)
}
