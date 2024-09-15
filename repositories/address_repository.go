package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Suggestion struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type AddressRepository interface {
	SearchAddress(query string, filter string) ([]Suggestion, error)
}

type externalAddressRepository struct {
	cachedCities []IBGEResponse
}

func NewExternalAddressRepository() AddressRepository {
	return &externalAddressRepository{}
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

func (repo *externalAddressRepository) fetchCities() error {
	if len(repo.cachedCities) > 0 {
		return nil
	}

	url := "https://servicodados.ibge.gov.br/api/v1/localidades/municipios"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&repo.cachedCities); err != nil {
		return err
	}
	return nil
}

func (repo *externalAddressRepository) SearchAddress(query string, filter string) ([]Suggestion, error) {
	err := repo.fetchCities()
	if err != nil {
		return nil, err
	}

	if len(query) < 3 {
		return nil, nil
	}

	query = strings.ToLower(query)
	filteredSuggestions := []Suggestion{}

	states := []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal",
		"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
		"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
		"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins", "Brasil",
	}

	for _, state := range states {
		if strings.HasPrefix(strings.ToLower(state), query) && (filter == "all" || filter == "state") {
			filteredSuggestions = append(filteredSuggestions, Suggestion{
				Text: state,
				Type: "state",
			})
		}
	}

	for _, city := range repo.cachedCities {
		cityName := strings.ToLower(city.Nome)
		if strings.HasPrefix(cityName, query) && (filter == "all" || filter == "city") {
			cityText := fmt.Sprintf("%s - %s", city.Nome, city.Microrregiao.Mesorregiao.UF.Sigla)
			filteredSuggestions = append(filteredSuggestions, Suggestion{
				Text: cityText,
				Type: "city",
			})
		}
	}

	return filteredSuggestions, nil
}
