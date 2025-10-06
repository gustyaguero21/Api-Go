package router

import (
	"api-go/internal/database"
	"api-go/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func URLMapping(r *gin.Engine){

	db,err:=database.DatabaseConn()
	if err!=nil{
		log.Fatal(err)
	}
	mr:=database.MoviesRepository{DB: db}
	
	if tableErr:=mr.CreateTables();tableErr!=nil{
		log.Fatal(tableErr)
	}
	actores := []models.Actor{
		{Nombre: "Tom", Apellido: "Hanks", Nacionalidad: "EEUU", Trayectoria: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Nombre: "Leonardo", Apellido: "DiCaprio", Nacionalidad: "EEUU", Trayectoria: time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Nombre: "Meryl", Apellido: "Streep", Nacionalidad: "EEUU", Trayectoria: time.Date(1975, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Nombre: "Robert", Apellido: "De Niro", Nacionalidad: "EEUU", Trayectoria: time.Date(1965, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Nombre: "Scarlett", Apellido: "Johansson", Nacionalidad: "EEUU", Trayectoria: time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC)},
}	

	if addActorErr:=mr.AddActors(actores);addActorErr!=nil{
		log.Fatal(addActorErr)
	}

	api:=r.Group("/movies")

	api.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
}