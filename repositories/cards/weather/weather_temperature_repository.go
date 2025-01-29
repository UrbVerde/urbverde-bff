// urbverde-bff/repositories/cards/weather/weather_temperature_repository.go
package repositories_cards_weather

import (
	"fmt"
	"math"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type WeatherTemperatureRepository interface {
	cards_shared.RepositoryBase
	LoadTemperatureData(city string, year string) ([]TemperatureDataItem, error)
}

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
	Subtitle *string `json:"subtitle,omitempty"`
	Value    string  `json:"value"`
}

type externalWeatherTemperatureRepository struct {
	geoserverURL string
}

func NewExternalWeatherTemperatureRepository() WeatherTemperatureRepository {
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

	return &externalWeatherTemperatureRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalWeatherTemperatureRepository) LoadYears(city string) ([]int, error) {
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

	temperatureProps := filtered.Properties.(TemperatureProperties)

	heat_island_value := int(math.Round(temperatureProps.C1))
	avg_temp_value := int(math.Round(temperatureProps.C2))
	amplitude_value := temperatureProps.H5b
	max_temp_value := int(math.Round(temperatureProps.C3))

	// var heat_island_subtitle string = " da média nacional de " // deve ser adicionado junto do dado nacional
	// var avg_temp_subtitle string = " da média nacional de "
	var amplitude_subtitle string = "É a diferença entre a temperatura mais quente e a mais fria"

	// tempLoadData(heat_island_value, avg_temp_value, &heat_island_subtitle, &avg_temp_subtitle)

	result := []TemperatureDataItem{
		{"Nível de ilha de calor", nil, strconv.Itoa(heat_island_value)},
		{"Temperatura média da superfície", nil, strconv.Itoa(avg_temp_value) + "°C"},
		{"Maior amplitude", &amplitude_subtitle, strconv.Itoa(amplitude_value) + "°C"},
		{"Temperatura máxima da superfície", nil, strconv.Itoa(max_temp_value) + "°C"},
	}

	return result, nil
}
