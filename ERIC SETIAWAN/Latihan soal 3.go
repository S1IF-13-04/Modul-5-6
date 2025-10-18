package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	hasil := 1
	for i := 1; i <= y; i++ {
		hasil *= x
	}
	fmt.Println(hasil)
}
