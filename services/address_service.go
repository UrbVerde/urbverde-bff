package services

import (
	"urbverde-api/repositories"
)

type AddressService interface {
	GetSuggestions(query string, filter string) ([]repositories.Suggestion, error)
}

type addressService struct {
	AddressRepository repositories.AddressRepository
}

func NewAddressService(ar repositories.AddressRepository) AddressService {
	return &addressService{
		AddressRepository: ar,
	}
}

func (as *addressService) GetSuggestions(query string, filter string) ([]repositories.Suggestion, error) {
	suggestions, err := as.AddressRepository.SearchAddress(query, filter)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}
