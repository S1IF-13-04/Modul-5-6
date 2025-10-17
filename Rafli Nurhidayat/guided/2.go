package main

import "fmt"

func main() {
	var iterasi, alas, tinggi, n int
	var luas float64
	fmt.Scan(&n)
	// for iterasi = 1; iterasi <= n; iterasi += 1 {
	for iterasi = n; iterasi >= 1; iterasi -= 1 {
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas*tinggi)
		fmt.Println("hasil:", luas)
	}
}
