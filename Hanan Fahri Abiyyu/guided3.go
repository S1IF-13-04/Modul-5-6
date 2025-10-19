package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Print("Input :")
	fmt.Scan(&a, &b)
	for i := 1; i <= b; i++ {
		hasil += a
	}
	fmt.Println(hasil)
}
