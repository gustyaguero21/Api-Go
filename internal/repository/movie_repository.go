package repository

import (
	"api-go/cmd/config"
	"api-go/internal/models"
	"database/sql"
)

type MoviesRepository struct{
	DB *sql.DB
}

func (mr *MoviesRepository)CreateTables() error{
	for _,query:=range(config.TableQueries){
		_,err:=mr.DB.Exec(query)
		if err!=nil{
			return  err
		}
	}
	return nil
}

func(mr *MoviesRepository)AddDirectors(directors []models.Director)error{
	for _,director:=range(directors){
		_,err:=mr.DB.Exec(config.AddDirectorsQuery,director.Nombre,director.Apellido,director.Nacionalidad,director.Trayectoria)
		if err!=nil{
			return  err
		}
	}
	return nil
}

func(mr *MoviesRepository)AddActors(actors []models.Actor)error{
	for _,actor:=range(actors){
		_,err:=mr.DB.Exec(config.AddActorsQuery,actor.Nombre,actor.Apellido,actor.Nacionalidad,actor.Trayectoria)
		if err!=nil{
			return  err
		}
	}
	return nil
}
// AGREGAR CHECK DE EXISTENCIAS