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
	dstns := []entity.Destination{}
	errs := dRepo.conn.Find(&dstns).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return dstns, errs
}

// Destination retrieves a destination by its id from the database
func (dRepo *DestinationGormRepo) Destination(id uint) (*entity.Destination, []error) {
	dstn := entity.Destination{}
	errs := dRepo.conn.First(&dstn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &dstn, errs
}

// UpdateDestination updates a given destination in the database
func (dRepo *DestinationGormRepo) UpdateDestination(destination *entity.Destination) (*entity.Destination, []error) {
	dstn := destination
	errs := dRepo.conn.Save(dstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}

// DeleteDestination deletes a given destination from the database
func (dRepo *DestinationGormRepo) DeleteDestination(id uint) (*entity.Destination, []error) {
	dstn, errs := dRepo.Destination(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = dRepo.conn.Delete(dstn, id).GetErrors()
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
