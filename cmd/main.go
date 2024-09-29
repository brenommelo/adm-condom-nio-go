package main

import (
	"github.com/brenommelo/adm-condominio-go/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	routes.SetupRouter()
}
