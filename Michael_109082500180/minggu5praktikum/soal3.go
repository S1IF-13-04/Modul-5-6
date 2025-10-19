package main
import "fmt"
func main() {
 var x,i,y float64
 var h float64 = 1  
 fmt.Print("Masukan angkanya:")
 fmt.Scan(&y, &x)

 for i=1; i<=x; i++{
	
	h = h * y
 }
 fmt.Printf("%g\n",h)
}
