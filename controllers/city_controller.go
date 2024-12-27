package controllers

import (
	"net/http"
	"urbverde-api/mock"

	"github.com/gin-gonic/gin"
)

type CityController struct{}

// NewCityController creates a new instance of CityController
func NewCityController() *CityController {
	return &CityController{}
}

// GetCityData returns the available categories and layers for a city
func (cc *CityController) GetCityData(c *gin.Context) {
	// For simplicity, return the mock categories and layers
	response := gin.H{
		"categories": mock.MockCityData.Categories,
	}

	c.JSON(http.StatusOK, response)
}
