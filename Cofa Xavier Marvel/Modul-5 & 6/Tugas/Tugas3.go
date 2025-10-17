package main

import (
	"fmt"
)

func main() {
	var x, n int
	// declare x and n as int
	fmt.Scan(&x, &n)
	// input n
	for i := 1; i < n; i++ {
		x = x + x
	}
	// for i from 1 to n, do the following:
	// add x to x
	// increment i by 1
	fmt.Print(x)

}
