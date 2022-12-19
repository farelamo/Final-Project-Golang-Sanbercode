package filmService

import "FinalProject/models"

type FilmService interface {
	AddFilm(film *models.AddFilm) (*models.Film, error)
	GetFilm() (*[]models.Film, error)
	GetFilmById(id int) (*models.Film, error)
	UpdateFilm(id int, film *models.Film) (int, error)
	DeleteFilm(id int) (int, error)
}