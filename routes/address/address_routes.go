// urbverde-bff/routes/address/address_routes.go
package address

import (
	"urbverde-api/controllers"
	"urbverde-api/repositories"
	"urbverde-api/services"

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
	addressRepo := repositories.NewExternalAddressRepository()
	addressService := services.NewAddressService(addressRepo)
	addressController := controllers.NewAddressController(addressService)

	rg.GET("/address/suggestions", addressController.GetSuggestions)
}
