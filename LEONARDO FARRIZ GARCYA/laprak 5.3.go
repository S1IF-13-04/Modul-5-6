package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Scan(&a, &b)
	for i := 0; i < b; i++ {
		hasil += a
	}
	fmt.Println(hasil)
}