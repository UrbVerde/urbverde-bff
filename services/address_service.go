// urbverde-bff/services/address_service.go
package services

import (
	"urbverde-api/repositories"
)

type AddressService interface {
	GetSuggestions(query string) ([]repositories.CityResponse, error)
}

type addressService struct {
	AddressRepository repositories.AddressRepository
}

// NewAddressService creates a new address service instance
func NewAddressService(ar repositories.AddressRepository) AddressService {
	return &addressService{
		AddressRepository: ar,
	}
}

// GetSuggestions retrieves address suggestions based on the query
func (as *addressService) GetSuggestions(query string) ([]repositories.CityResponse, error) {
	suggestions, err := as.AddressRepository.SearchAddress(query)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}
