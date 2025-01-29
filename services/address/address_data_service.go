// urbverde-bff/services/address/address_data_service.go
package services_address

import (
	repositories_address "urbverde-api/repositories/address"
)

type AddressDataService interface {
	GetLocationData(identifier string, locationType string) (*repositories_address.Location, error)
}

type addressDataService struct {
	AddressDataRepository repositories_address.AddressDataRepository
}

func NewAddressDataService(ar repositories_address.AddressDataRepository) AddressDataService {
	return &addressDataService{
		AddressDataRepository: ar,
	}
}

func (as *addressDataService) GetLocationData(identifier string, locationType string) (*repositories_address.Location, error) {
	return as.AddressDataRepository.GetLocationData(identifier, locationType)
}
