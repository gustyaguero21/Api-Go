package router

import (
	"api-go/internal/database"
	"api-go/internal/models"
	"api-go/internal/repository"
	"api-go/internal/services"
	"fmt"
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
	mr:=repository.MoviesRepository{DB: db}
	ms:=services.MovieService{MR: mr}


	if tableErr:=mr.CreateTables();tableErr!=nil{
		log.Fatal(tableErr)
	}
	
	directors:= []models.Director{
    {Nombre: "", Apellido: "Nolan", Nacionalidad: "Reino Unido", Trayectoria: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)},
    {Nombre: "Quentin", Apellido: "Tarantino", Nacionalidad: "EEUU", Trayectoria: time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)},
    {Nombre: "Martin", Apellido: "Scorsese", Nacionalidad: "EEUU", Trayectoria: time.Date(1967, 1, 1, 0, 0, 0, 0, time.UTC)},
    {Nombre: "James", Apellido: "Cameron", Nacionalidad: "Canadá", Trayectoria: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)},
}


	directores,directErr:=ms.CreateDirectors(&gin.Context{},directors)
	if directErr!=nil{
		log.Fatal(directErr)
	}

	fmt.Println(directores)

	api:=r.Group("/movies")

	api.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
}