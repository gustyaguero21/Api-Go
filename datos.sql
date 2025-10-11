-- SCRIPT: inserts_all_tables.sql
-- Contiene 20 registros por tabla, con ids explícitos para mantener integridad referencial.
USE movies;
-- (Opcional) asegurar que no fallen por referencias existentes:
SET FOREIGN_KEY_CHECKS=0;

-- GENEROS (20)
INSERT INTO Genero (id, nombre) VALUES
(1,'Acción'),
(2,'Drama'),
(3,'Crimen'),
(4,'Ciencia ficción'),
(5,'Thriller'),
(6,'Comedia'),
(7,'Romance'),
(8,'Animación'),
(9,'Fantasía'),
(10,'Terror'),
(11,'Documental'),
(12,'Biografía'),
(13,'Misterio'),
(14,'Aventura'),
(15,'Familiar'),
(16,'Musical'),
(17,'Guerra'),
(18,'Deportivo'),
(19,'Western'),
(20,'Histórico');

-- DIRECTORES (20)
INSERT INTO Director (id, nombre, apellido, nacionalidad, trayectoria) VALUES
(1,'Christopher','Nolan','Reino Unido','1997-01-01'),
(2,'Francis Ford','Coppola','Estados Unidos','1963-01-01'),
(3,'Bong Joon','Ho','Corea del Sur','1993-01-01'),
(4,'Steven','Spielberg','Estados Unidos','1971-01-01'),
(5,'Martin','Scorsese','Estados Unidos','1967-01-01'),
(6,'Quentin','Tarantino','Estados Unidos','1992-01-01'),
(7,'Alfonso','Cuarón','México','1990-01-01'),
(8,'Greta','Gerwig','Estados Unidos','2010-01-01'),
(9,'Denis','Villeneuve','Canadá','1998-01-01'),
(10,'James','Cameron','Canadá','1976-01-01'),
(11,'Peter','Jackson','Nueva Zelanda','1989-01-01'),
(12,'Taika','Waititi','Nueva Zelanda','2004-01-01'),
(13,'Ridley','Scott','Reino Unido','1979-01-01'),
(14,'David','Fincher','Estados Unidos','1992-01-01'),
(15,'Alejandro González','Iñárritu','México','1995-01-01'),
(16,'Kathryn','Bigelow','Estados Unidos','1981-01-01'),
(17,'Wes','Anderson','Estados Unidos','1996-01-01'),
(18,'Guy','Ritchie','Reino Unido','1995-01-01'),
(19,'Patty','Jenkins','Estados Unidos','2003-01-01'),
(20,'Robert','Zemeckis','Estados Unidos','1984-01-01');

-- ACTORES (20)
INSERT INTO Actor (id, nombre, apellido, nacionalidad, trayectoria) VALUES
(1,'Leonardo','DiCaprio','Estados Unidos','1991-01-01'),
(2,'Joseph','Gordon-Levitt','Estados Unidos','2000-01-01'),
(3,'Elliot','Page','Canadá','2005-01-01'),
(4,'Tom','Hardy','Reino Unido','2003-01-01'),
(5,'Marlon','Brando','Estados Unidos','1944-01-01'),
(6,'Al','Pacino','Estados Unidos','1969-01-01'),
(7,'Song','Kang-ho','Corea del Sur','1991-01-01'),
(8,'Lee','Sun-kyun','Corea del Sur','2003-01-01'),
(9,'Brad','Pitt','Estados Unidos','1987-01-01'),
(10,'Angelina','Jolie','Estados Unidos','1995-01-01'),
(11,'Scarlett','Johansson','Estados Unidos','2003-01-01'),
(12,'Robert','Downey Jr.','Estados Unidos','1985-01-01'),
(13,'Chris','Hemsworth','Australia','2009-01-01'),
(14,'Natalie','Portman','Estados Unidos','1994-01-01'),
(15,'Joaquin','Phoenix','Estados Unidos','2000-01-01'),
(16,'Meryl','Streep','Estados Unidos','1977-01-01'),
(17,'Tom','Hanks','Estados Unidos','1980-01-01'),
(18,'Emma','Stone','Estados Unidos','2010-01-01'),
(19,'Denzel','Washington','Estados Unidos','1981-01-01'),
(20,'Margot','Robbie','Australia','2011-01-01');

