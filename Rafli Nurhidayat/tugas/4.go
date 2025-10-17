package main

import "fmt"

func main() {
	var bilangan, iterasi, hasil int
	fmt.Scan(&bilangan)
	hasil = 1
	for iterasi = bilangan; iterasi >= 1; iterasi -= 1 {
		hasil = hasil * iterasi
	}
	fmt.Println(hasil)
}
