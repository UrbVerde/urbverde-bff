// urbverde-bff/repositories/cards/temperature/temperature_weather_repository_test.go
package repositories_cards_temperature

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"urbverde-api/utils"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadWeatherData(t *testing.T) {
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

		repo := NewExternalTemperatureWeatherRepository()

		results, err := repo.LoadWeatherData("3548906", "2020")

		assert.NoError(t, err)

		expected := []WeatherDataItem{
			{"Nível de ilha de calor", nil, "2"},
			{"Temperatura média da superfície", nil, "32°C"},
			{"Maior amplitude", utils.StringPtr("É a diferença entre a temperatura mais quente e a mais fria"), "9°C"},
			{"Temperatura máxima da superfície", nil, "42°C"},
		}

		assert.Equal(t, expected, results)
	})
}
