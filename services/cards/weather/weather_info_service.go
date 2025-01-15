// urbverde-bff/services/cards/weather/weather_info_service.go
package services_cards_weather

import (
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
)

type WeatherInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_weather.InfoDataItem, error)
}

type weatherInfoService struct {
	WeatherInfoRepository repositories_cards_weather.WeatherInfoRepository
}

func NewWeatherInfoService(ar repositories_cards_weather.WeatherInfoRepository) WeatherInfoService {
	return &weatherInfoService{
		WeatherInfoRepository: ar,
	}
}

func (as *weatherInfoService) LoadInfoData(city string) ([]repositories_cards_weather.InfoDataItem, error) {
	data, err := as.WeatherInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
