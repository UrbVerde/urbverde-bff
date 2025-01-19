// urbverde-bff/controllers/cards/weather/weather_info_controller.go
package controllers_cards_weather

import (
	"net/http"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

type WeatherInfoController struct {
	WeatherInfoService services_cards_weather.WeatherInfoService
}

func NewWeatherInfoController(service services_cards_weather.WeatherInfoService) *WeatherInfoController {
	return &WeatherInfoController{
		WeatherInfoService: service,
	}
}

func (ac *WeatherInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.WeatherInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
