package services

import (
	"api-go/internal/models"
	"api-go/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MovieService struct{
	MR repository.MoviesRepository
}

func(ms *MovieService)AddDirectorService(ctx *gin.Context,director models.Director)(models.Director,error){
	if director.Nombre == "" || director.Apellido == "" || director.Nacionalidad == "" || director.Trayectoria.IsZero() {
    	return models.Director{}, fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
	}

	if createDirectorErr:=ms.MR.AddDirectorRepository(director);createDirectorErr!=nil{
		return models.Director{},fmt.Errorf("error al crear un nuevo director. Error: %s",createDirectorErr)
	}

	return director,nil
}

func(ms *MovieService)AddActorService(ctx *gin.Context,actors models.Actor)(models.Actor,error){
	if actors.Nombre=="" || actors.Apellido=="" || actors.Nacionalidad == "" || actors.Trayectoria.IsZero() {
		return models.Actor{},fmt.Errorf("algunos de los parametros estan vacios o son invalidos")
	}

	if createActorErr:=ms.MR.AddActorRepository(actors);createActorErr!=nil{
		return models.Actor{},fmt.Errorf("error al crear un nuevo actor. Error: %s",createActorErr)
	}

	return actors,nil
}

