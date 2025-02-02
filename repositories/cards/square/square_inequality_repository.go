package repositories_cards_square

import (
	"fmt"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type SquareInequalityRepository interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]SquareInequalityDataItem, error)
}

type SquareInequalityProperties struct {
	Ano  int     `json:"ano"`
	H12a float64 `json:"h12a"` // Negros e indígenas
	H11a float64 `json:"h11a"` // Mulheres
	H10a float64 `json:"h10a"` // Crianças
	H9a  float64 `json:"h9a"`  // Idosos
}

type SquareInequalityDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"`
	Value    string  `json:"value"`
}

type externalSquareInequalityRepository struct {
	geoserverURL string
}

func NewExternalSquareInequalityRepository() SquareInequalityRepository {
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
		cards_shared.TypeName+"dados_pracas_por_municipio",
		cards_shared.CqlFilterPrefix,
	)

	return &externalSquareInequalityRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalSquareInequalityRepository) LoadYears(city string) ([]int, error) {
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

func (r *externalSquareInequalityRepository) LoadInequalityData(city string, year string) ([]SquareInequalityDataItem, error) {
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

		var inequalityProps SquareInequalityProperties
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

	inequalityProps := filtered.Properties.(SquareInequalityProperties)

	black_indigenous_value := strconv.FormatFloat(inequalityProps.H12a, 'f', 2, 64)
	women_value := strconv.FormatFloat(inequalityProps.H11a, 'f', 2, 64)
	children_value := strconv.FormatFloat(inequalityProps.H10a, 'f', 2, 64)
	elderly_value := strconv.FormatFloat(inequalityProps.H9a, 'f', 2, 64)

	var general_subtitle string = "Porcentagem vivendo fora da vizinhança das praças"

	result := []SquareInequalityDataItem{
		{"Negros e indígenas", &general_subtitle, black_indigenous_value + "%"},
		{"Mulheres", &general_subtitle, women_value + "%"},
		{"Crianças", &general_subtitle, children_value + "%"},
		{"Idosos", &general_subtitle, elderly_value + "%"},
	}

	return result, nil
}
