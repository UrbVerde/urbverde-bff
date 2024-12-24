// controllers\address_controller.go
package controllers

import (
	"net/http"
	"urbverde-api/services"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	AddressService services.AddressService
}

func NewAddressController(service services.AddressService) *AddressController {
	return &AddressController{
		AddressService: service,
	}
}

func (ac *AddressController) GetSuggestions(c *gin.Context) {
	query := c.Query("query")

	suggestions, err := ac.AddressService.GetSuggestions(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar sugestões",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, suggestions)
}
