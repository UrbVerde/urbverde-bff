// urbverde-bff/routes/categories/categories_routes.go
package categories

import (
	controllers_categories "urbverde-api/controllers/categories"
	repositories_categories "urbverde-api/repositories/categories"
	services_categories "urbverde-api/services/categories"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(rg *gin.RouterGroup) {
	categoriesRepo, err := repositories_categories.NewCategoriesRepository()
	if err != nil {
		panic(err)
	}
	categoriesService := services_categories.NewCategoriesService(categoriesRepo)
	categoriesController := controllers_categories.NewCategoriesController(categoriesService)

	rg.GET("/categories", categoriesController.GetCategories)
}
