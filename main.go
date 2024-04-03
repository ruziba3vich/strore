package main

import (
	"fmt"
	auth "strore/services"
	db "strore/database"
	check "strore/emailcheck"
	model "strore/models"
	register "strore/registerservice"
	searchservice "strore/searchService"

	ppg "github.com/k0kubun/pp"
)

func main() {
	user := model.Person{}
	status := 1
	var n int

	for status == 1 {

		fmt.Println("Quyidagi menu dagi imkoniyatlardan birini tanlang")
		fmt.Print("1 -> barcha foydalanuvchilarni ko'rish\n2 -> Foydalanuvchini qidirish\n3 -> Ro'yhatdan o'tish\n4 -> Profilga kirish : ")
		fmt.Scan(&n)

		for n > 4 || n < 1 {
			fmt.Println("Quyidagi menu dagi imkoniyatlardan birini tanlang")
			fmt.Print("1 -> barcha foydalanuvchilarni ko'rish\n2 -> Foydalanuvchini qidirish\n3 -> Ro'yhatdan o'tish\n4 -> Profilga kirish : ")
			fmt.Scan(&n)
		}

		if n == 1 {
			for _, person := range db.UserDb {
				for _, info := range person.PersonDTO() {
					ppg.Println(info)
				}
				fmt.Println()
				fmt.Println()
			}
		} else if n == 2 {
			var name string
			var surname string

			fmt.Print("Qidirilayotgan foydalanuvchining ismini kiriting : ")
			fmt.Scan(&name)
			fmt.Print("Qidirilayotgan foydalanuvchining familyasini kiriting : ")
			fmt.Scan(&surname)
			ok, person := searchservice.Search(name, surname)
			if ok {
				for _, info := range person.PersonDTO() {
					ppg.Println(info)
				}
				fmt.Println()
				fmt.Println()
			} else {
				fmt.Println("Ushbu ism familiyada foydalanuvchi topilmadi")
			}
		} else if n == 3 {
			var newPerson model.Person
			fmt.Print("Iltimos ismingizni kiriting : ")
			fmt.Scan(&newPerson.Name)
			fmt.Print("Iltimos familyangizni kiriting : ")
			fmt.Scan(&newPerson.Surname)
			fmt.Print("Iltimos yoshingizni kiriting : ")
			fmt.Scan(&newPerson.Age)
			fmt.Print("Iltimos jinsingizni tanlang kiriting : 1 -> erkak / 2 -> ayol : ")
			var gender int
			fmt.Scan(&gender)
			newPerson.SetGender(gender)
			fmt.Print("Iltimos telefon raqamingizni kiriting : ")
			fmt.Scan(&newPerson.Phone)
			fmt.Print("Iltimos yashash manzilingizni kiriting kiriting : ")
			fmt.Scan(&newPerson.Address)
			var email string
			fmt.Print("Iltimos profilingiz uchun emailingizni kiriting : ")
			fmt.Scan(& email)

			var notSignedIn bool = true

			for _, response := check.CHeckEmail(email); response; {
				fmt.Print("Bunday emai bizda mavjud, 1 boshqa email kiritish; -> 2 -> profilga kirish; ")
				fmt.Scan(& n)
				if n == 1 {
					fmt.Scan(&email)
				} else {
					signIn(&user)
					notSignedIn = false
					break
				}
			}

			if notSignedIn {
				newPerson.Email = email
				fmt.Print("Iltimos profilingiz uchun parol o'ylab toping : ")
				fmt.Scan(& newPerson.Password)
				registered := register.Register(newPerson)
				if registered {
					fmt.Println("Siz muvoffaqayatli ro'yhatdan o'tdingiz !")
				} else {
					fmt.Println("Ro'yhatdan o'tishda hatolik . . .")
				}
			}
		} else {
			signIn(&user)
		}
		fmt.Print("Yana qandaydir amal bajarasizmi ? 1 -> Ha/ 2 -> Yo'q : ")
		fmt.Scan(&status)
	}
	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}

func signIn(user *model.Person) {
	fmt.Print("Profilga kirish uchun emailingizni kiriting : ")
	var email string
	fmt.Scan(&email)
	var password string
	fmt.Print("Profilga kirish uchun parolingizni kiriting : ")
	fmt.Scan(&password)
	pp, entered, message := auth.Authenticate(email, password)
	if entered {
		fmt.Println(message)
		*user = pp
	} else {
		ppg.Println(message)
		wants := 0
		fmt.Print("Yana bir bor urunib ko'rasizmi ? : 1 -> ha/2 -> yo'q : ")
		fmt.Scan(&wants)
		if wants == 1 {
			signIn(user)
		}
	}
}
