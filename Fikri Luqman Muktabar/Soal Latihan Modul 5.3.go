package main

import "fmt"

func main() {
	var a, b int
	var hasil int = 1

	fmt.Print("Masukkan dua bilangan bulat positif: ")
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		hasil *= a
	}

	fmt.Printf("Hasil dari %d pangkat %d adalah %d\n", a, b, hasil)
}