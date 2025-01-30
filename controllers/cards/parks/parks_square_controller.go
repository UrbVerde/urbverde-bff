package controllers_cards_parks

import (
	"net/http"
	services_cards_parks "urbverde-api/services/cards/parks"

	"github.com/gin-gonic/gin"
)

type ParksSquareController struct {
	ParksSquareService services_cards_parks.ParksSquareService
}

func NewParksSquareController(service services_cards_parks.ParksSquareService) *ParksSquareController {
	return &ParksSquareController{
		ParksSquareService: service,
	}
}

func (ac *ParksSquareController) LoadSquareData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.ParksSquareService.LoadSquareData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de praças e parques",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.ParksSquareService.LoadYears(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar anos disponíveis de praças e parques",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
