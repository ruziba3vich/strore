package model

import (
	"fmt"
)

type Person struct {
	Name     string
	Surname  string
	Age      int
	Gender   bool
	Phone    string
	Address  string
	Password string
	Email    string
}

func (p Person) PersonDTO() (result []string) {
	result = append(result, fmt.Sprintf("Foydalanuvchining to'liq ismi %s %s", p.Name, p.Surname))
	if p.Gender {
		result = append(result, fmt.Sprintf("yoshi %d da, jinsi erkak", p.Age))
	} else {
		result = append(result, fmt.Sprintf("yoshi %d da, jinsi ayol, telefon raqami : %s", p.Age, p.Phone))
	}
	result = append(result, fmt.Sprintf("yashash manzili %s", p.Address))
	result = append(result, fmt.Sprintf("email : %s", p.Email))
	return result
}

func (p *Person) SetGender(g int) {
	if g == 0 {
		p.Gender = false
	} else {
		p.Gender = true
	}
}

func NewPerson(
	name string,
	surname string,
	age int,
	gender bool,
	phone string,
	address string,
	password string,
	email string) Person {
	return Person{
		Name:     name,
		Surname:  surname,
		Age:      age,
		Gender:   gender,
		Phone:    phone,
		Address:  address,
		Password: password,
		Email:    email,
	}
}

/*
HAMMASI FUNKSIYA SIFATIDA YARATILSIN !!!

Ro`yxardan o'tish:
	Foydalanuvchidan malumotlarini olamiz
Profilga kirish:
	email> password>   -> foydalanuvchini ma`lumotlari chiqarilsin
*/
