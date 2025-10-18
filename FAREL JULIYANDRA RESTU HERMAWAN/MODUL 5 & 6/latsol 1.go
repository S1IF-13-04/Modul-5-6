package main

import "fmt"

func main(){
	var a, i, sum int
	fmt.Print("masukan angka: ")
	fmt.Scan(&a)
	sum = 0
	for i = 1; i <= a; i++{
		sum += i
	}
	fmt.Print("keluaran :",sum)
}