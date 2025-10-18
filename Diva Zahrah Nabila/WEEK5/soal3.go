package main

import "fmt"

func main() {
	var a, b, i int
	fmt.Scan(&a, &b)

	hasil := 1
	for i = 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println(hasil)
}
