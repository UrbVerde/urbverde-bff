// urbverde-bff/repositories/categories/categories_repository.go
package repositories_categories

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed data/*.json
var dataFiles embed.FS

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

type ExclusiveLayer struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	IsActive        bool     `json:"isActive"`
	IsNew           bool     `json:"isNew"`
	AddToCategories []string `json:"addToCategories"`
}

type CityExclusiveLayers struct {
	Name   string           `json:"name"`
	Layers []ExclusiveLayer `json:"layers"`
}

type ExclusiveLayersMap map[int]CityExclusiveLayers

type CategoriesRepository interface {
	GetCategories(cityCode string) (*CategoriesResponse, error)
}

type categoriesRepository struct {
	baseCategories  CategoriesResponse
	exclusiveLayers ExclusiveLayersMap
}

func loadJSONFile(filename string, target interface{}) error {
	data, err := dataFiles.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filename, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("error unmarshaling %s: %w", filename, err)
	}

	return nil
}

func NewCategoriesRepository() (CategoriesRepository, error) {
	var baseCategories CategoriesResponse
	if err := loadJSONFile("data/base_categories.json", &baseCategories); err != nil {
		return nil, err
	}

	var exclusiveLayers ExclusiveLayersMap
	if err := loadJSONFile("data/exclusive_layers.json", &exclusiveLayers); err != nil {
		return nil, err
	}

	return &categoriesRepository{
		baseCategories:  baseCategories,
		exclusiveLayers: exclusiveLayers,
	}, nil
}

func (r *categoriesRepository) addExclusiveLayers(response *CategoriesResponse, cityCode string) {
	cityCodeInt, err := strconv.Atoi(cityCode)
	if err != nil {
		return
	}

	cityLayers, exists := r.exclusiveLayers[cityCodeInt]
	if !exists {
		return
	}

	// For each exclusive layer defined for this city
	for _, exclusiveLayer := range cityLayers.Layers {
		layer := Layer{
			ID:       exclusiveLayer.ID,
			Name:     exclusiveLayer.Name,
			IsActive: exclusiveLayer.IsActive,
			IsNew:    exclusiveLayer.IsNew,
		}

		// Add the layer to all specified categories
		for i := range response.Categories {
			for _, categoryID := range exclusiveLayer.AddToCategories {
				if response.Categories[i].ID == categoryID {
					// Insert at beginning of layers slice
					newLayers := make([]Layer, len(response.Categories[i].Layers)+1)
					newLayers[0] = layer
					copy(newLayers[1:], response.Categories[i].Layers)
					response.Categories[i].Layers = newLayers
				}
			}
		}
	}
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

	// // Check if it's a São Paulo state city (starts with 35)
	// isSaoPauloState := strings.HasPrefix(cityCode, "35")
	// if !isSaoPauloState {
	// 	// Filter categories for non-São Paulo cities
	// 	var filteredCategories []Category
	// 	for _, category := range response.Categories {
	// 		if category.ID == "parks" || category.ID == "census" || category.ID == "transport" {
	// 			filteredCategories = append(filteredCategories, category)
	// 		}
	// 	}
	// 	response.Categories = filteredCategories
	// 	return &response, nil
	// }

	// // Add any exclusive layers for this city
	// r.addExclusiveLayers(&response, cityCode)

	return &response, nil
}
