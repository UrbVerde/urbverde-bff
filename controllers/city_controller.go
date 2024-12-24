// controllers/city_controller.go
package controllers

import (
	"net/http"
	"urbverde-api/mock"

	"github.com/gin-gonic/gin"
)

type CityController struct{}

func NewCityController() *CityController {
	return &CityController{}
}

// GetCityBounds returns the city boundaries and code
func (cc *CityController) GetCityBounds(c *gin.Context) {
	cityName := c.Query("name")
	if cityName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "City name is required",
		})
		return
	}

	// For now, return the same mock data regardless of the city name
	response := gin.H{
		"cd_mun": mock.MockCityData.CdMun,
		"bounds": mock.MockCityData.Bounds,
	}

	c.JSON(http.StatusOK, response)
}

// GetCityData returns the available categories and layers
func (cc *CityController) GetCityData(c *gin.Context) {
	cityCode := c.Query("cd_mun")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "City code is required",
		})
		return
	}

	// For now, return the same categories and layers for any city code
	response := gin.H{
		"categories": mock.MockCityData.Categories,
	}

	c.JSON(http.StatusOK, response)
}
