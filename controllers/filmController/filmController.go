package filmController

import "github.com/gin-gonic/gin"

type FilmController interface {
	AddFilm(ctx *gin.Context)
	GetFilm(ctx *gin.Context)
	GetFilmById(ctx *gin.Context)
	UpdateFilm(ctx *gin.Context)
	DeleteFilm(ctx *gin.Context)
}