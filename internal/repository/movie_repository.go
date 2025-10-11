package repository

import (
	"api-go/cmd/config"
	"api-go/internal/models"
	"database/sql"
	"fmt"
)

type MoviesRepository struct{
	DB *sql.DB
}


func (mr *MoviesRepository) AddDirectorRepository(director models.Director) error {
	if mr.checkExistence("Director", director.Nombre, director.Apellido) {
		return fmt.Errorf("director already exists")
	}

	_, err := mr.DB.Exec(
		config.AddDirectorQuery,
		director.Nombre,
		director.Apellido,
		director.Nacionalidad,
		director.Trayectoria,
	)
	return err
}


func (mr *MoviesRepository) AddActorRepository(actor models.Actor) error {
	if mr.checkExistence("Actor", actor.Nombre, actor.Apellido) {
		return fmt.Errorf("actor already exists")
	}
	_, err := mr.DB.Exec(
			config.AddActorQuery,
			actor.Nombre,
			actor.Apellido,
			actor.Nacionalidad,
			actor.Trayectoria,
		)
		
	return err
}

func (mr *MoviesRepository)AddCastRepository(reparto models.Reparto)error{
	_,err:=mr.DB.Exec(config.AddCastQuery,reparto.PeliculaID,reparto.ActorID,reparto.Personaje)
	if err!=nil{
		return err
	}
	return nil
}



func (mr *MoviesRepository) checkExistence(tableName, nombre, apellido string) bool {
	query := fmt.Sprintf(config.CheckExistences, tableName)
	var id int
	err := mr.DB.QueryRow(query, nombre, apellido).Scan(&id)
	if err == sql.ErrNoRows {
		return false 
	}
	if err != nil {
		fmt.Println("error checking existences:", err)
		return false
	}
	return true 
}
