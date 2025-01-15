// urbverde-bff/services/categories_service.go
package services_categories

import (
	repositories_categories "urbverde-api/repositories/categories"
)

type CategoriesService interface {
	GetCategories(cityCode string) (*repositories_categories.CategoriesResponse, error)
}

type categoriesService struct {
	repo repositories_categories.CategoriesRepository
}

func NewCategoriesService(repo repositories_categories.CategoriesRepository) CategoriesService {
	return &categoriesService{
		repo: repo,
	}
}

func (s *categoriesService) GetCategories(cityCode string) (*repositories_categories.CategoriesResponse, error) {
	return s.repo.GetCategories(cityCode)
}
