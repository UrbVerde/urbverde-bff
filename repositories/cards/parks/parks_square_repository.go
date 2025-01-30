package repositories_cards_parks

import (
	"fmt"
	"math"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type ParksSquareRepository interface {
	cards_shared.RepositoryBase
	LoadSquareData(city string, year string) ([]SquareDataItem, error)
}

type SquareProperties struct {
	Ano int     `json:"ano"`
	A1  float64 `json:"a1"` // % moradores proximos a praças
	A4  float64 `json:"a4"` // Distancia média até as praças
	H6  float64 `json:"h6"` // Desigualdade de renda
	H7  float64 `json:"h7"` // Racismo ambiental
}

type SquareDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"`
	Value    string  `json:"value"`
}

type externalParksSquareRepository struct {
	geoserverURL string
}

func NewExternalParksSquareRepository() ParksSquareRepository {
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

	return &externalParksSquareRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalParksSquareRepository) LoadYears(city string) ([]int, error) {
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

func (r *externalParksSquareRepository) LoadSquareData(city string, year string) ([]SquareDataItem, error) {
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

		var squareProps SquareProperties
		if err := cards_shared.MapToStruct(props, &squareProps); err != nil {
			return nil, err
		}

		if squareProps.Ano == convYear {
			filtered = feature
			filtered.Properties = squareProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	squareProps := filtered.Properties.(SquareProperties)

	square_value := int(squareProps.A1)
	distance_value := int(math.Round(squareProps.A4))
	inequality_value := squareProps.H6
	racism_value := int(math.Round(squareProps.H7))

	// var square_subtitle string = " da média nacional de " // a média nacional deve ser incluída poosteriormente
	var inequality_subtitle string = "Moradores próximos a praças têm em média 15% mais de renda"
	var racism_subtitle string = "População negra ou indígena que vive fora da vizinhança das praças"

	result := []SquareDataItem{
		{"Moradores próximos a praças", nil, strconv.Itoa(square_value) + "%"},
		{"Distância média até as praças", nil, strconv.Itoa(distance_value) + " metros"},
		{"Desigualdade de renda", &inequality_subtitle, strconv.FormatFloat(inequality_value, 'f', 2, 64) + "x"},
		{"Racismo ambiental", &racism_subtitle, strconv.Itoa(racism_value) + "%"},
	}

	return result, nil
}
