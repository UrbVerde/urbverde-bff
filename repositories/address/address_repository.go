// urbverde-bff/repositories/address/address_repository.go
package repositories_address

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	// Normalize the query for API request
	normalizedAPIQuery := normalizeText(query)
	encodedQuery := url.QueryEscape(normalizedAPIQuery)
	apiURL := r.apiURL + "?nome=" + encodedQuery

	resp, err := http.Get(apiURL)
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
	lowercaseQuery := strings.ToLower(query)
	normalizedQuery := normalizeText(query)

	for _, city := range cities {
		// Only include cities from São Paulo state
		if city.Microrregiao.Mesorregiao.UF.Sigla != "SP" {
			continue
		}

		lowercaseCityName := strings.ToLower(city.Nome)
		normalizedCityName := normalizeText(city.Nome)

		// Skip cities that don't match either with or without accents
		if !strings.HasPrefix(lowercaseCityName, lowercaseQuery) &&
			!strings.HasPrefix(normalizedCityName, normalizedQuery) &&
			!strings.Contains(normalizedCityName, normalizedQuery) {
			continue
		}

		cityResponses = append(cityResponses, CityResponse{
			DisplayName: fmt.Sprintf("%s - %s", city.Nome, city.Microrregiao.Mesorregiao.UF.Sigla),
			CdMun:       city.ID,
		})
	}

	// Sort with custom comparator that prioritizes accent-matches
	sort.Slice(cityResponses, func(i, j int) bool {
		nameI := strings.Split(cityResponses[i].DisplayName, " - ")[0]
		nameJ := strings.Split(cityResponses[j].DisplayName, " - ")[0]

		lowercaseNameI := strings.ToLower(nameI)
		lowercaseNameJ := strings.ToLower(nameJ)

		// Check if either name starts with the query (with accents)
		startsWithI := strings.HasPrefix(lowercaseNameI, lowercaseQuery)
		startsWithJ := strings.HasPrefix(lowercaseNameJ, lowercaseQuery)

		// If one starts with the query and the other doesn't, prioritize the one that does
		if startsWithI != startsWithJ {
			return startsWithI
		}

		// If both or neither start with the query, sort alphabetically
		return cityResponses[i].DisplayName < cityResponses[j].DisplayName
	})

	return cityResponses, nil
}
