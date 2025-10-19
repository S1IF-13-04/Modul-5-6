package main

import "fmt"

func main() {
	var bil, f, faktorial int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	faktorial = 1
	for f = bil; f > 0; f-- {
		faktorial *= f
	}
	fmt.Print("hasil faktorial: ", faktorial)
}
