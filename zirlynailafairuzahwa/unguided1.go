package main
import "fmt"

func main(){
	var i, n, hasil int
	fmt.Scan(&n)
	for i = 1; i <= n; i++{
		hasil += i
	}
	fmt.Println(hasil)
}