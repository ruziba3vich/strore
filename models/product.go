package model

type Currency string

const (
	dollar Currency = "$"
	sum    Currency = "so'm"
)

type Product struct {
	name      string
	title     string
	price     int
	currency  Currency
	discount  int
	createdBy Admin
}

func CreateProduct(
	admin Admin,
	name string,
	title string,
	price int,
	currency Currency,
	discount int) Product {
	return Product{
		name:      name,
		title:     title,
		price:     price,
		currency:  currency,
		discount:  discount,
		createdBy: admin,
	}
}
