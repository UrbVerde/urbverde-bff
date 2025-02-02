// urbverde-bff/repositories/address/address_data_repository.go
package repositories_address

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AddressDataRepository interface {
	GetLocationData(identifier string, locationType string) (*Location, error)
}

type externalAddressDataRepository struct {
	stateData   LocationResponse
	countryData LocationResponse
	cityData    LocationResponse
}

func NewExternalAddressDataRepository() (AddressDataRepository, error) {
	_ = godotenv.Load()

	// Load state data
	stateData, err := loadJSONFile("repositories/address/data/states.json")
	if err != nil {
		return nil, fmt.Errorf("error loading state data: %w", err)
	}

	// Load country data
	countryData, err := loadJSONFile("repositories/address/data/country.json")
	if err != nil {
		return nil, fmt.Errorf("error loading country data: %w", err)
	}

	// Load city data
	cityData, err := loadJSONFile("repositories/address/data/cities.json")
	if err != nil {
		// Not returning error as city data might not be available yet
		fmt.Printf("Warning: city data not loaded: %v\n", err)
	}

	return &externalAddressDataRepository{
		stateData:   stateData,
		countryData: countryData,
		cityData:    cityData,
	}, nil
}

func loadJSONFile(filename string) (LocationResponse, error) {
	var response LocationResponse
	data, err := os.ReadFile(filename)
	if err != nil {
		return response, fmt.Errorf("error reading file %s: %w", filename, err)
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return response, fmt.Errorf("error unmarshaling %s: %w", filename, err)
	}

	return response, nil
}

func guessLocationType(identifier string) string {
	// Remove any non-numeric characters for code checking
	numericID := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, identifier)

	switch len(numericID) {
	case 0:
		return "country" // Likely a country code like "BR"
	case 2:
		return "state" // State codes in Brazil are 2 digits
	case 7:
		return "city" // City codes in Brazil are 7 digits
	default:
		// If we can't guess from the code length, we'll try to match by display name
		return ""
	}
}

func (r *externalAddressDataRepository) findByDisplayName(displayName string) (*Location, string, error) {
	// Try cities first as they're most specific
	for code, city := range r.cityData.Features {
		if strings.EqualFold(city.DisplayName, displayName) {
			return &city, code, nil
		}
	}

	// Try states
	for code, state := range r.stateData.Features {
		if strings.EqualFold(state.DisplayName, displayName) {
			return &state, code, nil
		}
	}

	// Try country
	for code, country := range r.countryData.Features {
		if strings.EqualFold(country.DisplayName, displayName) {
			return &country, code, nil
		}
	}

	return nil, "", fmt.Errorf("location not found for display name: %s", displayName)
}

func (r *externalAddressDataRepository) GetLocationData(identifier string, locationType string) (*Location, error) {
	// If type wasn't provided, try to guess it
	if locationType == "" {
		locationType = guessLocationType(identifier)
	}

	// If we still don't have a type or if it's empty string,
	// try to find by display name across all types
	if locationType == "" {
		location, _, err := r.findByDisplayName(identifier)
		if err != nil {
			return nil, fmt.Errorf("location not found: %w", err)
		}
		return location, nil
	}

	// If we have a type, look in the specific dataset
	var location *Location
	var found bool

	switch locationType {
	case "state":
		if loc, ok := r.stateData.Features[identifier]; ok {
			location = &loc
			found = true
		}
	case "country":
		if loc, ok := r.countryData.Features[identifier]; ok {
			location = &loc
			found = true
		}
	case "city":
		if loc, ok := r.cityData.Features[identifier]; ok {
			location = &loc
			found = true
		}
	default:
		return nil, fmt.Errorf("invalid location type: %s", locationType)
	}

	if !found {
		// If not found by code with the specified type, try display name as fallback
		location, _, err := r.findByDisplayName(identifier)
		if err != nil {
			return nil, fmt.Errorf("location not found for identifier %s and type %s", identifier, locationType)
		}
		return location, nil
	}

	return location, nil
}
