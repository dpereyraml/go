package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"package-oriented-design/internal"
	"package-oriented-design/internal/repository"
	"package-oriented-design/platform/tools"
	"time"

	"github.com/bootcamp-go/web/response"
)

// DefaultProduct is a struct that contains the handlers for the products
type DefaultProduct struct {
	/*
		ya no trabaja con esto

		product map[int]internal.Product
		lastID  int

		ahora trabaja con el servicio
	*/
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

/*
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
*/

// Methods GET
/* solucionar
func (h *DefaultProduct) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Pong!"))
		response.JSON(w, http.StatusOK, map[string]any{"response": "Pong!"})
	}
}

func (h *DefaultProduct) Products() http.HandlerFunc {
	productsData := repository.LoadProductsFromFile()
	// falta validaciones de errores
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, productsData)
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
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto no encontrado"})
			return
		}

		// Return the product as JSON
		response.JSON(w, http.StatusOK, product)
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
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto no encontrado"})
			return
		}

		// Return the product as JSON
		response.JSON(w, http.StatusOK, product)
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
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Producto/s no encontrado/s"})
			return
		}

		// Return the product as JSON
		response.JSON(w, http.StatusOK, products)
	}
}
*/
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

		// - validate request body fields
		/*
			if _, ok := bodyMap["name"]; !ok {
				response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body name is required"})
				return
			}
			if _, ok := bodyMap["quantity"]; !ok {
				response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body quantity is required"})
				return
			}
			if _, ok := bodyMap["code_value"]; !ok {
				response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body code_value is required"})
				return
			}
			if _, ok := bodyMap["expiration"]; !ok {
				response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body expiration is required"})
				return
			}
		*/

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
		// - get the lastID
		// - validate code_value|unique
		/*
			IMPORTANTE!!!
			code_val := bodyMap["code_value"]
			var idInt int
			productsData := repository.LoadProductsFromFile()
			for _, p := range productsData {
				fmt.Println("no funcionaba", p, code_val, idInt)
				if p.CodeValue == code_val {
					response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body code_value must be unique"})
					return
				}
				if p.ID > idInt {
					idInt = p.ID
				}
			}
		*/

		// Valid if the product was found
		/*
			if  idInt != 0 {
				// Product not found - increment the lastID
				idInt++
				h.lastID = idInt
			} else {
				// - increment the lastID
				h.lastID++
			}
		*/

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
		repository.SaveProductToFile(repository.ProductRep{})

		response.JSON(w, http.StatusCreated, map[string]interface{}{
			"message": "task created",
			"data":    data,
		})

	}
}
