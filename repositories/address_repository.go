package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AddressRepository interface {
	SearchAddress(query string) ([]string, error)
}

type IBGEResponse struct {
	Nome         string `json:"nome"`
	Microrregiao struct {
		Mesorregiao struct {
			UF struct {
				Sigla string `json:"sigla"`
			} `json:"UF"`
		} `json:"mesorregiao"`
	} `json:"microrregiao"`
}

type externalAddressRepository struct{}

func NewExternalAddressRepository() AddressRepository {
	return &externalAddressRepository{}
}

func (r *externalAddressRepository) SearchAddress(query string) ([]string, error) {

	url := fmt.Sprintf("https://servicodados.ibge.gov.br/api/v1/localidades/municipios?nome=%s", query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição. Código de status: %d", resp.StatusCode)
	}

	var cities []IBGEResponse
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	var cityNames []string
	qL := strings.ToLower(query)
	for _, city := range cities {
		cityName := strings.ToLower(city.Nome)
		if strings.HasPrefix(cityName, qL) {
			cityNames = append(cityNames, fmt.Sprintf("%s - %s", city.Nome, city.Microrregiao.Mesorregiao.UF.Sigla))
		}
	}

	return cityNames, nil
}
