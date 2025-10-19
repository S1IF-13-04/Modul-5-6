package main

import "fmt"

func main() {
	var bil, f, hasil int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)

	for f = bil; f > 0; f-- {
		hasil += f
	}
	fmt.Print("hasil: ", hasil)
}
