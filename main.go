package main

import (
	"urbverde-api/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Group("/api") // Prefixo de rota
	r.Run(":8080")
}
