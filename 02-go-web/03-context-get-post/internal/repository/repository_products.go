package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

/* QUITADO DEL HANDLER VER DE DISTRIBUIR SEGUN CORRESPONDA CON O FUERA DE REPOSITORY */

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
	/* products := LoadProductsFromFile()
	if products == nil {
		products = []Product{}
	} */
	var products []Product
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
