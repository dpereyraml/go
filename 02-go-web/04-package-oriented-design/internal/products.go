package internal

import "errors"

// struct default for product
type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var (
	// ErrproductNotFound is an error that is used when the product is not found
	ErrproductNotFound = errors.New("product not found")
	// ErrproductDuplicate is an error that is used when the product already exists
	ErrproductDuplicated = errors.New("product already exists")
	// ErrproductInvalidField is an error that is used when the product is invalid
	ErrproductInvalidField = errors.New("product is invalid")
	// ErrproductInternal is an error that is used when the product can't be saved
	ErrproductInternal = errors.New("product can't be processed")
)

// ProductsRepository is the interface for the products repository
type ProductsRepository interface {
	CreateProduct(product *Product) (err error)
	GetProductsById(id int) (*Product, error)
	GetProducts() []Product
	GetBySearchQuery(priceGT float64) ([]Product, error)
}

// ProductsService is the interface for the products service
type ProductsService interface {
	CreateProduct(product *Product) (err error)
	GetProductsById(id int) (Product, error)
	GetProducts() []Product
	GetBySearchQuery(priceGT float64) ([]Product, error)
}
