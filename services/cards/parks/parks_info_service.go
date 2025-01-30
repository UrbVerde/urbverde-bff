package services_cards_parks

import (
	repositories_cards_parks "urbverde-api/repositories/cards/parks"
)

type ParksInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_parks.InfoDataItem, error)
}

type parksInfoService struct {
	ParksInfoRepository repositories_cards_parks.ParksInfoRepository
}

func NewParksInfoService(ar repositories_cards_parks.ParksInfoRepository) ParksInfoService {
	return &parksInfoService{
		ParksInfoRepository: ar,
	}
}

func (as *parksInfoService) LoadInfoData(city string) ([]repositories_cards_parks.InfoDataItem, error) {
	data, err := as.ParksInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
