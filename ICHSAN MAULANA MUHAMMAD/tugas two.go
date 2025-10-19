package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const phi = 3.141592653589793
	for i := 0; i < n; i++ {
		var r, t float64
		fmt.Scan(&r, &t)
		volume := (1.0 / 3.0) * phi * r * r * t
		fmt.Println(volume)
	}
}
