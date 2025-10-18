package main

import "fmt"

func main() {
	var n int
	var total int = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Println(total)
}
