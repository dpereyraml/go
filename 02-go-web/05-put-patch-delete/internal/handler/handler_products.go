package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"put-patch-delete/internal"
	"put-patch-delete/platform/tools"
	"strconv"
	"time"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// DefaultProduct is a struct that contains the handlers for the products
type DefaultProduct struct {
	sv internal.ProductsService
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
func NewHandlerProducts(sv internal.ProductsService) *DefaultProduct {
	return &DefaultProduct{
		sv: sv,
	}
}

func (h *DefaultProduct) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Pong!"))
		response.JSON(w, http.StatusOK, map[string]any{"response": "Pong!"})
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
		productsData, err := h.sv.GetProductsById(idInt)
		if err != nil {
			fmt.Println("Error al obtener el producto:", err)
			http.Error(w, "Error al obtener el producto", http.StatusInternalServerError)
			return
		}

		// Valid if the product was found
		if productsData.ID == 0 {
			// Product not found
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto no encontrado"})
			return
		}

		// Return the product as JSON
		response.JSON(w, http.StatusOK, productsData)
	}
}

func (h *DefaultProduct) GetProducts() http.HandlerFunc {

	// falta validaciones de errores
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, h.sv.GetProducts())
	}
}

func (h *DefaultProduct) GetProductsBySearchQuery() http.HandlerFunc {
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
		productsData, err := h.sv.GetBySearchQuery(priceGt)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error al obtener el producto"})
			return
		}

		if productsData == nil {
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto/s no encontrado/s"})
			return
		}

		// Return the product as JSON
		response.JSON(w, http.StatusOK, productsData)
	}
}

// Methods POST
func (h *DefaultProduct) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// validate token
		/*   */

		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body 1"})
			return
		}

		// - validate request body parse into map[string]any
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body Unmarshal"})
			return
		}

		// validate requerid fields
		if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "expiration"); err != nil {
			var FieldError *tools.FieldError
			if errors.As(err, &FieldError) {
				response.JSON(w, http.StatusInternalServerError, map[string]any{"error": fmt.Sprintf("%s is required", FieldError.Field)})
				return
			}
		}

		expiration, _ := bodyMap["expiration"].(string)
		/* if !ok {
		 			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "expiration don't have the correct format, is not a string"})
					return
				} */

		_, errExt := time.Parse("02/01/2006", expiration) // REVISAR
		if errExt != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "expiration don't have the correct format"})

			return
		}

		if _, ok := bodyMap["price"]; !ok {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body price is required"})
			return
		}

		// - read into bytes
		var bodyProductRequest ProductRequest

		// Forma " Automática " - Using web package
		if err := json.Unmarshal(bytes, &bodyProductRequest); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid request body Unmarshal",
			})
			return
		}

		// process

		// validate the product
		if len(bodyProductRequest.Name) > 35 {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body name must be less than 35 characters"})
			return
		}
		// - validate the product fields for the business rules
		if bodyProductRequest.Name == "" {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body name is required"})
			return
		}
		// - serialize the request body into a Product struct
		product := internal.Product{
			// ID:          h.lastID,
			Name:        bodyProductRequest.Name,
			Quantity:    bodyProductRequest.Quantity,
			CodeValue:   bodyProductRequest.CodeValue,
			IsPublished: bodyProductRequest.IsPublished,
			Expiration:  bodyProductRequest.Expiration,
			Price:       bodyProductRequest.Price,
		}
		// store the prduct

		if err := h.sv.CreateProduct(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrproductDuplicated):
				response.JSON(w, http.StatusConflict, map[string]any{"message": "task already exists"})
			case errors.Is(err, internal.ErrproductInvalidField):
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": "task is invalid"})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}
		// ----- h.product[product.ID] = product

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

		// repository.SaveProductToFile(repository.Product(product))
		h.sv.CreateProduct(&internal.Product{})

		response.JSON(w, http.StatusCreated, map[string]interface{}{
			"message": "task created",
			"data":    data,
		})

	}
}

func (h *DefaultProduct) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// validate required fields
		if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "is_published", "expiration"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.Text(w, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field))
				return
			}

			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}
		// - parse json to struct (static)
		var body ProductRequest
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - validate the task
		if body.Name == "" || len(body.Name) > 25 {
			response.Text(w, http.StatusBadRequest, "Name is required and must be less than 25 characters")
			return
		}

		// process
		// - serialize the request body into a product
		product := internal.Product{
			ID:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		// - update the task
		if err := h.sv.UpdateProduct(product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrproductInvalidField):
				response.Text(w, http.StatusBadRequest, "product is invalid")
			case errors.Is(err, internal.ErrproductDuplicated):
				response.Text(w, http.StatusConflict, "product already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}
}

func (h *DefaultProduct) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete the task
		if err := h.sv.DeleteProduct(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product deleted",
		})
	}
}

func (h *DefaultProduct) UpdateProductPartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UpdateProductPartial")
		// request
		// - id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// parse map from json
		bodyMap := map[string]any{}
		if err := json.NewDecoder(r.Body).Decode(&bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// validate required fields
		/* if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "is_published", "expiration"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.Text(w, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field))
				return
			}

			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		} */

		// process
		// - update the product
		fmt.Println("bodyMap", bodyMap)
		if err := h.sv.UpdateProductPartial(id, bodyMap); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrproductInvalidField):
				response.Text(w, http.StatusBadRequest, "product is invalid")
			case errors.Is(err, internal.ErrproductDuplicated):
				// response.Text(w, http.StatusConflict, "product already exists 407 handler")
				response.Text(w, http.StatusConflict, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}

			return
		}

		// response
		response.JSON(w, http.StatusOK, "product updated")
	}
}
