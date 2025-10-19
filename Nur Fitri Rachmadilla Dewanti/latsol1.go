package main
import "fmt"

func main() {
	var j, n int
	var hasil int
	fmt.Scan(&n)
	hasil = 0
	for j = 1; j <= n; j+=1 {
		hasil = n * (n+1)/2
	}
	fmt.Println(hasil)
}