package router

import (
	"api-go/internal/controller"
	"api-go/internal/database"
	"api-go/internal/repository"
	"api-go/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func URLMapping(r *gin.Engine){

	db,err:=database.DatabaseConn()
	if err!=nil{
		log.Fatal(err)
	}

	database.CreateTables(db)


	mr:=repository.MoviesRepository{DB: db}
	ms:=services.MovieService{MR: mr}
	mc:=controller.MovieContoller{MS:ms}

	api:=r.Group("/movies")

	api.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})

	api.POST("/add-director",mc.AddDirectorController)
	api.POST("/add-actor",mc.AddActorController)
}