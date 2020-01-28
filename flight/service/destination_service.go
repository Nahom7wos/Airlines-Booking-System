package service

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

// DestinationService implements flight.DestinationService interface
type DestinationService struct {
	destinationRepo flight.DestinationRepository
}

// NewDestinationService will create new DestinationService object
func NewDestinationService(destRepo flight.DestinationRepository) flight.DestinationService {
	return &DestinationService{destinationRepo: destRepo}
}

// Destinations returns list of destinations
func (dServ *DestinationService) Destinations() ([]entity.Destination, []error) {

	dstns, errs := dServ.destinationRepo.Destinations()

	if len(errs) > 0 {
		return nil, errs
	}

	return dstns, nil
}

// Destination retrieves a destination by its id
func (dServ *DestinationService) Destination(id uint) (*entity.Destination, []error) {
	dstn, errs := dServ.destinationRepo.Destination(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}

// UpdateDestination updates a given destination
func (dServ *DestinationService) UpdateDestination(destination *entity.Destination) (*entity.Destination, []error) {
	dstn, errs := dServ.destinationRepo.UpdateDestination(destination)
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}

// DeleteDestination deletes a given destination
func (dServ *DestinationService) DeleteDestination(id uint) (*entity.Destination, []error) {
	dstn, errs := dServ.destinationRepo.DeleteDestination(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return dstn, errs
}

// StoreDestination persists new destination information
func (dServ *DestinationService) StoreDestination(destination *entity.Destination) (*entity.Destination, []error) {

	dstn, errs := dServ.destinationRepo.StoreDestination(destination)

	if len(errs) > 0 {
		return nil, errs
	}

	return dstn, nil
}