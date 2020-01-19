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
func (fRepo *FlightGormRepo) Flights() ([]entity.Flight, []error) {
	flts := []entity.Flight{}
	errs := fRepo.conn.Find(&flts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flts, errs
}

// Flight retrieves a flight by its id from the database
func (fRepo *FlightGormRepo) Flight(id uint) (*entity.Flight, []error) {
	flt := entity.Flight{}
	errs := fRepo.conn.First(&flt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &flt, errs
}

// UpdateFlight updates a given flight in the database
func (fRepo *FlightGormRepo) UpdateFlight(flight *entity.Flight) (*entity.Flight, []error) {
	flt := flight
	errs := fRepo.conn.Save(flt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// DeleteFlight deletes a given flight from the database
func (fRepo *FlightGormRepo) DeleteFlight(id uint) (*entity.Flight, []error) {
	flt, errs := fRepo.Flight(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = fRepo.conn.Delete(flt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}

// StoreFlight stores a given flight in the database
func (fRepo *FlightGormRepo) StoreFlight(flight *entity.Flight) (*entity.Flight, []error) {
	flt := flight
	errs := fRepo.conn.Create(flt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return flt, errs
}
