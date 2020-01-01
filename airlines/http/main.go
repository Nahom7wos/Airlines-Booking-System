package main

import (
	"database/sql"
	"html/template"
	"net/http"
	// import menu hadler
)

// connect database - previleges?

// handlefunc, start server
var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func book(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "book.layout", nil)
}
func checkin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}
func main() {

	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/book", book)
	mux.HandleFunc("/checkin", book)
	http.ListenAndServe(":8080", mux)

}
