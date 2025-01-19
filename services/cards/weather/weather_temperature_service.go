// urbverde-bff/services/cards/weather/weather_temperature_service.go
package services_cards_weather

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
)

type WeatherTemperatureService interface {
	cards_shared.RepositoryBase
	LoadTemperatureData(city string, year string) ([]repositories_cards_weather.TemperatureDataItem, error)
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

func (as *weatherTemperatureService) LoadTemperatureData(city string, year string) ([]repositories_cards_weather.TemperatureDataItem, error) {
	data, err := as.WeatherTemperatureRepository.LoadTemperatureData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
