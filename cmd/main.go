package main

import (
	"github.com/brenommelo/adm-condominio-go/internal/config"
	"github.com/brenommelo/adm-condominio-go/internal/routes"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.ConnectToDB()
	routes.SetupRouter()
}

func main() {

}
