package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJariAlasKerucut, tinggiKerucut, iterasi, n int
	var volume float64

	fmt.Scan(&n)
	for iterasi = 1; iterasi <= n; iterasi += 1 {
		fmt.Scan(&jariJariAlasKerucut, &tinggiKerucut)
		volume = (1.0 / 3.0) * math.Pi * math.Pow(float64(jariJariAlasKerucut), 2) * float64(tinggiKerucut)
		fmt.Println("Volume: ", volume)
	}
}
