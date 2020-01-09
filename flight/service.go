package flight

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// DestinationService specifies flight destination services
type DestinationService interface {
	Destinations() ([]entity.Destination, []error)
	StoreDestination(destination *entity.Destination) (*entity.Destination, []error)
}
