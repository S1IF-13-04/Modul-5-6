package main

import "fmt"

func main() {
	var n int
	var a int
	var hasiliterasi int

	fmt.Print("Input pertama: ")
	fmt.Scan(&n)
	fmt.Print("Input kedua: ")
	fmt.Scan(&a)
	
	for iterasi := 1; iterasi <= n; iterasi++ {
		hasiliterasi += a
	}
	fmt.Println(hasiliterasi)

}
