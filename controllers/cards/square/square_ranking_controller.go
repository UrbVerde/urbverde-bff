package controllers_cards_square

import (
	"net/http"
	services_cards_square "urbverde-api/services/cards/square"

	"github.com/gin-gonic/gin"
)

type SquareRankingController struct {
	SquareRankingService services_cards_square.SquareRankingService
}

func NewSquareRankingController(service services_cards_square.SquareRankingService) *SquareRankingController {
	return &SquareRankingController{
		SquareRankingService: service,
	}
}

func (ac *SquareRankingController) LoadRankingData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.SquareRankingService.LoadRankingData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de temperatura",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.SquareRankingService.LoadYears(city)
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
