package main

import "fmt"

func main() {
	var n, i, hasil int
	fmt.Print("input: ")
	fmt.Scan(&n)
	hasil = 0

	for i = 1; i <= n; i = i + 1 {
		hasil = hasil + i
	}
	fmt.Println(hasil)
}