package handler

import (
	"context-get-post/internal"
	"context-get-post/internal/repository"
	"context-get-post/platform/web"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// DefaultProduct is a struct that contains the handlers for the products
type DefaultProduct struct {
	product map[int]internal.Product
	lastID  int // lastID is the last ID used for a product used to generate new IDs
}

// Struct for JSON Products request body
type ProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// struct for return Product format JSON
type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// NewHandler returns a new DefaultProduct
func NewHandlerProducts(product map[int]internal.Product, lastID int) *DefaultProduct {
	// default values
	defaultProduct := make(map[int]internal.Product)
	defaultLastID := 0

	if product != nil {
		defaultProduct = product
	}
	if lastID != 0 {
		defaultLastID = lastID
	}

	return &DefaultProduct{
		product: defaultProduct,
		lastID:  defaultLastID,
	}
}

// Methods GET
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

// Methods POST
func (h *DefaultProduct) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// validate token
		token := r.Header.Get("Authorization")
		if token != "12345" {
			web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}
		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body 1"})
			return
		}

		// - validate request body parse into map[string]any
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body unm"})
			return
		}

		// - validate request body
		// ... se debe hacer para cada campo del body ...
		if _, ok := bodyMap["name"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body name is required"})
			return
		}

		// - read into bytes
		var bodyProductRequest ProductRequest

		// Forma " Automática " - Usingn web package
		if err := json.Unmarshal(bytes, &bodyProductRequest); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid request body 3",
			})
			return
		}

		/*
			Forma " A Mano "
			if err := json.NewDecoder(r.Body).Decode(&bodyProductRequest); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)

				// Dos formas distintas de devolver el mensaje dependiendo de lo que necesitemos
				// w.Write([]byte(`{"message": "invalid request body"}`))
				json.NewEncoder(w).Encode(map[string]any{"message": "invalid request body"})
				return
			}
		*/

		// process
		// - get the lastID
		var idInt int
		productsData := repository.LoadProductsFromFile()
		for _, p := range productsData {
			if p.ID > idInt {
				idInt = p.ID
			}
		}

		// Valid if the product was found
		if idInt != 0 {
			// Product not found - increment the lastID
			idInt++
			h.lastID = idInt
		} else {
			h.lastID++
		}

		// - increment the lastID

		// - serialize the request body into a Product struct
		product := internal.Product{
			ID:          h.lastID,
			Name:        bodyProductRequest.Name,
			Quantity:    bodyProductRequest.Quantity,
			CodeValue:   bodyProductRequest.CodeValue,
			IsPublished: bodyProductRequest.IsPublished,
			Expiration:  bodyProductRequest.Expiration,
			Price:       bodyProductRequest.Price,
		}

		// - validate the product fields for the business rules
		if product.Name == "" {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body name is required"})
			return
		}

		// store the prduct
		h.product[product.ID] = product

		// response
		// - return the product as JSON
		// response
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		repository.SaveProductToFile(repository.Product(product))

		web.ResponseJSON(w, http.StatusCreated, map[string]interface{}{
			"message": "task created",
			"data":    data,
		})

	}
}
