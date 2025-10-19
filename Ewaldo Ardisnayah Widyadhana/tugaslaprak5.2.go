package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var r, t float64
		fmt.Scan(&r, &t)
		fmt.Println((math.Pi * r * r * t) / 3)
	}
}
