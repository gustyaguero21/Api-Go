package router

import (
	"api-go/internal/database"
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


	// mr:=repository.MoviesRepository{DB: db}
	// ms:=services.MovieService{MR: mr}

	api:=r.Group("/movies")

	api.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
}