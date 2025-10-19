package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var jariJari, tinggi float64
		fmt.Scan(&jariJari, &tinggi)

		volume := (1.0 / 3.0) * math.Pi * jariJari * jariJari * tinggi
		fmt.Println(volume)
	}
}