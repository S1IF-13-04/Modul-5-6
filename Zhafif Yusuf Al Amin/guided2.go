package main
import "fmt"
func main() {
	var n int
	var a, t float64

	fmt.Print("Input batas: ")
	fmt.Scan(&n)

	
	for iterasi := 2; iterasi <= n+1; iterasi = iterasi + 1 {
		fmt.Print("Input alas dan tinggi: ")
		fmt.Scan(&a,&t)
		L := 0.5 * a * t
		fmt.Println(L)
	}

}
