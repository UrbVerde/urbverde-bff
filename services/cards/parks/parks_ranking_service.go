package services_cards_parks

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_parks "urbverde-api/repositories/cards/parks"
)

type ParksRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_parks.RankingData, error)
}

type parksRankingService struct {
	ParksRankingRepository repositories_cards_parks.ParksRankingRepository
}

func NewParksRankingService(ar repositories_cards_parks.ParksRankingRepository) ParksRankingService {
	return &parksRankingService{
		ParksRankingRepository: ar,
	}
}

func (as *parksRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.ParksRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *parksRankingService) LoadRankingData(city string, year string) ([]repositories_cards_parks.RankingData, error) {
	data, err := as.ParksRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
