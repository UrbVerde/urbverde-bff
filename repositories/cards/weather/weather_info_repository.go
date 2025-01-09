package repositories_cards_weather

import (
	"fmt"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type WeatherInfoRepository interface {
	LoadInfoData(city string, year string) ([]InfoDataItem, error)
}

type InfoProperties struct {
	Ano int     `json:"ano"`
	B1  float64 `json:"b1"` // % cobertura vegetal (vegetal)
	A1  float64 `json:"a1"` // Moradores prox a praças (square)
	B3  float64 `json:"b3"` // Desigualdade ambiental e social (square)
}

// Response JSON structure
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

func (r *externalWeatherInfoRepository) LoadInfoData(city string, year string) ([]InfoDataItem, error) {

	convYear, err := strconv.Atoi(year)
	if err != nil {
		return nil, fmt.Errorf("ano inválido: %w", err)
	}

	// square
	squareUrl := r.squareURL + city + "&outputFormat=application/json"
	squareData, err := cards_shared.FetchFromURL(squareUrl)
	if err != nil {
		return nil, err
	}

	// vegetal
	vegetalUrl := r.vegetalURL + city + "&outputFormat=application/json"
	vegetalData, err := cards_shared.FetchFromURL(vegetalUrl)
	if err != nil {
		return nil, err
	}

	Tdata := append(squareData.Features, vegetalData.Features...)
	data := &cards_shared.FeatureCollection{
		Features: Tdata,
	}

	var filtered cards_shared.Feature
	found := false
	for _, feature := range data.Features {
		props, ok := feature.Properties.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("tipo inesperado de propriedades")
		}

		var infoProps InfoProperties
		if err := cards_shared.MapToStruct(props, &infoProps); err != nil {
			return nil, err
		}

		if infoProps.Ano == convYear {
			filtered = feature
			filtered.Properties = infoProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	infoProps := filtered.Properties.(InfoProperties)

	fmt.Println(infoProps)

	result := []InfoDataItem{
		{"Média da cobertura vegetal", strconv.Itoa(int(infoProps.B1*100)) + "%"},
		{"Moradores próximos a praças", strconv.Itoa(int(infoProps.A1)) + "%"},
		{"Desigualdade ambiental e social", strconv.Itoa(int(infoProps.B3 * 100))},
	}

	return result, nil
}
