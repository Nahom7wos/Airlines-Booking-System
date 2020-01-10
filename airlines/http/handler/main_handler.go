package handler

import (
	"html/template"
	"net/http"
)

type MainHandler struct {
	tmpl *template.Template
}

// NewMainHandler initializes and returns new MainHandler
func NewMainHandler(T *template.Template) *MainHandler {
	return &MainHandler{tmpl: T}
}

// move these handlers to root_handler.go
func (rh *MainHandler) Index(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func (rh *MainHandler) Admin(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
func (rh *MainHandler) Book(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "book.layout", nil)
}
func (rh *MainHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}
func (rh *MainHandler) Flights(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "flights.layout", nil)
}
func (rh *MainHandler) Loyalty(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "loyalty.layout", nil)
}
