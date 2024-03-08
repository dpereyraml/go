package main

import (
	"context-get/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Ejercicio 1 : Iniciando el proyecto
Debemos crear un repositorio en github.com para poder subir nuestros avances.
Este repositorio es el que vamos a utilizar para llevar lo que realicemos durante las distintas prácticas de Go Web.

Primero debemos clonar el repositorio creado, luego iniciar nuestro proyecto de go con con el comando go mod init.
El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos de productos.
Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.
El archivo para trabajar es el siguiente:

Archivo JSON


*/

func main() {
	/*
		Vamos a levantar un servidor en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.

		Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
		Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
		Crear una ruta /products/:id que nos devuelva un producto por su id.
		Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
	*/
	r := chi.NewRouter()
	h := handler.NewHandlerProducts(nil, 0)

	/*
		// Ejemplo de Version "Manual"
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Pong!"))
		})
	*/

	r.Get("/ping", h.Ping())

	r.Get("/products", h.Products())

	r.Get("/products/{id}", h.ProductByID())

	r.Get("/products/search", h.GetBySearchQuery())

	http.ListenAndServe(":8080", r)

}
