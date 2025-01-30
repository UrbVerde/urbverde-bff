package controllers_cards_vegetation

import (
	"net/http"
	services_cards_vegetation "urbverde-api/services/cards/vegetation"

	"github.com/gin-gonic/gin"
)

type VegetationRankingController struct {
	VegetationRankingService services_cards_vegetation.VegetationRankingService
}

func NewVegetationRankingController(service services_cards_vegetation.VegetationRankingService) *VegetationRankingController {
	return &VegetationRankingController{
		VegetationRankingService: service,
	}
}

func (ac *VegetationRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetationRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetationRankingService.LoadYears(city)
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
