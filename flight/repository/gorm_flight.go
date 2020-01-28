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

//FlightDestination returns destination, and flight info
func (fRepo *FlightGormRepo) FlightDestination() ([]entity.FlightDestination, []error) {
	fltDstn := []entity.FlightDestination{}
	errs := fRepo.conn.Table("flights").Select("flights.id, flights.departure_date, destinations.name, destinations.price").Joins("inner join destinations on flights.destination_id = destinations.id").Where("flights.status = ?", "true").Scan(&fltDstn).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return fltDstn, nil
}
//FlightDestination returns plane, destination, and flight info
func (fRepo *FlightGormRepo) FlightInfo() ([]entity.FlightInfo, []error) {
	fltInfo := []entity.FlightInfo{}
	errs := fRepo.conn.Table("flights").Select("flights.id, flights.departure_date, flights.status, destinations.name as destination_name, planes.name as plane_name").Joins("inner join destinations on flights.destination_id = destinations.id").Joins("inner join planes on flights.plane_id = planes.id").Scan(&fltInfo).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return fltInfo, nil
}
//MyFlight returns destination, ticket, and flight info
func (fRepo *FlightGormRepo) MyFlight(userId uint) ([]entity.MyFlight, []error) {
	myFlt := []entity.MyFlight{}
	errs := fRepo.conn.Table("tickets").Select("tickets.id, destinations.name as destination_name, flights.departure_date, tickets.status").Joins("inner join users on tickets.user_id = users.id").Joins("inner join flights on tickets.flight_id = flights.id").Joins("inner join destinations on flights.destination_id = destinations.id").Where("user_id = ?", userId).Scan(&myFlt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return myFlt, nil
}



