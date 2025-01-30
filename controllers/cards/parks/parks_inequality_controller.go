package controllers_cards_parks

import (
	"net/http"
	services_cards_parks "urbverde-api/services/cards/parks"

	"github.com/gin-gonic/gin"
)

type ParksInequalityController struct {
	ParksInequalityService services_cards_parks.ParksInequalityService
}

func NewParksInequalityController(service services_cards_parks.ParksInequalityService) *ParksInequalityController {
	return &ParksInequalityController{
		ParksInequalityService: service,
	}
}

func (ac *ParksInequalityController) LoadInequalityData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.ParksInequalityService.LoadInequalityData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de desigualdade em parques e praças",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.ParksInequalityService.LoadYears(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar anos disponíveis de desigualdade",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
