package main

import "fmt"

func main() {
	var iterasi, b1, b2 int
	var hasil int
	fmt.Scan(&b1, &b2)
	for iterasi = 1; iterasi <= b2; iterasi = iterasi + 1 {
		hasil += b1
	}
	fmt.Println(hasil)
}
