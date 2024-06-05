package main

import (
	"log"
	"os"

	"github.com/jdsierrab3991/scripts/35-read-json/src/domain/database"
	"github.com/jdsierrab3991/scripts/35-read-json/src/infrastructure"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database, cleanup := database.New(os.Getenv("PG_URL"), os.Getenv("GCP_ACCOUNT"))
	if cleanup != nil {
		cleanup()
	}
	infrastructure.Run(database)
	log.Println("HEllo")
}
