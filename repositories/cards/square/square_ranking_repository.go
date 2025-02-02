package repositories_cards_square

import (
	"fmt"
	"os"
	"strconv"
	cards_shared "urbverde-api/repositories/cards"

	"github.com/joho/godotenv"
)

type SquareRankingRepository interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]RankingData, error)
}

type RankingProperties struct {
	Ano          int    `json:"ano"`
	NMMicro      string `json:"nm_micro"`     // Microregion name
	NRankMicro   int    `json:"n_rank_micro"` // Microregion total
	A1RankMicro  int    `json:"a1_rank_micro"`
	A2RankMicro  int    `json:"a2_rank_micro"`
	A3RankMicro  int    `json:"a3_rank_micro"`
	A4RankMicro  int    `json:"a4_rank_micro"`
	NMMeso       string `json:"nm_meso"`     // Mesoregion name
	NRankMeso    int    `json:"n_rank_meso"` // Mesoregion total
	A1RankMeso   int    `json:"a1_rank_meso"`
	A2RankMeso   int    `json:"a2_rank_meso"`
	A3RankMeso   int    `json:"a3_rank_meso"`
	A4RankMeso   int    `json:"a4_rank_meso"`
	NMEstado     string `json:"nm_estado"`     // State name
	NRankEstado  int    `json:"n_rank_estado"` // State total
	A1RankEstado int    `json:"a1_rank_estado"`
	A2RankEstado int    `json:"a2_rank_estado"`
	A3RankEstado int    `json:"a3_rank_estado"`
	A4RankEstado int    `json:"a4_rank_estado"`
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

type externalSquareRankingRepository struct {
	geoserverURL string
}

func NewExternalSquareRankingRepository() SquareRankingRepository {
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
		cards_shared.TypeName+"dados_pracas_por_municipio",
		cards_shared.CqlFilterPrefix,
	)

	return &externalSquareRankingRepository{
		geoserverURL: geoserverURL,
	}
}

func (r *externalSquareRankingRepository) LoadYears(city string) ([]int, error) {
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

func (r *externalSquareRankingRepository) LoadRankingData(city string, year string) ([]RankingData, error) {
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
		rankingProps.A1RankMicro = int(props["a1_rank_micro"].(float64))
		rankingProps.A2RankMicro = int(props["a2_rank_micro"].(float64))
		rankingProps.A3RankMicro = int(props["a3_rank_micro"].(float64))
		rankingProps.A4RankMicro = int(props["a4_rank_micro"].(float64))

		// Meso
		rankingProps.NMMeso = props["nm_meso"].(string)
		rankingProps.NRankMeso = int(props["n_rank_meso"].(float64))
		rankingProps.A1RankMeso = int(props["a1_rank_meso"].(float64))
		rankingProps.A2RankMeso = int(props["a2_rank_meso"].(float64))
		rankingProps.A3RankMeso = int(props["a3_rank_meso"].(float64))
		rankingProps.A4RankMeso = int(props["a4_rank_meso"].(float64))

		// Estado
		rankingProps.NMEstado = "São Paulo" // definindo manualmente (temporario)
		rankingProps.NRankEstado = 645
		rankingProps.A1RankEstado = int(props["a1_rank_estado"].(float64))
		rankingProps.A2RankEstado = int(props["a2_rank_estado"].(float64))
		rankingProps.A3RankEstado = int(props["a3_rank_estado"].(float64))
		rankingProps.A4RankEstado = int(props["a4_rank_estado"].(float64))

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
				{Type: "População atendida pelas praças", Number: rankingProps.A1RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Área de praças por habitante", Number: rankingProps.A2RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Área ocupada pela vizinhança das praças", Number: rankingProps.A3RankEstado, Of: rankingProps.NRankEstado},
				{Type: "Distribuição das praças na cidade", Number: rankingProps.A4RankEstado, Of: rankingProps.NRankEstado},
			},
		},
		{
			Title:    "Municipios da Mesorregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da mesorregião de %s", rankingProps.NRankMeso, rankingProps.NMMeso),
			Items: []RankingDataItem{
				{Type: "População atendida pelas praças", Number: rankingProps.A1RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Área de praças por habitante", Number: rankingProps.A2RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Área ocupada pela vizinhança das praças", Number: rankingProps.A3RankMeso, Of: rankingProps.NRankMeso},
				{Type: "Distribuição das praças na cidade", Number: rankingProps.A4RankMeso, Of: rankingProps.NRankMeso},
			},
		},
		{
			Title:    "Municipios da Microrregião",
			Subtitle: fmt.Sprintf("Posição do seu município entre os %d da microrregião de %s", rankingProps.NRankMicro, rankingProps.NMMicro),
			Items: []RankingDataItem{
				{Type: "População atendida pelas praças", Number: rankingProps.A1RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Área de praças por habitante", Number: rankingProps.A2RankMicro, Of: rankingProps.NRankMicro},
				{Type: "Área ocupada pela vizinhança das praças", Number: rankingProps.A3RankMicro, Of: rankingProps.NRankMeso},
				{Type: "Distribuição das praças na cidade", Number: rankingProps.A4RankMicro, Of: rankingProps.NRankMicro},
			},
		},
	}

	return result, nil
}
