package router

import (
	"api-go/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func URLMapping(r *gin.Engine){

	database.DatabaseConn()

	api:=r.Group("/movies")

	api.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
}