-- PRODUCTORAS (20)
INSERT INTO Productora (id, nombre, pais, anio_fundacion) VALUES
(1,'Warner Bros.','Estados Unidos',1923),
(2,'Paramount Pictures','Estados Unidos',1912),
(3,'20th Century Studios','Estados Unidos',1935),
(4,'Universal Pictures','Estados Unidos',1912),
(5,'Sony Pictures','Estados Unidos',1987),
(6,'Netflix','Estados Unidos',1997),
(7,'Amazon Studios','Estados Unidos',2010),
(8,'A24','Estados Unidos',2012),
(9,'Lionsgate','Estados Unidos',1997),
(10,'Walt Disney Pictures','Estados Unidos',1923),
(11,'Focus Features','Estados Unidos',2002),
(12,'Miramax','Estados Unidos',1979),
(13,'Columbia Pictures','Estados Unidos',1918),
(14,'DreamWorks Pictures','Estados Unidos',1994),
(15,'Pixar Animation Studios','Estados Unidos',1986),
(16,'Studio Ghibli','Japón',1985),
(17,'Annapurna Pictures','Estados Unidos',2011),
(18,'Blumhouse Productions','Estados Unidos',2000),
(19,'Metro-Goldwyn-Mayer (MGM)','Estados Unidos',1924),
(20,'Illumination','Estados Unidos',2007);

-- PLATAFORMAS (20)
INSERT INTO Plataforma (id, nombre, tipo, sitio_web) VALUES
(1,'Netflix','Streaming','https://www.netflix.com'),
(2,'Amazon Prime Video','Streaming','https://www.primevideo.com'),
(3,'Disney+','Streaming','https://www.disneyplus.com'),
(4,'HBO Max','Streaming','https://www.hbomax.com'),
(5,'Hulu','Streaming','https://www.hulu.com'),
(6,'Apple TV+','Streaming','https://tv.apple.com'),
(7,'Paramount+','Streaming','https://www.paramountplus.com'),
(8,'Peacock','Streaming','https://www.peacocktv.com'),
(9,'Crunchyroll','Streaming','https://www.crunchyroll.com'),
(10,'MUBI','Streaming','https://mubi.com'),
(11,'YouTube Movies','VOD','https://www.youtube.com/movies'),
(12,'Google Play Movies','VOD','https://play.google.com/store/movies'),
(13,'iTunes','VOD','https://www.apple.com/itunes'),
(14,'Filmin','Streaming','https://www.filmin.es'),
(15,'Cine Colombia','Cine','https://www.cinecolombia.com'),
(16,'Rakuten TV','Streaming','https://rakuten.tv'),
(17,'Starz','Streaming','https://www.starz.com'),
(18,'Canal+','Cine/TV','https://www.canalplus.com'),
(19,'Sky','TV de pago','https://www.sky.com'),
(20,'Tubi','Streaming','https://tubitv.com');

-- PELICULAS (20)
-- Cada pelicula referencia genero_id, director_id, productora_id, plataforma_id
INSERT INTO Pelicula (id, titulo, anio_estreno, duracion, genero_id, director_id, productora_id, plataforma_id) VALUES
(1,'Inception',2010,148,4,1,1,1),
(2,'The Godfather',1972,175,3,2,2,15),
(3,'Parasite',2019,132,2,3,8,1),
(4,'Jurassic Park',1993,127,14,4,4,15),
(5,'The Wolf of Wall Street',2013,180,2,5,9,2),
(6,'Pulp Fiction',1994,154,6,6,12,2),
(7,'Gravity',2013,91,4,7,3,6),
(8,'Little Women',2019,135,2,8,11,3),
(9,'Dune',2021,155,4,9,3,3),
(10,'Avatar',2009,162,9,10,4,3),
(11,'The Lord of the Rings: The Fellowship of the Ring',2001,178,9,11,13,15),
(12,'Thor: Love and Thunder',2022,119,1,12,5,2),
(13,'Blade Runner 2049',2017,164,4,13,14,4),
(14,'Babel',2006,143,2,15,11,1),
(15,'The Hurt Locker',2008,131,17,16,9,4),
(16,'The Grand Budapest Hotel',2014,99,6,17,7,10),
(17,'Sherlock Holmes',2009,128,13,18,5,7),
(18,'Wonder Woman',2017,141,1,19,3,3),
(19,'Forrest Gump',1994,142,2,20,2,11),
(20,'Back to the Future',1985,116,14,20,14,15);

