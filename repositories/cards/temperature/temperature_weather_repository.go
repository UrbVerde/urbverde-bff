// urbverde-bff/repositories/cards/temperature/temperature_temperature_repository.go
package repositories_cards_temperature

import (
	"fmt"
	"math"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type TemperatureWeatherRepository interface {
	cards_shared.RepositoryBase
	LoadWeatherData(city string, year string) ([]WeatherDataItem, error)
}

type WeatherProperties struct {
	Ano int     `json:"ano"`
	C1  float64 `json:"c1"`  // Nível de Ilha de Calor
	H5b int     `json:"h5b"` // Amplitude
	C2  float64 `json:"c2"`  // Temperatura Média
	C3  float64 `json:"c3"`  // Temperatura Máxima
}

type WeatherDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"`
	Value    string  `json:"value"`
}

type externalTemperatureWeatherRepository struct {
	geoserverURL string
}

func NewExternalTemperatureWeatherRepository() TemperatureWeatherRepository {
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
		cards_shared.TypeName+"dados_temperatura_por_municipio",
		cards_shared.CqlFilterPrefix,
	)

	return &externalTemperatureWeatherRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalTemperatureWeatherRepository) LoadYears(city string) ([]int, error) {
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

func tempLoadData(v1 int, v2 int, sub1 *string, sub2 *string) {
	cards_shared.AuxLoadSubtitles(v1, 0, sub1)
	cards_shared.AuxLoadSubtitles(v2, 0, sub2)
}

func (r *externalTemperatureWeatherRepository) LoadWeatherData(city string, year string) ([]WeatherDataItem, error) {
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

		var weatherProps WeatherProperties
		if err := cards_shared.MapToStruct(props, &weatherProps); err != nil {
			return nil, err
		}

		if weatherProps.Ano == convYear {
			filtered = feature
			filtered.Properties = weatherProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	weatherProps := filtered.Properties.(WeatherProperties)

	heat_island_value := int(math.Round(weatherProps.C1))
	avg_temp_value := int(math.Round(weatherProps.C2))
	amplitude_value := weatherProps.H5b
	max_temp_value := int(math.Round(weatherProps.C3))

	// var heat_island_subtitle string = " da média nacional de " // deve ser adicionado junto do dado nacional
	// var avg_temp_subtitle string = " da média nacional de "
	var amplitude_subtitle string = "É a diferença entre a temperatura mais quente e a mais fria"

	// tempLoadData(heat_island_value, avg_temp_value, &heat_island_subtitle, &avg_temp_subtitle)

	result := []WeatherDataItem{
		{"Nível de ilha de calor", nil, strconv.Itoa(heat_island_value)},
		{"Temperatura média da superfície", nil, strconv.Itoa(avg_temp_value) + "°C"},
		{"Maior amplitude", &amplitude_subtitle, strconv.Itoa(amplitude_value) + "°C"},
		{"Temperatura máxima da superfície", nil, strconv.Itoa(max_temp_value) + "°C"},
	}

	return result, nil
}
