package models

import "time"

type Director struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Nacionalidad    string `json:"nacionalidad"`
	Trayectoria time.Time `json:"trayectoria"`
}

type Actor struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Nacionalidad    string `json:"nacionalidad"`
	Trayectoria time.Time `json:"trayectoria"`
}

type Reparto struct {
	ID         int    `json:"id"`
	PeliculaID int    `json:"pelicula_id"`
	ActorID    int    `json:"actor_id"`
	Personaje        string `json:"personaje"`
}

type Genero struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type Calificacion struct {
	ID         int     `json:"id"`
	PeliculaID int     `json:"pelicula_id"`
	Puntuacion float64 `json:"puntuacion"`
}

type Taquilla struct {
	ID         int     `json:"id"`
	PeliculaID int     `json:"pelicula_id"`
	Recaudacion float64 `json:"recaudacion"`
	Pais        string  `json:"pais"`
}

type Presupuesto struct {
	ID         int     `json:"id"`
	PeliculaID int     `json:"pelicula_id"`
	Monto      float64 `json:"monto"`
}

type Productora struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Pais          string `json:"pais"`
	AnioFundacion int    `json:"anio_fundacion"`
}

type Plataforma struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Tipo     string `json:"tipo"`
	SitioWeb string `json:"sitio_web"`
}

type Pelicula struct {
	ID           int     `json:"id"`
	Titulo       string  `json:"titulo"`
	AnioEstreno  int     `json:"anio_estreno"`
	Duracion     int     `json:"duracion"`
	GeneroID     int     `json:"genero_id"`
	DirectorID   int     `json:"director_id"`
	ProductoraID int     `json:"productora_id"`
	PlataformaID int     `json:"plataforma_id"`
}

type Expediente struct{
	ID int   `json:"id"`
	NumeroExpediente string `json:"numero_expediente"`
	PeliculaID int `json:"pelicula_id"`
	FechaCreacion time.Time `json:"fecha_creacion"`
	Descripcion string `json:"descripcion"`
	Estado string `json:"estado"`
}

