package services_cards_vegetation

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetation "urbverde-api/repositories/cards/vegetation"
)

type VegetationCoverService interface {
	cards_shared.RepositoryBase
	LoadCoverData(city string, year string) ([]repositories_cards_vegetation.CoverDataItem, error)
}

type vegetationCoverService struct {
	VegetationCoverRepository repositories_cards_vegetation.VegetationCoverRepository
}

func NewVegetationCoverService(ar repositories_cards_vegetation.VegetationCoverRepository) VegetationCoverService {
	return &vegetationCoverService{
		VegetationCoverRepository: ar,
	}
}

func (as *vegetationCoverService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetationCoverRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetationCoverService) LoadCoverData(city string, year string) ([]repositories_cards_vegetation.CoverDataItem, error) {
	data, err := as.VegetationCoverRepository.LoadCoverData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
