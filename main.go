package main

import (
	"urbverde-api/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
