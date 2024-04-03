package services

import (
	db "strore/database"
	model "strore/models"
)

func AddProduct(
	admin model.Admin,
	name string,
	title string,
	price int,
	currency model.Currency,
	discount int,
	numberOfProducts int) {
	db.ProductsDb[model.CreateProduct(
		admin, name,
		title, price,
		currency, discount)] = numberOfProducts
}
