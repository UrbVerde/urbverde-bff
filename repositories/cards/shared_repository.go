// urbverde-bff/repositories/cards/shared_repository.go
package cards_shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

const (
	WfsService      = "WFS"
	WfsVersion      = "1.0.0"
	WfsRequest      = "GetFeature"
	TypeName        = "urbverde:"
	CqlFilterPrefix = "CQL_FILTER=cd_mun="
)

type RepositoryBase interface {
	LoadYears(city string) ([]int, error)
}

// Dynamic properties
type Properties interface{}

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Properties Properties `json:"properties"`
}

func FetchFromURL(url string) (*FeatureCollection, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição. Código de status: %d", resp.StatusCode)
	}

	var data FeatureCollection
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	return &data, nil
}

func MapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func LoadYears(url string, processProperties func(map[string]interface{}) (int, error)) ([]int, error) {
	data, err := FetchFromURL(url)
	if err != nil {
		return nil, err
	}

	yearsMap := make(map[int]bool)

	for _, feature := range data.Features {
		props, ok := feature.Properties.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected property type")
		}

		year, err := processProperties(props)
		if err != nil {
			return nil, err
		}
		yearsMap[year] = true
	}

	var years []int
	for year := range yearsMap {
		years = append(years, year)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	return years, nil
}
