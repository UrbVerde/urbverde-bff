package controllers_cards_vegetal

import (
	"net/http"
	services_cards_vegetal "urbverde-api/services/cards/vegetal"

	"github.com/gin-gonic/gin"
)

type VegetalRankingController struct {
	VegetalRankingService services_cards_vegetal.VegetalRankingService
}

func NewVegetalRankingController(service services_cards_vegetal.VegetalRankingService) *VegetalRankingController {
	return &VegetalRankingController{
		VegetalRankingService: service,
	}
}

func (ac *VegetalRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetalRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetalRankingService.LoadYears(city)
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
