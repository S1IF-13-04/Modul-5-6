package main

import "fmt"

func main() {
	var n, m int
	var alas, tinggi, luas float64
	fmt.Scan(&n)

	for m = 1; m <= n; m = m + 1 {
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * alas * tinggi
		fmt.Println(luas)
	}
}
