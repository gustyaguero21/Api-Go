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


func (mr *MoviesRepository) AddDirectors(directors []models.Director) error {
	for _, director := range directors {
		if !mr.checkExistence("Director", director.Nombre) {
			_, err := mr.DB.Exec(
				config.AddDirectorsQuery,
				director.Nombre,
				director.Apellido,
				director.Nacionalidad,
				director.Trayectoria,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mr *MoviesRepository) AddActors(actors []models.Actor) error {
	for _, actor := range actors {
		if !mr.checkExistence("Actor", actor.Nombre) {
			_, err := mr.DB.Exec(
				config.AddActorsQuery,
				actor.Nombre,
				actor.Apellido,
				actor.Nacionalidad,
				actor.Trayectoria,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mr *MoviesRepository) checkExistence(tableName, name string) bool {
	query := fmt.Sprintf(config.CheckExistences, tableName)

	var exists string
	err := mr.DB.QueryRow(query, name).Scan(&exists)

	if err == sql.ErrNoRows {
		return false 
	}
	if err != nil {
		return false
	}
	return true
}
