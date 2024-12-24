// mock/data_store.go
package mock

import "urbverde-api/models"

var MockCityData = models.CityData{
	CdMun: "3548906", // Default to São Carlos
	Name:  "São Carlos",
	Bounds: models.Bounds{
		Type: "Feature",
		Properties: models.BoundProps{
			Name:    "São Carlos",
			AreaKm2: 1136.907,
		},
		Geometry: models.Geometry{
			Type: "Polygon",
			Coordinates: [][][]float64{
				{
					{-47.7644, -21.8581},
					{-47.9843, -21.8684},
					{-48.0547, -22.0989},
					{-47.8741, -22.1561},
					{-47.7528, -22.0851},
					{-47.7644, -21.8581},
				},
			},
		},
	},
	Categories: []models.Category{
		{
			ID:   "vegetacao",
			Name: "Vegetação",
			Icon: "/assets/icons/pineTree.svg",
			Layers: []models.Layer{
				{ID: "pcv", Name: "Coberta vegetal (PCV)", IsActive: false},
				{ID: "icv", Name: "Cobertura vegetal por habitante (ICV)", IsActive: false},
				{ID: "idsa", Name: "Desigualdade sociambiental (IDSA)", IsActive: false},
				{ID: "cvp", Name: "Cobertura vegetal por pixel", IsActive: false},
				{ID: "ndvi", Name: "Vigor da vegetação (NDVI)", IsActive: false},
			},
		},
		{
			ID:   "clima",
			Name: "Clima",
			Icon: "/assets/icons/sunBehindeCloud.svg",
			Layers: []models.Layer{
				{ID: "temp_sup", Name: "Temperatura de superfície", IsActive: false},
				{ID: "temp_max", Name: "Temperatura máxima de superfície", IsActive: false},
				{ID: "ilha_calor", Name: "Nível de exposição à ilha de calor", IsActive: false},
			},
		},
		{
			ID:   "parques",
			Name: "Parques e Praças",
			Icon: "/assets/icons/bike.svg",
			Layers: []models.Layer{
				{ID: "dist_parques", Name: "Distribuição dos parques", IsActive: false},
				{ID: "area_verde", Name: "Área verde por habitante", IsActive: false},
				{ID: "acess_areas", Name: "Acessibilidade a áreas verdes", IsActive: false},
			},
		},
	},
}
