# Go-Gin Base
Proyecto cuyo fin es tener lo básico para arrancar una api para iniciar otros proyectos.

- Cuenta con un docker-compose que arranca una base de datos mysql.
## Requisitos previos
- Es necesario tener Go instalado
- Es necesario tener una base de datos funcionando
## Instalar dependencias:

- ``Go mod tidy``

## Arrancar el servidor

- ``go run ./cmd/api/main.go``

## Personalización

- Cambiar los ajustes por defecto en el fichero ``./cmd/api/bootstrap/bootstrap.go``
- En caso de usar una base de datos que no sea mysql, descargar los paquetes necesarios y reescribir la conexión usando
el driver correspondiente.