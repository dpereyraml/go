/*
	Crear un programa que cumpla los siguiente puntos:

	1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
	2. Tener un slice global de Product llamado Products instanciado con valores.
	3. 2 métodos asociados a la estructura Product: Save(), GetAll().
	El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
	El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
	4. Una función getById() al cual se le deberá pasar un INT como parámetro
	y retorna el producto correspondiente al parámetro pasado.
	5. Ejecutar al menos una vez cada método y función definido desde main().

*/

package main

import "fmt"

type Product struct { // 1
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{ // 2
	{
		ID:          1,
		Name:        "Laptop Lenovo",
		Price:       799.99,
		Description: "Powerful laptop with Intel Core i7 processor and 16GB RAM.",
		Category:    "Electronics",
	},
	{
		ID:          2,
		Name:        "Smartphone Samsung",
		Price:       699.99,
		Description: "Latest Samsung smartphone with 6.5-inch display and triple camera setup.",
		Category:    "Electronics",
	},
}

func main() {
	fmt.Println("Estado inicial")
	fmt.Println(Products)

	var prod = []Product{
		{
			ID:          3,
			Name:        "Wireless Headphones",
			Price:       149.99,
			Description: "High-quality wireless headphones with noise-canceling feature.",
			Category:    "Electronics",
		},
	}
	Save(prod)
	GetAll()
	elementoEspecifico := getById(3)
	fmt.Println("elementoEspecifico", elementoEspecifico)
}

// 3
func Save(prod []Product) {
	Products = append(Products, prod...)
	fmt.Println("Productos Guardados")
	fmt.Println(Products)
}

func GetAll() {
	for _, v := range Products {

		fmt.Println("Listado de Productos")
		fmt.Println(v)
	}
}

// 4

func getById(id int) Product {
	for _, v := range Products {
		if v.ID == id {
			/* var result = []Product{
				{ID: v.ID,
					Name:        v.Name,
					Price:       v.Price,
					Description: v.Description,
					Category:    v.Category},
			} */
			return v
		}
	}
	return Product{}
}
