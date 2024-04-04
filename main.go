package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	db "strore/database"
	"strore/emailcheck"
	model "strore/models"
	auth "strore/services"

	ppg "github.com/k0kubun/pp"
)

var shift int = 47

// var global_user any

var role int

// var authorize func() func()

func init() {
	data := authenticate()
	if len(data) > 0 {
		if data[0] == "USER" {
			role = 1
		} else {
			role = 2
		}
	}

	userDbSL, _ := ReadFile("database/users.txt")
	for _, d := range userDbSL {
		lineOfData := getDataFromLine(d)
		email := lineOfData[0] + lineOfData[1]
		age, _ := strconv.Atoi(lineOfData[4])
		gender := lineOfData[5] == "true"
		newUser := model.Person {
			Name: lineOfData[2],
			Surname: lineOfData[3],
			Age: age,
			Gender: gender,
			Phone: lineOfData[6],
			Address: lineOfData[7],
			Password: lineOfData[8],
			Email: lineOfData[9],
		}
		db.UserDb[email] = newUser
	}

	adminsDbSL, _ := ReadFile("database/admins.txt")

	for _, d := range adminsDbSL {
		lineOfData := getDataFromLine(d)
		newAdmin := model.Admin {
			Username: lineOfData[0],
			Password: lineOfData[1],
			Name: lineOfData[2],
			Surname: lineOfData[3],
			Phone_number: lineOfData[4],
		}
		db.AdminsDb[newAdmin.Username] = newAdmin
	}

	productsDbSL, _ := ReadFile("database/products.txt")

	for _, d := range productsDbSL {
		lineOfData := getDataFromLine(d)
		
		price, _ := strconv.Atoi(lineOfData[3])
		var cur model.Currency
		if lineOfData[4] == "$" {
			cur = model.Dollar
		} else {
			cur = model.Sum
		}
		numberOfProducts, _ := strconv.Atoi(lineOfData[5])
		newProduct := model.CreateProduct(
								db.AdminsDb[lineOfData[0]],
								lineOfData[1],
								getTitle(lineOfData[2]),
								price, cur)
		db.ProductsDb[newProduct] = numberOfProducts
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
				auth.CreateNewAdmin(newUsername, newPassword, newName, newSurname, newPhone)
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
				auth.AddProduct(user, newName, newTitle, newPrice, cur, 0, newCount)
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
			cart := map[model.Product] int {}
			ordn := 0
			for ordn != -1 {
				fmt.Print("Sotib olmoqchi bo'lgan maxsulotingizni yuqoridagi\nro'yhatdagi tartib raqamini kiriting yoki haridni to'xtatish uchun -1 : ")
				fmt.Scan(& ordn)
				num := 0
				if ordn != -1 {
					fmt.Println("ushbu maxsulotdan nechta sotib olmoqchisiz ? : ")
					fmt.Scan(& num)
					cart[products[ordn - 1]] += num
				}
			}
			
			ok, pro := auth.BuyProduct(user, cart)
			if ok {
				fmt.Println("Haridingiz muvoffaqayatli amalga oshirildi")
			} else {
				fmt.Println("Maxsulot sotib olishda xatolik", pro.GetName())
			}
			
		default:
			isAdmin := 0
			fmt.Print("1 -> Adminga kirish / 2 -> Userga kirish : 3 -> Ro'yhatdan o'tish : ")
			fmt.Scan(& isAdmin)
			if isAdmin == 1 {
				var adminUsername string
				var adminPassword string
				fmt.Print("username kiriting : ")
				fmt.Scan(& adminUsername)
				fmt.Print("password kiriting : ")
				fmt.Scan(& adminPassword)

				yy, ok := signInAsAdmin(adminUsername, adminPassword)
				if ok {
					role = 2
					u = yy
				} else {
					ppg.Println("Noto'g'ri login yoki parol !")
				}
			} else if isAdmin == 2 {
				uuu := model.Person {}
				signIn(& uuu)
				u = uuu
			} else {
				var user_ model.Person
				fmt.Print("Ismingizni kiriting : ")
				fmt.Scan(& user_.Name)
				fmt.Print("Familyangizni kiriting : ")
				fmt.Scan(& user_.Name)
				fmt.Print("Yoshingizni kiriting : ")
				fmt.Scan(& user_.Age)
				fmt.Print("Telefon raqamingizni kiriting : ")
				fmt.Scan(& user_.Phone)
				fmt.Print("Emailingizni kiriting : ")
				fmt.Scan(& user_.Email)
				fmt.Print("Profil uchun parol o'ylab toping : ")
				fmt.Scan(& user_.Name)
				var gen int
				fmt.Print("Jinsingizni tanlang 1 -> Erkak / 2 -> Ayol : ")
				fmt.Scan(& gen)
				user_.Gender = gen == 1
				ok := signUp(user_)

				if ok {
					fmt.Println("Muvoffaqayatli ro'yhatdan o'tdingiz !")
				} else {
					ppg.Println("Ushbu e-mail tizim ro'yhatida mavjud !")
				}
			}
		}
		ppg.Print("Yana qandaydir amal bajarasizmi 1 -> ha / 2 -> yo'q : ")
		fmt.Scan(& status)
	}
	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}

func signInAsAdmin(username string, password string) (a model.Admin, got bool) {
	admin, ok := db.AdminsDb[username]
	if ok {
		password = caesarEncrypt(password, shift)
		if admin.Password == password {
			return admin, true
		} else {
			return a, false
		}
	}
	return a, got
}

func getTitle(s string) (str string) {
	ss := strings.Split(s, "_")
	for _, word := range ss {
		str += word + " "
	}
	return str
}

func getDataFromLine(s string) [] string {
	return strings.Split(s, ".")
}

func ReadFile(s string) (res [] string, err error) {
	file, err := os.Open(s)
    if err != nil {
        return res, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		res = append(res, line)
    }
	return res, nil
}

func signIn(user * model.Person) {
	fmt.Print("Profilga kirish uchun emailingizni kiriting : ")
	var email string
	fmt.Scan(&email)
	var password string
	fmt.Print("Profilga kirish uchun parolingizni kiriting : ")
	fmt.Scan(&password)
	_, entered, message := auth.Authenticate(email, password)
	if entered {
		fmt.Println(message)
		// user = pp
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

func signUp (u model.Person) bool {
	_, ok := emailcheck.CHeckEmail(u.Email)
	if ! ok {
		db.UserDb[u.Email] = u
		return true
	}
	return false
}
