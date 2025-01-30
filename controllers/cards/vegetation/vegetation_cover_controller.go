package controllers_cards_vegetation

import (
	"net/http"
	services_cards_vegetation "urbverde-api/services/cards/vegetation"

	"github.com/gin-gonic/gin"
)

type VegetationCoverController struct {
	VegetationCoverService services_cards_vegetation.VegetationCoverService
}

func NewVegetationCoverController(service services_cards_vegetation.VegetationCoverService) *VegetationCoverController {
	return &VegetationCoverController{
		VegetationCoverService: service,
	}
}

func (ac *VegetationCoverController) LoadCoverData(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")

	if year != "" {
		data, err := ac.VegetationCoverService.LoadCoverData(city, year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao buscar dados de cobertura vegetal",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	} else {
		data, err := ac.VegetationCoverService.LoadYears(city)
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
