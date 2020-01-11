package flight

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// DestinationService specifies flight destination services
type DestinationService interface {
	Destinations() ([]entity.Destination, []error)
	StoreDestination(destination *entity.Destination) (*entity.Destination, []error)
}

// PlaneService specifies flight plane services
type PlaneService interface {
	Planes() ([]entity.Plane, []error)
	StorePlane(plane *entity.Plane) (*entity.Plane, []error)
}

// FlightService specifies flight services
type FlightService interface {
	Flights() ([]entity.Flight, []error)
	StoreFlight(flight *entity.Flight) (*entity.Flight, []error)
}