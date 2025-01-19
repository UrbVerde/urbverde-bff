// urbverde-bff/repositories/cards/weather/weather_temperature_repository_test.go
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

func TestLoadTemperatureData(t *testing.T) {
	_ = godotenv.Load()

	t.Run("Valid response", func(t *testing.T) {
		// Mock server to simulate the Geoserver response
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"features": [
					{
						"properties": {
							"ano": 2020,
							"c1": 2.0,
							"h5b": 9,
							"c2": 32.0,
							"c3": 42.0
						}
					}
				]
			}`))
		}))
		defer server.Close()

		os.Setenv("GEOSERVER_URL", server.URL+"/")

		repo := NewExternalWeatherTemperatureRepository()

		results, err := repo.LoadTemperatureData("3548906", "2020")

		assert.NoError(t, err)

		expected := []TemperatureDataItem{
			{"Nível de ilha de calor", utils.StringPtr("Acima da média nacional de 0"), "2"},
			{"Temperatura média da superfície", utils.StringPtr("Acima da média nacional de 0"), "32°C"},
			{"Maior amplitude", utils.StringPtr("É a diferença entre a temperatura mais quente e a mais fria"), "9°C"},
			{"Temperatura máxima da superfície", nil, "42°C"},
		}

		assert.Equal(t, expected, results)
	})
}
