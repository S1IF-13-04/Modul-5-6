package main

import (
	"fmt"
	"math"
)

func main() {
	var n, jari, tinggi int
	var volume float64

	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&jari, &tinggi)
		volume = (1.0 / 3.0) * math.Pi * math.Pow(float64(jari), 2) * float64(tinggi)
		fmt.Println(volume)
	}
}
