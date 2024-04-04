package services

import (
	db "strore/database"
	model "strore/models"
)

func BuyProduct(user model.Person, products map [model.Product] int) (ok bool, p model.Product) {
	for product, n := range products {
		if db.ProductsDb[product] + 1 > n {
			db.ProductsDb[product] -= n
		} else {
			return false, product
		}
	}
	return true, p
}
