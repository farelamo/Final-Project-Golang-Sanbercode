package filmService

import (
	"FinalProject/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type FilmImpl struct {
	DB *sql.DB
}

func NewFilmService(DB *sql.DB) FilmService {
	return &FilmImpl {
		DB: DB,
	}
}


func (f *FilmImpl) GetFilm() (*[]models.Film, error) {
	var films = []models.Film{}

	sql 		:= `SELECT * FROM films`
	rows, err 	:= f.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var film = models.Film{}

		err = rows.Scan(&film.Id, &film.Title, &film.Desc, &film.Genre, &film.Year, &film.Actor, &film.CreatedAt, &film.UpdatedAt,)
		if err != nil {
			return nil, err
		}

		films = append(films, film)
	}

	return &films, nil
}

func (f *FilmImpl) AddFilm(film *models.AddFilm) (*models.Film, error) {
	var newFilm = models.Film{}

	if film.Title == "" {
		return nil, errors.New("Title Must Be Filled")
	}else if film.Desc == "" {
		return nil, errors.New("Description Must Be Filled")
	}else if film.Genre == "" {
		return nil, errors.New("Genre Must Be Filled")
	}else if strconv.Itoa(film.Year) == "" {
		return nil, errors.New("Year Must Be Filled")
	}else if film.Actor == "" {
		return nil, errors.New("Actor Must Be Filled")
	}

	sql := `INSERT INTO films (title, description, genre, year, actor) VALUES ($1, $2, $3, $4, $5) Returning *`
	err := f.DB.QueryRow(sql, film.Title, film.Desc, film.Genre, film.Year, film.Actor).Scan(
		&newFilm.Id, &newFilm.Title, &newFilm.Desc, &newFilm.Genre, &newFilm.Year, 
		&newFilm.Actor, &newFilm.CreatedAt, &newFilm.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newFilm, nil
}

func (f *FilmImpl) GetFilmById(id int) (*models.Film, error) {
	var film = models.Film{}
	sql := `SELECT * FROM films WHERE id=($1)`
	err := f.DB.QueryRow(sql, id).Scan(&film.Id, &film.Title, &film.Desc, &film.Genre, &film.Year, &film.Actor, &film.CreatedAt, &film.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &film, err
}

func (f *FilmImpl) UpdateFilm(id int, film *models.Film) (int, error) {
	if film.Title == "" {
		return 0, errors.New("Title Must Be Filled")
	}else if film.Desc == "" {
		return 0, errors.New("Description Must Be Filled")
	}else if film.Genre == "" {
		return 0, errors.New("Genre Must Be Filled")
	}else if strconv.Itoa(film.Year) == "" {
		return 0, errors.New("Year Must Be Filled")
	}else if film.Actor == "" {
		return 0, errors.New("Actor Must Be Filled")
	}

	sqlStatement := `UPDATE films SET title=$2, description=$3, genre=$4, year=$5, actor=$6 WHERE id=$1;`
	
	result, err := f.DB.Exec(sqlStatement, id, film.Title,  film.Desc, film.Genre, film.Year, film.Actor,)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating film record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the film, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (f *FilmImpl) DeleteFilm(id int) (int, error) {
	sql := `DELETE FROM films WHERE id=$1;`
	res, err := f.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete film record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the film, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