-- REPARTO (20) -- asigno 1-2 actores por pelicula (mínimo 20 registros en total)
INSERT INTO Reparto (id, pelicula_id, actor_id, personaje) VALUES
(1,1,1,'Dom Cobb'),
(2,1,2,'Arthur'),
(3,2,5,'Don Vito Corleone'),
(4,2,6,'Michael Corleone'),
(5,3,7,'Kim Ki-taek'),
(6,3,8,'Park Dong-ik'),
(7,4,17,'Dr. Alan Grant'),
(8,5,1,'Jordan Belfort'),
(9,6,12,'Jules Winnfield'),
(10,6,15,'Vincent Vega'),
(11,7,14,'Dr. Ryan Stone'),
(12,9,13,'Paul Atreides'),
(13,10,9,'Jake Sully'),
(14,11,11,'Frodo Baggins'),
(15,12,13,'Thor'),
(16,13,14,'K'),
(17,14,3,'Ishmael'),
(18,15,19,'Sergente Eldridge'),
(19,16,18,'Agatha'),
(20,20,20,'Marty McFly');

-- CALIFICACIONES (20) -- puntuaciones tipo IMDB/Rotten (ejemplos)
INSERT INTO Calificacion (id, pelicula_id, puntuacion) VALUES
(1,1,8.8),
(2,2,9.2),
(3,3,8.6),
(4,4,8.1),
(5,5,8.2),
(6,6,8.9),
(7,7,7.7),
(8,8,7.8),
(9,9,8.0),
(10,10,7.8),
(11,11,8.8),
(12,12,6.4),
(13,13,8.0),
(14,14,7.5),
(15,15,7.6),
(16,16,8.1),
(17,17,7.6),
(18,18,7.4),
(19,19,8.8),
(20,20,8.5);

-- TAQUILLA (20) -- recaudaciones aproximadas (USD)
INSERT INTO Taquilla (id, pelicula_id, recaudacion, pais) VALUES
(1,1,829895144.00,'USA'),
(2,2,246120974.00,'USA'),
(3,3,258795251.00,'Corea del Sur'),
(4,4,1020000000.00,'USA'),
(5,5,392000000.00,'USA'),
(6,6,213928762.00,'USA'),
(7,7,723192705.00,'USA'),
(8,8,218000000.00,'USA'),
(9,9,402000000.00,'USA'),
(10,10,2787965087.00,'USA'),
(11,11,887887904.00,'Worldwide'),
(12,12,760000000.00,'USA'),
(13,13,259239658.00,'USA'),
(14,14,30500000.00,'USA'),
(15,15,50000000.00,'USA'),
(16,16,172900000.00,'USA'),
(17,17,524000000.00,'USA'),
(18,18,821847012.00,'USA'),
(19,19,678226465.00,'USA'),
(20,20,388000000.00,'USA');

-- PRESUPUESTOS (20) -- montos aproximados (USD)
INSERT INTO Presupuesto (id, pelicula_id, monto) VALUES
(1,1,160000000.00),
(2,2,6000000.00),
(3,3,11400000.00),
(4,4,630000000.00),
(5,5,100000000.00),
(6,6,8000000.00),
(7,7,100000000.00),
(8,8,40000000.00),
(9,9,165000000.00),
(10,10,237000000.00),
(11,11,93000000.00),
(12,12,250000000.00),
(13,13,150000000.00),
(14,14,25000000.00),
(15,15,15000000.00),
(16,16,25000000.00),
(17,17,90000000.00),
(18,18,149000000.00),
(19,19,55000000.00),
(20,20,19000000.00);

