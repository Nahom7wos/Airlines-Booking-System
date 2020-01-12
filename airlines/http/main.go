package main

import (
	"github.com/Nahom7wos/Airlines-Booking-System/airlines/http/handler"
	mrepim "github.com/Nahom7wos/Airlines-Booking-System/flight/repository"
	msrvim "github.com/Nahom7wos/Airlines-Booking-System/flight/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"net/http"
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

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:Postgre_1@localhost/airlinesdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	// createTables(dbconn)

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	destinationRepo := mrepim.NewDestinationGormRepo(dbconn)
	destinationServ := msrvim.NewDestinationService(destinationRepo)

	planeRepo := mrepim.NewDestinationGormRepo(dbconn)
	planeServ := msrvim.NewDestinationService(destinationRepo)

	mainHandler := handler.NewMainHandler(tmpl, destinationServ)
	destinationHandler := handler.NewDestinationHandler(tmpl, destinationServ)
	planeHandler := handler.NewPlaneHandler(tmpl, destinationServ)

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
	mux.HandleFunc("/admin/destination", destinationHandler.Destination)
	mux.HandleFunc("/admin/destination/create", destinationHandler.DestinationStore)
	mux.HandleFunc("/admin/plane", planeHandler.Plane)
	mux.HandleFunc("/admin/plane/create", planeHandler.PlaneStore)

	http.ListenAndServe(":8080", mux)

}
