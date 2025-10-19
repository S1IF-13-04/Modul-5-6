package main

import "fmt"

func main() {
	var baris, l int
	var jariAlas, Tinggi, Volume float64
	const PI = 3.1415926535897932384626433
	fmt.Print("Baris: ")
	fmt.Scan(&baris)

	for l = baris; l > 0; l-- {
		fmt.Println("Jari-jari dan Tinggi: ")
		fmt.Scan(&jariAlas, &Tinggi)
		Volume = PI * jariAlas * jariAlas * Tinggi / 3
		fmt.Println("Volume: ", Volume)
	}
}
