package main

import (
	"github.com/Nahom7wos/Airlines-Booking-System/airlines/http/handler"
	"html/template"
	"net/http"
	// import book hadler
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

//admin
// func Plane(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "admin.plane.layout", nil)
// }
// func PlaneCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {

// 	}
// 	tmpl.ExecuteTemplate(w, "admin.plane.create.layout", nil)
// }
// func Destination(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "admin.destination.layout", nil)
// }
// func DestinationCreate(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "admin.destination.create.layout", nil)
// }

//create dbConn
//pass to repo
//create a repo
//pass to service

//pass tmpl and service to
//adminHandler
//menuHandler

func main() {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	mainHandler := handler.NewMainHandler(tmpl)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", mainHandler.Index)
	mux.HandleFunc("/admin", mainHandler.Admin)
	mux.HandleFunc("/book", mainHandler.Book)
	mux.HandleFunc("/checkin", mainHandler.Checkin)
	mux.HandleFunc("/flights", mainHandler.Flights)
	mux.HandleFunc("/loyalty", mainHandler.Loyalty)

	//admin paths
	// mux.HandleFunc("/admin/flight", MainHandler.Admin)
	// mux.HandleFunc("/admin/flight/create", MainHandler.Admin)
	// mux.HandleFunc("/admin/destination", Destination)
	// mux.HandleFunc("/admin/destination/create", DestinationCreate)
	// mux.HandleFunc("/admin/plane", Plane)
	// mux.HandleFunc("/admin/plane/create", PlaneCreate)

	http.ListenAndServe(":8080", mux)

}
