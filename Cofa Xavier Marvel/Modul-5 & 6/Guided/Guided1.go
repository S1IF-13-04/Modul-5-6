package main

import "fmt"

func main() {

	var j, h, k float64
	fmt.Scan(&h, &k)

	for j = h; j <= k; j += 1 {
		fmt.Print(j, " ")

	}
}

// print from h to k with a space in between
// if k < h, do not print anything
// both h and k are float64
// input h and k in that order
// example: if h = 2 and k = 5, the output is 2 3 4 5
// example: if h = 5 and k = 2, the output is (nothing)
