package repository

import (
	"fmt"
	"put-patch-delete/internal"
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

func (p *ProductRep) UpdateProduct(product internal.Product) (err error) {
	_, ok := (*p).db[(product).ID]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	p.db[(product).ID] = product
	return
}

func (p *ProductRep) UpdateProductPartial(id int, fields map[string]any) (err error) {
	fmt.Println("fields", fields)
	fmt.Println("id", id)
	// check if the product exists
	product, ok := (*p).db[id]
	fmt.Println("product", product)
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	// iterate over the fields
	for field, value := range fields {
		// update the field
		switch field {
		case "name":
			name, ok := value.(string)
			if !ok {
				err = internal.ErrproductInvalidField
				return
			}
			// check if the name already exists
			for _, product := range (*p).db {
				if product.Name == name {
					err = internal.ErrproductDuplicated
					return
				}
			}
			product.Name = name
		/* case "quantity":
		quantity, ok := value.(int)
		if !ok {
			err = internal.ErrproductInvalidField
			return
		}
		// check if the quantity already exists
		for _, product := range (*p).db {
			if product.Quantity == quantity {
				err = internal.ErrproductDuplicated
				return
			}
		}
		product.Quantity = quantity */
		case "code_value":
			codeValue, ok := value.(string)
			if !ok {
				err = internal.ErrproductInvalidField
				return
			}
			// check if the codeValue already exists
			for _, product := range (*p).db {
				if product.CodeValue == codeValue {
					err = internal.ErrproductDuplicated
					return
				}
			}
			product.CodeValue = codeValue

		/* case "is_published":
			isPublished, ok := value.(bool)
			if !ok {
				err = internal.ErrproductInvalidField
				return
			}
			// check if the isPublished already exists
			for _, product := range (*p).db {
				if product.IsPublished == isPublished {
					err = internal.ErrproductDuplicated
					return
				}
			}
			product.IsPublished = isPublished

		case "expiration":
			expiration, ok := value.(string)
			if !ok {
				err = internal.ErrproductInvalidField
				return
			}
			// check if the expiration already exists
			for _, product := range (*p).db {
				if product.Expiration == expiration {
					err = internal.ErrproductDuplicated
					return
				}
			}
			product.Expiration = expiration

		case "price":
			price, ok := value.(float64)
			if !ok {
				err = internal.ErrproductInvalidField
				return
			}
			// check if the price already exists
			for _, product := range (*p).db {
				if product.Price == price {
					err = internal.ErrproductDuplicated
					return
				}
			}
			product.Price = price */

		default:

		}
	}

	// update the product
	(*p).db[id] = product

	return

}

func (p *ProductRep) DeleteProduct(id int) (err error) {
	// check if the product exists
	_, ok := (*p).db[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	// delete the product
	delete(p.db, id)
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
