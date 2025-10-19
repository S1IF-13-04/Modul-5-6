package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&n) 

	for i := 0; i < n; i++ {
		var alas, tinggi float64
		fmt.Scan(&alas, &tinggi)
		luas := (alas * tinggi) / 2
		fmt.Println(luas)
	}
}
