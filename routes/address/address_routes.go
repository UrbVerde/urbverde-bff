// urbverde-bff/routes/address/address_routes.go
package address

import (
	controllers_address "urbverde-api/controllers/address"
	repositories_address "urbverde-api/repositories/address"
	services_address "urbverde-api/services/address"

	"github.com/gin-gonic/gin"
)

type AddressSuggestion struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// @Summary Retorna sugestões de endereço
// @Description Retorna sugestões baseadas nos dados fornecidos
// @Tags address
// @Accept json
// @Produce json
// @Param query query string true "Texto para buscar sugestões"
// @Success 200 {array} AddressSuggestion
// @Failure 400 {object} ErrorResponse
// @Router /address/suggestions [get]
func SetupAddressRoutes(rg *gin.RouterGroup) {
	addressRepo := repositories_address.NewExternalAddressRepository()
	addressService := services_address.NewAddressService(addressRepo)
	addressController := controllers_address.NewAddressController(addressService)

	rg.GET("/address/suggestions", addressController.GetSuggestions)
}
