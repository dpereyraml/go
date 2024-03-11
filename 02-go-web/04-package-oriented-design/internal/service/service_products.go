package service

import "package-oriented-design/internal"

func NewProductSS(rp internal.ProductsRepository) *Product {
	return &Product{
		rp: rp,
	}
}

// the interface service works with the interface repository - union
type Product struct {
	rp internal.ProductsRepository
}

/*
	func (p *Product) Ping() any {
		return p.rp.Ping()
	}

	func (p *Product) LoadProducts() []internal.Product {
		return p.rp.LoadProducts()
	}

	func (p *Product) Products() []internal.Product {
		return p.rp.Products()
	}

	func (p *Product) ProductByID(id int) (internal.Product, error) {
		return p.rp.ProductByID(id)
	}

	func (p *Product) GetBySearchQuery(query string) ([]internal.Product, error) {
		return p.rp.GetBySearchQuery(query)
	}
*/
func (p *Product) CreateProduct(product *internal.Product) error {
	return p.rp.CreateProduct(product)
}
