package main

import (
	"html/template"
	"net/http"
	// import menu hadler
)

// connect database - previleges?

// handlefunc, start server, (can't declare := outsite func)
var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func main() {
	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)
}
