package flight

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// DestinationService specifies flight destination services
type DestinationService interface {
	Destinations() ([]entity.Destination, []error)
	Destination(id uint) (*entity.Destination, []error)
	UpdateDestination(destination *entity.Destination) (*entity.Destination, []error)
	DeleteDestination(id uint) (*entity.Destination, []error)
	StoreDestination(destination *entity.Destination) (*entity.Destination, []error)
}

// PlaneService specifies flight plane services
type PlaneService interface {
	Planes() ([]entity.Plane, []error)
	Plane(id uint) (*entity.Plane, []error)
	UpdatePlane(plane *entity.Plane) (*entity.Plane, []error)
	DeletePlane(id uint) (*entity.Plane, []error)
	StorePlane(plane *entity.Plane) (*entity.Plane, []error)
}

// FlightService specifies flight services
type FlightService interface {
	Flights() ([]entity.Flight, []error)
	Flight(id uint) (*entity.Flight, []error)
	UpdateFlight(flight *entity.Flight) (*entity.Flight, []error)
	DeleteFlight(id uint) (*entity.Flight, []error)
	StoreFlight(flight *entity.Flight) (*entity.Flight, []error)
	FlightDestination() ([]entity.FlightDestination, []error)
	FlightInfo() ([]entity.FlightInfo, []error)
	MyFlight(userId uint) ([]entity.MyFlight, []error)
}
