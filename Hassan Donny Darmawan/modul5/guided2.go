package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan nilai pertama: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var alas, tinggi, luas float64
		fmt.Scanln(&alas, &tinggi)
		luas = 0.5 * alas * tinggi
		fmt.Println("luasnya", luas)
	}
}