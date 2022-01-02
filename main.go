package main

import (
	"log"

	"github.com/informeai/gogql/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error: %v\n", err.Error())
	}
	router := routes.NewRouter()
	if err := router.Start(); err != nil {
		log.Fatalf("error: %v\n", err.Error())
	}
}
