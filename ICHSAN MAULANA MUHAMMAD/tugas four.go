package main

import "fmt"

func main() {
	var n, f, hasil int

	fmt.Scan(&n)

	hasil = 1
	for f = n; f > 1; f-- {
		hasil *= f
	}

	fmt.Print("", hasil)
}
