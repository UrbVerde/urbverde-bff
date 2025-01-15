package services_cards_vegetal

import (
	cards_shared "urbverde-api/repositories/cards"
	repositories_cards_vegetal "urbverde-api/repositories/cards/vegetal"
)

type VegetalCoverService interface {
	cards_shared.RepositoryBase
	LoadCoverData(city string, year string) ([]repositories_cards_vegetal.CoverDataItem, error)
}

type vegetalCoverService struct {
	VegetalCoverRepository repositories_cards_vegetal.VegetalCoverRepository
}

func NewVegetalCoverService(ar repositories_cards_vegetal.VegetalCoverRepository) VegetalCoverService {
	return &vegetalCoverService{
		VegetalCoverRepository: ar,
	}
}

func (as *vegetalCoverService) LoadYears(city string) ([]int, error) {
	data, err := as.VegetalCoverRepository.LoadYears(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *vegetalCoverService) LoadCoverData(city string, year string) ([]repositories_cards_vegetal.CoverDataItem, error) {
	data, err := as.VegetalCoverRepository.LoadCoverData(city, year)
	if err != nil {
		return nil, err
	}

	return data, nil
}
