package services_cards_square

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_square "urbverde-api/repositories/cards/square"
)

type SquareInequalityService interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]repositories_cards_square.SquareInequalityDataItem, error)
}

type squareInequalityService struct {
	SquareInequalityRepository repositories_cards_square.SquareInequalityRepository
}

func NewSquareInequalityService(ar repositories_cards_square.SquareInequalityRepository) SquareInequalityService {
	return &squareInequalityService{
		SquareInequalityRepository: ar,
	}
}

func (as *squareInequalityService) LoadYears(city string) ([]int, error) {
	data, err := as.SquareInequalityRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *squareInequalityService) LoadInequalityData(city string, year string) ([]repositories_cards_square.SquareInequalityDataItem, error) {
	data, err := as.SquareInequalityRepository.LoadInequalityData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
