# personas_crud
API de gestión de personas

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `personas_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/personas_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `API_PERSONAS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `PERSONAS_CRUD__PGUSER`: Usuario de la base de datos
 - `PERSONAS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `PERSONAS_CRUD__PGURLS`: Host de conexión
 - `PERSONAS_CRUD__PGDB`: Nombre de la base de datos
 - `PERSONAS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_PERSONAS_HTTP_PORT=8083 PERSONAS_CRUD__PGUSER=user PERSONAS_CRUD__PGPASS=password PERSONAS_CRUD__PGURLS=localhost PERSONAS_CRUD__PGDB=bd PERSONAS_CRUD__SCHEMA=schema_new bee run

## MODELO
![image](https://github.com/udistrital/personas_crud/blob/develop/modelo_personas_crud.png).
