package main

import "fmt"

func main() {
	var n, i, jumlah int

	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	jumlah = 0
	for i = 1; i <= n; i++ {
		jumlah += i
	}

	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah %d\n", n, jumlah)
}