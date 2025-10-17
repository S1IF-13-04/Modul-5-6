package main

import "fmt"

func main() {
	var b1, b2 int
	var iterasi int
	fmt.Scan(&b1, &b2)

	for iterasi = b1; iterasi <= b2; iterasi += 1 {
		fmt.Print(iterasi, " ")
	}
}
