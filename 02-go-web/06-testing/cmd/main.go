package main

import (
	"fmt"
	"net/http"
	"testing-integration/internal/handler"

	"github.com/go-chi/chi/v5"
)

/*
	Ejercicio 1: Añadir un producto
	En esta ocasión vamos a añadir un producto al slice cargado enmemoria.
	Dentro de la ruta /products añadimos el método POST,
	al  cual  vamos  a  enviar  en  el  cuerpo  de  la  request  el  nuevoproducto.
	El mismo tiene ciertas restricciones, conozcámoslas:
	1. No  es  necesario  pasar  el  Id,
	al  momento  de  añadirlo  sedebe inferir del estado de la lista de productos,
	vericandoque no se repitan ya que debe ser un campo único.
	2. Ningún  dato  puede  estar  vacío,  exceptuando
	is_published (vacío indica un valor false).
	3. El campo code_value debe ser único para cada producto.
	4. Los tipos de datos deben coincidir con los denidos en elplanteo del problema.5.
	La  fecha  de  vencimiento  debe  tener  el  formato:
	XX/XX/XXXX,  además  debemos  vericar  que  día,  mes  yaño sean valores válidos.
	Recordá: si una consulta está mal formulada por parte del cliente,
	el status code cae en los 4XX.

	Ejercicio 2: Traer el producto
	Realiza una consulta a un método GET con el id del producto reciénañadido,
	tené  en  cuenta  que  la  lista  de  producto  se  encuentracargada en la memoria,
	si terminás la ejecución del programa esteproducto no estará en la próxima ejecución
*/

func main() {
	/*
		 Si no existiera el "constructuror"
		 h := handler.NewHandlerProducts(nil, 0)
		 se podría hacer de la siguiente manera:
		 h := &handler.DefaultProduct{
			 product: map[int]internal.Product{},
			 lastID:  0,
			}
	*/

	// handler
	h := handler.NewHandlerProducts(nil, 0)
	//router
	r := chi.NewRouter()
	r.Get("/ping", h.Ping())

	r.Get("/products", h.Products())

	r.Get("/products/{id}", h.ProductByID())

	r.Get("/products/search", h.GetBySearchQuery())

	r.Post("/products", h.CreateProduct())

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		fmt.Println(err)
		return
	}

}
