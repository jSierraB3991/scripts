package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)

func DatabaseConnect() {
	var err error
	DbConn, err = gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed connect to database")
	}
	fmt.Println("Connection Openned to database")
}
