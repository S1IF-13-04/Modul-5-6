package main
import "fmt"

func main () {
	var n,i int
	fmt.Scan(&n)
	hasil :=0
	for i = 1; i <= n ; i++ {
		hasil = hasil + i
	}
	fmt.Print(hasil)
}