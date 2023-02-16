package main

import (
	"github.com/joho/godotenv"
	"log"
	"webscraping-league-of-legends-v2/cblol/infra/http/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	router := routes.AddRoutesWithAuthentication()
	router.Run(":8989")
}
