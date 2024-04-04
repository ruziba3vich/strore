package main

import "fmt"

func main() {
	// Pointer yordamida string ni joyida (funksiya return qimidi) qaytaruvchi funktsiyani ishlab chiqing.
	// Ex: hello => olleh
	myStr := "hello"
	ReverseString(&myStr)
	fmt.Println(myStr)
}

func ReverseString(s *string) {
	var str string
	
	for i := len(*s) - 1; i > -1; i -- {
		str += string((*s)[i])
	}
	*s = str
}
