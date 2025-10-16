package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var jariJari, tinggi int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Scan(&jariJari, &tinggi)

		r := float64(jariJari)
		t := float64(tinggi)
		volume := (1.0 / 3.0) * math.Pi * r * r * t

		fmt.Println(volume)
	}
}
