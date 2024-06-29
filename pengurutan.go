package main

import (
	"fmt"
)

// bikin jadi func coba
func main() {
	var arr [5]int
	var temp int = 0

	// pembacaan elemen array
	fmt.Printf("Masukkan elemen array: \n")
	for i := 0; i <= 4; i++ {
		fmt.Printf("Elemen array ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// pengurutan elemen array secara menaik
	for i := 0; i <= 4; i++ {
		for j := 0; j < 5-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	// pencetakan elemen array setelah diurutkan
	fmt.Println("Array setelah diurutkan secara menaik:")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", arr[i])
	}

	// pengurutan elemen array secara menurun
	for i := 0; i <= 4; i++ {
		for j := 4; j >= i+1; j-- {
			if arr[j] > arr[j-1] {
				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}

	// pencetakan elemen array setelah diurutkan
	fmt.Println("\nArray setelah diurutkan secara menurun:")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
