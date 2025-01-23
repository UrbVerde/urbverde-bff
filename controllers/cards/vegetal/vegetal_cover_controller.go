package controllers_cards_vegetal

import (
	"net/http"
	services_cards_vegetal "urbverde-api/services/cards/vegetal"

	"github.com/gin-gonic/gin"
)

type VegetalCoverController struct {
	VegetalCoverService services_cards_vegetal.VegetalCoverService
}

func NewVegetalCoverController(service services_cards_vegetal.VegetalCoverService) *VegetalCoverController {
	return &VegetalCoverController{
		VegetalCoverService: service,
	}
}

func (ac *VegetalCoverController) LoadCoverData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetalCoverService.LoadCoverData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura vegetal",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetalCoverService.LoadYears(city)
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
