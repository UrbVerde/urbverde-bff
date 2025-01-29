// urbverde-bff/controllers/address/address_data_controller.go
package controllers_address

import (
	"net/http"
	services_address "urbverde-api/services/address"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message" example:"Erro ao processar a solicitação"`
	Error   string `json:"error" example:"MISSING_PARAMETERS"`
}

type AddressDataController struct {
	AddressDataService services_address.AddressDataService
}

func NewAddressDataController(service services_address.AddressDataService) *AddressDataController {
	return &AddressDataController{
		AddressDataService: service,
	}
}

// GetLocationData returns detailed location data
// @Summary Retorna dados de localização
// @Description Retorna dados detalhados de localização
// @Tags address
// @Accept json
// @Produce json
// @Param code query string false "Código da localização"
// @Param name query string false "Nome ou nome de exibição da localização"
// @Param type query string false "Tipo da localização (state/city/country)"
// @Success 200 {object} repositories_address.Location
// @Failure 400 {object} ErrorResponse
// @Router /address/data [get]
func (ac *AddressDataController) GetLocationData(c *gin.Context) {
	identifier := c.Query("code")
	if identifier == "" {
		identifier = c.Query("name") // Allow lookup by name if code is not provided
	}

	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Either location code or name is required",
			"error":   "MISSING_IDENTIFIER",
		})
		return
	}

	locationType := c.Query("type") // Optional, will be guessed if not provided

	locationData, err := ac.AddressDataService.GetLocationData(identifier, locationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching location data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, locationData)
}
