package config

// database queries
const (
	CheckDBQuery  = "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"
	CreateDBQuery = "CREATE DATABASE IF NOT EXISTS %s"
	UseDBQuery    = "USE %s"
)

var TableQueries = []string{
    `CREATE TABLE IF NOT EXISTS Director (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        apellido VARCHAR(100) NOT NULL,
        nacionalidad VARCHAR(50),
        trayectoria DATE
    );`,

    `CREATE TABLE IF NOT EXISTS Actor (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        apellido VARCHAR(100) NOT NULL,
        nacionalidad VARCHAR(50),
        trayectoria DATE
    );`,

    `CREATE TABLE IF NOT EXISTS Genero (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(50) NOT NULL
    );`,

    `CREATE TABLE IF NOT EXISTS Productora (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        pais VARCHAR(50),
        anio_fundacion INT
    );`,

    `CREATE TABLE IF NOT EXISTS Plataforma (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        tipo VARCHAR(50),
        sitio_web VARCHAR(150)
    );`,

    `CREATE TABLE IF NOT EXISTS Pelicula (
        id INT AUTO_INCREMENT PRIMARY KEY,
        titulo VARCHAR(150) NOT NULL,
        anio_estreno INT,
        duracion INT,
        genero_id INT,
        director_id INT,
        productora_id INT,
        plataforma_id INT,
        FOREIGN KEY (genero_id) REFERENCES Genero(id),
        FOREIGN KEY (director_id) REFERENCES Director(id),
        FOREIGN KEY (productora_id) REFERENCES Productora(id),
        FOREIGN KEY (plataforma_id) REFERENCES Plataforma(id)
    );`,

    `CREATE TABLE IF NOT EXISTS Reparto (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pelicula_id INT NOT NULL,
        actor_id INT NOT NULL,
        personaje VARCHAR(100),
        FOREIGN KEY (pelicula_id) REFERENCES Pelicula(id),
        FOREIGN KEY (actor_id) REFERENCES Actor(id)
    );`,

    `CREATE TABLE IF NOT EXISTS Calificacion (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pelicula_id INT NOT NULL,
        puntuacion DECIMAL(3,1),
        FOREIGN KEY (pelicula_id) REFERENCES Pelicula(id)
    );`,

    `CREATE TABLE IF NOT EXISTS Taquilla (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pelicula_id INT NOT NULL,
        recaudacion DECIMAL(15,2),
        pais VARCHAR(50),
        FOREIGN KEY (pelicula_id) REFERENCES Pelicula(id)
    );`,

    `CREATE TABLE IF NOT EXISTS Presupuesto (
        id INT AUTO_INCREMENT PRIMARY KEY,
        pelicula_id INT NOT NULL,
        monto DECIMAL(15,2),
        FOREIGN KEY (pelicula_id) REFERENCES Pelicula(id)
    );`,

    `CREATE TABLE IF NOT EXISTS Expedientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    numero_expediente VARCHAR(50) NOT NULL,
    pelicula_id INT NOT NULL,
    fecha_creacion DATETIME NOT NULL,
    descripcion TEXT,
    estado ENUM('en_proceso', 'proximo_estreno', 'estrenada', 'archivada'),
    FOREIGN KEY (pelicula_id) REFERENCES Pelicula(id)
    );`,

}

var(
    AddDirectorsQuery=`INSERT INTO Director (nombre,apellido,nacionalidad,trayectoria) VALUES (?,?,?,?);`
    AddActorsQuery=`INSERT INTO Actor (nombre,apellido,nacionalidad,trayectoria) VALUES (?,?,?,?);`
    CheckExistences=`SELECT * FROM %s WHERE nombre = ? LIMIT 1;`
)