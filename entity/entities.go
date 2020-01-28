package entity

// Flight includes destination and plane
type Flight struct {
	ID            uint
	DepartureDate string `gorm:"type:varchar(255);not null"`
	Status        bool
	DestinationID uint
	PlaneID       uint
}

// Destination details
type Destination struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Image       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Price       uint   `gorm:"not null"`
}

// Plane details with unique names
type Plane struct {
	ID       uint
	Name     string `gorm:"type:varchar(255);not null;unique"`
	Capacity uint   `gorm:"not null"`
	Status   bool
}

// Ticket info with Checkin Status
type Ticket struct {
	ID       uint
	Status   bool
	FlightID uint
	UserID   uint
}

// Login represents registered users
type Login struct {
	ID       uint
	Username string `gorm:"type:varchar(255);not null; unique"`
	Password string `gorm:"type:varchar(255);not null"`
	UserID 	 uint
}

// User info
type User struct {
	ID           uint
	FullName     string     `gorm:"type:varchar(255);not null"`
	Email        string     `gorm:"type:varchar(255);not null; unique"`
	Passport     string     `gorm:"type:varchar(255);not null; unique"`
	Registered   bool       
	RoleID        uint     
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

// Loyalty to be redeemed
type Loyalty struct {
	ID     uint
	Status bool
	Point  uint
	UserID uint
}


//Custom entities 

// FlightDestination details
type FlightDestination struct {
	ID          uint
	DepartureDate string 
	Name        string 
	Price       uint  
}
// FlightInfo holds related plane and destination name
type FlightInfo struct {
	ID          uint
	DepartureDate string 
	Status        bool
	DestinationName        string 
	PlaneName       string 
}

// MyFlight holdes user, ticket and flight destination info
type MyFlight struct {
	ID		uint
	DestinationName string
	DepartureDate string
	Status bool
}