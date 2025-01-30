// urbverde-bff/controllers/cards/temperature/temperature_ranking_controller.go
package controllers_cards_temperature

import (
	"net/http"
	services_cards_temperature "urbverde-api/services/cards/temperature"

	"github.com/gin-gonic/gin"
)

type TemperatureRankingController struct {
	TemperatureRankingService services_cards_temperature.TemperatureRankingService
}

func NewTemperatureRankingController(service services_cards_temperature.TemperatureRankingService) *TemperatureRankingController {
	return &TemperatureRankingController{
		TemperatureRankingService: service,
	}
}

func (ac *TemperatureRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.TemperatureRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.TemperatureRankingService.LoadYears(city)
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
