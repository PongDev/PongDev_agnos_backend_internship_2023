package main

import (
	"log"

	"github.com/PongDev/PongDev_agnos_backend_internship_2023/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	app := app.NewApp()
	app.SetupRouter()
	app.SetupDatabase()
	app.Run()
}
