package loyalty

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// LoyaltyRepository specifies flight loyalty database operations
type LoyaltyRepository interface {
	Loyaltys() ([]entity.Loyalty, []error)
	Loyalty(id uint) (*entity.Loyalty, []error)
	UpdateLoyalty(menu *entity.Loyalty) (*entity.Loyalty, []error)
	DeleteLoyalty(id uint) (*entity.Loyalty, []error)
	StoreLoyalty(loyalty *entity.Loyalty) (*entity.Loyalty, []error)
}
