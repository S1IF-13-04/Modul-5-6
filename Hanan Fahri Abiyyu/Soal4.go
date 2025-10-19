package main

import "fmt"

func main() {
	var n, total int
	fmt.Print("Masukkan angka :")
	fmt.Scan(&n)
	total = 1
	for bilangan := 1; bilangan <= n; bilangan++{
		total = total * bilangan
	}
	fmt.Printf("%d", total)
}
