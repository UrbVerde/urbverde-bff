// urbverde-bff/controllers/cards/temperature/temperature_info_controller.go
package controllers_cards_temperature

import (
	"net/http"
	services_cards_temperature "urbverde-api/services/cards/temperature"

	"github.com/gin-gonic/gin"
)

type TemperatureInfoController struct {
	TemperatureInfoService services_cards_temperature.TemperatureInfoService
}

func NewTemperatureInfoController(service services_cards_temperature.TemperatureInfoService) *TemperatureInfoController {
	return &TemperatureInfoController{
		TemperatureInfoService: service,
	}
}

func (ac *TemperatureInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.TemperatureInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
