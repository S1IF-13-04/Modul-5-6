package main
import (
"fmt"
"math"
)

func main () {
	var n,r,t,i int
	var v float64

	fmt.Scan(&n)

	for i = 1; i<=n;i++ {
		fmt.Scan(&r,&t)
		v = (math.Pi/3.0) * math.Pow(float64(r),2) * float64(t)
		fmt.Printf("%g\n",v)
	}
}