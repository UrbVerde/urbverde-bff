// urbverde-bff/repositories/address_repository.go
package repositories

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

// IBGEResponse represents the raw response from IBGE API
type IBGEResponse struct {
	ID           int    `json:"id"`   // City ID (cd_mun)
	Nome         string `json:"nome"` // City name
	Microrregiao struct {
		Mesorregiao struct {
			UF struct {
				Sigla string `json:"sigla"` // State abbreviation
			} `json:"UF"`
		} `json:"mesorregiao"`
	} `json:"microrregiao"`
}

// CityResponse represents our formatted response
type CityResponse struct {
	DisplayName string `json:"display_name"` // What user sees: "City Name - ST"
	CdMun       int    `json:"cd_mun"`       // City ID for internal use
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

func (r *externalAddressRepository) SearchAddress(query string) ([]CityResponse, error) {
	url := r.apiURL

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
	qL := strings.ToLower(query)
	for _, city := range cities {
		cityName := strings.ToLower(city.Nome)
		if strings.HasPrefix(cityName, qL) {
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
