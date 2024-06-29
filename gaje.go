package main

import "fmt"

const max int = 123

type arr [max]int
type shagy [max]trec

type trec struct {
	v1     int
	v2, v3 int
	v4     int
}

func TambahData(T *arr, n *int) {
	var a int
	a = 0
	*n = 0
	fmt.Scan(&T[a])
	for T[a] != 9999 {
		a++
		*n++
		fmt.Scan(&T[a])
	}

}

func CariSekuensial(T arr, v int, n int) int {
	/*{Mengembalikan lokasi di mana nilai v berada di dalam array tab,
	  atau -1  apabila v tidak ditemukan}
	*/
	var hasil, a int
	a = 0
	hasil = -1
	for a < n {
		if T[a] == v {
			hasil = a
		}
		a++
	}
	return hasil
}

func NilaiMinimum(T arr, n int) int {
	/*{Mengembalikan lokasi nilai minimum yang terdapat pada array
	  yang  berisi n bilangan bulat}
	*/
	var hasil, a, akhir int
	a = 1
	hasil = T[0]
	akhir = 0
	for a < n {
		if T[a] < hasil {
			hasil = T[a]
			akhir = a
		}
		a++
	}
	return akhir
}

func NilaiRerata(T arr, n int) int {
	/*{Mengembalikan nilai rata-rata dari seluruh bilangan
	  yang terdapat pada  array tab yang berisi n bilangan bulat}
	*/
	var rerata, a int
	a = 0
	rerata = 0
	for a < n {
		rerata = rerata + T[a]
		a++
	}

	return rerata
}
func TerurutA(T *arr, n int) {
	/*{I.S. terdefinisi array tab yang berisi n bilangan bulat.
	  F.S. array tab terurut membesar menggunakan algoritma selection sort}
	*/
	var sesi, idx, i, temp int

	sesi = 1
	for sesi <= n-1 {
		idx = sesi - 1
		i = sesi
		for i < n {
			if T[idx] < T[i] {
				idx = i
			}
			i++
		}
		temp = T[sesi-1]
		T[sesi-1] = T[idx]
		T[idx] = temp
		sesi++
	}

}

func TerurutB(T *arr, n int) {
	/*{I.S. terdefinisi array tab yang berisi n bilangan bulat.
	  F.S. array tab terurut mengecil menggunakan algoritma insertion sort}
	*/
	var tahap, tampung, i int
	//algoritma
	tahap = 1
	for tahap <= n-1 {
		i = tahap
		tampung = T[tahap]
		for i > 0 && tampung < T[i-1] {
			T[i] = T[i-1]
			i = i - 1
		}
		T[i] = tampung
		tahap = tahap + 1

	}
}

func CariCepat(T arr, n, v int) int {
	/*{Mengembalikan lokasi di mana nilai v berada di dalam array
	  yang  berisi n bilangan bulat dan terurut mengecil,
	  atau -1 apabila v tidak  ditemukan.
	  Gunakan algoritma pencarian biner/belah tengah}
	*/
	var left, right, mid, x_cari, hasil int
	left = 1
	right = n
	mid = (left + right) / 2
	hasil = -1

	for left <= right && hasil == -1 {
		if T[mid] == x_cari {
			hasil = mid
		} else if T[mid] < x_cari {
			left = mid + 1
		} else if T[mid] > x_cari {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return hasil
}
func Shaggy(T shagy, n int) {
	var i int
	var found1, found2, found bool
	found = false
	i = 2
	for i < n && !found {
		found1 = (T[i-1].v1 == T[i].v2)
		found2 = (T[i].v3 == T[i].v4)
		found = found1 && found2
		i = i + 1
	}
	if found {
		fmt.Println("Ada Shaggy disana")
	} else {
		fmt.Println("tidak ada shaggy dissana")
	}
}
func main() {
	var T arr
	var T1 shagy
	var hasil trec
	var n, v int
	fmt.Scanln(&v)
	TambahData(&T, &n)
	CariSekuensial(T, v, n)
	hasil.v1 = T[NilaiMinimum(T, n)]
	hasil.v2 = NilaiRerata(T, n)
	hasil.v3 = NilaiRerata(T, n) / n
	TerurutA(&T, n)
	hasil.v4 = T[0]
	TerurutB(&T, n)
	CariCepat(T, n, v)
	fmt.Println("nilai terkecil :", hasil.v1)
	fmt.Println("jumlah :", hasil.v2)
	fmt.Println("rerata :", hasil.v3)
	fmt.Println("nilai terbesar:", hasil.v4)
	Shaggy(T1, n)
}
