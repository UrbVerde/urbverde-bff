// urbverde-bff/controllers/address_controller.go
package controllers_address

import (
	"net/http"
	services_address "urbverde-api/services/address"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	AddressService services_address.AddressService
}

func NewAddressController(service services_address.AddressService) *AddressController {
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
