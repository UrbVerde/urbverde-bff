// urbverde-bff/main.go
package main

import (
	"urbverde-api/routes"
)

// @title UrbVerde BFF
// @version 1.0
// @description API para fornecer sugestões de endereço e outros serviços relacionados
// @contact.name API Support
// @contact.url http://www.urbverde.com.br/
// @contact.email comunica.urbverde@usp.br
// @license.name ???
// @license.url ???
// @host localhost:8080
// @BasePath /v1
func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
