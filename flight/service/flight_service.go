package service

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

// FlightService implements flight.FlightService interface
type FlightService struct {
	flightRepo flight.FlightRepository
}

// NewFlightService will create new FlightService object
func NewFlightService(fRepo flight.FlightRepository) flight.FlightService {
	return &FlightService{flightRepo: fRepo}
}

// Flights returns list of flights
func (fServ *FlightService) Flights() ([]entity.Flight, []error) {

	flights, errs := fServ.flightRepo.Flights()

	if len(errs) > 0 {
		return nil, errs
	}

	return flights, nil
}

// StoreFlight persists new flight information
func (fServ *FlightService) StoreFlight(flight *entity.Flight) (*entity.Flight, []error) {

	flt, errs := fServ.flightRepo.StoreFlight(flight)

	if len(errs) > 0 {
		return nil, errs
	}

	return flt, nil
}
