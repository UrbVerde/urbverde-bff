package services_cards_vegetation

import (
	repositories_cards_vegetation "urbverde-api/repositories/cards/vegetation"
)

type VegetationInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_vegetation.InfoDataItem, error)
}

type vegetationInfoService struct {
	VegetationInfoRepository repositories_cards_vegetation.VegetationInfoRepository
}

func NewVegetationInfoService(ar repositories_cards_vegetation.VegetationInfoRepository) VegetationInfoService {
	return &vegetationInfoService{
		VegetationInfoRepository: ar,
	}
}

func (as *vegetationInfoService) LoadInfoData(city string) ([]repositories_cards_vegetation.InfoDataItem, error) {
	data, err := as.VegetationInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
