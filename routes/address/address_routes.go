package address

import (
	"urbverde-api/controllers"
	"urbverde-api/repositories"
	"urbverde-api/routes/tracker"
	"urbverde-api/services"

	"github.com/gin-gonic/gin"
)

func SetupAddressRoutes(rg *gin.RouterGroup) {
	// Inicializa o repositório e serviço
	addressRepo := repositories.NewExternalAddressRepository()
	addressService := services.NewAddressService(addressRepo)
	addressController := controllers.NewAddressController(addressService)

	// Define a rota para buscar sugestões
	rg.GET("/address/suggestions", addressController.GetSuggestions)
	tracker.AddEndpoint("GET", "/api/v1/address/suggestions", "Retorna sugestões de endereço")
}
