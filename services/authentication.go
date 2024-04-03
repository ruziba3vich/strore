package services

import (
	m "strore/models"
	db "strore/database"
)

func Authenticate(email string, password string) (m.Person, bool, string) {
	dummy := m.NewPerson("", "", 0, false, "", "", "", "")
	person, ok := db.UserDb[email]
	if ok {
		if person.Password == password {
			return person, ok, "Siz muvoffaqiyatli ravishda profilingizga kirdingiz !"
		} else {
			return dummy, ok, "Siz kiritgan parol noto'g'ri"
		}
	} else {
		return dummy, ok, "Bunday email bilan foydalanuvchi topilmadi !"
	}
}
