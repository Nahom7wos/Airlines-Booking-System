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
func NewDestinationService(destRepo flight.DestinationRepository) flight.DestiantionService {
	return &DestinationService{destinationRepo: destRepo}
}

// Destinations returns list of destinations
func (dServ *DestinationService) Destinations() ([]entity.Destination, []error) {

	destinations, errs := dServ.destinationRepo.Destinations()

	if len(errs) > 0 {
		return nil, errs
	}

	return destinations, nil
}

// StoreDestination persists new destination information
func (dServ *DestinationService) StoreDestination(destination *entity.Destination) (*entity.Destination, []error) {

	dstn, errs := dServ.destinationRepo.StoreDestination(destination)

	if len(errs) > 0 {
		return nil, errs
	}

	return dstn, nil
}