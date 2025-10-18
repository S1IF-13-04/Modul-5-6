package main
import(
	"fmt"
	"math"
)
func main (){
	var a, i int
	var volume, r, t float64
	fmt.Print("masukan angka: ")
	fmt.Scan(&a)
	for i = 1; i <= a; i++ {
		fmt.Print("masukan r dan t: ")
		fmt.Scan(&r, &t)
		volume = (math.Pi * r * r * t) / 3
		fmt.Printf("keluaran: %.15f\n",volume)
	}
}