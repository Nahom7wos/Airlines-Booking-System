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

	flts, errs := fServ.flightRepo.Flights()

	if len(errs) > 0 {
		return nil, errs
	}

	return flts, nil
}

// Flight retrieves a flight by its id
func (fServ *FlightService) Flight(id uint) (*entity.Flight, []error) {
	flt, errs := fServ.flightRepo.Flight(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// UpdateFlight updates a given flight
func (fServ *FlightService) UpdateFlight(flight *entity.Flight) (*entity.Flight, []error) {
	flt, errs := fServ.flightRepo.UpdateFlight(flight)
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// DeleteFlight deletes a given flight
func (fServ *FlightService) DeleteFlight(id uint) (*entity.Flight, []error) {
	flt, errs := fServ.flightRepo.DeleteFlight(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// StoreFlight persists new flight information
func (fServ *FlightService) StoreFlight(flight *entity.Flight) (*entity.Flight, []error) {

	flt, errs := fServ.flightRepo.StoreFlight(flight)

	if len(errs) > 0 {
		return nil, errs
	}

	return flt, nil
}

// FlightDestination returns list of flight and destination
func (fServ *FlightService) FlightDestination() ([]entity.FlightDestination, []error) {

	fltDstn, errs := fServ.flightRepo.FlightDestination()
	if len(errs) > 0 {
		return nil, errs
	}

	return fltDstn, nil
}

// FlightInfo returns list of flight, plane, and destination
func (fServ *FlightService) FlightInfo() ([]entity.FlightInfo, []error) {

	fltDstn, errs := fServ.flightRepo.FlightInfo()
	if len(errs) > 0 {
		return nil, errs
	}

	return fltDstn, nil
}

// MyFlight retrieves a flight by its userid
func (fServ *FlightService) MyFlight(userId uint) ([]entity.MyFlight, []error) {
	myFlts, errs := fServ.flightRepo.MyFlight(userId)
	if len(errs) > 0 {
		return nil, errs
	}
	return myFlts, errs
}