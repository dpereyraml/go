package handler

import (
	"context-get/internal"
	"context-get/internal/repository"
	"context-get/platform/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// DefaultProduct is a handler with a map of users as data
type DefaultProduct struct {
	product map[int]internal.Product
}

// NewHandler returns a new DefaultProduct
func NewHandlerProducts(product map[int]internal.Product, lastID int) *DefaultProduct {
	// default values
	defaultProduct := make(map[int]internal.Product)

	if product != nil {
		defaultProduct = product
	}

	return &DefaultProduct{
		product: defaultProduct,
	}
}

// Ping returns a handler for the Ping / route
func (h *DefaultProduct) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		/*
			// validate token
			token := r.Header.Get("Authorization")
			if token != "12345" {
				web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
				return
			}
		*/
		// w.Write([]byte("Pong!"))
		web.ResponseJSON(w, http.StatusOK, map[string]any{"response": "Pong!"})
	}
}

// Products returns a handler for the Products / Products
func (h *DefaultProduct) Products() http.HandlerFunc {
	productsData := repository.LoadProductsFromFile()
	// productsJSON := repository.ToJSON(productsData)
	// quitar el ToJson?
	// falta validaciones de errores
	return func(w http.ResponseWriter, r *http.Request) {
		web.ResponseJSON(w, http.StatusOK, productsData)
	}
}

func (h *DefaultProduct) FindPricesGreaterThan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtainer el ID del producto de la URL
		productId := chi.URLParam(r, "search")

		// Convert the ID to an integer
		idInt, err := strconv.Atoi(productId)
		if err != nil {
			fmt.Println("Error al convertir ID a entero:", err)
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Load the products from the file
		productsData := repository.LoadProductsFromFile()

		// Search the product by ID
		var product repository.Product
		for _, p := range productsData {
			if p.ID == idInt {
				product = p
				break
			}
		}

		// Valid if the product was found
		if product.ID == 0 {
			// Product not found
			web.ResponseJSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto no encontrado"})
			return
		}

		// Return the product as JSON
		web.ResponseJSON(w, http.StatusOK, product)
	}
}
func (h *DefaultProduct) ProductByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtainer el ID del producto de la URL
		productId := chi.URLParam(r, "id")

		// Convert the ID to an integer
		idInt, err := strconv.Atoi(productId)
		if err != nil {
			fmt.Println("Error al convertir ID a entero:", err)
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Load the products from the file
		productsData := repository.LoadProductsFromFile()

		// Search the product by ID
		var product repository.Product
		for _, p := range productsData {
			if p.ID == idInt {
				product = p
				break
			}
		}

		// Valid if the product was found
		if product.ID == 0 {
			// Product not found
			web.ResponseJSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto no encontrado"})
			return
		}

		// Return the product as JSON
		web.ResponseJSON(w, http.StatusOK, product)
	}
}

func (h *DefaultProduct) GetBySearchQuery() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the priceGt from the query params
		priceParam := r.URL.Query().Get("priceGt")

		// conert the priceGt to an float64
		priceGt, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			fmt.Println("Error al convertir priceGt a float64:", err)
			http.Error(w, "priceGt inválido", http.StatusBadRequest)
			return
		}
		// Load the products from the file
		productsData := repository.LoadProductsFromFile()

		// Search the product by price higher than priceGt
		var products []repository.Product

		for _, p := range productsData {
			if p.Price > priceGt {
				products = append(products, p)
			}
		}

		// Valid if the product was found
		if products == nil {
			// Product not found
			web.ResponseJSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto/s no encontrado/s"})
			return
		}

		// Return the product as JSON
		web.ResponseJSON(w, http.StatusOK, products)
	}
}
