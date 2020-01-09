package flght

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// DestinationRepository specifies flight destination database operations
type DestinationRepository interface {
	Destinations() ([]entity.Destination, []error)
	StoreDestination(destination *entity.Destination) (*entity.Destination, []error)
}
