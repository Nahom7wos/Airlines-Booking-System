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
func main() {

	//Database conneection
	dbconn, err := sql.Open("postgres", "")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/book", book)
	http.ListenAndServe(":8080", mux)
}
