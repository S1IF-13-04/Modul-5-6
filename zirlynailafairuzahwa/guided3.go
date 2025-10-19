package main
import "fmt"

func main(){
	var a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = 0
	for iterasi := 1; iterasi <= b; iterasi += 1 {
		hasil += a
	}
	fmt.Println(hasil)
}