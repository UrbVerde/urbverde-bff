package address

import (
	"urbverde-api/controllers"
	"urbverde-api/repositories"
	"urbverde-api/routes/tracker"
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
	// Inicializa o repositório e serviço
	addressRepo := repositories.NewExternalAddressRepository()
	addressService := services.NewAddressService(addressRepo)
	addressController := controllers.NewAddressController(addressService)

	// Define a rota para buscar sugestões
	rg.GET("/address/suggestions", addressController.GetSuggestions)
	tracker.AddEndpoint("GET", "/api/v1/address/suggestions", "Retorna sugestões de endereço")
}
