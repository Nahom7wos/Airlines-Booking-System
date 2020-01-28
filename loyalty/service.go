package ticket

// TicketService specifies flight ticket services
type TicketService interface {
	Tickets() ([]entity.Ticket, []error)
	Ticket(id uint) (*entity.Ticket, []error)
	UpdateTicket(menu *entity.Ticket) (*entity.Ticket, []error)
	DeleteTicket(id uint) (*entity.Ticket, []error)
	StoreTicket(ticket *entity.Ticket) (*entity.Ticket, []error)
}