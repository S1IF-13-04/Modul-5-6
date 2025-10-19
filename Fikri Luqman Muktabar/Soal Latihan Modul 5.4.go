package main

import "fmt"

func main() {
	var n int
	var hasil uint64 = 1

	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus non-negatif!")
	} else {
		for i := 1; i <= n; i++ {
			hasil *= uint64(i)
		}
		fmt.Printf("Hasil faktorial dari %d adalah %d\n", n, hasil)
	}
}