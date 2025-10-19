package main

import "fmt"

func main() {
	var a, b, c, hasil int

	fmt.Scan(&a, &b)

	for c = 0; c < a; c++ {
		hasil += b
	}
	fmt.Print(hasil)
}
