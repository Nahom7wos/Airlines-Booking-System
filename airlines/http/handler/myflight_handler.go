package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"github.com/Nahom7wos/Airlines-Booking-System/book"
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"github.com/Nahom7wos/Airlines-Booking-System/user"


)

type MyFlightHandler struct {
	tmpl      *template.Template
	ticketSrv book.TicketService
	flightSrv flight.FlightService
	userSrv  user.UserService
}

// NewMyFlightHandler initializes and returns new MyFlightHandler
func NewMyFlightHandler(T *template.Template, tSrv book.TicketService, fSrv flight.FlightService, uSrv  user.UserService ) *MyFlightHandler {
	return &MyFlightHandler{tmpl: T, ticketSrv: tSrv, flightSrv: fSrv, userSrv: uSrv}
}

// go call user update, go update ticket flight_id as well
//  go delete ticket by id
func (mfh *MyFlightHandler) UpdateMyFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// userid get ticket, send <ticket>
		// userid get user
		//flightDestinaiton to display, value of ticket(flight_id)
		//usrId
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		tkt, errs := mfh.ticketSrv.Ticket(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}
		fltDstn, errs := mfh.flightSrv.FlightDestination()
		if len(errs) > 0 {
			panic(errs)
		}
		usr, errs := mfh.userSrv.User(tkt.UserID)
		if len(errs) > 0 {
			panic(errs)
		}
		data := struct {
			User      *entity.User // gets user info
			FlightDestination []entity.FlightDestination // gets flight.id as ID to choose from
			Ticket       *entity.Ticket // gets flight_id for unchanged value
		}{
			usr,
			fltDstn,
			tkt,
		}

		mfh.tmpl.ExecuteTemplate(w, "myflight.update.layout", data)

	} else if r.Method == http.MethodPost {
		tkt := &entity.Ticket{}
		usr := &entity.User{}

		id, _ := strconv.Atoi(r.FormValue("id"))
		usrId, _ := strconv.Atoi(r.FormValue("usrId"))
		fltID, _ := strconv.Atoi(r.FormValue("flightDestination"))
		
		usr.ID = uint(usrId)
		usr.FullName = r.FormValue("fullName")
		usr.Email = r.FormValue("email")
		usr.Passport = r.FormValue("passport")

		_, errs := mfh.userSrv.UpdateUser(usr)
		if len(errs) > 0 {
			panic(errs)
		}

		tkt.ID = uint(id)
		tkt.FlightID = uint(fltID)
		//user can't change Ticket's UserID
		_, errs := mfh.ticketSrv.UpdateTicket(tkt) 
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/myflight", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/myflight", http.StatusSeeOther)
	}
}