package main

import "fmt"

func main() {
	var n int64

	fmt.Print("Masukkan bilangan bulat non-negatif (n): ")
	fmt.Scan(&n)
	var faktorial int64 = 1

	for i := int64(1); i <= n; i++ {
		faktorial *= i 
	}
	
	fmt.Printf("Faktorial dari %d (%d!) adalah: %d\n", n, n, faktorial)

}