# personas_crud

--Api de personas con CI--
CI deploy with lambda - S3
Drone 0.8 
Personas_crud master/develop
Documentation in process..

## Requirements
Go version >= 1.?.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/personas_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_PERSONAS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `PERSONAS_CRUD__PGUSER`: Usuario de la base de datos
 - `PERSONAS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `PERSONAS_CRUD__PGURLS`: Host de conexión
 - `PERSONAS_CRUD__PGDB`: Nombre de la base de datos
 - `PERSONAS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_PERSONAS_HTTP_PORT=8083 PERSONAS_CRUD__PGUSER=user PERSONAS_CRUD__PGPASS=password PERSONAS_CRUD__PGURLS=localhost PERSONAS_CRUD__PGDB=udistrital_core_db PERSONAS_CRUD__SCHEMA=core_new bee run

## MODELO


