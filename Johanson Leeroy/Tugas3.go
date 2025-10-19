package main

import "fmt"

func main() {
	var b1, b2, f, hasil int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b1)
	fmt.Print("Pangkat: ")
	fmt.Scan(&b2)
	hasil = 1
	for f = b2; f > 0; f-- {
		hasil *= b1
	}
	fmt.Print("hasis: ", hasil)
}
