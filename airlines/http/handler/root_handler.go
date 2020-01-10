package handler

import (
	"html/template"
	"net/http"
)

type RootHandler struct {
	tmpl *template.Template
}

// NewMenuHandler initializes and returns new MenuHandler
func NewRootHandler(T *template.Template) *RootHandler {
	return &RootHandler{tmpl: T}
}

// move these handlers to root_handler.go
func (rh *RootHandler) Index(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func (rh *RootHandler) Admin(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
func (rh *RootHandler) Book(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "book.layout", nil)
}
func (rh *RootHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}
func (rh *RootHandler) Flights(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "flights.layout", nil)
}
func (rh *RootHandler) Loyalty(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "loyalty.layout", nil)
}
