package models

import "time"

type Studio struct {
	Id 			int 		`json:"id,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	FilmId 		int 		`json:"film_id,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}

type AddStudio struct {
	Name 		string 		`json:"name,omitempty"`
	FilmId 		int 		`json:"film_id,omitempty"`
}

type ResultStudioById struct {
	Id 			int 		`json:"id,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	FilmId 		*Film 		`json:"film_id,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}