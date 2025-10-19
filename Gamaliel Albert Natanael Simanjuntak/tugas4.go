package main
import "fmt"
func main() {
	var n int64
	fmt.Scan(&n)
	var hasil int64 = 1
	var i int64
	for i = 1; i <= n; i++ {
		hasil *= i
	}
	fmt.Println(hasil)
}