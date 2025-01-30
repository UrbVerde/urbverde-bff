package services_cards_parks

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_parks "urbverde-api/repositories/cards/parks"
)

type ParksInequalityService interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]repositories_cards_parks.ParksInequalityDataItem, error)
}

type parksInequalityService struct {
	ParksInequalityRepository repositories_cards_parks.ParksInequalityRepository
}

func NewParksInequalityService(ar repositories_cards_parks.ParksInequalityRepository) ParksInequalityService {
	return &parksInequalityService{
		ParksInequalityRepository: ar,
	}
}

func (as *parksInequalityService) LoadYears(city string) ([]int, error) {
	data, err := as.ParksInequalityRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *parksInequalityService) LoadInequalityData(city string, year string) ([]repositories_cards_parks.ParksInequalityDataItem, error) {
	data, err := as.ParksInequalityRepository.LoadInequalityData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
