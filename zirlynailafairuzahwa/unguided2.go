package main
import (
	"fmt"
	"math"
)

func main(){
	var kerucut, jariJari, tinggi int
	fmt.Scan(&kerucut)
	for i := 1; i <= kerucut; i++ {
		fmt.Scan(&jariJari, &tinggi)
		volume := (1.0 / 3.0) * math.Pi * float64(jariJari) * float64(jariJari) * float64(tinggi)
		fmt.Println(volume)
	}
}