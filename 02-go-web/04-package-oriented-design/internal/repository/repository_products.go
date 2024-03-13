package repository

import (
	"fmt"
	"package-oriented-design/internal"
)

func NewProductRepository(db map[int]internal.Product, lastID int) *ProductRep {
	if db == nil {
		db = make(map[int]internal.Product)
	}
	return &ProductRep{
		db:     db,
		lastID: lastID,
	}
}

type ProductRep struct {
	db     map[int]internal.Product
	lastID int
}

func (p *ProductRep) CreateProduct(product *internal.Product) (err error) {
	// increment lastID and assign it to the product ID
	(*p).lastID++
	(*product).ID = (*p).lastID
	(*p).db[(*product).ID] = *product
	return
}

func (p *ProductRep) GetProductsById(id int) (product *internal.Product, err error) {

	productRetrieved, ok := p.db[id]
	if !ok {
		return nil, fmt.Errorf("product not found")

	} else {
		return &productRetrieved, nil

	}
}

func (p *ProductRep) GetProducts() []internal.Product {
	products := make([]internal.Product, 0, len(p.db))
	for _, product := range p.db {
		products = append(products, product)
	}
	return products
}

func (p *ProductRep) GetBySearchQuery(priceGT float64) ([]internal.Product, error) {
	products := make([]internal.Product, 0, len(p.db))
	for _, product := range p.db {
		if product.Price > priceGT {
			products = append(products, product)
		}
	}
	return products, nil
}

/* REVISAR * /

// add product to JSON file
func SaveProductToFile(product ProductRep) []ProductRep {
	products := LoadProductsFromFile()
	if products == nil {
		products = make([]ProductRep, 0)
	}
	products = append(products, product)
	bytes, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = os.WriteFile("../docs/products.json", bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return products
}

// return to slice of product from JSON file (CREO)
func LoadProductsFromFile() []ProductRep {
	var products []ProductRep
	raw, err := os.ReadFile("../docs/products.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	errJson := json.Unmarshal(raw, &products)
	if errJson != nil {
		fmt.Println(errJson)
		return nil
	}

	return products
}

func (p *ProductRep) Ping() any {
	return p.
}

func (p *ProductRep) LoadProducts() []internal.Product {
	return p.rp.LoadProducts()
}

func (p *ProductRep) Products() []internal.Product {
	return p.rp.Products()
}

func (p *ProductRep) ProductByID(id int) (internal.Product, error) {
	return p.rp.ProductByID(id)
}

func (p *Product) GetBySearchQuery(query string) ([]internal.Product, error) {
	return p.rp.GetBySearchQuery(query)
}



/*
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// return to string bytes of product
func (p Product) toString() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

// return to slice of product from JSON file (CREO)
func LoadProductsFromFile() []Product {
	var products []Product
	raw, err := os.ReadFile("../docs/products.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	errJson := json.Unmarshal(raw, &products)
	if errJson != nil {
		fmt.Println(errJson)
		return nil
	}

	return products
}

// add product to JSON file
func SaveProductToFile(product Product) []Product {
	products := LoadProductsFromFile()
	if products == nil {
		products = []Product{}
	}

	products = append(products, product)

	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = os.WriteFile("../docs/products.json", data, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return products
}

// return JSON from slice of product
func ToJSON(products []Product) []byte {
	bytes, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bytes
}
*/
