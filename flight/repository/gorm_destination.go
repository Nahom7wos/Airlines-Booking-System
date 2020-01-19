package repository

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"github.com/jinzhu/gorm"
)

// DestinationGormRepo implements the flight.DestinationRepository interface
type DestinationGormRepo struct {
	conn *gorm.DB
}

// NewDestinationGormRepo will create a new DestinationGormRepo object
func NewDestinationGormRepo(db *gorm.DB) flight.DestinationRepository {
	return &DestinationGormRepo{conn: db}
}

// Destinations returns all destinations stored in the database
func (dRepo *DestinationGormRepo) Destinations() ([]entity.Destination, []error) {
	dstn := []entity.Destination{}
	errs := dRepo.conn.Find(&dstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}

// StoreDestination stores a given destination in the database
func (dRepo *DestinationGormRepo) StoreDestination(destination *entity.Destination) (*entity.Destination, []error) {
	dstn := destination
	errs := dRepo.conn.Create(dstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}
