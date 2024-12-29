package repositories_cards_weather

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type WeatherHeatRepository interface {
	cards_shared.RepositoryBase
	LoadHeatData(city string, year string) ([]HeatDataItem, error)
}

// Defina as propriedades específicas para este repositório
type HeatProperties struct {
	Ano  int     `json:"ano"`
	H12b float64 `json:"h12b"` // Negros e indígenas
	H11b float64 `json:"h11b"` // Mulheres
	H10b float64 `json:"h10b"` // Crianças
	H9b  float64 `json:"h9b"`  // Idosos
}

// Response JSON structure
type HeatDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"` // Omitir caso seja nil
	Value    string  `json:"value"`
}

type externalWeatherHeatRepository struct {
	geoserverURL string
}

// Constructor
func NewExternalWeatherHeatRepository() WeatherHeatRepository {
	_ = godotenv.Load()

	geoserverURL := os.Getenv("GEOSERVER_WEATHER_URL")
	if geoserverURL == "" {
		panic("A variável de ambiente GEOSERVER_WEATHER_URL não está definida")
	}

	return &externalWeatherHeatRepository{
		geoserverURL: geoserverURL,
	}
}

// LoadYears retrieves all unique years from the data
func (r *externalWeatherHeatRepository) LoadYears(city string) ([]int, error) {
	url := r.geoserverURL + city + "&outputFormat=application/json"

	data, err := cards_shared.FetchFromURL(url)
	if err != nil {
		return nil, err
	}

	yearsMap := make(map[int]bool)
	for _, feature := range data.Features {

		props, ok := feature.Properties.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("tipo inesperado de propriedades")
		}

		var heatProps HeatProperties
		if err := cards_shared.MapToStruct(props, &heatProps); err != nil {
			return nil, err
		}

		yearsMap[heatProps.Ano] = true
	}

	var years []int
	for year := range yearsMap {
		years = append(years, year)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	return years, nil
}

var subtitle string = "Porcentagem vivendo nas regiões mais quentes"

// LoadData retrieves weather heat data for a specific year
func (r *externalWeatherHeatRepository) LoadHeatData(city string, year string) ([]HeatDataItem, error) {
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

		var heatProps HeatProperties
		if err := cards_shared.MapToStruct(props, &heatProps); err != nil {
			return nil, err
		}

		if heatProps.Ano == convYear {
			filtered = feature
			filtered.Properties = heatProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	// Cast para HeatProperties
	heatProps := filtered.Properties.(HeatProperties)

	// Values
	black_indigenous_percentage := int(math.Round(heatProps.H12b * 100))
	women_percentage := int(math.Round(heatProps.H11b * 100))
	children_percentage := int(math.Round(heatProps.H10b * 100))
	senior_percentage := int(math.Round(heatProps.H9b * 100))

	result := []HeatDataItem{
		{"Negros e indígenas afetados", &subtitle, strconv.Itoa(black_indigenous_percentage) + "%"},
		{"Mulheres afetadas", &subtitle, strconv.Itoa(women_percentage) + "%"},
		{"Crianças afetadas", &subtitle, strconv.Itoa(children_percentage) + "%"},
		{"Idosos afetados", &subtitle, strconv.Itoa(senior_percentage) + "%"},
	}

	return result, nil
}
