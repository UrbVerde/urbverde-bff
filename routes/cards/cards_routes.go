package cards

import (
	controllers_cards_weather "urbverde-api/controllers/cards/weather"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
	services_cards_weather "urbverde-api/services/cards/weather"

	"github.com/gin-gonic/gin"
)

type CardsDataItem struct {
	Title    string  `json:"title" example:"Nível de ilha de calor"`
	Subtitle *string `json:"subtitle,omitempty" example:"Abaixo da média nacional de 0"`
	Value    string  `json:"value" example:"25°C"`
}

type RankingDataItem struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
	Of     int    `json:"of"`
}

type RankingData struct {
	Title    string            `json:"title" example:"Municipios do Estado"`
	Subtitle *string           `json:"subtitle,omitempty" example:"Posição do seu município entre os 645 do Estado de São Paulo"`
	Items    []RankingDataItem `json:"items"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Erro ao processar a solicitação"`
	Code    int    `json:"code" example:"400"`
}

// SetupCardsRoutes
func SetupCardsRoutes(rg *gin.RouterGroup) {
	// Weather
	setupTemperatureRoutes(rg)
	setupHeatRoutes(rg)
	setupRankingRoutes(rg)
	setupInfoRoutes(rg)
	setupRankingRoutes(rg)
}

// @Summary Retorna dados de temperatura
// @Description Retorna os dados de temperatura para o município e ano fornecidos
// @Tags cards
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/weather/temperature [get]
func setupTemperatureRoutes(rg *gin.RouterGroup) {
	tempeRepo := repositories_cards_weather.NewExternalWeatherTemperatureRepository()
	tempeService := services_cards_weather.NewWeatherTemperatureService(tempeRepo)
	tempeController := controllers_cards_weather.NewWeatherTemperatureController(tempeService)

	rg.GET("/cards/weather/temperature", tempeController.LoadTemperatureData)
}

// @Summary Retorna dados de calor extremo
// @Description Retorna os dados de calor extremo para o município e ano fornecidos
// @Tags cards
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/weather/heat [get]
func setupHeatRoutes(rg *gin.RouterGroup) {
	heatRepo := repositories_cards_weather.NewExternalWeatherHeatRepository()
	heatService := services_cards_weather.NewWeatherHeatService(heatRepo)
	heatController := controllers_cards_weather.NewWeatherHeatController(heatService)

	rg.GET("/cards/weather/heat", heatController.LoadHeatData)
}

// @Summary Retorna dados de ranking de clima
// @Description Retorna os dados de ranking em clima para o município e ano fornecidos
// @Tags cards
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/weather/ranking [get]
func setupRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_weather.NewExternalWeatherRankingRepository()
	rankService := services_cards_weather.NewWeatherRankingService(rankRepo)
	rankController := controllers_cards_weather.NewWeatherRankingController(rankService)

	rg.GET("/cards/weather/ranking", rankController.LoadRankingData)
}

func setupInfoRoutes(rg *gin.RouterGroup) {
	infoRepo := repositories_cards_weather.NewExternalWeatherInfoRepository()
	infoService := services_cards_weather.NewWeatherInfoService(infoRepo)
	infoController := controllers_cards_weather.NewWeatherInfoController(infoService)

	rg.GET("/cards/weather/info", infoController.LoadInfoData)
}
