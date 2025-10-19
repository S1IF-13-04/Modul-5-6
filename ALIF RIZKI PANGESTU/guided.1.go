package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("masukan angka dari var a dan b: ")
	fmt.Scanln(&a, &b)
	for i := a; i <= b; i++ {
		fmt.Println(i)

	}
}
