// urbverde-bff/controllers/cards/temperature/temperature_weather_controller.go
package controllers_cards_temperature

import (
	"net/http"
	services_cards_temperature "urbverde-api/services/cards/temperature"

	"github.com/gin-gonic/gin"
)

type TemperatureWeatherController struct {
	TemperatureWeatherService services_cards_temperature.TemperatureWeatherService
}

func NewTemperatureWeatherController(service services_cards_temperature.TemperatureWeatherService) *TemperatureWeatherController {
	return &TemperatureWeatherController{
		TemperatureWeatherService: service,
	}
}

func (ac *TemperatureWeatherController) LoadWeatherData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.TemperatureWeatherService.LoadWeatherData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.TemperatureWeatherService.LoadYears(city)
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
