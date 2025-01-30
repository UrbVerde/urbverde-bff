// urbverde-bff/routes/cards/cards_routes.go
package cards

import (
	controllers_cards_temperature "urbverde-api/controllers/cards/temperature"
	repositories_cards_temperature "urbverde-api/repositories/cards/temperature"
	services_cards_temperature "urbverde-api/services/cards/temperature"

	controllers_cards_vegetation "urbverde-api/controllers/cards/vegetation"
	repositories_cards_vegetation "urbverde-api/repositories/cards/vegetation"
	services_cards_vegetation "urbverde-api/services/cards/vegetation"

	controllers_cards_parks "urbverde-api/controllers/cards/parks"
	repositories_cards_parks "urbverde-api/repositories/cards/parks"
	services_cards_parks "urbverde-api/services/cards/parks"

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

func SetupCardsRoutes(rg *gin.RouterGroup) {
	SetupTemperatureRoutes(rg)
	SetupVegetationRoutes(rg)
	SetupParksRoutes(rg)
}

func SetupTemperatureRoutes(rg *gin.RouterGroup) {
	setupWeatherRoutes(rg)
	setupHeatRoutes(rg)
	setupTemperatureRankingRoutes(rg)
	setupTemperatureInfoRoutes(rg)
}

func SetupVegetationRoutes(rg *gin.RouterGroup) {
	setupCoverRoutes(rg)
	setupVegetationInequalityRoutes(rg)
	setupVegetationRankingRoutes(rg)
	setupVegetationInfoRoutes(rg)
}

func SetupParksRoutes(rg *gin.RouterGroup) {
	setupParksSquareRoutes(rg)
	setupParksInequalityRoutes(rg)
	setupParksRankingRoutes(rg)
	setupParksInfoRoutes(rg)
}

// @Summary Retorna dados de temperatura
// @Description Retorna os dados de temperatura para o município e ano fornecidos
// @Tags cards/temperature
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/temperature/weather [get]
func setupWeatherRoutes(rg *gin.RouterGroup) {
	weatherRepo := repositories_cards_temperature.NewExternalTemperatureWeatherRepository()
	weatherService := services_cards_temperature.NewTemperatureWeatherService(weatherRepo)
	weatherController := controllers_cards_temperature.NewTemperatureWeatherController(weatherService)

	rg.GET("/cards/temperature/weather", weatherController.LoadWeatherData)
}

// @Summary Retorna dados de calor extremo
// @Description Retorna os dados de calor extremo para o município e ano fornecidos
// @Tags cards/temperature
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/temperature/heat [get]
func setupHeatRoutes(rg *gin.RouterGroup) {
	heatRepo := repositories_cards_temperature.NewExternalTemperatureHeatRepository()
	heatService := services_cards_temperature.NewTemperatureHeatService(heatRepo)
	heatController := controllers_cards_temperature.NewTemperatureHeatController(heatService)

	rg.GET("/cards/temperature/heat", heatController.LoadHeatData)
}

// @Summary Retorna dados de ranking de clima
// @Description Retorna os dados de ranking em clima para o município e ano fornecidos
// @Tags cards/temperature
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/temperature/ranking [get]
func setupTemperatureRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_temperature.NewExternalTemperatureRankingRepository()
	rankService := services_cards_temperature.NewTemperatureRankingService(rankRepo)
	rankController := controllers_cards_temperature.NewTemperatureRankingController(rankService)

	rg.GET("/cards/temperature/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais
// @Description Retorna dados adicionais para o município fornecido
// @Tags cards/temperature
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/temperature/info [get]
func setupTemperatureInfoRoutes(rg *gin.RouterGroup) {
	infoRepo := repositories_cards_temperature.NewExternalTemperatureInfoRepository()
	infoService := services_cards_temperature.NewTemperatureInfoService(infoRepo)
	infoController := controllers_cards_temperature.NewTemperatureInfoController(infoService)

	rg.GET("/cards/temperature/info", infoController.LoadInfoData)
}

// @Summary Retorna dados relacionados à cobertura vegetal
// @Description Retorna dados relacionados à cobertura vegetal para o município fornecido
// @Tags cards/vegetation
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetation/cover [get]
func setupCoverRoutes(rg *gin.RouterGroup) {
	coverRepo := repositories_cards_vegetation.NewExternalVegetationCoverRepository()
	coverService := services_cards_vegetation.NewVegetationCoverService(coverRepo)
	coverController := controllers_cards_vegetation.NewVegetationCoverController(coverService)

	rg.GET("/cards/vegetation/cover", coverController.LoadCoverData)
}

// @Summary Retorna dados relacionados à desigualdade ambiental e a vegetação
// @Description Retorna dados relacionados à desigualdade ambiental e a vegetação para o município fornecido
// @Tags cards/vegetation
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetation/inequality [get]
func setupVegetationInequalityRoutes(rg *gin.RouterGroup) {
	inequalityRepo := repositories_cards_vegetation.NewExternalVegetationInequalityRepository()
	inequalityService := services_cards_vegetation.NewVegetationInequalityService(inequalityRepo)
	inequalityController := controllers_cards_vegetation.NewVegetationInequalityController(inequalityService)

	rg.GET("/cards/vegetation/inequality", inequalityController.LoadInequalityData)
}

// @Summary Retorna dados de ranking
// @Description Retorna dados para a construção do ranking de desigualdade ambiental e a vegetação
// @Tags cards/vegetation
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetation/ranking [get]
func setupVegetationRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_vegetation.NewExternalVegetationRankingRepository()
	rankService := services_cards_vegetation.NewVegetationRankingService(rankRepo)
	rankController := controllers_cards_vegetation.NewVegetationRankingController(rankService)

	rg.GET("/cards/vegetation/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais para a vegetação
// @Description Retorna dados adicionais para a camada
// @Tags cards/vegetation
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetation/info [get]
func setupVegetationInfoRoutes(rg *gin.RouterGroup) {
	infoRepo := repositories_cards_vegetation.NewExternalVegetationInfoRepository()
	infoService := services_cards_vegetation.NewVegetationInfoService(infoRepo)
	infoController := controllers_cards_vegetation.NewVegetationInfoController(infoService)

	rg.GET("/cards/vegetation/info", infoController.LoadInfoData)
}

// @Summary Retorna dados dos parques e praças
// @Description Retorna dados de parques e praças para a camada
// @Tags cards/parks
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/parks/square [get]
func setupParksSquareRoutes(rg *gin.RouterGroup) {
	parksRepo := repositories_cards_parks.NewExternalParksSquareRepository()
	parksService := services_cards_parks.NewParksSquareService(parksRepo)
	parksController := controllers_cards_parks.NewParksSquareController(parksService)

	rg.GET("/cards/parks/square", parksController.LoadSquareData)
}

// @Summary Retorna dados sobre desigualdade
// @Description Retorna dados de desigualdade para a camada
// @Tags cards/parks
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/parks/inequality [get]
func setupParksInequalityRoutes(rg *gin.RouterGroup) {
	inequalityRepo := repositories_cards_parks.NewExternalParksInequalityRepository()
	inequalityService := services_cards_parks.NewParksInequalityService(inequalityRepo)
	inequalityController := controllers_cards_parks.NewParksInequalityController(inequalityService)

	rg.GET("/cards/parks/inequality", inequalityController.LoadInequalityData)
}

// @Summary Retorna dados de ranking
// @Description Retorna dados para a construção do ranking de praças e parques
// @Tags cards/parks
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/parks/ranking [get]
func setupParksRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_parks.NewExternalParksRankingRepository()
	rankService := services_cards_parks.NewParksRankingService(rankRepo)
	rankController := controllers_cards_parks.NewParksRankingController(rankService)

	rg.GET("/cards/parks/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais para a camada de praças e parques
// @Description Retorna dados adicionais para a camada
// @Tags cards/parks
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/parks/info [get]
func setupParksInfoRoutes(rg *gin.RouterGroup) {
	parksRepo := repositories_cards_parks.NewExternalParksInfoRepository()
	parksService := services_cards_parks.NewParksInfoService(parksRepo)
	parksController := controllers_cards_parks.NewParksInfoController(parksService)

	rg.GET("/cards/parks/info", parksController.LoadInfoData)
}
