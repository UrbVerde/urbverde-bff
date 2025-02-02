package controllers_cards_square

import (
	"net/http"
	services_cards_square "urbverde-api/services/cards/square"

	"github.com/gin-gonic/gin"
)

type SquareParksController struct {
	SquareParksService services_cards_square.SquareParksService
}

func NewSquareParksController(service services_cards_square.SquareParksService) *SquareParksController {
	return &SquareParksController{
		SquareParksService: service,
	}
}

func (ac *SquareParksController) LoadParksData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.SquareParksService.LoadParksData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura square",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.SquareParksService.LoadYears(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar anos disponíveis de cobertura square",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
