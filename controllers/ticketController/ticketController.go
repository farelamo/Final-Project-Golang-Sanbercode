package ticketController

import "github.com/gin-gonic/gin"

type TicketController interface {
	AddTicket(ctx *gin.Context)
	GetTicket(ctx *gin.Context)
	GetTicketById(ctx *gin.Context)
	UpdateTicket(ctx *gin.Context)
	DeleteTicket(ctx *gin.Context)
}