package domain

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(ID string, name string, price float64) *Product {
	return &Product{
		ID:    ID,
		Name:  name,
		Price: price,
	}
}
