// urbverde-bff/repositories/address_repository_test.go
package repositories

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSearchAddress(t *testing.T) {
	_ = godotenv.Load()

	t.Run("Valid response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{"nome": "São Paulo", "microrregiao": {"mesorregiao": {"UF": {"sigla": "SP"}}}},
				{"nome": "São Pedro", "microrregiao": {"mesorregiao": {"UF": {"sigla": "SP"}}}}
			]`))
		}))
		defer server.Close()

		os.Setenv("IBGE_API_URL", server.URL)

		repo := NewExternalAddressRepository()
		results, err := repo.SearchAddress("São")

		assert.NoError(t, err)

		expected := []CityResponse{
			{DisplayName: "São Paulo - SP", CdMun: 3550308},
			{DisplayName: "São Pedro - SP", CdMun: 3550407},
		}

		var actualStrings []string
		for _, result := range results {
			actualStrings = append(actualStrings, result.DisplayName)
		}

		var expectedStrings []string
		for _, city := range expected {
			expectedStrings = append(expectedStrings, city.DisplayName)
		}

		assert.Equal(t, expectedStrings, actualStrings)
	})
}
