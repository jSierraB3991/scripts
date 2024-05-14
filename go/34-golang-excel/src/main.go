package main

import (
	"context"
	"log"
	"os"

	"github.com/jdsierrab3991/scripts/34-golang-excel/domain/database"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure"
	"github.com/jdsierrab3991/scripts/34-golang-excel/infrastructure/repository"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := database.New(os.Getenv("POSTGRE_URL"))
	repo := repository.InitiateRepo(db, context.Background())
	database.AutoMigrate(repo)

	read := infrastructure.NewReadExcelData(repo)
	homeData := os.Getenv("HOME_DATA")
	documents := readDir(homeData)
	err := read.Run(homeData, documents)
	if err != nil {
		log.Fatal(err)
	}
}

func readDir(homeData string) []string {
	files, err := os.ReadDir(homeData)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result
}
