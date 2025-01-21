package repositories_cards_vegetal

import (
	"fmt"
	"math"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type VegetalCoverRepository interface {
	cards_shared.RepositoryBase
	LoadCoverData(city string, year string) ([]CoverDataItem, error)
}

type CoverProperties struct {
	Ano  int     `json:"ano"`
	B1   float64 `json:"b1"`   // % da cobertura vegetal
	B1h1 float64 `json:"b1h1"` // Área vegetada em campos de futebol
	B1h3 float64 `json:"b1h3"` // Variação Min
	B1h4 float64 `json:"b1h4"` // Variação Max
}

// Response JSON structure
type CoverDataItem struct {
	Title    string  `json:"title"`
	Subtitle *string `json:"subtitle,omitempty"` // Omitir caso seja nil
	Value    string  `json:"value"`
}

type externalVegetalCoverRepository struct {
	geoserverURL string
}

func NewExternalVegetalCoverRepository() VegetalCoverRepository {
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
		cards_shared.TypeName+"vegetacao_highlights_data",
		cards_shared.CqlFilterPrefix,
	)

	return &externalVegetalCoverRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalVegetalCoverRepository) LoadYears(city string) ([]int, error) {
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

func tempLoadCoverData(v1 int, sub1 *string) {
	cards_shared.AuxLoadSubtitles(v1, 0, sub1) // a media nacional deve ser incluida posteriormente
}

func (r *externalVegetalCoverRepository) LoadCoverData(city string, year string) ([]CoverDataItem, error) {
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

		var coverProps CoverProperties
		if err := cards_shared.MapToStruct(props, &coverProps); err != nil {
			return nil, err
		}

		if coverProps.Ano == convYear {
			filtered = feature
			filtered.Properties = coverProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	coverProps := filtered.Properties.(CoverProperties)

	avg_cover_value := int(coverProps.B1 * 100)
	futebol_cover_value := int(math.Round(coverProps.B1h1))
	cover_max_value := int(math.Round(coverProps.B1h4))
	cover_min_value := int(math.Round(coverProps.B1h3))

	var avg_cover_subtitle string = " da média nacional de "
	var futebol_cover_subtitle string = "* Um campo equivale à 6.400 metros quadrados"

	tempLoadCoverData(avg_cover_value, &avg_cover_subtitle)

	result := []CoverDataItem{
		{"A área vegetada é igual a", &futebol_cover_subtitle, strconv.Itoa(futebol_cover_value) + " campos de futebol*"},
		{"Média da cobertura vegetal", &avg_cover_subtitle, strconv.Itoa(avg_cover_value) + "%"},
		{"A cobertura vegetal na cidade varia entre", nil, strconv.Itoa(cover_min_value) + "% a " + strconv.Itoa(cover_max_value) + "%"},
	}

	return result, nil
}
