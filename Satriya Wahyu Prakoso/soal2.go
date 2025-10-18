package main
import (
	"fmt"
	"math"
)
func main() {
	var r, t, n int
	fmt.Print("Masukan jumlah kerucut: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i = i + 1 {
		fmt.Print("Masukan jari-jari alas dan tinggi: ")
		fmt.Scan(&r, &t)
		volume := math.Pi * 1.0 / 3.0 * float64(r) * float64(r) * float64(t)
		fmt.Println("Volume: ", volume)
	}
}