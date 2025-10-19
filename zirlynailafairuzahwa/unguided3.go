package main
import "fmt"

func main(){
	var bilangan1, bilangan2, hasil int
	fmt.Scan(&bilangan1, &bilangan2)
	hasil = 1
	for i := 1; i <= bilangan2; i++{
		hasil *= bilangan1
	}
	fmt.Println(hasil)
}