package main

import (
	"fmt"
)

func main() {
	var bilangannya int
	// var suku5, suku4, suku3, suku2, suku1 int
	//var hasilnya int
	// var hasil string
	// var kata string

	fmt.Println("berapa bilangannya?")
	fmt.Scan(&bilangannya)
	fmt.Println("hasilnya:")

	//kata := fmt.Sprint(bilangannya)
	i := len(fmt.Sprint(bilangannya))
	m := 10
	k := 1

	for i > 0 {
		//hasilnya = (bilangannya % m) / k
		fmt.Print(((bilangannya % m) / k), " ")

		m *= 10
		k *= 10
		i -= 1
	}

	fmt.Print("\n")
	a := len(fmt.Sprint(bilangannya))
	b := "apaan si"
	c := len(b)
	fmt.Println(a, c)
	// suku5 = (n % 10) / 1
	// suku4 = (n % 100) / 10
	// suku3 = (n % 1000) / 100
	// suku2 = (n % 10000) / 1000
	// suku1 = (n % 100000) / 10000

	// hasil = fmt.Sprint(suku5, suku4, suku3, suku2, suku1)

	// fmt.Print("\n", hasil)
}
