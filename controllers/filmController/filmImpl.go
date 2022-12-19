package filmController

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	. "FinalProject/services/filmService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FilmControllerImpl struct {
	filmService FilmService
}

func NewFilmController(filmService FilmService) FilmController {
	return &FilmControllerImpl{
		filmService: filmService,
	}
}

func (f *FilmControllerImpl) AddFilm(ctx *gin.Context){
	var addFilm models.AddFilm

	if err := ctx.ShouldBindJSON(&addFilm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	
	film, err := f.filmService.AddFilm(&addFilm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"data"	 : film,
	})
}

func (f *FilmControllerImpl) GetFilm(ctx *gin.Context) {
	films, err := f.filmService.GetFilm()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data"	 : films,
	})
}

func (f *FilmControllerImpl) GetFilmById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	film, err := f.filmService.GetFilmById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data" 	 : film,
	})
}

func (f *FilmControllerImpl) UpdateFilm(ctx *gin.Context) {
	var film models.Film

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&film); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status"  : false,
			"message" : err,
		})
		return
	}

	count, err := f.filmService.UpdateFilm(id, &film)
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

func (f *FilmControllerImpl) DeleteFilm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	count, err := f.filmService.DeleteFilm(id)
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


