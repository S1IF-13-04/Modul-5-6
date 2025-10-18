package main

import "fmt"

func main() {
	var a, i int
	fmt.Scan(&a)

	hasil := 1
	for i = 1; i <= a; i++ {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}
