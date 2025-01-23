package services_cards_vegetal

import (
	repositories_cards_vegetal "urbverde-api/repositories/cards/vegetal"
)

type VegetalInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_vegetal.InfoDataItem, error)
}

type vegetalInfoService struct {
	VegetalInfoRepository repositories_cards_vegetal.VegetalInfoRepository
}

func NewVegetalInfoService(ar repositories_cards_vegetal.VegetalInfoRepository) VegetalInfoService {
	return &vegetalInfoService{
		VegetalInfoRepository: ar,
	}
}

func (as *vegetalInfoService) LoadInfoData(city string) ([]repositories_cards_vegetal.InfoDataItem, error) {
	data, err := as.VegetalInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
