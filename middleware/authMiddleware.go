package middleware

import (
	"errors"
	"net/http"
	"FinalProject/config"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()

	if !ok {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("authentication not found"))
		return
	}

	sql := `SELECT * FROM users WHERE username = $1 AND password = $2;`
	res, err := config.DB.Exec(sql, username, password)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if count == 0 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("username atau password salah"))
		return
	}

	ctx.Next()
}
