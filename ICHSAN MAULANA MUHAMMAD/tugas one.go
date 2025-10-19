package main

import "fmt"

func main() {
	var n, i, jumlah int
	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		jumlah = jumlah + i
	}
	fmt.Println(jumlah)
}
