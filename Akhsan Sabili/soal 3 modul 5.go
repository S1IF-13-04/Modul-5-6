package main

import "fmt"

func main() {
	var n, m, hasil int
	fmt.Scan(&n, &m)
	hasil = 1
	for i := 1; i <= m; i++ {
		hasil = hasil * n
	}
fmt.Println(hasil)
}
