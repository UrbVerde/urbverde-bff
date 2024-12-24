// services\address_service.go
package services

import (
	"urbverde-api/repositories"
)

type AddressService interface {
	GetSuggestions(query string) ([]string, error)
}

type addressService struct {
	AddressRepository repositories.AddressRepository
}

// Função para criar um novo serviço de endereço.
func NewAddressService(ar repositories.AddressRepository) AddressService {
	return &addressService{
		AddressRepository: ar,
	}
}

// Função que chama o repositório para buscar as sugestões de endereços.
func (as *addressService) GetSuggestions(query string) ([]string, error) {
	suggestions, err := as.AddressRepository.SearchAddress(query)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}
