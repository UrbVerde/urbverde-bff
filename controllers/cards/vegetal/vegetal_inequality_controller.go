package controllers_cards_vegetal

import (
	"net/http"
	services_cards_vegetal "urbverde-api/services/cards/vegetal"

	"github.com/gin-gonic/gin"
)

type VegetalInequalityController struct {
	VegetalInequalityService services_cards_vegetal.VegetalInequalityService
}

func NewVegetalInequalityController(service services_cards_vegetal.VegetalInequalityService) *VegetalInequalityController {
	return &VegetalInequalityController{
		VegetalInequalityService: service,
	}
}

func (ac *VegetalInequalityController) LoadInequalityData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetalInequalityService.LoadInequalityData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura vegetal",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetalInequalityService.LoadYears(city)
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
