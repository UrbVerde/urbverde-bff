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

type WeatherTemperatureRepository interface {
	cards_shared.RepositoryBase
	LoadTemperatureData(city string, year string) ([]TemperatureDataItem, error)
}

// Defina as propriedades específicas para este repositório
type TemperatureProperties struct {
	Ano int     `json:"ano"`
	C1  float64 `json:"c1"`  // Nível de Ilha de Calor
	H5b int     `json:"h5b"` // Amplitude
	C2  float64 `json:"c2"`  // Temperatura Média
	C3  float64 `json:"c3"`  // Temperatura Máxima
}

// Response JSON structure
type TemperatureDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"` // Omitir caso seja nil
	Value    string  `json:"value"`
}

type externalWeatherTemperatureRepository struct {
	geoserverURL string
}

// Constructor
func NewExternalWeatherTemperatureRepository() WeatherTemperatureRepository {
	_ = godotenv.Load()

	geoserverURL := os.Getenv("GEOSERVER_WEATHER_URL")
	if geoserverURL == "" {
		panic("A variável de ambiente GEOSERVER_WEATHER_URL não está definida")
	}

	return &externalWeatherTemperatureRepository{
		geoserverURL: geoserverURL,
	}
}

// LoadYears retrieves all unique years from the data
func (r *externalWeatherTemperatureRepository) LoadYears(city string) ([]int, error) {
	url := r.geoserverURL + city + "&outputFormat=application/json"

	data, err := cards_shared.FetchFromURL(url)
	if err != nil {
		return nil, err
	}

	yearsMap := make(map[int]bool)
	for _, feature := range data.Features {
		// Realiza o cast para as propriedades específicas
		props, ok := feature.Properties.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("tipo inesperado de propriedades")
		}

		var temperatureProps TemperatureProperties
		if err := cards_shared.MapToStruct(props, &temperatureProps); err != nil {
			return nil, err
		}

		yearsMap[temperatureProps.Ano] = true
	}

	var years []int
	for year := range yearsMap {
		years = append(years, year)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	return years, nil
}

// Result
// Nível de ilha de calor
// Temperatura média da superfície
// Maior amplitude
// Temperatura máxima da superfície

var heat_island_subtitle string = " da média nacional de "
var avg_temp_subtitle string = " da média nacional de "
var amplitude_subtitle string = "É a diferença entre a temperatura mais quente e a mais fria"

func auxLoadSubtitles(value int, avg int, subtitle *string) {
	if subtitle == nil {
		return
	}

	if value < avg {
		*subtitle = "Abaixo" + *subtitle + strconv.Itoa(avg)
	} else if value > avg {
		*subtitle = "Acima" + *subtitle + strconv.Itoa(avg)
	} else {
		*subtitle = "Está na média nacional de " + strconv.Itoa(avg)
	}
}

func tempLoadData(v1 int, v2 int) {
	auxLoadSubtitles(v1, 0, &heat_island_subtitle)
	auxLoadSubtitles(v2, 0, &avg_temp_subtitle)
}

// LoadData retrieves weather temperature data for a specific year
func (r *externalWeatherTemperatureRepository) LoadTemperatureData(city string, year string) ([]TemperatureDataItem, error) {
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

		var temperatureProps TemperatureProperties
		if err := cards_shared.MapToStruct(props, &temperatureProps); err != nil {
			return nil, err
		}

		if temperatureProps.Ano == convYear {
			filtered = feature
			filtered.Properties = temperatureProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	// Cast para TemperatureProperties
	temperatureProps := filtered.Properties.(TemperatureProperties)

	// Values
	heat_island_value := int(math.Round(temperatureProps.C1))
	avg_temp_value := int(math.Round(temperatureProps.C2))
	amplitude_value := temperatureProps.H5b
	max_temp_value := int(math.Round(temperatureProps.C3))

	tempLoadData(heat_island_value, avg_temp_value)

	result := []TemperatureDataItem{
		{"Nível de ilha de calor", &heat_island_subtitle, strconv.Itoa(heat_island_value)},
		{"Temperatura média da superfície", &avg_temp_subtitle, strconv.Itoa(avg_temp_value) + "°C"},
		{"Maior amplitude", &amplitude_subtitle, strconv.Itoa(amplitude_value) + "°C"},
		{"Temperatura máxima da superfície", nil, strconv.Itoa(max_temp_value) + "°C"},
	}

	return result, nil
}
