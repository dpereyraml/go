package service

import "put-patch-delete/internal"

func NewProductSS(rp internal.ProductsRepository) *Product {
	return &Product{
		rp: rp,
	}
}

// the interface service works with the interface repository - union
type Product struct {
	rp internal.ProductsRepository
}

func (p *Product) CreateProduct(product *internal.Product) error {
	return p.rp.CreateProduct(product)
}

func (p *Product) UpdateProduct(product internal.Product) error {
	return p.rp.UpdateProduct(product)
}

func (p *Product) DeleteProduct(id int) error {
	return p.rp.DeleteProduct(id)
}

func (p *Product) UpdateProductPartial(id int, fields map[string]interface{}) (err error) {
	err = p.rp.UpdateProductPartial(id, fields)
	return
}

func (p *Product) GetProductsById(id int) (internal.Product, error) {
	productRetrieved, err := p.rp.GetProductsById(id)
	if err != nil {
		return internal.Product{}, err
	}
	return *productRetrieved, nil
}

func (p *Product) GetProducts() []internal.Product {
	return p.rp.GetProducts()
}

func (p *Product) GetBySearchQuery(priceGT float64) ([]internal.Product, error) {
	return p.rp.GetBySearchQuery(priceGT)
}
