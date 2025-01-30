package services_cards_vegetation

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetation "urbverde-api/repositories/cards/vegetation"
)

type VegetationInequalityService interface {
	cards_shared.RepositoryBase
	LoadInequalityData(city string, year string) ([]repositories_cards_vegetation.InequalityDataItem, error)
}

type vegetationInequalityService struct {
	VegetationInequalityRepository repositories_cards_vegetation.VegetationInequalityRepository
}

func NewVegetationInequalityService(ar repositories_cards_vegetation.VegetationInequalityRepository) VegetationInequalityService {
	return &vegetationInequalityService{
		VegetationInequalityRepository: ar,
	}
}

func (as *vegetationInequalityService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetationInequalityRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetationInequalityService) LoadInequalityData(city string, year string) ([]repositories_cards_vegetation.InequalityDataItem, error) {
	data, err := as.VegetationInequalityRepository.LoadInequalityData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
