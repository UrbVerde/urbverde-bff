package repositories_cards_vegetation

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type VegetationInfoRepository interface {
	LoadInfoData(city string) ([]InfoDataItem, error)
}

type InfoProperties struct {
	Ano int     `json:"ano"`
	C2  float64 `json:"c2"` // Temperatura média da superfície (temperature)
	A1  float64 `json:"a1"` // Moradores prox a praças (square)
	A4  float64 `json:"a4"` // Distancia média até as praças (square)
}

type InfoDataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type externalVegetationInfoRepository struct {
	temperatureURL string
	squareURL      string
}

func NewExternalVegetationInfoRepository() VegetationInfoRepository {
	_ = godotenv.Load()

	geoserverURL := os.Getenv("GEOSERVER_URL")
	if geoserverURL == "" {
		panic("A variável de ambiente GEOSERVER_URL não está definida")
	}

	temperatureURL := fmt.Sprintf("%sows?service=%s&version=%s&request=%s&typeName=%s&%s",
		geoserverURL,
		cards_shared.WfsService,
		cards_shared.WfsVersion,
		cards_shared.WfsRequest,
		cards_shared.TypeName+"dados_temperatura_por_municipio",
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

	return &externalVegetationInfoRepository{
		temperatureURL: temperatureURL,
		squareURL:      squareURL,
	}
}

func (r *externalVegetationInfoRepository) loadGeneralData(city string) (*cards_shared.FeatureCollection, error) {
	squareUrl := r.squareURL + city + "&outputFormat=application/json"
	squareData, err := cards_shared.FetchFromURL(squareUrl)
	if err != nil {
		return nil, err
	}

	temperatureUrl := r.temperatureURL + city + "&outputFormat=application/json"
	temperatureData, err := cards_shared.FetchFromURL(temperatureUrl)
	if err != nil {
		return nil, err
	}

	Tdata := append(squareData.Features, temperatureData.Features...)
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

func (r *externalVegetationInfoRepository) LoadInfoData(city string) ([]InfoDataItem, error) {
	data, err := r.loadGeneralData(city)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar dados: %w", err)
	}

	data = sortGeneralData(data)

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

	latestC2, foundC2 := findLatestNonZero(data.Features, "c2")
	latestA1, foundA1 := findLatestNonZero(data.Features, "a1")
	latestA4, foundA4 := findLatestNonZero(data.Features, "a4")

	if !foundC2 && !foundA1 && !foundA4 {
		return nil, fmt.Errorf("nenhum dado válido encontrado")
	}

	result := []InfoDataItem{
		{"Temperatura média da superfície", strconv.Itoa(int(latestC2)) + "°C"},
		{"Moradores próximos a praças", strconv.Itoa(int(latestA1)) + "%"},
		{"Distancia média até as praças", strconv.Itoa(int(latestA4)) + " metros"},
	}

	return result, nil
}
