package main
import "fmt"
func main() {
	var a, i int
	var hasil int = 1
	fmt.Print("masukan a: ")
	fmt.Scan(&a)
	for i = 1; i <= a; i++ {
		hasil = hasil * i
	}
	fmt.Println("Hasil faktorial:", hasil)

}