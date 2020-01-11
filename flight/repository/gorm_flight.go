package repository

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"github.com/jinzhu/gorm"
)

// FlightGormRepo implements the flight.FlightRepository interface
type FlightGormRepo struct {
	conn *gorm.DB
}

// NewFlightGormRepo will create a new FlightGormRepo object
func NewFlightGormRepo(db *gorm.DB) flight.FlightRepository {
	return &FlightGormRepo{conn: db}
}

// Flights returns all flights stored in the database
func (pRepo *FlightGormRepo) Flights() ([]entity.Flight, []error) {
	flt := []entity.Flight{}
	errs := pRepo.conn.Find(&flt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// StoreFlight stores a given flight in the database
func (pRepo *FlightGormRepo) StoreFlight(flight *entity.Flight) (*entity.Flight, []error) {
	flt := flight
	errs := pRepo.conn.Create(flt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}
