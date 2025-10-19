package main

import "fmt"

func main() {
	var n, i int
	var hasil int
	fmt.Scan(&n)
	for i = 0; i <= n; i++ {
		hasil = hasil + i
	}
	fmt.Println(hasil)
}
