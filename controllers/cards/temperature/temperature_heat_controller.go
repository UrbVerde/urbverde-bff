// urbverde-bff/controllers/cards/temperature/temperature_heat_controller.go
package controllers_cards_temperature

import (
	"net/http"
	services_cards_temperature "urbverde-api/services/cards/temperature"

	"github.com/gin-gonic/gin"
)

type TemperatureHeatController struct {
	TemperatureHeatService services_cards_temperature.TemperatureHeatService
}

func NewTemperatureHeatController(service services_cards_temperature.TemperatureHeatService) *TemperatureHeatController {
	return &TemperatureHeatController{
		TemperatureHeatService: service,
	}
}

func (ac *TemperatureHeatController) LoadHeatData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.TemperatureHeatService.LoadHeatData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.TemperatureHeatService.LoadYears(city)
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
