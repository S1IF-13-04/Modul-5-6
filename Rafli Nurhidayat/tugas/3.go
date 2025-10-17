package main

import "fmt"

func main() {
	var b1, b2, iterasi, hasil int

	fmt.Scan(&b1, &b2)
	hasil = 1
	for iterasi = 1; iterasi <= b2; iterasi += 1 {
		hasil = hasil * b1
	}
	fmt.Println(hasil)
}
