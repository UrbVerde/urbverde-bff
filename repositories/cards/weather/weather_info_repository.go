// urbverde-bff/repositories/cards/weather/weather_info_repository.go
package repositories_cards_weather

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type WeatherInfoRepository interface {
	LoadInfoData(city string) ([]InfoDataItem, error)
}

type InfoProperties struct {
	Ano int     `json:"ano"`
	B1  float64 `json:"b1"` // % cobertura vegetal (vegetal)
	A1  float64 `json:"a1"` // Moradores prox a praças (square)
	B3  float64 `json:"b3"` // Desigualdade ambiental e social (vegetal)
}

type InfoDataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type externalWeatherInfoRepository struct {
	vegetalURL string
	squareURL  string
}

func NewExternalWeatherInfoRepository() WeatherInfoRepository {
	_ = godotenv.Load()

	geoserverURL := os.Getenv("GEOSERVER_URL")
	if geoserverURL == "" {
		panic("A variável de ambiente GEOSERVER_URL não está definida")
	}

	vegetalURL := fmt.Sprintf("%sows?service=%s&version=%s&request=%s&typeName=%s&%s",
		geoserverURL,
		cards_shared.WfsService,
		cards_shared.WfsVersion,
		cards_shared.WfsRequest,
		cards_shared.TypeName+"vegetacao_highlights_data",
		cards_shared.CqlFilterPrefix,
	)

	squareURL := fmt.Sprintf("%sows?service=%s&version=%s&request=%s&typeName=%s&%s",
		geoserverURL,
		cards_shared.WfsService,
		cards_shared.WfsVersion,
		cards_shared.WfsRequest,
		cards_shared.TypeName+"dados_pracas_por_municipio",
		cards_shared.CqlFilterPrefix,
	)

	return &externalWeatherInfoRepository{
		vegetalURL: vegetalURL,
		squareURL:  squareURL,
	}
}

func (r *externalWeatherInfoRepository) loadGeneralData(city string) (*cards_shared.FeatureCollection, error) {
	squareUrl := r.squareURL + city + "&outputFormat=application/json"
	squareData, err := cards_shared.FetchFromURL(squareUrl)
	if err != nil {
		return nil, err
	}

	vegetalUrl := r.vegetalURL + city + "&outputFormat=application/json"
	vegetalData, err := cards_shared.FetchFromURL(vegetalUrl)
	if err != nil {
		return nil, err
	}

	Tdata := append(squareData.Features, vegetalData.Features...)
	data := &cards_shared.FeatureCollection{
		Features: Tdata,
	}

	return data, nil
}

func sortGeneralData(data *cards_shared.FeatureCollection) *cards_shared.FeatureCollection {
	sort.Slice(data.Features, func(i, j int) bool {
		propsI := data.Features[i].Properties.(map[string]interface{})
		propsJ := data.Features[j].Properties.(map[string]interface{})

		anoI, okI := propsI["ano"].(float64)
		anoJ, okJ := propsJ["ano"].(float64)
		if !okI || !okJ {
			panic("campo 'ano' não é do tipo float64")
		}

		return int(anoI) > int(anoJ)
	})

	return data
}

func (r *externalWeatherInfoRepository) LoadInfoData(city string) ([]InfoDataItem, error) {
	data, err := r.loadGeneralData(city)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar dados: %w", err)
	}

	data = sortGeneralData(data)

	// busca o valor mais recente != 0
	findLatestNonZero := func(features []cards_shared.Feature, key string) (float64, bool) {
		for _, feature := range features {
			props, ok := feature.Properties.(map[string]interface{})
			if !ok {
				continue
			}

			value, ok := props[key].(float64)
			if ok && value != 0 {
				return value, true
			}
		}
		return 0, false
	}

	latestB1, foundB1 := findLatestNonZero(data.Features, "b1")
	latestA1, foundA1 := findLatestNonZero(data.Features, "a1")
	latestB3, foundB3 := findLatestNonZero(data.Features, "b3")

	if !foundB1 && !foundA1 && !foundB3 {
		return nil, fmt.Errorf("nenhum dado válido encontrado")
	}

	result := []InfoDataItem{
		{"Média da cobertura vegetal", strconv.Itoa(int(latestB1*100)) + "%"},
		{"Moradores próximos a praças", strconv.Itoa(int(latestA1)) + "%"},
		{"Desigualdade ambiental e social", strconv.Itoa(int(latestB3 * 100))},
	}

	return result, nil
}
