package main
import "fmt"

func main(){
	var a,n int
	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan Pangkat: ")
	fmt.Scan(&n)

	h:=1
	for i:=1;	i<=n;	i++{
		h *= a
	}
	fmt.Print(h)
}