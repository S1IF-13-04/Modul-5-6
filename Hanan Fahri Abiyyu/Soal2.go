package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i, r, t int
	fmt.Print("Masukkan jumlah kerucut yang akan dihitung: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&r, &t)
		jarijari := float64(r)
		tinggi := float64(t)
		volume := (1.0 / 3.0) * math.Pi * jarijari * jarijari * tinggi
		fmt.Println(volume)
	}
}
