package main

import "fmt"

func main() {
	var bilangan, iterasi, hasil  int
	fmt.Scan(&bilangan)

	for iterasi = bilangan; iterasi >= 1; iterasi -= 1 {
		hasil += iterasi
	}
	fmt.Println(hasil)
}
