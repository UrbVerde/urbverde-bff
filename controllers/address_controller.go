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
	query := c.Query("q")
	filter := c.Query("filter")

	// Define acceptable filter values
	validFilters := map[string]bool{
		"all":   true,
		"city":  true,
		"state": true,
	}

	// Set default filter if not provided or invalid
	if !validFilters[filter] {
		filter = "all"
	}

	suggestions, err := ac.AddressService.GetSuggestions(query, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar sugestões",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, suggestions)
}
