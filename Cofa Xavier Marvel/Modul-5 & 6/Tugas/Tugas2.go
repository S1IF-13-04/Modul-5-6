package main

import (
	"fmt"
	"math"
)

func main() {
	var n, radius, height int
	fmt.Scan(&n)
	// input n
	for i := 1; i <= n; i++ {
		fmt.Scan(&radius, &height)
		volume := math.Pi * math.Pow(float64(radius), 2) * (float64(height) / 3)
		fmt.Printf("Volume : %.14f\n", volume)
	}
	// for i from 1 to n, do the following:
	// input a and b
	// calculate the volume of the cylinder using the formula π * r^2 * h/3
	// print the volume with 14 decimal places
}
