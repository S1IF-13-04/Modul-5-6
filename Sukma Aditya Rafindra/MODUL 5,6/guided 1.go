package main

import "fmt"

func main() {
    var a, b int

	fmt.Print("masukan 2 angka: ")
    fmt.Scan(&a, &b)

    for i := a; i <= b; i++ {
        fmt.Print(i, " ")
   }
}
