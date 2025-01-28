package controllers_cards_square

import (
	"net/http"
	services_cards_square "urbverde-api/services/cards/square"

	"github.com/gin-gonic/gin"
)

type SquareInequalityController struct {
	SquareInequalityService services_cards_square.SquareInequalityService
}

func NewSquareInequalityController(service services_cards_square.SquareInequalityService) *SquareInequalityController {
	return &SquareInequalityController{
		SquareInequalityService: service,
	}
}

func (ac *SquareInequalityController) LoadInequalityData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.SquareInequalityService.LoadInequalityData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura square",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.SquareInequalityService.LoadYears(city)
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
