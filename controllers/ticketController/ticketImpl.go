package ticketController

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	. "FinalProject/services/ticketService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketControllerImpl struct {
	ticketService TicketService
}

func NewTicketController(ticketService TicketService) TicketController {
	return &TicketControllerImpl{
		ticketService: ticketService,
	}
}

func (t *TicketControllerImpl) AddTicket(ctx *gin.Context){
	var addTicket models.AddTicket

	if err := ctx.ShouldBindJSON(&addTicket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"message": err,
		})
		return
	}
	
	film, err := t.ticketService.AddTicket(&addTicket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, film)
}

func (t *TicketControllerImpl) GetTicket(ctx *gin.Context) {
	ticket, err := t.ticketService.GetTicket()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status": true,
		"data"	: ticket,
	})
}

func (t *TicketControllerImpl) GetTicketById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"message": "Invalid request",
		})
		return
	}

	ticket, err := t.ticketService.GetTicketById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status": false,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
			"Status": true,
			"data"  : ticket,
		},
	)
}

func (t *TicketControllerImpl) UpdateTicket(ctx *gin.Context) {
	var ticket models.Ticket

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	count, err := t.ticketService.UpdateTicket(id, &ticket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})
}

func (t *TicketControllerImpl) DeleteTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	count, err := t.ticketService.DeleteTicket(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})	
}


