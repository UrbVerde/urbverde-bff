package controllers_cards_weather

import (
	"net/http"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

type WeatherRankingController struct {
	WeatherRankingService services_cards_weather.WeatherRankingService
}

func NewWeatherRankingController(service services_cards_weather.WeatherRankingService) *WeatherRankingController {
	return &WeatherRankingController{
		WeatherRankingService: service,
	}
}

func (ac *WeatherRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.WeatherRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.WeatherRankingService.LoadYears(city)
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
