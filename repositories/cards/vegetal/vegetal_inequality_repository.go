package repositories_cards_vegetal

import (
	"fmt"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type VegetalInequalityRepository interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]InequalityDataItem, error)
}

type InequalityProperties struct {
	Ano  int     `json:"ano"`
	B3h2 float64 `json:"b3h2"` // % moradores pouca vegetação
	B3   float64 `json:"b3"`   // Desigualdade ambiental e social
	// X float64 `json:"x"` // Vigor
}

type InequalityDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"`
	Value    string  `json:"value"`
}

type externalVegetalInequalityRepository struct {
	geoserverURL string
}

func NewExternalVegetalInequalityRepository() VegetalInequalityRepository {
	_ = godotenv.Load()

	geoserverURL := os.Getenv("GEOSERVER_URL")
	if geoserverURL == "" {
		panic("A variável de ambiente GEOSERVER_URL não está definida")
	}

	geoserverURL = fmt.Sprintf("%sows?service=%s&version=%s&request=%s&typeName=%s&%s",
		geoserverURL,
		cards_shared.WfsService,
		cards_shared.WfsVersion,
		cards_shared.WfsRequest,
		cards_shared.TypeName+"vegetacao_highlights_data",
		cards_shared.CqlFilterPrefix,
	)

	return &externalVegetalInequalityRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalVegetalInequalityRepository) LoadYears(city string) ([]int, error) {
	url := r.geoserverURL + city + "&outputFormat=application/json"

	processProperties := func(props map[string]interface{}) (int, error) {
		year, ok := props["ano"].(float64)
		if !ok {
			return 0, fmt.Errorf("year not found or invalid type")
		}
		return int(year), nil
	}

	return cards_shared.LoadYears(url, processProperties)
}

func tempLoadInequalityData(v1 int, sub1 *string) {
	cards_shared.AuxLoadSubtitles(v1, 0, sub1) // a media nacional deve ser incluida posteriormente
}

func (r *externalVegetalInequalityRepository) LoadInequalityData(city string, year string) ([]InequalityDataItem, error) {
	url := r.geoserverURL + city + "&outputFormat=application/json"

	data, err := cards_shared.FetchFromURL(url)
	if err != nil {
		return nil, err
	}

	convYear, err := strconv.Atoi(year)
	if err != nil {
		return nil, fmt.Errorf("ano inválido: %w", err)
	}

	var filtered cards_shared.Feature
	found := false
	for _, feature := range data.Features {
		props, ok := feature.Properties.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("tipo inesperado de propriedades")
		}

		var inequalityProps InequalityProperties
		if err := cards_shared.MapToStruct(props, &inequalityProps); err != nil {
			return nil, err
		}

		if inequalityProps.Ano == convYear {
			filtered = feature
			filtered.Properties = inequalityProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	inequalityProps := filtered.Properties.(InequalityProperties)
	residents_inequality_value := int(inequalityProps.B3h2)
	environmental_inequality_value := int(inequalityProps.B3 * 100)
	// vegetation_vigor_value := int(math.Round(inequalityProps.B1h4))

	var residents_inequality_subtitle = "Porcentagem vivendo nas regiões menos vegetadas"
	var environmental_inequality_subtitle string = " da média nacional de "
	// var vegetation_vigor_subtitle string = "Indica a média da saúde da vegetação"

	tempLoadInequalityData(environmental_inequality_value, &environmental_inequality_subtitle)

	result := []InequalityDataItem{
		{"Moradores em áreas de pouca vegetação", &residents_inequality_subtitle, strconv.Itoa(residents_inequality_value) + "%"},
		{"Desigualdade ambiental e social (IDSA)", &environmental_inequality_subtitle, strconv.Itoa(environmental_inequality_value)},
		// {"Vigor da vegetação (NDVI)", &vegetation_vigor_subtitle, "*Dado em construção*"},
	}

	return result, nil
}
