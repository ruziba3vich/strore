package emailcheck

import (
	db "strore/database"
	model "strore/models"
)

func CHeckEmail(email string) (model.Person, bool) {
	a, b := db.UserDb[email]
	return a, b
}
