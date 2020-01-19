package handler

import (
	"html/template"
	"net/http"

	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

type FlightHandler struct {
	tmpl      *template.Template
	flightSrv flight.FlightService
}

// NewFlightHandler initializes and returns new FlightHandler
func NewFlightHandler(T *template.Template, pSrv flight.FlightService) *FlightHandler {
	return &FlightHandler{tmpl: T, flightSrv: pSrv}
}

// Flight displays all the flights in the database
func (dh *FlightHandler) Flight(w http.ResponseWriter, r *http.Request) {

}

// FlightStore creates new flight in the database
func (dh *FlightHandler) FlightStore(w http.ResponseWriter, r *http.Request) {

}
