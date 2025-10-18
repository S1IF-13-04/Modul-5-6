package main
import "fmt"
func main() {
	var n,total int
	fmt.Print("Masukan bilangan bulat non negatif : ")
	fmt.Scan(&n)
	total = 1
	for i := n; i >= 1; i -= 1 {
		total = total * i
	}
	fmt.Print("Hasil faktorial : ", total)
}