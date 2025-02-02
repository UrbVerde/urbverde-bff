package services_cards_square

import (
	repositories_cards_square "urbverde-api/repositories/cards/square"
)

type SquareInfoService interface {
	LoadInfoData(city string) ([]repositories_cards_square.InfoDataItem, error)
}

type squareInfoService struct {
	SquareInfoRepository repositories_cards_square.SquareInfoRepository
}

func NewSquareInfoService(ar repositories_cards_square.SquareInfoRepository) SquareInfoService {
	return &squareInfoService{
		SquareInfoRepository: ar,
	}
}

func (as *squareInfoService) LoadInfoData(city string) ([]repositories_cards_square.InfoDataItem, error) {
	data, err := as.SquareInfoRepository.LoadInfoData(city)
	if err != nil {
		return nil, err
	}

	return data, nil
}
