package main

import (
	"fmt"
)

var n, i int
var hasil, hasilatas, hasilbawah int

func faktorial(n int, faktorialnya *int) {
	*faktorialnya = 1
	i = 1

	for i <= n {
		*faktorialnya = *faktorialnya * i
		i += 1
	}
}

func permutasi(a, b int) int {
	faktorial(a, &hasilatas)
	faktorial(a-b, &hasilbawah)

	hasil = hasilatas / hasilbawah
	return hasil
}

func kombinasi(a, b int) int {
	var hasilbawah2 int

	faktorial(a, &hasilatas)
	faktorial(b, &hasilbawah)
	faktorial(a-b, &hasilbawah2)

	hasil = (hasilatas / (hasilbawah * hasilbawah2))
	return hasil
}

func main() {
	var a int
	fmt.Print("Hello World!\n\n")
	faktorial(5, &a)
	fmt.Print(permutasi(6, 4), "\n")
	fmt.Print(kombinasi(6, 4), "\n")
	fmt.Print(a)
}
