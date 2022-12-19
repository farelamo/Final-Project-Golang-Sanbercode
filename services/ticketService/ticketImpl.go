package ticketService

import (
	"FinalProject/services/studioService"
	"FinalProject/services/userService"
	"FinalProject/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type TicketImpl struct {
	DB *sql.DB
}

func NewTicketService(DB *sql.DB) TicketService {
	return &TicketImpl {
		DB: DB,
	}
}


func (t *TicketImpl) GetTicket() (*[]models.TiketResultById, error) {
	var tickets = []models.TiketResultById{}

	sql 		:= `SELECT * FROM tickets`
	rows, err 	:= t.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket = models.Ticket{}

		err = rows.Scan(&ticket.Id, &ticket.Type, &ticket.UserId, &ticket.StudioId, &ticket.CreatedAt, &ticket.UpdatedAt,)
		if err != nil {
			return nil, err
		}

		studioService 		:= studioService.NewStudioService(t.DB)
		getStudio, err 	:= studioService.GetStudioById(ticket.StudioId)
		if err != nil {
			return nil, err
		}

		userService 	:= userService.NewUserService(t.DB)
		getUser, err 	:= userService.GetUserById(ticket.UserId)
		if err != nil{
			return nil, err
		}

		result := models.TiketResultById{
			ticket.Id, ticket.Type, getUser, getStudio,
			ticket.CreatedAt, ticket.UpdatedAt,
		}

		tickets = append(tickets, result)
	}

	return &tickets, nil
}

func (t *TicketImpl) AddTicket(ticket *models.AddTicket) (*models.Ticket, error) {
	var newTicket = models.Ticket{}

	if ticket.Type == "" {
		return nil, errors.New("Title Must Be Filled")
	}else if strconv.Itoa(ticket.UserId) == "" {
		return nil, errors.New("user Id Must Be Filled")
	}else if strconv.Itoa(ticket.StudioId) == "" {
		return nil, errors.New("Studi Id Must Be Filled")
	}

	studioService 		:= studioService.NewStudioService(t.DB)
	_, err 	:= studioService.GetStudioById(ticket.StudioId)

	if err != nil {
		return nil, errors.New("There's no data studio with id " + strconv.Itoa(ticket.StudioId))
	}

	sql := `INSERT INTO tickets (type, user_id, studio_id) VALUES ($1, $2, $3) Returning *`
	err = t.DB.QueryRow(sql, ticket.Type, ticket.UserId, ticket.StudioId,).Scan(
		&newTicket.Id, &newTicket.Type, &newTicket.UserId, &newTicket.StudioId, &newTicket.CreatedAt, &newTicket.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newTicket, nil
}

func (t *TicketImpl) GetTicketById(id int) (*models.TiketResultById, error) {
	var ticket = models.Ticket{}

	sql := `SELECT * FROM tickets WHERE id=($1)`
	err := t.DB.QueryRow(sql, id).Scan(&ticket.Id, &ticket.Type, &ticket.UserId, &ticket.StudioId, &ticket.CreatedAt, &ticket.UpdatedAt,)
	if err != nil {
		return nil, err		
	}

	studioService 	:= studioService.NewStudioService(t.DB)
	getStudio, err 	:= studioService.GetStudioById(ticket.StudioId)
	if err != nil{
		return nil, err
	}

	userService 	:= userService.NewUserService(t.DB)
	getUser, err 	:= userService.GetUserById(ticket.UserId)
	if err != nil{
		return nil, err
	}

	result := models.TiketResultById{ticket.Id, ticket.Type, getUser, getStudio, ticket.CreatedAt, ticket.UpdatedAt}

	return &result, err
}

func (t *TicketImpl) UpdateTicket(id int, ticket *models.Ticket) (int, error) {
	if ticket.Type == "" {
		return 0, errors.New("Type Must Be Filled")
	}else if strconv.Itoa(ticket.UserId) == "" {
		return 0, errors.New("User Id Must Be Filled")
	}else if strconv.Itoa(ticket.StudioId) == "" {
		return 0, errors.New("Studio Id Must Be Filled")
	}

	sqlStatement := `UPDATE tickets SET type=$2, user_id=$3, studio_id=$4 WHERE id=$1;`
	
	result, err := t.DB.Exec(sqlStatement, id, ticket.Type, ticket.UserId, ticket.StudioId,)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating ticket record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the ticket, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (t *TicketImpl) DeleteTicket(id int) (int, error) {
	sql := `DELETE FROM tickets WHERE id=$1;`
	res, err := t.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete ticket record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the ticket, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


