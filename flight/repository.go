package flight

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// DestinationRepository specifies flight destination database operations
type DestinationRepository interface {
	Destinations() ([]entity.Destination, []error)
	StoreDestination(destination *entity.Destination) (*entity.Destination, []error)
}

// PlaneRepository specifies flight plane database operations
type PlaneRepository interface {
	Planes() ([]entity.Plane, []error)
	StorePlane(plane *entity.Plane) (*entity.Plane, []error)
}

// FlightRepository specifies flight database operations
type FlightRepository interface {
	Flighs() ([]entity.Flight, []error)
	StoreFlight(flight *entity.Flight) (*entity.Flight, []error)
}