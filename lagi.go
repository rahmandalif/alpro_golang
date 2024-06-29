package main

import (
	"fmt"
)

func main() { //program cek 2 angka genap dari inputan bilangan antara 1000 samape 9999
	var bilangan, cek1, cek2 int
	var ngecek1, ngecek2, ngecek3, ngecek4 int
	var hasil1, hasil2, hasil3, hasil4 bool
	var jadinya bool

	fmt.Print("masukkan bilangan antara 1111 sampai 9999\n")
	fmt.Scan(&bilangan)

	cek1 = bilangan / 100   // ini untuk ngehasilin 2 bilangan awal dari inputan bilangan
	ngecek1 = cek1 / 10     // ini untuk ngehasilin bilangan awal dari cek1
	ngecek2 = cek1 % 10     // ini untuk ngehasilin bilangan akhir dari cek1
	hasil1 = ngecek1%2 == 0 // ini untuk nentuin bilangan ngecek1 genap atau bukan
	hasil2 = ngecek2%2 == 0 // ini juga

	cek2 = bilangan % 100   // ini untuk ngehasilin 2 bilangan akhir dari inputan bilangan
	ngecek3 = cek2 / 10     // ini untuk ngehasilin bilangan awal dari cek2
	ngecek4 = cek2 % 10     // ini untuk ngehasilin bilangan akhir dari cek2
	hasil3 = ngecek3%2 == 0 // ini sama kayak baris 18
	hasil4 = ngecek4%2 == 0 // ini juga sama

	jadinya = (hasil1 || hasil2) && (hasil3 || hasil4) // ini buat hasil akhirnya

	fmt.Println(ngecek1, ngecek2, ngecek3, ngecek4)
	fmt.Println("apakah bilangan ini mengandung 2 angka genap?")
	fmt.Print(jadinya) // di print deh

}
