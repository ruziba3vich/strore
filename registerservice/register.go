package register

import (
	db "strore/database"
	model "strore/models"
)


func Register(p model.Person) (ok bool) {
	if _, ok := db.UserDb[p.Email]; ok {
		return false
	}
	db.UserDb[p.Email] = p
	return true
}
