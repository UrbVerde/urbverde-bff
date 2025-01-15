// urbverde-bff/repositories/categories/categories_repository.go
package repositories_categories

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Layer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	IsNew    bool   `json:"isNew"`
}

type Category struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Icon   string  `json:"icon"`
	Layers []Layer `json:"layers"`
}

type CategoriesResponse struct {
	Categories []Category `json:"categories"`
}

type CategoriesRepository interface {
	GetCategories(cityCode string) (*CategoriesResponse, error)
}

type categoriesRepository struct {
	baseCategories CategoriesResponse
}

func NewCategoriesRepository() (CategoriesRepository, error) {
	// This would typically come from a database or external service
	categoriesJson := `{
		"categories": [
			{
				"id": "climate",
				"name": "Clima",
				"icon": "Snowy Sunny Day.svg",
				"layers": [
					{"id": "surface_temp", "name": "Temperatura de superfície", "isActive": false, "isNew": true},
					{"id": "max_surface_temp", "name": "Temperatura máxima de superfície", "isActive": false, "isNew": false},
					{"id": "heat_island", "name": "Nível de exposição à ilha de calor", "isActive": false, "isNew": false}
				]
			},
			{
				"id": "vegetation",
				"name": "Vegetação",
				"icon": "Oak Tree.svg",
				"layers": [
					{"id": "pcv", "name": "Coberta vegetal (PCV)", "isActive": false, "isNew": false},
					{"id": "icv", "name": "Cobertura vegetal por habitante (ICV)", "isActive": false, "isNew": false}
				]
			},
			{
				"id": "parks",
				"name": "Parques e Praças",
				"icon": "Nature.svg",
				"layers": [
					{"id": "park_distribution", "name": "Distribuição dos parques", "isActive": false, "isNew": false},
					{"id": "green_area_per_capita", "name": "Área verde por habitante", "isActive": false, "isNew": false}
				]
			},
			{
				"id": "census",
				"name": "Censo",
				"icon": "bi bi-people",
				"layers": [
					{"id": "population", "name": "População", "isActive": false, "isNew": false},
					{"id": "revenue", "name": "Renda per capita", "isActive": false, "isNew": false}
				]
			},
			{
				"id": "transport",
				"name": "Mobilidade",
				"icon": "bi bi-bicycle",
				"layers": [
					{"id": "bike_lanes", "name": "Ciclovias", "isActive": false, "isNew": true}
				]
			}
		]
	}`

	var baseCategories CategoriesResponse
	if err := json.Unmarshal([]byte(categoriesJson), &baseCategories); err != nil {
		return nil, fmt.Errorf("error unmarshaling base categories: %w", err)
	}

	return &categoriesRepository{
		baseCategories: baseCategories,
	}, nil
}

func (r *categoriesRepository) GetCategories(cityCode string) (*CategoriesResponse, error) {
	// Deep clone the base categories
	categoriesBytes, err := json.Marshal(r.baseCategories)
	if err != nil {
		return nil, fmt.Errorf("error cloning categories: %w", err)
	}

	var response CategoriesResponse
	if err := json.Unmarshal(categoriesBytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling categories clone: %w", err)
	}

	// Check if it's a São Paulo state city (starts with 35)
	isSaoPauloState := strings.HasPrefix(cityCode, "35")
	if !isSaoPauloState {
		// Filter categories for non-São Paulo cities
		var filteredCategories []Category
		for _, category := range response.Categories {
			if category.ID == "parks" || category.ID == "census" || category.ID == "transport" {
				filteredCategories = append(filteredCategories, category)
			}
		}
		response.Categories = filteredCategories
		return &response, nil
	}

	// Handle special cases for specific cities
	switch cityCode {
	case "3534708": // Ourinhos-SP
		exclusiveLayer := Layer{
			ID:       "tree_inventory",
			Name:     "Inventário de Árvores",
			IsActive: false,
			IsNew:    true,
		}

		// Add to vegetation and parks categories
		for i, category := range response.Categories {
			if category.ID == "vegetation" || category.ID == "parks" {
				newLayers := append([]Layer{exclusiveLayer}, category.Layers...)
				response.Categories[i].Layers = newLayers
			}
		}

	case "3548906": // São Carlos
		exclusiveLayer := Layer{
			ID:       "fruit_trees",
			Name:     "Pés de Fruta",
			IsActive: false,
			IsNew:    true,
		}

		// Add to agriculture category
		for i, category := range response.Categories {
			if category.ID == "agriculture" {
				newLayers := append([]Layer{exclusiveLayer}, category.Layers...)
				response.Categories[i].Layers = newLayers
			}
		}

	case "3552205": // Sorocaba
		exclusiveLayer := Layer{
			ID:       "innundation_points",
			Name:     "Pontos de inundação",
			IsActive: false,
			IsNew:    true,
		}

		// Add to water category
		for i, category := range response.Categories {
			if category.ID == "water" {
				newLayers := append([]Layer{exclusiveLayer}, category.Layers...)
				response.Categories[i].Layers = newLayers
			}
		}
	}

	return &response, nil
}
