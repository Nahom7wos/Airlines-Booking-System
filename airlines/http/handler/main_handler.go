package handler

import (
	"html/template"
	"net/http"

	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

type MainHandler struct {
	tmpl           *template.Template
	destinationSrv flight.DestinationService
	flightSrv      flight.FlightService
}

// NewMainHandler initializes and returns new MainHandler
func NewMainHandler(T *template.Template, dSrv flight.DestinationService, fSrv flight.FlightService) *MainHandler {
	return &MainHandler{tmpl: T, destinationSrv: dSrv, flightSrv: fSrv}
}

// move these handlers to main_handler.go
func (mh *MainHandler) Index(w http.ResponseWriter, r *http.Request) {
	destinations, errs := mh.destinationSrv.Destinations()
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "index.layout", destinations)
}
func (mh *MainHandler) Admin(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
func (mh *MainHandler) Book(w http.ResponseWriter, r *http.Request) {
	flights, errs := mh.flightSrv.Flights()
	if errs != nil {
		panic(errs)
	}
	mh.tmpl.ExecuteTemplate(w, "book.layout", flights)
}
func (rh *MainHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	rh.tmpl.ExecuteTemplate(w, "checkin.layout", nil)
}
func (rh *MainHandler) Flights(w http.ResponseWriter, r *http.Request) {
	//check the flight on ticket
	//get flight by id
	rh.tmpl.ExecuteTemplate(w, "flights.layout", nil)
}
func (rh *MainHandler) Loyalty(w http.ResponseWriter, r *http.Request) {
	// get loyalty by user
	rh.tmpl.ExecuteTemplate(w, "loyalty.layout", nil)
}
