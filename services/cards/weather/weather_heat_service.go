package services_cards_weather

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
)

type WeatherHeatService interface {
	cards_shared.RepositoryBase
	LoadHeatData(city string, year string) ([]repositories_cards_weather.HeatDataItem, error)
}

type weatherHeatService struct {
	WeatherHeatRepository repositories_cards_weather.WeatherHeatRepository
}

func NewWeatherHeatService(ar repositories_cards_weather.WeatherHeatRepository) WeatherHeatService {
	return &weatherHeatService{
		WeatherHeatRepository: ar,
	}
}

func (as *weatherHeatService) LoadYears(city string) ([]int, error) {
	data, err := as.WeatherHeatRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *weatherHeatService) LoadHeatData(city string, year string) ([]repositories_cards_weather.HeatDataItem, error) {
	data, err := as.WeatherHeatRepository.LoadHeatData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
