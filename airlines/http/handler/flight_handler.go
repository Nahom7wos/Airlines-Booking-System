package handler

import (
	"html/template"
	"net/http"
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

type FlightHandler struct {
	tmpl      *template.Template
	destinationSrv flight.DestinationService
	planeSrv flight.PlaneService
	flightSrv flight.FlightService
}

// NewFlightHandler initializes and returns new FlightHandler
func NewFlightHandler(T *template.Template, pSrv flight.PlaneService, dSrv flight.DestinationService, fSrv flight.FlightService) *FlightHandler {
	return &FlightHandler{tmpl: T, planeSrv: pSrv, destinationSrv: dSrv, flightSrv: fSrv}
}

// Flight displays all the flights in the database
func (fh *FlightHandler) Flight(w http.ResponseWriter, r *http.Request) {
	
}

// FlightStore creates new flight in the database
func (fh *FlightHandler) FlightStore(w http.ResponseWriter, r *http.Request) {
	destinations, errs := fh.destinationSrv.Destinations()
	if errs != nil {
		panic(errs)
	}
	planes, errs2 := fh.planeSrv.Planes()
	if errs2 != nil {
		panic(errs2)
	}
	data := struct {
        Destination []entity.Destination
        Plane []entity.Plane
    } {
        destinations,
        planes,
    }
	fh.tmpl.ExecuteTemplate(w, "admin.flight.layout", data)
}
