package main

import (
	"fmt"
)

const NMAX int = 128

type taabNya [NMAX]int

func cariMaksimum(data taabNya, a int, iy *int) {
	nilaiMaks := 0

	for i := 0; i < a; i++ {
		if data[i] > nilaiMaks {
			nilaiMaks = data[i]
			*iy = i
		}
	}
}

func tukarBilangan(x, y *int) {
	sementara := 0

	sementara = *x
	*x = *y
	*y = sementara
}

func main() {
	a, b := 4, 7
	tukarBilangan(&a, &b)

	fmt.Println(a, b)
}
