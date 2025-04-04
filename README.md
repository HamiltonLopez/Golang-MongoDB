# Microservicio en Golang con MongoDB

## Descripción

Este proyecto consiste en un microservicio desarrollado en Golang que implementa operaciones CRUD (Crear, Leer, Actualizar, Eliminar) para la entidad `estudiantes`. Utiliza MongoDB como base de datos, está contenerizado con Docker y orquestado mediante Docker Compose.

El objetivo principal es demostrar la implementación de un servicio RESTful siguiendo una arquitectura modular, junto con buenas prácticas de desarrollo, pruebas, contenerización y gestión de bases de datos.

## Características Principales

* **API RESTful:** Endpoints para operaciones CRUD sobre la entidad `estudiantes`.
* **Arquitectura Modular:** Separación de responsabilidades en Controladores, Servicios y Repositorios (DAO). [source: 4]
* **Base de Datos:** MongoDB integrada usando el driver oficial `mongo-go-driver`. [source: 5]
* **Contenerización:** Dockerfile optimizado con Multi-Stage Build. [source: 13, 14]
* **Orquestación:** `docker-compose.yml` para gestionar el microservicio y la base de datos MongoDB. [source: 17]
* **Persistencia de Datos:** Volumen Docker para la base de datos MongoDB, asegurando que los datos no se pierdan al detener/reiniciar contenedores. [source: 6]
* **Pruebas:**
    * Pruebas Unitarias (utilizando `Testify/GoMock`). [source: 9]
    * Pruebas de Integración para validar el flujo CRUD completo. [source: 10]
    * Reporte de Cobertura de Código. [source: 10]

## Tecnologías Utilizadas

* **Lenguaje:** Golang (`1.20-alpine`)
* **Base de Datos:** MongoDB
* **Driver MongoDB:** `mongo-go-driver`
* **Contenerización:** Docker
* **Orquestación:** Docker Compose
* **Testing:** `testing` (nativo de Go), `Testify/GoMock`
* **API Client:** Postman

## Prerrequisitos

* Go (`Versión, ej: 1.18+`) instalado: [https://golang.org/dl/](https://golang.org/dl/)
* Docker instalado: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
* Docker Compose instalado: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)
* Git: [https://git-scm.com/downloads](https://git-scm.com/downloads)
* Postman (opcional, para probar API): [https://www.postman.com/downloads/](https://www.postman.com/downloads/)
