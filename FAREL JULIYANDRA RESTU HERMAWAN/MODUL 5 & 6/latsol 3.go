package main

import "fmt"

func main(){
	var a, b, i int
	var hasil int = 1
	fmt.Print("masukan a dan b : ")
	fmt.Scan(&a,&b)

	for i = 1; i <= b; i++ {
		hasil = hasil * a
	}
	fmt.Println("Hasil : ", hasil)
}