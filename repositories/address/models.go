// urbverde-bff/repositories/address/models.go
package repositories_address

// CityResponse represents our formatted response for suggestions
// @Description City suggestion response model
type CityResponse struct {
	DisplayName string `json:"display_name" example:"São Paulo - SP"` // What user sees: "City Name - ST"
	CdMun       int    `json:"cd_mun" example:"3550308"`              // City ID for internal use
	Type        string `json:"type" example:"city"`                   // Type of location
}

// Location represents a detailed location response
// @Description Detailed location data model
type Location struct {
	Name        string                `json:"name" example:"São Paulo"`
	DisplayName string                `json:"display_name" example:"São Paulo - SP"`
	Type        string                `json:"type" example:"city"`
	Code        string                `json:"code" example:"3550308"`
	State       string                `json:"state,omitempty" example:"SP"`
	StateName   string                `json:"state_name,omitempty" example:"São Paulo"`
	Country     string                `json:"country" example:"Brasil"`
	CountryCode string                `json:"country_code" example:"BR"`
	CenterOpts  map[string]CenterOpts `json:"center_options"`
}

// CenterOpts represents location center and zoom options
type CenterOpts struct {
	Lat  float64   `json:"lat" example:"-23.5505"`
	Lng  float64   `json:"lng" example:"-46.6333"`
	Zoom int       `json:"zoom" example:"10"`
	BBox []float64 `json:"bbox" example:"-46.8264,-24.0082,-46.3652,-23.3566"`
}

// LocationResponse represents the JSON structure for location data files
type LocationResponse struct {
	Features map[string]Location `json:"features"`
}

// IBGEResponse represents the raw response from IBGE API
type IBGEResponse struct {
	ID           int    `json:"id"`   // City ID (cd_mun)
	Nome         string `json:"nome"` // City name
	Microrregiao struct {
		Mesorregiao struct {
			UF struct {
				Sigla string `json:"sigla"` // State abbreviation
			} `json:"UF"`
		} `json:"mesorregiao"`
	} `json:"microrregiao"`
}
