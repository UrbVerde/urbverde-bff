package controllers_cards_vegetation

import (
	"net/http"
	services_cards_vegetation "urbverde-api/services/cards/vegetation"

	"github.com/gin-gonic/gin"
)

type VegetationInfoController struct {
	VegetationInfoService services_cards_vegetation.VegetationInfoService
}

func NewVegetationInfoController(service services_cards_vegetation.VegetationInfoService) *VegetationInfoController {
	return &VegetationInfoController{
		VegetationInfoService: service,
	}
}

func (ac *VegetationInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.VegetationInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
