package main

import (
	"html/template"
	"net/http"
	"github.com/Nahom7wos/Airlines-Booking-System/airlines/http/handler"
	// import book hadler
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

rootHandler := handler.NewRootHandler(tmpl)


//admin
func Plane(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "admin.plane.layout", nil)
}
func PlaneCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	tmpl.ExecuteTemplate(w, "admin.plane.create.layout", nil)
}
func Destination(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "admin.destination.layout", nil)
}
func DestinationCreate(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "admin.destination.create.layout", nil)
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
	mux.HandleFunc("/", rootHandler.Index)
	mux.HandleFunc("/admin", rootHandler.Admin)
	mux.HandleFunc("/book", rootHandler.Book)
	mux.HandleFunc("/checkin", rootHandler.Checkin)
	mux.HandleFunc("/flights", rootHandler.Flights)
	mux.HandleFunc("/loyalty", rootHandler.Loyalty)

	//admin paths
	mux.HandleFunc("/admin/flight", admin)
	mux.HandleFunc("/admin/flight/create", admin)
	mux.HandleFunc("/admin/destination", Destination)
	mux.HandleFunc("/admin/destination/create", DestinationCreate)
	mux.HandleFunc("/admin/plane", Plane)
	mux.HandleFunc("/admin/plane/create", PlaneCreate)

	http.ListenAndServe(":8080", mux)

}
