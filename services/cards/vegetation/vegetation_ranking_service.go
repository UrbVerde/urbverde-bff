package services_cards_vegetation

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetation "urbverde-api/repositories/cards/vegetation"
)

type VegetationRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_vegetation.RankingData, error)
}

type vegetationRankingService struct {
	VegetationRankingRepository repositories_cards_vegetation.VegetationRankingRepository
}

func NewVegetationRankingService(ar repositories_cards_vegetation.VegetationRankingRepository) VegetationRankingService {
	return &vegetationRankingService{
		VegetationRankingRepository: ar,
	}
}

func (as *vegetationRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetationRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetationRankingService) LoadRankingData(city string, year string) ([]repositories_cards_vegetation.RankingData, error) {
	data, err := as.VegetationRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
