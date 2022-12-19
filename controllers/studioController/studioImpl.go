package studioController

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	. "FinalProject/services/studioService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudioControllerImpl struct {
	studioService StudioService
}

func NewStudioController(studioService StudioService) StudioController {
	return &StudioControllerImpl{
		studioService: studioService,
	}
}

func (s *StudioControllerImpl) AddStudio(ctx *gin.Context){
	var addStudio models.AddStudio

	if err := ctx.ShouldBindJSON(&addStudio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	
	studio, err := s.studioService.AddStudio(&addStudio)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status": true,
		"data"  : studio,
	})
}

func (s *StudioControllerImpl) GetStudio(ctx *gin.Context) {
	studios, err := s.studioService.GetStudio()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status": true,
		"data"  : studios,
	})
}

func (s *StudioControllerImpl) GetStudioById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": false,
			"message": "Invalid request",
		})
		return
	}

	studio, err := s.studioService.GetStudioById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status": false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status": true,
		"data"  : studio,
	})
}

func (s *StudioControllerImpl) UpdateStudio(ctx *gin.Context) {
	var studio models.Studio

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&studio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	count, err := s.studioService.UpdateStudio(id, &studio)
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

func (s *StudioControllerImpl) DeleteStudio(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	count, err := s.studioService.DeleteStudio(id)
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


