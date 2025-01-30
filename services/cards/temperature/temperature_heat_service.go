// urbverde-bff/services/cards/temperature/temperature_heat_service.go
package services_cards_temperature

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_temperature "urbverde-api/repositories/cards/temperature"
)

type TemperatureHeatService interface {
	cards_shared.RepositoryBase
	LoadHeatData(city string, year string) ([]repositories_cards_temperature.HeatDataItem, error)
}

type temperatureHeatService struct {
	TemperatureHeatRepository repositories_cards_temperature.TemperatureHeatRepository
}

func NewTemperatureHeatService(ar repositories_cards_temperature.TemperatureHeatRepository) TemperatureHeatService {
	return &temperatureHeatService{
		TemperatureHeatRepository: ar,
	}
}

func (as *temperatureHeatService) LoadYears(city string) ([]int, error) {
	data, err := as.TemperatureHeatRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *temperatureHeatService) LoadHeatData(city string, year string) ([]repositories_cards_temperature.HeatDataItem, error) {
	data, err := as.TemperatureHeatRepository.LoadHeatData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
