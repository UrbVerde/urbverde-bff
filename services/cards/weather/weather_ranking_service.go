// urbverde-bff/services/cards/weather/weather_ranking_service.go
package services_cards_weather

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
)

type WeatherRankingService interface {
	cards_shared.RepositoryBase
	LoadRankingData(city string, year string) ([]repositories_cards_weather.RankingData, error)
}

type weatherRankingService struct {
	WeatherRankingRepository repositories_cards_weather.WeatherRankingRepository
}

func NewWeatherRankingService(ar repositories_cards_weather.WeatherRankingRepository) WeatherRankingService {
	return &weatherRankingService{
		WeatherRankingRepository: ar,
	}
}

func (as *weatherRankingService) LoadYears(city string) ([]int, error) {
	data, err := as.WeatherRankingRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *weatherRankingService) LoadRankingData(city string, year string) ([]repositories_cards_weather.RankingData, error) {
	data, err := as.WeatherRankingRepository.LoadRankingData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
