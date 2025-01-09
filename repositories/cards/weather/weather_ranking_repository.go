package repositories_cards_weather

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type WeatherRankingRepository interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]RankingData, error)
}

// RankingProperties represents the properties used for ranking
type RankingProperties struct {
	Ano          int    `json:"ano"`
	NMMicro      string `json:"nm_micro"`     // Microregion name
	NRankMicro   int    `json:"n_rank_micro"` // Microregion total
	C1RankMicro  int    `json:"c1_rank_micro"`
	C2RankMicro  int    `json:"c2_rank_micro"`
	C3RankMicro  int    `json:"c3_rank_micro"`
	NMMeso       string `json:"nm_meso"`     // Mesoregion name
	NRankMeso    int    `json:"n_rank_meso"` // Mesoregion total
	C1RankMeso   int    `json:"c1_rank_meso"`
	C2RankMeso   int    `json:"c2_rank_meso"`
	C3RankMeso   int    `json:"c3_rank_meso"`
	NMEstado     string `json:"nm_estado"`     // State name
	NRankEstado  int    `json:"n_rank_estado"` // State total
	C1RankEstado int    `json:"c1_rank_estado"`
	C2RankEstado int    `json:"c2_rank_estado"`
	C3RankEstado int    `json:"c3_rank_estado"`
}

type RankingDataItem struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
	Of     int    `json:"of"`
}

// Response JSON structure
type RankingData struct {
	Title    string            `json:"title"`
	Subtitle string            `json:"subtitle"`
	Items    []RankingDataItem `json:"items"`
}

type externalWeatherRankingRepository struct {
	geoserverURL string
}

func NewExternalWeatherRankingRepository() WeatherRankingRepository {
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

	return &externalWeatherRankingRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalWeatherRankingRepository) LoadYears(city string) ([]int, error) {
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

		var rankingProps RankingProperties
		if err := cards_shared.MapToStruct(props, &rankingProps); err != nil {
			return nil, err
		}

		yearsMap[rankingProps.Ano] = true
	}

	var years []int
	for year := range yearsMap {
		years = append(years, year)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	return years, nil
}

func (r *externalWeatherRankingRepository) LoadRankingData(city string, year string) ([]RankingData, error) {
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

		// Map values
		var rankingProps RankingProperties
		rankingProps.Ano = int(props["ano"].(float64))

		// Micro
		rankingProps.NMMicro = props["nm_micro"].(string)
		rankingProps.NRankMicro = int(props["n_rank_micro"].(float64))
		rankingProps.C1RankMicro = int(props["c1_rank_micro"].(float64))
		rankingProps.C2RankMicro = int(props["c2_rank_micro"].(float64))
		rankingProps.C3RankMicro = int(props["c3_rank_micro"].(float64))

		// Meso
		rankingProps.NMMeso = props["nm_meso"].(string)
		rankingProps.NRankMeso = int(props["n_rank_meso"].(float64))
		rankingProps.C1RankMeso = int(props["c1_rank_meso"].(float64))
		rankingProps.C2RankMeso = int(props["c2_rank_meso"].(float64))
		rankingProps.C3RankMeso = int(props["c3_rank_meso"].(float64))

		// Estado
		rankingProps.NMEstado = "São Paulo" // definindo manualmente (temporario)
		rankingProps.NRankEstado = 645
		rankingProps.C1RankEstado = int(props["c1_rank_estado"].(float64))
		rankingProps.C2RankEstado = int(props["c2_rank_estado"].(float64))
		rankingProps.C3RankEstado = int(props["c3_rank_estado"].(float64))

		if rankingProps.Ano == convYear {
			filtered = feature
			filtered.Properties = rankingProps
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ano %d não encontrado nos dados", convYear)
	}

	rankingProps := filtered.Properties.(RankingProperties)

	result := []RankingData{
		{
			Title:    "Municipios do Estado",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d do Estado de %s", rankingProps.NRankEstado, rankingProps.NMEstado),
			Items: []RankingDataItem{
				{Type: "Nível de ilha de calor", Number: rankingProps.C1RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Temperatura média da superfície", Number: rankingProps.C2RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Temperatura máxima da superfície", Number: rankingProps.C3RankEstado, Of: rankingProps.NRankEstado},
			},
		},
		{
			Title:    "Municipios da Mesorregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da mesorregião de %s", rankingProps.NRankMeso, rankingProps.NMMeso),
			Items: []RankingDataItem{
				{Type: "Nível de ilha de calor", Number: rankingProps.C1RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Temperatura média da superfície", Number: rankingProps.C2RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Temperatura máxima da superfície", Number: rankingProps.C3RankMeso, Of: rankingProps.NRankMeso},
			},
		},
		{
			Title:    "Municipios da Microrregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da microrregião de %s", rankingProps.NRankMicro, rankingProps.NMMicro),
			Items: []RankingDataItem{
				{Type: "Nível de ilha de calor", Number: rankingProps.C1RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Temperatura média da superfície", Number: rankingProps.C2RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Temperatura máxima da superfície", Number: rankingProps.C3RankMicro, Of: rankingProps.NRankMicro},
			},
		},
	}

	return result, nil
}
