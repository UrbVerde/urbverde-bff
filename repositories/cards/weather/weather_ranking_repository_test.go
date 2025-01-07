package repositories_cards_weather

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadRankingData(t *testing.T) {
	_ = godotenv.Load()

	t.Run("Valid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"features": [
					{
						"properties": {
							"ano": 2020,
							"nm_micro": "São Carlos",
							"n_rank_micro": 6,
							"c1_rank_micro": 6,
							"c2_rank_micro": 3,
							"c3_rank_micro": 3,
							"nm_meso": "Araraquara",
							"n_rank_meso": 21,
							"c1_rank_meso": 21,
							"c2_rank_meso": 7,
							"c3_rank_meso": 16,
							"n_rank_estado": 645,
							"c1_rank_estado": 612,
							"c2_rank_estado": 406,
							"c3_rank_estado": 546
						}
					}
				]
			}`))
		}))
		defer server.Close()

		os.Setenv("GEOSERVER_URL", server.URL+"/")

		repo := NewExternalWeatherRankingRepository()

		results, err := repo.LoadRankingData("3548906", "2020")

		assert.NoError(t, err)
		expected := []RankingData{
			{
				Title:    "Municipios do Estado",
				Subtitle: "Posição do seu município entre os 645 do Estado de São Paulo",
				Items: []RankingDataItem{
					{Type: "Nível de ilha de calor", Number: 612, Of: 645},
					{Type: "Temperatura média da superfície", Number: 406, Of: 645},
					{Type: "Temperatura máxima da superfície", Number: 546, Of: 645},
				},
			},
			{
				Title:    "Municipios da Mesorregião",
				Subtitle: "Posição do seu município entre os 21 da mesorregião de Araraquara",
				Items: []RankingDataItem{
					{Type: "Nível de ilha de calor", Number: 21, Of: 21},
					{Type: "Temperatura média da superfície", Number: 7, Of: 21},
					{Type: "Temperatura máxima da superfície", Number: 16, Of: 21},
				},
			},
			{
				Title:    "Municipios da Microrregião",
				Subtitle: "Posição do seu município entre os 6 da microrregião de São Carlos",
				Items: []RankingDataItem{
					{Type: "Nível de ilha de calor", Number: 6, Of: 6},
					{Type: "Temperatura média da superfície", Number: 3, Of: 6},
					{Type: "Temperatura máxima da superfície", Number: 3, Of: 6},
				},
			},
		}

		assert.Equal(t, expected, results)
	})
}
