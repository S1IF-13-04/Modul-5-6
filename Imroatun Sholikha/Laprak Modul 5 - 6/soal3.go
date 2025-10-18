package main

import "fmt"

func main() {
	var x, y, i int
	var hasil int
	fmt.Print("input: ")
	fmt.Scan(&x, &y)

	hasil = 1
	for i = 1; i <= y; i = i + 1 {
		hasil = hasil * x
	}
	fmt.Println(hasil)
}
