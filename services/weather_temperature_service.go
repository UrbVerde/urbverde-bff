package services

import (
	"urbverde-api/repositories"
)

type WeatherTemperatureService interface {
	LoadYears(city string) ([]int, error)
	LoadData(city string, year string) ([]repositories.DataItem, error)
}

type weatherTemperatureService struct {
	WeatherTemperatureRepository repositories.WeatherTemperatureRepository
}

func NewWeatherTemperatureService(ar repositories.WeatherTemperatureRepository) WeatherTemperatureService {
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

func (as *weatherTemperatureService) LoadData(city string, year string) ([]repositories.DataItem, error) {
	data, err := as.WeatherTemperatureRepository.LoadData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
