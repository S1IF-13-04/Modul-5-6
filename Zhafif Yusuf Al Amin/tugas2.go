package main

import (
"fmt"
"math"
)
func main() {
	var n int
	var r,t float64
	fmt.Scan(&n)
	
	for iterasi := 1; iterasi <= n; iterasi = iterasi + 1 {
		fmt.Scan(&r,&t)
		V:=  (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println(V)
	}
	

}
