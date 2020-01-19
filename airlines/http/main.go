package main

import (
	"html/template"
	"net/http"

	"github.com/Nahom7wos/Airlines-Booking-System/airlines/http/handler"
	frepim "github.com/Nahom7wos/Airlines-Booking-System/flight/repository"
	fsrvim "github.com/Nahom7wos/Airlines-Booking-System/flight/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Destination{}, &entity.Plane{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	//createTables(dbconn)

	dbconn, err := gorm.Open("postgres", "postgres://postgres:Postgre_1@localhost/airlinesdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	flightRepo := frepim.NewFlightGormRepo(dbconn)
	flightServ := fsrvim.NewFlightService(flightRepo)
	
	destinationRepo := frepim.NewDestinationGormRepo(dbconn)
	destinationServ := fsrvim.NewDestinationService(destinationRepo)

	planeRepo := frepim.NewDestinationGormRepo(dbconn)
	planeServ := fsrvim.NewDestinationService(destinationRepo)

	mainHandler := handler.NewMainHandler(tmpl, destinationServ, flightServ)
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
