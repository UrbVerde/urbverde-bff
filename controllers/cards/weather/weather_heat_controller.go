// urbverde-bff/controllers/cards/weather/weather_heat_controller.go
package controllers_cards_weather

import (
	"net/http"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

type WeatherHeatController struct {
	WeatherHeatService services_cards_weather.WeatherHeatService
}

func NewWeatherHeatController(service services_cards_weather.WeatherHeatService) *WeatherHeatController {
	return &WeatherHeatController{
		WeatherHeatService: service,
	}
}

func (ac *WeatherHeatController) LoadHeatData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.WeatherHeatService.LoadHeatData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.WeatherHeatService.LoadYears(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar anos disponíveis de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
