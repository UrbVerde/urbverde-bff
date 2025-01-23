package services_cards_vegetal

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetal "urbverde-api/repositories/cards/vegetal"
)

type VegetalRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_vegetal.RankingData, error)
}

type vegetalRankingService struct {
	VegetalRankingRepository repositories_cards_vegetal.VegetalRankingRepository
}

func NewVegetalRankingService(ar repositories_cards_vegetal.VegetalRankingRepository) VegetalRankingService {
	return &vegetalRankingService{
		VegetalRankingRepository: ar,
	}
}

func (as *vegetalRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetalRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetalRankingService) LoadRankingData(city string, year string) ([]repositories_cards_vegetal.RankingData, error) {
	data, err := as.VegetalRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
