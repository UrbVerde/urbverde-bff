package services_cards_parks

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_parks "urbverde-api/repositories/cards/parks"
)

type ParksSquareService interface {
	cards_shared.RepositoryBase
	LoadSquareData(city string, year string) ([]repositories_cards_parks.SquareDataItem, error)
}

type parksSquareService struct {
	ParksSquareRepository repositories_cards_parks.ParksSquareRepository
}

func NewParksSquareService(ar repositories_cards_parks.ParksSquareRepository) ParksSquareService {
	return &parksSquareService{
		ParksSquareRepository: ar,
	}
}

func (as *parksSquareService) LoadYears(city string) ([]int, error) {
	data, err := as.ParksSquareRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *parksSquareService) LoadSquareData(city string, year string) ([]repositories_cards_parks.SquareDataItem, error) {
	data, err := as.ParksSquareRepository.LoadSquareData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
