// urbverde-bff/services/address_service.go
package services_address

import (
	repositories_address "urbverde-api/repositories/address"
)

type AddressService interface {
	GetSuggestions(query string) ([]repositories_address.CityResponse, error)
}

type addressService struct {
	AddressRepository repositories_address.AddressRepository
}

func NewAddressService(ar repositories_address.AddressRepository) AddressService {
	return &addressService{
		AddressRepository: ar,
	}
}

func (as *addressService) GetSuggestions(query string) ([]repositories_address.CityResponse, error) {
	suggestions, err := as.AddressRepository.SearchAddress(query)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}