// urbverde-bff/controllers/categories/categories_controller.go
package controllers

import (
	"net/http"
	services_categories "urbverde-api/services/categories"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	CategoriesService services_categories.CategoriesService
}

func NewCategoriesController(service services_categories.CategoriesService) *CategoriesController {
	return &CategoriesController{
		CategoriesService: service,
	}
}

// @Summary Retorna categorias disponíveis
// @Description Retorna as categorias e camadas disponíveis para o município
// @Tags categories
// @Accept json
// @Produce json
// @Param city query string true "Código do município"
// @Success 200 {object} repositories_categories.CategoriesResponse
// @Failure 400 {object} ErrorResponse
// @Router /categories [get]
func (cc *CategoriesController) GetCategories(c *gin.Context) {
	cityCode := c.Query("code")

	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "City code is required",
			"error":   "MISSING_CODE",
		})
		return
	}

	categories, err := cc.CategoriesService.GetCategories(cityCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching categories",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}
