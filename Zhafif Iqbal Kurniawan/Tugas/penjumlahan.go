package main

import "fmt"

func main() {
	var n int
	var jumlah int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		jumlah += i
	}

	fmt.Println(jumlah)
}
