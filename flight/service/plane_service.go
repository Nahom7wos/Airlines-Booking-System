package service

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
)

// PlaneService implements flight.PlaneService interface
type PlaneService struct {
	planeRepo flight.PlaneRepository
}

// NewPlaneService will create new PlaneService object
func NewPlaneService(pRepo flight.PlaneRepository) flight.PlaneService {
	return &PlaneService{planeRepo: pRepo}
}

// Planes returns list of planes
func (pServ *PlaneService) Planes() ([]entity.Plane, []error) {

	planes, errs := pServ.planeRepo.Planes()

	if len(errs) > 0 {
		return nil, errs
	}

	return planes, nil
}

// StorePlane persists new plane information
func (pServ *PlaneService) StorePlane(plane *entity.Plane) (*entity.Plane, []error) {

	pln, errs := pServ.planeRepo.StorePlane(plane)

	if len(errs) > 0 {
		return nil, errs
	}

	return pln, nil
}
