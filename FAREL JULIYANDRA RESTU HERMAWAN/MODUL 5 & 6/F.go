package main

import "fmt"

func main(){
	var t int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&t)
	for i := 1; i <= t; i ++ {
		for b := 1 ; b <=t-i; b++ {			
			fmt.Print(" ")
		}
		for s := 1; s <= i; s++ {
			fmt.Print(" *")
		}
		fmt.Println()

	}
}
