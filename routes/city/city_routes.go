// routes/city/city_routes.go
package city

import (
	"urbverde-api/controllers"
	"urbverde-api/routes/tracker"

	"github.com/gin-gonic/gin"
)

func SetupCityRoutes(rg *gin.RouterGroup) {
	cityController := controllers.NewCityController()

	// City boundaries endpoint
	rg.GET("/city/bounds", cityController.GetCityBounds)
	tracker.AddEndpoint("GET", "/api/v1/city/bounds", "Retorna os dados de localização e código do município")

	// City data endpoint
	rg.GET("/data", cityController.GetCityData)
	tracker.AddEndpoint("GET", "/api/v1/data", "Retorna as camadas e categorias disponíveis para o código do município")
}
