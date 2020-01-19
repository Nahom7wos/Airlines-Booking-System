package repository

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/ticket"
	"github.com/jinzhu/gorm"
)

// TicketGormRepo implements the flight.TicketRepository interface
type TicketGormRepo struct {
	conn *gorm.DB
}

// NewTicketGormRepo will create a new TicketGormRepo object
func NewTicketGormRepo(db *gorm.DB) flight.TicketRepository {
	return &TicketGormRepo{conn: db}
}

// Tickets returns all tickets stored in the database
func (tRepo *TicketGormRepo) Tickets() ([]entity.Ticket, []error) {
	tkts := []entity.Ticket{}
	errs := tRepo.conn.Find(&tkts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tkts, errs
}

// Ticket retrieves a ticket by its id from the database
func (tRepo *TicketGormRepo) Ticket(id uint) (*entity.Ticket, []error) {
	tkt := entity.Ticket{}
	errs := tRepo.conn.First(&tkt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &tkt, errs
}

// UpdateTicket updates a given ticket in the database
func (tRepo *TicketGormRepo) UpdateTicket(ticket *entity.Ticket) (*entity.Ticket, []error) {
	tkt := ticket
	errs := tRepo.conn.Save(tkt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// DeleteTicket deletes a given ticket from the database
func (tRepo *TicketGormRepo) DeleteTicket(id uint) (*entity.Ticket, []error) {
	tkt, errs := tRepo.Ticket(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = tRepo.conn.Delete(tkt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// StoreTicket stores a given ticket in the database
func (tRepo *TicketGormRepo) StoreTicket(ticket *entity.Ticket) (*entity.Ticket, []error) {
	tkt := ticket
	errs := tRepo.conn.Create(tkt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}
