package main

import "fmt"

func main() {
	for baris := 0; baris <= 4; baris++ {
		for kolom := 0; kolom <= 4; kolom++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
