// urbverde-bff/services/cards/temperature/temperature_temperature_service.go
package services_cards_temperature

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_temperature "urbverde-api/repositories/cards/temperature"
)

type TemperatureWeatherService interface {
	cards_shared.RepositoryBase
	LoadWeatherData(city string, year string) ([]repositories_cards_temperature.WeatherDataItem, error)
}

type temperatureWeatherService struct {
	TemperatureWeatherRepository repositories_cards_temperature.TemperatureWeatherRepository
}

func NewTemperatureWeatherService(ar repositories_cards_temperature.TemperatureWeatherRepository) TemperatureWeatherService {
	return &temperatureWeatherService{
		TemperatureWeatherRepository: ar,
	}
}

func (as *temperatureWeatherService) LoadYears(city string) ([]int, error) {
	data, err := as.TemperatureWeatherRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *temperatureWeatherService) LoadWeatherData(city string, year string) ([]repositories_cards_temperature.WeatherDataItem, error) {
	data, err := as.TemperatureWeatherRepository.LoadWeatherData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
