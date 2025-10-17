package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	// declare x as int
	// input x

	for i := x - 1; i >= 1; i-- {
		x = x * i
	}
	// for i from x-1 to 1, do the following:
	// decrement i by 1
	// multiply x by i
	// until i is less than 1

	fmt.Print(x)
}