-- EXPEDIENTES (20)
-- Para varios expedientes uso estado 'proximo_estreno' para títulos que en 2025 aparecen como próximos o recientes; otros quedan 'en_proceso' o 'estrenada'/'archivada'
INSERT INTO Expedientes (id, numero_expediente, pelicula_id, fecha_creacion, descripcion, estado) VALUES
(1,'EXP-2025-001',9,'2025-09-30 09:00:00','Expediente para "Dune" — revisión de distribución y licencia streaming.','estrenada'),
(2,'EXP-2025-002',12,'2025-10-01 10:00:00','Expediente relacionado con la licencia internacional de "Thor: Love and Thunder".','estrenada'),
(3,'EXP-2025-003',3,'2025-10-02 11:15:00','Expediente de "Parasite" — festival y catálogo.','estrenada'),
(4,'EXP-2025-004',1,'2025-10-03 08:30:00','Expediente técnico de restauración digital de "Inception".','estrenada'),
(5,'EXP-2025-005',4,'2025-10-04 14:20:00','Expediente de contratación de proyecciones educativas para "Jurassic Park".','estrenada'),
(6,'EXP-2025-006',5,'2025-10-05 09:45:00','Expediente: evento especial con "The Wolf of Wall Street".','estrenada'),
(7,'EXP-2025-007',16,'2025-09-15 12:00:00','Expediente "The Grand Budapest Hotel" — permiso de exhibición.','estrenada'),
(8,'EXP-2025-008',20,'2025-09-20 13:30:00','Expediente archivo retro: "Back to the Future".','archivada'),
(9,'EXP-2025-009',14,'2025-10-06 16:00:00','Expediente para reedición y remaster de "Babel".','en_proceso'),
(10,'EXP-2025-010',17,'2025-10-07 17:10:00','Expediente: derechos para remake de "Sherlock Holmes".','en_proceso'),
-- Ahora agrego expedientes para próximos estrenos 2025 (ejemplos públicos: Tron: Ares, Roofman, After the Hunt, The Smashing Machine, Gabby\'s Dollhouse)
(11,'EXP-2025-011',21,'2025-10-01 08:00:00','Expediente: próximo estreno "TRON: Ares" — gestión de prensa y pases de prensa.','proximo_estreno'),
(12,'EXP-2025-012',22,'2025-10-02 08:30:00','Expediente: próximo estreno "Roofman" — licencias y distribución regional.','proximo_estreno'),
(13,'EXP-2025-013',23,'2025-10-03 09:00:00','Expediente: "After the Hunt" — coordinación con cines y plataformas.','proximo_estreno'),
(14,'EXP-2025-014',24,'2025-10-04 09:30:00','Expediente: "The Smashing Machine" — prensa y aprobación de materiales.','proximo_estreno'),
(15,'EXP-2025-015',25,'2025-10-05 10:00:00','Expediente: "Gabby''s Dollhouse: The Movie" — control de materiales infantiles.','proximo_estreno'),
(16,'EXP-2025-016',6,'2025-10-06 10:30:00','Expediente sobre restauración y ficha técnica de "Pulp Fiction".','estrenada'),
(17,'EXP-2025-017',8,'2025-10-07 11:00:00','Expediente: derechos de "Little Women" para exhibición escolar.','en_proceso'),
(18,'EXP-2025-018',15,'2025-10-08 11:30:00','Expediente: "The Hurt Locker" — permiso para proyección.','archivada'),
(19,'EXP-2025-019',13,'2025-10-09 12:00:00','Expediente: "Blade Runner 2049" — licencia de remake/serie.','en_proceso'),
(20,'EXP-2025-020',11,'2025-10-10 12:30:00','Expediente: "The Lord of the Rings" — coordinación de maratón.','en_proceso');

-- Observación:
-- Para los expedientes que referencian peliculas con id > 20 (ej. 21..25), a continuación insertaré esas peliculas "próximas" en la tabla Pelicula.
-- (Dejo FOREIGN_KEY_CHECKS activo más abajo después de crear esas películas)

-- PELICULAS ADICIONALES (próximos estrenos) -> ids 21..25
INSERT INTO Pelicula (id, titulo, anio_estreno, duracion, genero_id, director_id, productora_id, plataforma_id) VALUES
(21,'TRON: Ares',2025,120,4,9,1,3),
(22,'Roofman',2025,110,14,12,2,2),
(23,'After the Hunt',2025,105,5,8,3,1),
(24,'The Smashing Machine',2025,115,1,1,17,2),
(25,'Gabby''s Dollhouse: The Movie',2025,85,15,14,6,6);

-- REPARTO para las peliculas 21..25 (añado 5 registros más, total Reparto 25 registros)
INSERT INTO Reparto (id, pelicula_id, actor_id, personaje) VALUES
(21,21,13,'Ares (protagonista)'),
(22,22,9,'Protagonista'),
(23,23,18,'Personaje principal'),
(24,24,1,'Cameo'),
(25,25,19,'Voz de personaje');

-- CALIFICACIONES adicionales (21..25)
INSERT INTO Calificacion (id, pelicula_id, puntuacion) VALUES
(21,21,0.0),
(22,22,0.0),
(23,23,0.0),
(24,24,0.0),
(25,25,0.0);

-- TAQUILLA adicionales (21..25) -- inicialmente 0.00 porque son próximos estrenos
INSERT INTO Taquilla (id, pelicula_id, recaudacion, pais) VALUES
(21,21,0.00,'TBD'),
(22,22,0.00,'TBD'),
(23,23,0.00,'TBD'),
(24,24,0.00,'TBD'),
(25,25,0.00,'TBD');

-- PRESUPUESTOS adicionales (21..25) -- estimados
INSERT INTO Presupuesto (id, pelicula_id, monto) VALUES
(21,21,150000000.00),
(22,22,40000000.00),
(23,23,30000000.00),
(24,24,50000000.00),
(25,25,15000000.00);

-- Reactivar controles FK
SET FOREIGN_KEY_CHECKS=1;

-- FIN DEL SCRIPT
