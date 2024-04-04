package model

type Currency string

const (
	Dollar Currency = "$"
	Sum    Currency = "so'm"
)


type Product struct {
	name      string
	title     string
	price     int
	currency  Currency
	createdBy Admin
}


func CreateProduct(
	admin Admin,
	name string,
	title string,
	price int,
	currency Currency) Product {
	return Product{
		name:      name,
		title:     title,
		price:     price,
		currency:  currency,
		createdBy: admin,
	}
}


func (p Product) GetName () string {
	return p.name
}
