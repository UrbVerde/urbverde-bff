package controllers_cards_vegetation

import (
	"net/http"
	services_cards_vegetation "urbverde-api/services/cards/vegetation"

	"github.com/gin-gonic/gin"
)

type VegetationInequalityController struct {
	VegetationInequalityService services_cards_vegetation.VegetationInequalityService
}

func NewVegetationInequalityController(service services_cards_vegetation.VegetationInequalityService) *VegetationInequalityController {
	return &VegetationInequalityController{
		VegetationInequalityService: service,
	}
}

func (ac *VegetationInequalityController) LoadInequalityData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetationInequalityService.LoadInequalityData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura vegetal",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetationInequalityService.LoadYears(city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar anos disponíveis de cobertura vegetal",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
