package searchservice

import (
	db "strore/database"
	m "strore/models"
)

func Search(name string, surname string) (bool, m.Person) {
	dummy := m.NewPerson("", "", 0, false, "", "", "", "")
	for _, person := range db.UserDb {
		if person.Name == name && person.Surname == surname {
			return true, person
		}
	}
	return false, dummy
}
