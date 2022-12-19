package models

import "time"

type Film struct {
	Id 			int 		`json:"id,omitempty"`
	Title 		string 		`json:"title,omitempty"`
	Desc 		string 		`json:"description,omitempty"`
	Genre 		string 		`json:"genre,omitempty"`
	Year 		int 		`json:"year,omitempty"`
	Actor 		string 		`json:"actor,omitempty"`
	CreatedAt   time.Time 	`json:"created_at,omitempty"`
	UpdatedAt   time.Time 	`json:"updated_at,omitempty"`
}

type AddFilm struct {
	Title 		string 		`json:"title,omitempty"`
	Desc 		string 		`json:"description,omitempty"`
	Genre 		string 		`json:"genre,omitempty"`
	Year 		int 		`json:"year,omitempty"`
	Actor 		string 		`json:"actor,omitempty"`
}