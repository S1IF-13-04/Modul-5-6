package main

import (
	"fmt"
	"math"
)

func main() {
	var n, jari, tinggi, i int
	var volume float64
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i = i + 1 {
		fmt.Print("Input jari dan tinggi: ")
		fmt.Scan(&jari, &tinggi)
		volume = (1.0 / 3.0) * math.Pi * float64(jari*jari*tinggi)
		fmt.Println(volume)

	}

}
