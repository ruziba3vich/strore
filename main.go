package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	model "strore/models"
	auth "strore/services"
	adminService "strore/services"
	ppg "github.com/k0kubun/pp"
	db "strore/database"
)

var shift int = 47

var global_user any

var role int

// var makeUser func() func()(model.Person)
// var makeAdmin func() func()(model.Admin)

var authorize func() func()

func init() {
	data := authenticate()
	if len(data) > 0 {
		if data[0] == "USER" {
			role = 1
		} else {
			role = 2
		}
	}
}

func main() {
	var u interface{}
	if role == 1 {
		u = model.Person{}
	} else if role == 2 {
		u = model.Admin{}
	}

	status := 1
	var n int

	for status == 1 {

		switch user := u.(type) {
		case model.Admin:
			ppg.Print("1 -> Yangi `product` yaratish\n2 -> Yangi `admin` yaratish\n    ->  ")
			fmt.Scan(&n)
			if n == 2 {
				var newName string
				fmt.Print("Yangi yaratmoqchi bo'lgan admin ismi : ")
				fmt.Scan(& newName)
				var newSurname string
				fmt.Print("Yangi yaratmoqchi bo'lgan admin familyasi : ")
				fmt.Scan(& newSurname)
				var newUsername string
				fmt.Print("Yangi yaratmoqchi bo'lgan admin uchun username : ")
				fmt.Scan(& newUsername)
				var newPassword string
				fmt.Print("Yangi yaratmoqchi bo'lgan admin uchun parol : ")
				fmt.Scan(& newPassword)
				var newPhone string
				fmt.Print("Yangi yaratmoqchi bo'lgan admin tel raqami : ")
				fmt.Scan(& newPhone)
				adminService.CreateNewAdmin(user, newUsername, newPassword, newName, newSurname, newPhone)
			} else {
				fmt.Print("Yangi yaratilmoqchi bo'lgan product nomi : ")
				var newName string
				fmt.Scan(& newName)
				fmt.Print("Yangi yaratilmoqchi bo'lgan product valyutasi 1 -> $ / 2 -> so'm : ")
				var newCurrency int
				fmt.Scan(& newCurrency)
				var cur model.Currency
				if newCurrency == 1 {
					cur = model.Dollar
				} else {
					cur = model.Sum
				}
				fmt.Print("Yangi yaratilmoqchi bo'lgan product narxi : ")
				var newPrice int
				fmt.Scan(& newPrice)
				fmt.Print("Yangi yaratilmoqchi bo'lgan product nomi : ")
				var newTitle string
				fmt.Scanln(& newTitle)
				fmt.Print("Yangi yaratilmoqchi bo'lgan product soni : ")
				var newCount int
				fmt.Scanln(& newCount)
				adminService.AddProduct(user, newName, newTitle, newPrice, cur, 0, newCount)
				fmt.Println("Siz yangi product qo'shdingiz !")
			}

		case model.Person:
			i := 1
			products := [] model.Product {}
			for key, val := range db.ProductsDb {
				fmt.Println(i, key, val, "ta")
				i ++
				products = append(products, key)
			}
			fmt.Print("Sotib olmoqchi bo'lgan maxsulotingizni yuqoridagi\nro'yhatdagi tartib raqamini kiriting : ")
			ordn := 0
			fmt.Scan(& ordn)
			num := 0
			fmt.Println("ushbu maxsulotdan nechta sotib olmoqchisiz ? : ")
			fmt.Scan(& num)
			adminService.BuyProduct()
			
		default:

		}

		// fmt.Println("Quyidagi menu dagi imkoniyatlardan birini tanlang")
		// fmt.Print("1 -> barcha foydalanuvchilarni ko'rish\n2 -> Foydalanuvchini qidirish\n3 -> Ro'yhatdan o'tish\n4 -> Profilga kirish : ")
		// fmt.Scan(&n)

		// for n > 4 || n < 1 {
		// 	fmt.Println("Quyidagi menu dagi imkoniyatlardan birini tanlang")
		// 	fmt.Print("1 -> barcha foydalanuvchilarni ko'rish\n2 -> Foydalanuvchini qidirish\n3 -> Ro'yhatdan o'tish\n4 -> Profilga kirish : ")
		// 	fmt.Scan(&n)
		// }

		// if n == 1 {
		// 	for _, person := range db.UserDb {
		// 		for _, info := range person.PersonDTO() {
		// 			ppg.Println(info)
		// 		}
		// 		fmt.Println()
		// 		fmt.Println()
		// 	}
		// } else if n == 2 {
		// 	var name string
		// 	var surname string

		// 	fmt.Print("Qidirilayotgan foydalanuvchining ismini kiriting : ")
		// 	fmt.Scan(&name)
		// 	fmt.Print("Qidirilayotgan foydalanuvchining familyasini kiriting : ")
		// 	fmt.Scan(&surname)
		// 	ok, person := searchservice.Search(name, surname)
		// 	if ok {
		// 		for _, info := range person.PersonDTO() {
		// 			ppg.Println(info)
		// 		}
		// 		fmt.Println()
		// 		fmt.Println()
		// 	} else {
		// 		fmt.Println("Ushbu ism familiyada foydalanuvchi topilmadi")
		// 	}
		// } else if n == 3 {
		// 	var newPerson model.Person
		// 	fmt.Print("Iltimos ismingizni kiriting : ")
		// 	fmt.Scan(&newPerson.Name)
		// 	fmt.Print("Iltimos familyangizni kiriting : ")
		// 	fmt.Scan(&newPerson.Surname)
		// 	fmt.Print("Iltimos yoshingizni kiriting : ")
		// 	fmt.Scan(&newPerson.Age)
		// 	fmt.Print("Iltimos jinsingizni tanlang kiriting : 1 -> erkak / 2 -> ayol : ")
		// 	var gender int
		// 	fmt.Scan(&gender)
		// 	newPerson.SetGender(gender)
		// 	fmt.Print("Iltimos telefon raqamingizni kiriting : ")
		// 	fmt.Scan(&newPerson.Phone)
		// 	fmt.Print("Iltimos yashash manzilingizni kiriting kiriting : ")
		// 	fmt.Scan(&newPerson.Address)
		// 	var email string
		// 	fmt.Print("Iltimos profilingiz uchun emailingizni kiriting : ")
		// 	fmt.Scan(&email)

		// 	var notSignedIn bool = true

		// 	for _, response := check.CHeckEmail(email); response; {
		// 		fmt.Print("Bunday emai bizda mavjud, 1 boshqa email kiritish; -> 2 -> profilga kirish; ")
		// 		fmt.Scan(&n)
		// 		if n == 1 {
		// 			fmt.Scan(&email)
		// 		} else {
		// 			signIn(&user)
		// 			notSignedIn = false
		// 			break
		// 		}
		// 	}

		// 	if notSignedIn {
		// 		newPerson.Email = email
		// 		fmt.Print("Iltimos profilingiz uchun parol o'ylab toping : ")
		// 		fmt.Scan(&newPerson.Password)
		// 		registered := register.Register(newPerson)
		// 		if registered {
		// 			fmt.Println("Siz muvoffaqayatli ro'yhatdan o'tdingiz !")
		// 		} else {
		// 			fmt.Println("Ro'yhatdan o'tishda hatolik . . .")
		// 		}
		// 	}
		// } else {
		// 	signIn(user)
		// }
		// fmt.Print("Yana qandaydir amal bajarasizmi ? 1 -> Ha/ 2 -> Yo'q : ")
		// fmt.Scan(&status)
	}
	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}

func signIn(user model.Person) {
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

func authenticate() (data []string) {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := scanner.Text()
	data = strings.Split(line, ".")
	for i := range data {
		data[i] = caesarDecrypt(data[i], shift)
	}

	return data
}

func caesarEncrypt(plainText string, shift int) string {
	shiftedAlphabet := make(map[byte]byte)

	for i := byte('A'); i <= byte('Z'); i++ {
		shiftedAlphabet[i] = byte((int(i-'A')+shift)%26 + 'A')
	}
	cipherText := make([]byte, len(plainText))
	for i := 0; i < len(plainText); i++ {
		char := plainText[i]
		if char >= 'A' && char <= 'Z' {
			cipherText[i] = shiftedAlphabet[char]
		} else if char >= 'a' && char <= 'z' {
			upperChar := byte(strings.ToUpper(string(char))[0])
			cipherText[i] = byte(shiftedAlphabet[upperChar] + 32)
		} else {
			cipherText[i] = char
		}
	}
	return string(cipherText)
}

func caesarDecrypt(cipherText string, shift int) string {
	return caesarEncrypt(cipherText, 26-shift)
}
