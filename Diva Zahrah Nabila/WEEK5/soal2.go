package main

import "fmt"

func main() {
	var n int
	const pi = 3.1415926535
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var r, t float64
		fmt.Scan(&r, &t)
		fmt.Println((1.0 / 3.0) * pi * r * r * t)
	}

}
