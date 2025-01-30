package controllers_cards_parks

import (
	"net/http"
	services_cards_parks "urbverde-api/services/cards/parks"

	"github.com/gin-gonic/gin"
)

type ParksRankingController struct {
	ParksRankingService services_cards_parks.ParksRankingService
}

func NewParksRankingController(service services_cards_parks.ParksRankingService) *ParksRankingController {
	return &ParksRankingController{
		ParksRankingService: service,
	}
}

func (ac *ParksRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.ParksRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.ParksRankingService.LoadYears(city)
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
