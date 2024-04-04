package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	// Ushbu code da hatolik bor, shu ni izlab tog'rilash kerak
	// Kutiliyotgan natija: brand: apple

	var c computer
	change(&c, "apple")
	fmt.Printf("brand: %s\n", *(c.brand))
}

func change(c *computer, brand string) {
	c.brand = &brand
}
