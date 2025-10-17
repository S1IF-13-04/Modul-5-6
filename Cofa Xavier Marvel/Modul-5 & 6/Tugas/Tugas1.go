package main

import "fmt"

func main() {
	var m, n int
	// declare m and n as int
	fmt.Scan(&n)
	// input n
	for i := 0; i <= n; i++ {
		m = m + i
	}
	// for i from 0 to n, do the following:
	// add i to m
	// increment i by 1
	fmt.Print(m)
}
