package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, t float64
		fmt.Scan(&a, &t)
		luas := 0.5 * a * t
		fmt.Println(luas)
	}
}
