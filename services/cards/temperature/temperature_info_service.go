// urbverde-bff/services/cards/temperature/temperature_info_service.go
package services_cards_temperature

import (
	repositories_cards_temperature "urbverde-api/repositories/cards/temperature"
)

type TemperatureInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_temperature.InfoDataItem, error)
}

type temperatureInfoService struct {
	TemperatureInfoRepository repositories_cards_temperature.TemperatureInfoRepository
}

func NewTemperatureInfoService(ar repositories_cards_temperature.TemperatureInfoRepository) TemperatureInfoService {
	return &temperatureInfoService{
		TemperatureInfoRepository: ar,
	}
}

func (as *temperatureInfoService) LoadInfoData(city string) ([]repositories_cards_temperature.InfoDataItem, error) {
	data, err := as.TemperatureInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
