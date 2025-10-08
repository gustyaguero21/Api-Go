package services

import (
	"api-go/internal/models"
	"api-go/internal/repository"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MovieService struct{
	MR repository.MoviesRepository
}

func(ms *MovieService)CreateDirectors(ctx *gin.Context,directors []models.Director)([]models.Director,error){
	for _,dato:=range(directors){
		if dato.Nombre=="" || dato.Apellido=="" || dato.Nacionalidad == "" {
			return nil,fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
		}
		validData:=strconv.Itoa(time.Time.Year(dato.Trayectoria))
		if validData==""{
			return nil,fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
		}
	}

	if createDirectorErr:=ms.MR.AddDirectors(directors);createDirectorErr!=nil{
		return nil,fmt.Errorf("error al crear un nuevo director. Error: %s",createDirectorErr)
	}

	return directors,nil
}

func(ms *MovieService)CreateActors(ctx *gin.Context,actors []models.Actor)([]models.Actor,error){
	for _,dato:=range(actors){
		if dato.Nombre=="" || dato.Apellido=="" || dato.Nacionalidad == "" {
			return nil,fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
		}
		validData:=strconv.Itoa(time.Time.Year(dato.Trayectoria))
		if validData==""{
			return nil,fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
		}
	}

	if createActorsErr:=ms.MR.AddActors(actors);createActorsErr!=nil{
		return nil,fmt.Errorf("error al crear un nuevo director. Error: %s",createActorsErr)
	}

	return actors,nil
}

