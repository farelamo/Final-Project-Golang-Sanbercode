package models

import (
	"time"
)

type Ticket struct {
	Id 			int 		`json:"id,omitempty"`
	Type 		string 		`json:"type,omitempty"`
	UserId 		int 		`json:"user_id,omitempty"`
	StudioId 	int 		`json:"studio_id,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}

type AddTicket struct {
	Type 		string 		`json:"type,omitempty"`
	UserId 		int 		`json:"user_id,omitempty"`
	StudioId 	int 		`json:"studio_id,omitempty"`
}

type TiketResultById struct {
	Id 			int 				`json:"id"`
	Type 		string 				`json:"type"`
	UserId 		*User 				`json:"user_id"`
	Studio 	    *ResultStudioById 	`json:"studio_id"`
	CreatedAt   time.Time 			`json:"created_at,omitempty"`
	UpdatedAt   time.Time 			`json:"updated_at,omitempty"`
}