package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

type FlightHandler struct {
	tmpl           *template.Template
	destinationSrv flight.DestinationService
	planeSrv       flight.PlaneService
	flightSrv      flight.FlightService
}

// NewFlightHandler initializes and returns new FlightHandler
func NewFlightHandler(T *template.Template, pSrv flight.PlaneService, dSrv flight.DestinationService, fSrv flight.FlightService) *FlightHandler {
	return &FlightHandler{tmpl: T, planeSrv: pSrv, destinationSrv: dSrv, flightSrv: fSrv}
}

// Flight displays all the flights in the database
func (fh *FlightHandler) Flight(w http.ResponseWriter, r *http.Request) {
	flights, errs := fh.flightSrv.FlightInfo()
	if errs != nil {
		panic(errs)
	}
	fh.tmpl.ExecuteTemplate(w, "admin.flight.layout", flights)
}

// StoreFlight creates new flight in the database
func (fh *FlightHandler) StoreFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		flt := &entity.Flight{}
		dstnID, _ := strconv.Atoi(r.FormValue("flightDestination"))
		plnID, _ := strconv.Atoi(r.FormValue("flightPlane"))
		status, _ := strconv.ParseBool(r.FormValue("flightStatus"))

		flt.DestinationID = uint(dstnID)
		flt.PlaneID = uint(plnID)
		flt.DepartureDate = r.FormValue("flightDeparture")
		flt.Status = status

		_, errs := fh.flightSrv.StoreFlight(flt)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/flight", http.StatusSeeOther)
	} else {
		destinations, errs := fh.destinationSrv.Destinations()
		if errs != nil {
			panic(errs)
		}
		planes, errs := fh.planeSrv.Planes()
		if errs != nil {
			panic(errs)
		}
		data := struct {
			Destination []entity.Destination
			Plane       []entity.Plane
		}{
			destinations,
			planes,
		}
		fh.tmpl.ExecuteTemplate(w, "admin.flight.create.layout", data)
	}
}

// get flight by id
// get all destinations and planes
// UpdateFlight handle requests on /admin/destination/update
func (fh *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		flight, errs := fh.flightSrv.Flight(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}
		destinations, errs := fh.destinationSrv.Destinations()
		if errs != nil {
			panic(errs)
		}
		planes, errs := fh.planeSrv.Planes()
		if errs != nil {
			panic(errs)
		}
		data := struct {
			Flight      *entity.Flight
			Destination []entity.Destination
			Plane       []entity.Plane
		}{
			flight,
			destinations,
			planes,
		}

		fh.tmpl.ExecuteTemplate(w, "admin.flight.update.layout", data)

	} else if r.Method == http.MethodPost {

		flt := &entity.Flight{}
		dstnID, _ := strconv.Atoi(r.FormValue("flightDestination"))
		plnID, _ := strconv.Atoi(r.FormValue("flightPlane"))
		status, _ := strconv.ParseBool(r.FormValue("flightStatus"))

		flt.DestinationID = uint(dstnID)
		flt.PlaneID = uint(plnID)
		flt.DepartureDate = r.FormValue("flightDeparture")
		flt.Status = status

		_, errs := fh.flightSrv.StoreFlight(flt)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/flight", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/flight", http.StatusSeeOther)
	}

}

// DeleteFlight handles requests on route /admin/flight/delete
func (fh *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}

		_, errs := fh.flightSrv.DeleteFlight(uint(id))
		if len(errs) > 0 {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/flight", http.StatusSeeOther)
}
