package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func dbConnection() *sql.DB {
	driver := "mysql"
	user := "root"
	passwd := "chrrot"
	dbName := "go_mysql"
	ipDb := "127.0.0.1"

	conecction, err := sql.Open(driver, user+":"+passwd+"@tcp("+ipDb+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return conecction
}

func Start(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "start", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/create", Create)
	log.Println("Server running!!!")
	http.ListenAndServe(":8080", nil)
}
