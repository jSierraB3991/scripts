package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// sqlServerUrl="sqlserver://sa:FEVRIPS1*@localhost:1433?database=rips-local"
func NewSqlServerConnection(sqlServerUrl string) *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(sqlServerUrl), &gorm.Config{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db

}
