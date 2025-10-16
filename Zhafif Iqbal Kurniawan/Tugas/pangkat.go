package main

import "fmt"

func main() {
	var b1, pangkat int
	fmt.Scan(&b1, &pangkat)

	hasil := 1
	for i := 0; i < pangkat; i++ {
		hasil = hasil * b1
	}

	fmt.Println(hasil)
}
