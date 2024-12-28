package controllers

import (
	"fmt"
	"net/http"
	"urbverde-api/services"

	"github.com/gin-gonic/gin"
)

type WeatherTemperatureController struct {
	WeatherTemperatureService services.WeatherTemperatureService
}

func NewWeatherTemperatureController(service services.WeatherTemperatureService) *WeatherTemperatureController {
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
