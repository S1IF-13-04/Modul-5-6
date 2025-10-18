package main
import "fmt"

func main () {
	var a,b,i int
	fmt.Print("masukan 2 bilangan yang akan dikalikan: ")
	fmt.Scanln(&a,&b)
	hasil := 0

	for i = 1; i <= b; i += 1 {
		hasil += a
	}
	fmt.Print("maka hasil perkaliannya adalah: ", hasil)
}