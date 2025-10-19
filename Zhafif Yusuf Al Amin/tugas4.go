package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	hasil := 1
	for iterasi := n; iterasi >= 1; iterasi-- {
		hasil = hasil * iterasi
	}
	fmt.Println(hasil)
	
}
