package main

import "fmt"

// this program calculates the area of n triangles based on user input
// input the number of triangles, followed by the base and height of each triangle
// output the area of each triangle

func main() {
	var base, height, n, h int
	var area float64
	// declare base, height, n, h as int
	// declare area as float64

	fmt.Scan(&n)
	// input n
	// n is the number of triangles

	for j := h; j < n; j += 1 {
		fmt.Scan(&base, &height)
		area = 0.5 * float64(base*height)
		fmt.Printf("Area : %.2f", area)
	}
	// for j from 1 to n, do the following:
	// input base and height
	// calculate the area of the triangle using the formula 0.5 * base * height
	// print the area
}
