package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_KEY"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("API_SECRET"))

	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	alpacaClient := alpaca.NewClient(common.Credentials())
	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(*acct)
}
