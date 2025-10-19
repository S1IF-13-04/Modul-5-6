package main
import "fmt"
func main() {
 var x,i,h,y,t float64
 var v = 3.14
 fmt.Print("Masukan angkanya:")
 fmt.Scan(&x)

 for i=1; i<=x; i++{
	fmt.Print("masukan jari jari dan tinggi: ")
	fmt.Scan(&y, &t)

	h = (1.0/3.0) * v * y * y * t

	fmt.Printf("%g\n",h)
 }

}
