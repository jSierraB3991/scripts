package main

import (
	"log"
	"net/http"

	"github.com/OkabRitarou/08-mysql-gorm/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
