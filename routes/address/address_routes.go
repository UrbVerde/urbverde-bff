// urbverde-bff/routes/address/address_routes.go
package address

import (
	controllers_address "urbverde-api/controllers/address"
	repositories_address "urbverde-api/repositories/address"
	services_address "urbverde-api/services/address"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message" example:"Erro ao processar a solicitação"`
	Error   string `json:"error" example:"MISSING_PARAMETERS"`
}

func SetupAddressRoutes(rg *gin.RouterGroup) {
	// Setup suggestions endpoint
	addressRepo := repositories_address.NewExternalAddressRepository()
	addressService := services_address.NewAddressService(addressRepo)
	addressController := controllers_address.NewAddressController(addressService)

	// Setup data endpoint
	addressDataRepo, err := repositories_address.NewExternalAddressDataRepository()
	if err != nil {
		panic(err)
	}
	addressDataService := services_address.NewAddressDataService(addressDataRepo)
	addressDataController := controllers_address.NewAddressDataController(addressDataService)

	rg.GET("/address/suggestions", addressController.GetSuggestions)
	rg.GET("/address/data", addressDataController.GetLocationData)
}
