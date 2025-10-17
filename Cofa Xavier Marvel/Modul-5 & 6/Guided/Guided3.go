package main

import "fmt"

func main() {
	var j, v1, v2, result int

	// declare j, v1, v2, result as int

	fmt.Scan(&v1, &v2)
	// input v1 and v2 in that order
	result = 0
	// initialize result to 0
	for j = 1; j <= v2; j += 1 {
		result = result + v1
	}
	// for j from 1 to v2, do the following:
	// add v1 to result
	// increment j by 1

	fmt.Println(result)
	// print result
}
