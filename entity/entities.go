package entity

import "time"

// Flight includes destination and plane
type Flight struct {
	ID            uint
	DepartureDate time.Time
	Status        string
	DestinationID []Destination `gorm:"many2one:flight_destination"`
	PlaneID       []Plane       `gorm:"one2one:flight_plane"`
}

// Destination details
type Destination struct {
	ID          uint
	Name        string `gorm:"type:varchar(255)"`
	Price       float32
	Description string
	Image       string `gorm:"type:varchar(255)"`
}

// Plane details with unique names
type Plane struct {
	ID       uint
	Name     string `gorm:"type:varchar(255);unique"`
	Capacity uint
}

// Ticket info with Checkin Status
type Ticket struct {
	ID       uint
	Status   string
	FlightID uint
	UserID   uint
}

// Role repesents application user roles
type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

// CreditCard details
type CreditCard struct {
	ID              uint
	CardNumber      string
	ExpirationMonth uint
	ExpirationYear  uint
}

// User info
type User struct {
	ID           uint
	FullName     string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(255);not null; unique"`
	Phone        string `gorm:"type:varchar(100);not null; unique"`
	Password     string `gorm:"type:varchar(255)"`
	Passport     string
	CreditCardID []CreditCard
	Roles        []Role   `gorm:"many2many:user_roles"`
	TicketID     []Ticket `gorm:"one2many:user_tickets"`
}

// Loyalty to be redeemed
type Loyalty struct {
	ID     uint
	Status string
	UserID uint
}
