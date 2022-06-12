package routes

import (
	"github.com/OkabRitarou/08-mysql-gorm/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoute = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
