package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Print("masukan bilangan pertma: ")
	fmt.Scan(&a)
	fmt.Print("masukan bilangan ke dua: ")
	fmt.Scan(&b)
	hasil = 0

	for i := 0; i < b; i++ {
		hasil += a
	}
	fmt.Printf("hasil perkalian: %d\n", hasil)

}
