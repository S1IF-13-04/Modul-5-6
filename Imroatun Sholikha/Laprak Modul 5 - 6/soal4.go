package main

import "fmt"

func main() {
	var x, i int
	var hasil int
	fmt.Print("input: ")
	fmt.Scan(&x)

	hasil = 1
	for i = 1; i <= x; i = i + 1 {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}
