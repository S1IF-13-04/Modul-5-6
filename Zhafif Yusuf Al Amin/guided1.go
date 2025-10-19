package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Input: ")
	fmt.Scan(&m,&n)
	
	for iterasi := m; iterasi <= n; iterasi = iterasi + 1 {
		fmt.Println(iterasi)
	}
}
