package book

import "github.com/Nahom7wos/Airlines-Booking-System/entity"

// TicketRepository specifies flight ticket database operations
type TicketRepository interface {
	Tickets() ([]entity.Ticket, []error)
	Ticket(id uint) (*entity.Ticket, []error)
	UpdateTicket(menu *entity.Ticket) (*entity.Ticket, []error)
	DeleteTicket(id uint) (*entity.Ticket, []error)
	StoreTicket(ticket *entity.Ticket) (*entity.Ticket, []error)
}
