package entity

import "time"

// Flight includes destination and plane
type Flight struct {
	ID            uint
	DepartureDate time.Time
	Status        bool
	DestinationID Destination
	PlaneID       Plane
}

// Destination details
type Destination struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Image       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Price       float32
}

// Plane details with unique names
type Plane struct {
	ID       uint
	Name     string `gorm:"type:varchar(255);not null;unique"`
	Capacity uint `gorm:"not null"`
}

// Ticket info with Checkin Status
type Ticket struct {
	ID       uint
	Status   bool
	FlightID uint
	UserID   uint
}

// Role repesents application user roles
type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

// Login represents registered users
type Login struct {
	ID       uint
	Username string `gorm:"type:varchar(255);not null; unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

// CreditCard details
type CreditCard struct {
	ID              uint
	CardNumber      string `gorm:"type:varchar(255);not null; unique"`
	ExpirationMonth uint 
	ExpirationYear  uint
}

// User info
type User struct {
	ID           uint
	FullName     string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(255);not null; unique"`
	Passport     string `gorm:"type:varchar(255);not null; unique"`
	Registered   bool
	CreditCardID CreditCard // One-To-One relationship
	Roles        []Role     `gorm:"many2many:user_roles"`
}

// Loyalty to be redeemed
type Loyalty struct {
	ID     uint
	Status bool
	UserID uint
}
