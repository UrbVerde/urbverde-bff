// urbverde-bff/repositories/address/address_repository.go
package repositories_address

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

type AddressRepository interface {
	SearchAddress(query string) ([]CityResponse, error)
}

type externalAddressRepository struct {
	apiURL string
}

func NewExternalAddressRepository() AddressRepository {
	_ = godotenv.Load()

	apiURL := os.Getenv("IBGE_API_URL")
	if apiURL == "" {
		panic("IBGE_API_URL environment variable not set")
	}

	return &externalAddressRepository{
		apiURL: apiURL,
	}
}

// improved matching for city names with accents
var replacer = strings.NewReplacer(
	"á", "a", "à", "a", "ã", "a", "â", "a",
	"é", "e", "ê", "e",
	"í", "i",
	"ó", "o", "ô", "o", "õ", "o",
	"ú", "u",
	"ç", "c",
)

func normalizeText(s string) string {
	return replacer.Replace(strings.ToLower(s))
}

func (r *externalAddressRepository) SearchAddress(query string) ([]CityResponse, error) {
	url := r.apiURL + "?nome=" + query

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	var cities []IBGEResponse
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var cityResponses []CityResponse
	normalizedQuery := normalizeText(query)

	for _, city := range cities {
		normalizedCityName := normalizeText(city.Nome)
		if strings.HasPrefix(normalizedCityName, normalizedQuery) {
			cityResponses = append(cityResponses, CityResponse{
				DisplayName: fmt.Sprintf("%s - %s", city.Nome, city.Microrregiao.Mesorregiao.UF.Sigla),
				CdMun:       city.ID,
			})
		}
	}

	// Sort cityResponses alphabetically by DisplayName
	sort.Slice(cityResponses, func(i, j int) bool {
		return cityResponses[i].DisplayName < cityResponses[j].DisplayName
	})

	return cityResponses, nil
}
