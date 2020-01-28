package repository

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/flight"
	"github.com/jinzhu/gorm"
)

// PlaneGormRepo implements the flight.PlaneRepository interface
type PlaneGormRepo struct {
	conn *gorm.DB
}

// NewPlaneGormRepo will create a new PlaneGormRepo object
func NewPlaneGormRepo(db *gorm.DB) flight.PlaneRepository {
	return &PlaneGormRepo{conn: db}
}

// Planes returns all planes stored in the database
func (pRepo *PlaneGormRepo) Planes() ([]entity.Plane, []error) {
	plns := []entity.Plane{}
	errs := pRepo.conn.Find(&plns).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return plns, errs
}

// Plane retrieves a plane by its id from the database
func (pRepo *PlaneGormRepo) Plane(id uint) (*entity.Plane, []error) {
	pln := entity.Plane{}
	errs := pRepo.conn.First(&pln, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &pln, errs
}

// UpdatePlane updates a given plane in the database
func (pRepo *PlaneGormRepo) UpdatePlane(plane *entity.Plane) (*entity.Plane, []error) {
	pln := plane
	errs := pRepo.conn.Save(pln).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}

// DeletePlane deletes a given plane from the database
func (pRepo *PlaneGormRepo) DeletePlane(id uint) (*entity.Plane, []error) {
	pln, errs := pRepo.Plane(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = pRepo.conn.Delete(pln, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}

// StorePlane stores a given plane in the database
func (pRepo *PlaneGormRepo) StorePlane(plane *entity.Plane) (*entity.Plane, []error) {
	pln := plane
	errs := pRepo.conn.Create(pln).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pln, errs
}
