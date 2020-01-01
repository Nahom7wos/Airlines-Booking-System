package main

import (
	"html/template"
	"net/http"

	// import book hadler
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

// move these handlers to book_handler.go
func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func book(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "book.layout", nil)
}
func checkin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}

//create dbConn
//pass to repo
//create a repo
//pass to service

//pass tmpl and service to
//adminHandler
//menuHandler

func main() {

	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/book", book)
	mux.HandleFunc("/checkin", book)
	http.ListenAndServe(":8080", mux)

	//admin paths
}
