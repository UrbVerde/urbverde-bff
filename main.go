package main

import (
	"urbverde-api/routes"
)

// @title UrbVerde BFF
// @version 1.0
// @description API para fornecer sugestões de endereço e outros serviços relacionados
// @host localhost:8080
// @BasePath /v1
func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
