package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var alas, tinggi float64
		fmt.Scan(&alas, &tinggi)

		luas := 0.5 * alas * tinggi
		fmt.Println(luas)
	}
}
