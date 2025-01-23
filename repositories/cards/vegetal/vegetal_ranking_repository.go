package repositories_cards_vegetal

import (
	"fmt"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type VegetalRankingRepository interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]RankingData, error)
}

type RankingProperties struct {
	Ano          int    `json:"ano"`
	NMMicro      string `json:"nm_micro"`     // Microregion name
	NRankMicro   int    `json:"n_rank_micro"` // Microregion total
	B1RankMicro  int    `json:"b1_rank_micro"`
	B2RankMicro  int    `json:"b2_rank_micro"`
	B3RankMicro  int    `json:"b3_rank_micro"`
	NMMeso       string `json:"nm_meso"`     // Mesoregion name
	NRankMeso    int    `json:"n_rank_meso"` // Mesoregion total
	B1RankMeso   int    `json:"b1_rank_meso"`
	B2RankMeso   int    `json:"b2_rank_meso"`
	B3RankMeso   int    `json:"b3_rank_meso"`
	NMEstado     string `json:"nm_estado"`     // State name
	NRankEstado  int    `json:"n_rank_estado"` // State total
	B1RankEstado int    `json:"b1_rank_estado"`
	B2RankEstado int    `json:"b2_rank_estado"`
	B3RankEstado int    `json:"b3_rank_estado"`
}

type RankingDataItem struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
	Of     int    `json:"of"`
}

type RankingData struct {
	Title    string            `json:"title"`
	Subtitle string            `json:"subtitle"`
	Items    []RankingDataItem `json:"items"`
}

type externalVegetalRankingRepository struct {
	geoserverURL string
}

func NewExternalVegetalRankingRepository() VegetalRankingRepository {
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

	return &externalVegetalRankingRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalVegetalRankingRepository) LoadYears(city string) ([]int, error) {
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

func (r *externalVegetalRankingRepository) LoadRankingData(city string, year string) ([]RankingData, error) {
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
		rankingProps.B1RankMicro = int(props["b1_rank_micro"].(float64))
		rankingProps.B2RankMicro = int(props["b2_rank_micro"].(float64))
		rankingProps.B3RankMicro = int(props["b3_rank_micro"].(float64))

		// Meso
		rankingProps.NMMeso = props["nm_meso"].(string)
		rankingProps.NRankMeso = int(props["n_rank_meso"].(float64))
		rankingProps.B1RankMeso = int(props["b1_rank_meso"].(float64))
		rankingProps.B2RankMeso = int(props["b2_rank_meso"].(float64))
		rankingProps.B3RankMeso = int(props["b3_rank_meso"].(float64))

		// Estado
		rankingProps.NMEstado = "São Paulo" // definindo manualmente (temporario)
		rankingProps.NRankEstado = 645
		rankingProps.B1RankEstado = int(props["b1_rank_estado"].(float64))
		rankingProps.B2RankEstado = int(props["b2_rank_estado"].(float64))
		rankingProps.B3RankEstado = int(props["b3_rank_estado"].(float64))

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
				{Type: "Percentual de cobertura vegetal (PCV)", Number: rankingProps.B1RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Índice de cobertura vegetal (ICV)", Number: rankingProps.B2RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Nível de desigualdade socioambiental (IDSA)", Number: rankingProps.B3RankEstado, Of: rankingProps.NRankEstado},
			},
		},
		{
			Title:    "Municipios da Mesorregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da mesorregião de %s", rankingProps.NRankMeso, rankingProps.NMMeso),
			Items: []RankingDataItem{
				{Type: "Percentual de cobertura vegetal (PCV)", Number: rankingProps.B1RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Índice de cobertura vegetal (ICV)", Number: rankingProps.B2RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Nível de desigualdade socioambiental (IDSA)", Number: rankingProps.B3RankMeso, Of: rankingProps.NRankMeso},
			},
		},
		{
			Title:    "Municipios da Microrregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da microrregião de %s", rankingProps.NRankMicro, rankingProps.NMMicro),
			Items: []RankingDataItem{
				{Type: "Percentual de cobertura vegetal (PCV)", Number: rankingProps.B1RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Índice de cobertura vegetal (ICV)", Number: rankingProps.B2RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Nível de desigualdade socioambiental (IDSA)", Number: rankingProps.B3RankMicro, Of: rankingProps.NRankMicro},
			},
		},
	}

	return result, nil
}
