// urbverde-bff/controllers/address/address_controller.go
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

// GetSuggestions returns city suggestions based on query
// @Summary Retorna sugestões de endereço
// @Description Retorna sugestões baseadas nos dados fornecidos
// @Tags address
// @Accept json
// @Produce json
// @Param query query string true "Texto para buscar sugestões"
// @Success 200 {array} repositories_address.CityResponse
// @Failure 400 {object} ErrorResponse
// @Router /address/suggestions [get]
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
