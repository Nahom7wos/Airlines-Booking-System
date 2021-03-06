package service

import (
	"github.com/Nahom7wos/Airlines-Booking-System/entity"
	"github.com/Nahom7wos/Airlines-Booking-System/book"
)

// TicketService implements book.TicketService interface
type TicketService struct {
	ticketRepo book.TicketRepository
}

// NewTicketService will create new TicketService object
func NewTicketService(destRepo book.TicketRepository) book.TicketService {
	return &TicketService{ticketRepo: destRepo}
}

// Tickets returns list of tickets
func (tServ *TicketService) Tickets() ([]entity.Ticket, []error) {

	tkts, errs := tServ.ticketRepo.Tickets()

	if len(errs) > 0 {
		return nil, errs
	}

	return tkts, nil
}

// Ticket retrieves a ticket by its id
func (tServ *TicketService) Ticket(id uint) (*entity.Ticket, []error) {
	tkt, errs := tServ.ticketRepo.Ticket(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// UpdateTicket updates a given ticket
func (tServ *TicketService) UpdateTicket(ticket *entity.Ticket) (*entity.Ticket, []error) {
	tkt, errs := tServ.ticketRepo.UpdateTicket(ticket)
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// DeleteTicket deletes a given ticket
func (tServ *TicketService) DeleteTicket(id uint) (*entity.Ticket, []error) {
	tkt, errs := tServ.ticketRepo.DeleteTicket(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// StoreTicket persists new ticket information
func (tServ *TicketService) StoreTicket(ticket *entity.Ticket) (*entity.Ticket, []error) {

	tkt, errs := tServ.ticketRepo.StoreTicket(ticket)

	if len(errs) > 0 {
		return nil, errs
	}

	return tkt, nil
}

// Ticket retrieves a ticket by its id
func (tServ *TicketService) TicketByUserId(userId uint) (*entity.Ticket, []error) {
	tkt, errs := tServ.ticketRepo.TicketByUserId(userId)
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}

// UpdateTicketStatus updates a given ticket status 
func (tServ *TicketService) UpdateTicketStatus(ticket *entity.Ticket, status bool) (*entity.Ticket, []error) {
	tkt, errs := tServ.ticketRepo.UpdateTicketStatus(ticket, status)
	if len(errs) > 0 {
		return nil, errs
	}
	return tkt, errs
}