package cards

import (
	"urbverde-api/controllers"
	"urbverde-api/repositories"
	"urbverde-api/services"

	"github.com/gin-gonic/gin"
)

func SetupCardsRoutes(rg *gin.RouterGroup) {
	// Weather

	// Temperature
	tempeRepo := repositories.NewExternalWeatherTemperatureRepository()
	tempeService := services.NewWeatherTemperatureService(tempeRepo)
	tempeController := controllers.NewWeatherTemperatureController(tempeService)

	rg.GET("/cards/weather/temperature", tempeController.LoadTemperatureData)
	// http://localhost:8080/api/v1/cards/weather/temperature?city=3549003&year=2022
}
