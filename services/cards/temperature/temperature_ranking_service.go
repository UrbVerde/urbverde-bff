// urbverde-bff/services/cards/temperature/temperature_ranking_service.go
package services_cards_temperature

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_temperature "urbverde-api/repositories/cards/temperature"
)

type TemperatureRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_temperature.RankingData, error)
}

type temperatureRankingService struct {
	TemperatureRankingRepository repositories_cards_temperature.TemperatureRankingRepository
}

func NewTemperatureRankingService(ar repositories_cards_temperature.TemperatureRankingRepository) TemperatureRankingService {
	return &temperatureRankingService{
		TemperatureRankingRepository: ar,
	}
}

func (as *temperatureRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.TemperatureRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *temperatureRankingService) LoadRankingData(city string, year string) ([]repositories_cards_temperature.RankingData, error) {
	data, err := as.TemperatureRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
