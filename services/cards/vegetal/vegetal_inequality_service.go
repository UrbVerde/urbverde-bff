package services_cards_vegetal

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetal "urbverde-api/repositories/cards/vegetal"
)

type VegetalInequalityService interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]repositories_cards_vegetal.InequalityDataItem, error)
}

type vegetalInequalityService struct {
	VegetalInequalityRepository repositories_cards_vegetal.VegetalInequalityRepository
}

func NewVegetalInequalityService(ar repositories_cards_vegetal.VegetalInequalityRepository) VegetalInequalityService {
	return &vegetalInequalityService{
		VegetalInequalityRepository: ar,
	}
}

func (as *vegetalInequalityService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetalInequalityRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetalInequalityService) LoadInequalityData(city string, year string) ([]repositories_cards_vegetal.InequalityDataItem, error) {
	data, err := as.VegetalInequalityRepository.LoadInequalityData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
