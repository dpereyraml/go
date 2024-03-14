package main

import (
	"fmt"
	"net/http"
	"put-patch-delete/internal/handler"
	"put-patch-delete/internal/repository"
	"put-patch-delete/internal/service"

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
	router.Get("/ping", hd.Ping())
	router.Route("/products", func(r chi.Router) {
		// POST /tasks
		r.Post("/", hd.CreateProduct())
		r.Get("/", hd.GetProducts())
		r.Get("/{id}", hd.ProductByID())
		r.Get("/search", hd.GetProductsBySearchQuery())
		r.Put("/{id}", hd.UpdateProduct())
		r.Patch("/{id}", hd.UpdateProductPartial())
		r.Delete("/{id}", hd.DeleteProduct())

	})

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}

}
