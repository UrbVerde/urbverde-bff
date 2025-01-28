package controllers_cards_square

import (
	"net/http"
	services_cards_square "urbverde-api/services/cards/square"

	"github.com/gin-gonic/gin"
)

type SquareInfoController struct {
	SquareInfoService services_cards_square.SquareInfoService
}

func NewSquareInfoController(service services_cards_square.SquareInfoService) *SquareInfoController {
	return &SquareInfoController{
		SquareInfoService: service,
	}
}

func (ac *SquareInfoController) LoadInfoData(c *gin.Context) {
	city := c.Query("city")

	data, err := ac.SquareInfoService.LoadInfoData(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar dados de temperatura",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
