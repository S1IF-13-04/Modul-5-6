package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Input pertama dan kedua: ")
	fmt.Scan(&x,&y)

	hasiliterasi := 1
	for iterasi := 1; iterasi <= y; iterasi++ {
		hasiliterasi = hasiliterasi * x
		
	}
fmt.Println(hasiliterasi)
}