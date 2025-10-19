package main

import "fmt"

func main() {
	var n,hasil int
	fmt.Scan(&n)
	
	for iterasi := 1; iterasi <= n; iterasi = iterasi + 1 {
		hasil += iterasi
	}
	fmt.Println(hasil)

}
