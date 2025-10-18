package main
import "fmt"

func main () {
	var i,a,b,hasil int
	fmt.Scan(&a,&b)
	hasil=1
	for i=1;i<=b;i++ {
		hasil *= a
	}
	fmt.Print(hasil)
}