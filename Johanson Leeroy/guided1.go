package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b)

	for c = a; c <= b; c++ {
		fmt.Print(c, " ")
	}
}
