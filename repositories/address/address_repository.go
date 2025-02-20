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

	var exactMatches []CityResponse
	var prefixMatches []CityResponse
	var otherMatches []CityResponse

	normalizedQuery := normalizeText(query)

	for _, city := range cities {
		normalizedCityName := normalizeText(city.Nome)
		displayName := fmt.Sprintf("%s - %s", city.Nome, city.Microrregiao.Mesorregiao.UF.Sigla)

		if strings.HasPrefix(normalizedCityName, normalizedQuery) {
			if normalizedCityName == normalizedQuery {
				// Exact match
				exactMatches = append(exactMatches, CityResponse{
					DisplayName: displayName,
					CdMun:       city.ID,
				})
			} else {
				// Prefix match
				prefixMatches = append(prefixMatches, CityResponse{
					DisplayName: displayName,
					CdMun:       city.ID,
				})
			}
		} else if strings.Contains(normalizedCityName, normalizedQuery) {
			// Contains but not prefix
			otherMatches = append(otherMatches, CityResponse{
				DisplayName: displayName,
				CdMun:       city.ID,
			})
		}
	}

	// Sort each category alphabetically
	sortByDisplayName := func(responses []CityResponse) {
		sort.Slice(responses, func(i, j int) bool {
			return responses[i].DisplayName < responses[j].DisplayName
		})
	}

	sortByDisplayName(exactMatches)
	sortByDisplayName(prefixMatches)
	sortByDisplayName(otherMatches)

	// Combine results with priority: exact -> prefix -> contains
	result := append(exactMatches, prefixMatches...)
	result = append(result, otherMatches...)

	return result, nil
}
