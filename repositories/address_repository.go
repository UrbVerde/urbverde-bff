// urbverde-bff/repositories/address_repository.go
package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
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

type externalAddressRepository struct {
	apiURL string
}

func NewExternalAddressRepository() AddressRepository {
	_ = godotenv.Load()

	apiURL := os.Getenv("IBGE_API_URL")
	if apiURL == "" {
		panic("A variável de ambiente IBGE_API_URL não está definida")
	}

	return &externalAddressRepository{
		apiURL: apiURL,
	}
}

func (r *externalAddressRepository) SearchAddress(query string) ([]string, error) {
	url := r.apiURL

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
