package main

import (
	"fmt"
	"net/http"
	"package-oriented-design/internal/handler"
	"package-oriented-design/internal/repository"
	"package-oriented-design/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {

	// dependencies
	// - repository
	rp := repository.NewProductRepository(nil, 0)
	// - service
	sv := service.NewProductSS(rp) // comente en rp los gets
	// - handler
	hd := handler.NewHandlerProducts(sv)
	// - router
	router := chi.NewRouter()
	router.Route("/products", func(r chi.Router) {
		// POST /tasks
		r.Post("/", hd.CreateProduct())
	})
	/*
		// handler
		h := handler.NewHandlerProducts(nil, 0)
		//router
		r := chi.NewRouter()
		r.Get("/ping", h.Ping())

		r.Get("/products", h.Products())

		r.Get("/products/{id}", h.ProductByID())

		r.Get("/products/search", h.GetBySearchQuery())

		r.Post("/products", h.CreateProduct())
	*/

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}

}
