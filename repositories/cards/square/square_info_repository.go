package repositories_cards_square

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type SquareInfoRepository interface {
	LoadInfoData(city string) ([]InfoDataItem, error)
}

type InfoProperties struct {
	Ano  int     `json:"ano"`
	C2   float64 `json:"c2"`   // Temperatura média da superfície (temperature)
	B1h1 float64 `json:"b1h1"` // Média da cobertura vegetal (vegetal)
	B3   float64 `json:"b3"`   // Desigualdade ambiental e social (vegetal))
}

type InfoDataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type externalSquareInfoRepository struct {
	temperatureURL string
	vegetalURL     string
}

func NewExternalSquareInfoRepository() SquareInfoRepository {
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

	vegetalURL := fmt.Sprintf("%sows?service=%s&version=%s&request=%s&typeName=%s&%s",
		geoserverURL,
		cards_shared.WfsService,
		cards_shared.WfsVersion,
		cards_shared.WfsRequest,
		cards_shared.TypeName+"vegetacao_highlights_data",
		cards_shared.CqlFilterPrefix,
	)

	return &externalSquareInfoRepository{
		temperatureURL: temperatureURL,
		vegetalURL:     vegetalURL,
	}
}

func (r *externalSquareInfoRepository) loadGeneralData(city string) (*cards_shared.FeatureCollection, error) {
	vegetalUrl := r.vegetalURL + city + "&outputFormat=application/json"
	vegetalData, err := cards_shared.FetchFromURL(vegetalUrl)
	if err != nil {
		return nil, err
	}

	temperatureUrl := r.temperatureURL + city + "&outputFormat=application/json"
	temperatureData, err := cards_shared.FetchFromURL(temperatureUrl)
	if err != nil {
		return nil, err
	}

	Tdata := append(vegetalData.Features, temperatureData.Features...)
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

func (r *externalSquareInfoRepository) LoadInfoData(city string) ([]InfoDataItem, error) {
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
	latestB1h1, foundB1h1 := findLatestNonZero(data.Features, "b1h1")
	latestB3, foundB3 := findLatestNonZero(data.Features, "b3")

	if !foundC2 && !foundB1h1 && !foundB3 {
		return nil, fmt.Errorf("nenhum dado válido encontrado")
	}

	result := []InfoDataItem{
		{"Temperatura média da superfície", strconv.Itoa(int(latestC2)) + "°C"},
		{"Média da cobertura vegetal", strconv.Itoa(int(latestB1h1)) + "%"},
		{"Desigualdade ambiental e social", strconv.Itoa(int(latestB3))},
	}

	return result, nil
}
