package services_cards_weather

import (
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
)

type WeatherTemperatureService interface {
	LoadYears(city string) ([]int, error)
	LoadData(city string, year string) ([]repositories_cards_weather.DataItem, error)
}

type weatherTemperatureService struct {
	WeatherTemperatureRepository repositories_cards_weather.WeatherTemperatureRepository
}

func NewWeatherTemperatureService(ar repositories_cards_weather.WeatherTemperatureRepository) WeatherTemperatureService {
	return &weatherTemperatureService{
		WeatherTemperatureRepository: ar,
	}
}

func (as *weatherTemperatureService) LoadYears(city string) ([]int, error) {
	data, err := as.WeatherTemperatureRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *weatherTemperatureService) LoadData(city string, year string) ([]repositories_cards_weather.DataItem, error) {
	data, err := as.WeatherTemperatureRepository.LoadData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
