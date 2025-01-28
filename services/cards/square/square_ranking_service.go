package services_cards_square

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_square "urbverde-api/repositories/cards/square"
)

type SquareRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_square.RankingData, error)
}

type squareRankingService struct {
	SquareRankingRepository repositories_cards_square.SquareRankingRepository
}

func NewSquareRankingService(ar repositories_cards_square.SquareRankingRepository) SquareRankingService {
	return &squareRankingService{
		SquareRankingRepository: ar,
	}
}

func (as *squareRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.SquareRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *squareRankingService) LoadRankingData(city string, year string) ([]repositories_cards_square.RankingData, error) {
	data, err := as.SquareRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
