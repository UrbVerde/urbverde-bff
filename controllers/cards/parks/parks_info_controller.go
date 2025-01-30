package controllers_cards_parks

import (
	"net/http"
	services_cards_parks "urbverde-api/services/cards/parks"

	"github.com/gin-gonic/gin"
)

type ParksInfoController struct {
	ParksInfoService services_cards_parks.ParksInfoService
}

func NewParksInfoController(service services_cards_parks.ParksInfoService) *ParksInfoController {
	return &ParksInfoController{
		ParksInfoService: service,
	}
}

func (ac *ParksInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.ParksInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
