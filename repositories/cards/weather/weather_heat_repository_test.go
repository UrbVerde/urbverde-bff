// urbverde-bff/repositories/cards/weather/weather_heat_repository_test.go
package repositories_cards_weather

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"urbverde-api/utils"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadHeatData(t *testing.T) {
	_ = godotenv.Load()

	t.Run("Valid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"features": [
					{
						"properties": {
							"ano": 2020,
							"h12b": 0.25,
							"h11b": 0.39,
							"h10b": 0.15,
							"h9b": 0.12
						}
					}
				]
			}`))
		}))
		defer server.Close()

		os.Setenv("GEOSERVER_URL", server.URL+"/")

		repo := NewExternalWeatherHeatRepository()

		results, err := repo.LoadHeatData("3548906", "2020")

		assert.NoError(t, err)

		expected := []HeatDataItem{
			{"Negros e indígenas afetados", utils.StringPtr("Porcentagem vivendo nas regiões mais quentes"), "25%"},
			{"Mulheres afetadas", utils.StringPtr("Porcentagem vivendo nas regiões mais quentes"), "39%"},
			{"Crianças afetadas", utils.StringPtr("Porcentagem vivendo nas regiões mais quentes"), "15%"},
			{"Idosos afetados", utils.StringPtr("Porcentagem vivendo nas regiões mais quentes"), "12%"},
		}

		assert.Equal(t, expected, results)
	})
}
