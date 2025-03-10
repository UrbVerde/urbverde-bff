// urbverde-bff/routes/cards/cards_routes.go
package cards

import (
	controllers_cards_weather "urbverde-api/controllers/cards/weather"
	repositories_cards_weather "urbverde-api/repositories/cards/weather"
	services_cards_weather "urbverde-api/services/cards/weather"

	controllers_cards_vegetal "urbverde-api/controllers/cards/vegetal"
	repositories_cards_vegetal "urbverde-api/repositories/cards/vegetal"
	services_cards_vegetal "urbverde-api/services/cards/vegetal"

	controllers_cards_square "urbverde-api/controllers/cards/square"
	repositories_cards_square "urbverde-api/repositories/cards/square"
	services_cards_square "urbverde-api/services/cards/square"

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
	SetupWeatherRoutes(rg)
	SetupVegetalRoutes(rg)
	SetupSquareRoutes(rg)
}

func SetupWeatherRoutes(rg *gin.RouterGroup) {
	setupTemperatureRoutes(rg)
	setupHeatRoutes(rg)
	setupWeatherRankingRoutes(rg)
	setupWeatherInfoRoutes(rg)
}

func SetupVegetalRoutes(rg *gin.RouterGroup) {
	setupCoverRoutes(rg)
	setupVegetalInequalityRoutes(rg)
	setupVegetalRankingRoutes(rg)
	setupVegetalInfoRoutes(rg)
}

func SetupSquareRoutes(rg *gin.RouterGroup) {
	setupSquareParksRoutes(rg)
	setupSquareInequalityRoutes(rg)
	setupSquareRankingRoutes(rg)
	setupSquareInfoRoutes(rg)
}

// @Summary Retorna dados de temperatura
// @Description Retorna os dados de temperatura para o município e ano fornecidos
// @Tags cards/weather
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
// @Tags cards/weather
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
// @Tags cards/weather
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/weather/ranking [get]
func setupWeatherRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_weather.NewExternalWeatherRankingRepository()
	rankService := services_cards_weather.NewWeatherRankingService(rankRepo)
	rankController := controllers_cards_weather.NewWeatherRankingController(rankService)

	rg.GET("/cards/weather/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais
// @Description Retorna dados adicionais para o município fornecido
// @Tags cards/weather
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/weather/info [get]
func setupWeatherInfoRoutes(rg *gin.RouterGroup) {
	infoRepo := repositories_cards_weather.NewExternalWeatherInfoRepository()
	infoService := services_cards_weather.NewWeatherInfoService(infoRepo)
	infoController := controllers_cards_weather.NewWeatherInfoController(infoService)

	rg.GET("/cards/weather/info", infoController.LoadInfoData)
}

// @Summary Retorna dados relacionados à cobertura vegetal
// @Description Retorna dados relacionados à cobertura vegetal para o município fornecido
// @Tags cards/vegetal
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetal/cover [get]
func setupCoverRoutes(rg *gin.RouterGroup) {
	coverRepo := repositories_cards_vegetal.NewExternalVegetalCoverRepository()
	coverService := services_cards_vegetal.NewVegetalCoverService(coverRepo)
	coverController := controllers_cards_vegetal.NewVegetalCoverController(coverService)

	rg.GET("/cards/vegetal/cover", coverController.LoadCoverData)
}

// @Summary Retorna dados relacionados à desigualdade ambiental e a vegetação
// @Description Retorna dados relacionados à desigualdade ambiental e a vegetação para o município fornecido
// @Tags cards/vegetal
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetal/inequality [get]
func setupVegetalInequalityRoutes(rg *gin.RouterGroup) {
	inequalityRepo := repositories_cards_vegetal.NewExternalVegetalInequalityRepository()
	inequalityService := services_cards_vegetal.NewVegetalInequalityService(inequalityRepo)
	inequalityController := controllers_cards_vegetal.NewVegetalInequalityController(inequalityService)

	rg.GET("/cards/vegetal/inequality", inequalityController.LoadInequalityData)
}

// @Summary Retorna dados de ranking
// @Description Retorna dados para a construção do ranking de desigualdade ambiental e a vegetação
// @Tags cards/vegetal
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetal/ranking [get]
func setupVegetalRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_vegetal.NewExternalVegetalRankingRepository()
	rankService := services_cards_vegetal.NewVegetalRankingService(rankRepo)
	rankController := controllers_cards_vegetal.NewVegetalRankingController(rankService)

	rg.GET("/cards/vegetal/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais para a vegetação
// @Description Retorna dados adicionais para a camada
// @Tags cards/vegetal
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/vegetal/info [get]
func setupVegetalInfoRoutes(rg *gin.RouterGroup) {
	infoRepo := repositories_cards_vegetal.NewExternalVegetalInfoRepository()
	infoService := services_cards_vegetal.NewVegetalInfoService(infoRepo)
	infoController := controllers_cards_vegetal.NewVegetalInfoController(infoService)

	rg.GET("/cards/vegetal/info", infoController.LoadInfoData)
}

// @Summary Retorna dados dos parques e praças
// @Description Retorna dados de parques e praças para a camada
// @Tags cards/square
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/square/parks [get]
func setupSquareParksRoutes(rg *gin.RouterGroup) {
	parksRepo := repositories_cards_square.NewExternalSquareParksRepository()
	parksService := services_cards_square.NewSquareParksService(parksRepo)
	parksController := controllers_cards_square.NewSquareParksController(parksService)

	rg.GET("/cards/square/parks", parksController.LoadParksData)
}

// @Summary Retorna dados sobre desigualdade
// @Description Retorna dados de desigualdade para a camada
// @Tags cards/square
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/square/inequality [get]
func setupSquareInequalityRoutes(rg *gin.RouterGroup) {
	inequalityRepo := repositories_cards_square.NewExternalSquareInequalityRepository()
	inequalityService := services_cards_square.NewSquareInequalityService(inequalityRepo)
	inequalityController := controllers_cards_square.NewSquareInequalityController(inequalityService)

	rg.GET("/cards/square/inequality", inequalityController.LoadInequalityData)
}

// @Summary Retorna dados de ranking
// @Description Retorna dados para a construção do ranking de praças e parques
// @Tags cards/square
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Param year query string false "Ano dos dados"
// @Success 200 {object} []RankingData
// @Failure 400 {object} ErrorResponse
// @Router /cards/square/ranking [get]
func setupSquareRankingRoutes(rg *gin.RouterGroup) {
	rankRepo := repositories_cards_square.NewExternalSquareRankingRepository()
	rankService := services_cards_square.NewSquareRankingService(rankRepo)
	rankController := controllers_cards_square.NewSquareRankingController(rankService)

	rg.GET("/cards/square/ranking", rankController.LoadRankingData)
}

// @Summary Retorna dados adicionais para a camada de praças e parques
// @Description Retorna dados adicionais para a camada
// @Tags cards/square
// @Accept json
// @Produce json
// @Param city query string true "Código de município"
// @Success 200 {object} []CardsDataItem
// @Failure 400 {object} ErrorResponse
// @Router /cards/square/info [get]
func setupSquareInfoRoutes(rg *gin.RouterGroup) {
	squareRepo := repositories_cards_vegetal.NewExternalVegetalInfoRepository()
	squareService := services_cards_vegetal.NewVegetalInfoService(squareRepo)
	squareController := controllers_cards_vegetal.NewVegetalInfoController(squareService)

	rg.GET("/cards/square/info", squareController.LoadInfoData)
}
