package router

import (
	"api-go/internal/controller"
	"api-go/internal/database"
	"api-go/internal/repository"
	"api-go/internal/services"
	"log"

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

	api:=r.Group("/add-data")

	api.POST("/add-director",mc.AddDirectorController)
	api.POST("/add-actor",mc.AddActorController)
	api.POST("/add-cast",mc.AddCastController)

	
}