package controllers_cards_weather

import (
	"fmt"
	"net/http"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

type WeatherTemperatureController struct {
	WeatherTemperatureService services_cards_weather.WeatherTemperatureService
}

func NewWeatherTemperatureController(service services_cards_weather.WeatherTemperatureService) *WeatherTemperatureController {
	return &WeatherTemperatureController{
		WeatherTemperatureService: service,
	}
}

func (ac *WeatherTemperatureController) LoadTemperatureData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")
	fmt.Println(city, year)

	if year != "" {
		data, err := ac.WeatherTemperatureService.LoadData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.WeatherTemperatureService.LoadYears(city)
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
