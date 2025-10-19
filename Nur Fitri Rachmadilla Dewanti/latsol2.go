package main
import (
"fmt"
"math"
)

func main() {
	var j, n int
	var r, t float64
	fmt.Scan(&n)

	for j = 1; j <=n; j+=1 { 
        fmt.Scan(&r, &t) 
        volume := (1.0/3.0) * math.Pi * r * r * t
        fmt.Println(volume)
    }
}