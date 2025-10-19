package main

import "fmt"

func main() {
	var bilangan, pangkat, i, Total int
	fmt.Print("Masukkan Bilangan :")
	fmt.Scan(&bilangan)
	fmt.Print("Masukkan Pangkat :")
	fmt.Scan(&pangkat)
	Total = 1
	for i = 0; i < pangkat; i++{
		Total = Total * bilangan
	}
	fmt.Printf("%d", Total)
}
