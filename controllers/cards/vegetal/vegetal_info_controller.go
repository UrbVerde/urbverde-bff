package controllers_cards_vegetal

import (
	"net/http"
	services_cards_vegetal "urbverde-api/services/cards/vegetal"

	"github.com/gin-gonic/gin"
)

type VegetalInfoController struct {
	VegetalInfoService services_cards_vegetal.VegetalInfoService
}

func NewVegetalInfoController(service services_cards_vegetal.VegetalInfoService) *VegetalInfoController {
	return &VegetalInfoController{
		VegetalInfoService: service,
	}
}

func (ac *VegetalInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.VegetalInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
