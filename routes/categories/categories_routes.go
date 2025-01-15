// urbverde-bff/routes/categories/categories_routes.go
package categories

import (
	controllers_categories "urbverde-api/controllers/categories"
	repositories_categories "urbverde-api/repositories/categories"
	services_categories "urbverde-api/services/categories"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
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
func SetupCategoriesRoutes(rg *gin.RouterGroup) {
	categoriesRepo, err := repositories_categories.NewCategoriesRepository()
	if err != nil {
		panic(err)
	}
	categoriesService := services_categories.NewCategoriesService(categoriesRepo)
	categoriesController := controllers_categories.NewCategoriesController(categoriesService)

	rg.GET("/categories", categoriesController.GetCategories)
}
