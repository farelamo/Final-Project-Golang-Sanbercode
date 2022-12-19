package studioController

import "github.com/gin-gonic/gin"

type StudioController interface {
	AddStudio(ctx *gin.Context)
	GetStudio(ctx *gin.Context)
	GetStudioById(ctx *gin.Context)
	UpdateStudio(ctx *gin.Context)
	DeleteStudio(ctx *gin.Context)
}