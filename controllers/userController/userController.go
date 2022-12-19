package userController

import "github.com/gin-gonic/gin"

type UserController interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}