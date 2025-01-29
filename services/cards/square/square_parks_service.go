package services_cards_square

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_square "urbverde-api/repositories/cards/square"
)

type SquareParksService interface {
	cards_shared.RepositoryBase
	LoadParksData(city string, year string) ([]repositories_cards_square.ParksDataItem, error)
}

type squareParksService struct {
	SquareParksRepository repositories_cards_square.SquareParksRepository
}

func NewSquareParksService(ar repositories_cards_square.SquareParksRepository) SquareParksService {
	return &squareParksService{
		SquareParksRepository: ar,
	}
}

func (as *squareParksService) LoadYears(city string) ([]int, error) {
	data, err := as.SquareParksRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *squareParksService) LoadParksData(city string, year string) ([]repositories_cards_square.ParksDataItem, error) {
	data, err := as.SquareParksRepository.LoadParksData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
