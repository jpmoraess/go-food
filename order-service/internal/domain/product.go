package domain

import "github.com/google/uuid"

type Product struct {
	id    uuid.UUID
	name  string
	price float64
}

func newProduct(ID uuid.UUID, name string, price float64) *Product {
	return &Product{
		id:    ID,
		name:  name,
		price: price,
	}
}

func (p *Product) ID() uuid.UUID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}
