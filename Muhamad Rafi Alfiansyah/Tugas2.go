package main
import (
	"fmt" 
	"math"
)

func main(){
	var r, t, n float64
	fmt.Scan(&n)
	
	v := 0.0
	for i:=1.0;	i<=n;	i++{
		fmt.Scan(&r,&t)
		v = (1.0/3.0) * math.Pi * r * r * t
		fmt.Print(v)
	}
}