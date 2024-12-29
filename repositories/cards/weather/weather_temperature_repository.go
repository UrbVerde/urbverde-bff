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

var sub_ilha_calor string = " da média nacional de "
var sub_temp_media string = " da média nacional de "
var sub_amplitude string = "É a diferença entre a temperatura mais quente e a mais fria"

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
	auxLoadSubtitles(v1, 0, &sub_ilha_calor)
	auxLoadSubtitles(v2, 0, &sub_temp_media)
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
	v_ilha_calor := int(math.Round(temperatureProps.C1))
	v_tempe_media := int(math.Round(temperatureProps.C2))
	v_amplitude := temperatureProps.H5b
	v_tempe_max := int(math.Round(temperatureProps.C3))

	tempLoadData(v_ilha_calor, v_tempe_media)

	result := []TemperatureDataItem{
		{"Nível de ilha de calor", &sub_ilha_calor, strconv.Itoa(v_ilha_calor)},
		{"Temperatura média da superfície", &sub_temp_media, strconv.Itoa(v_tempe_media) + "°C"},
		{"Maior amplitude", &sub_amplitude, strconv.Itoa(v_amplitude) + "°C"},
		{"Temperatura máxima da superfície", nil, strconv.Itoa(v_tempe_max) + "°C"},
	}

	return result, nil
}
