package cards

import (
	controllers_cards_weather "urbverde-api/controllers/cards/weather"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

func SetupCardsRoutes(rg *gin.RouterGroup) {
	// Weather

	// Temperature
	tempeRepo := repositories_cards_weather.NewExternalWeatherTemperatureRepository()
	tempeService := services_cards_weather.NewWeatherTemperatureService(tempeRepo)
	tempeController := controllers_cards_weather.NewWeatherTemperatureController(tempeService)

	rg.GET("/cards/weather/temperature", tempeController.LoadTemperatureData)
	// http://localhost:8080/api/v1/cards/weather/temperature?city=3549003&year=2022

	// Heat
	heatRepo := repositories_cards_weather.NewExternalWeatherHeatRepository()
	heatService := services_cards_weather.NewWeatherHeatService(heatRepo)
	heatController := controllers_cards_weather.NewWeatherHeatController(heatService)

	rg.GET("/cards/weather/heat", heatController.LoadHeatData)
	// http://localhost:8080/api/v1/cards/weather/heat?city=3549003&year=2022
}
