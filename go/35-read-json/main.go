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
	database := database.NewSqlServerConnection(os.Getenv("SQL_SERVER_URL"))
	infrastructure.Run(database)
	log.Println("HEllo")
}
