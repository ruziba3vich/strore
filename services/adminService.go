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
		currency)] = numberOfProducts
}

func CreateNewAdmin(
	username string,
	password string,
	name string,
	surname string,
	phone string) {
	db.AdminsDb[username] = model.Admin{
		Username:     username,
		Password:     password,
		Name:         name,
		Surname:      surname,
		Phone_number: phone,
	}
}
