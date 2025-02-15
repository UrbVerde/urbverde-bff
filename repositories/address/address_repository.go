// urbverde-bff/repositories/address/address_repository.go
package repositories_address

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	data, err := os.ReadFile("repositories/address/data/cities.json")
	if err != nil {
		return nil, fmt.Errorf("error reading cities file: %w", err)
	}

	var citiesData LocationResponse
	if err := json.Unmarshal(data, &citiesData); err != nil {
		return nil, fmt.Errorf("error parsing cities data: %w", err)
	}

	var results []CityResponse
	normalizedQuery := normalizeText(query)

	for code, city := range citiesData.Features {
		normalizedDisplayName := normalizeText(city.DisplayName)
		if strings.Contains(normalizedDisplayName, normalizedQuery) {
			cdMun, err := strconv.Atoi(code)
			if err != nil {
				continue
			}

			results = append(results, CityResponse{
				DisplayName: city.DisplayName,
				CdMun:       cdMun,
				Type:        "city",
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].DisplayName < results[j].DisplayName
	})

	return results, nil
}
