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

	plns, errs := pServ.planeRepo.Planes()

	if len(errs) > 0 {
		return nil, errs
	}

	return plns, nil
}

// Plane retrieves a plane by its id
func (pServ *PlaneService) Plane(id uint) (*entity.Plane, []error) {
	pln, errs := pServ.planeRepo.Plane(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}

// UpdatePlane updates a given plane
func (pServ *PlaneService) UpdatePlane(plane *entity.Plane) (*entity.Plane, []error) {
	pln, errs := pServ.planeRepo.UpdatePlane(plane)
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}

// DeletePlane deletes a given plane
func (pServ *PlaneService) DeletePlane(id uint) (*entity.Plane, []error) {
	pln, errs := pServ.planeRepo.DeletePlane(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}

// StorePlane persists new plane information
func (pServ *PlaneService) StorePlane(plane *entity.Plane) (*entity.Plane, []error) {

	pln, errs := pServ.planeRepo.StorePlane(plane)

	if len(errs) > 0 {
		return nil, errs
	}

	return pln, nil
}
