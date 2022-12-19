package ticketService

import "FinalProject/models"

type TicketService interface {
	AddTicket(ticket *models.AddTicket) (*models.Ticket, error)
	GetTicket() (*[]models.TiketResultById, error)
	GetTicketById(id int) (*models.TiketResultById, error)
	UpdateTicket(id int, ticket *models.Ticket) (int, error)
	DeleteTicket(id int) (int, error)
}