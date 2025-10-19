package main
import "fmt"

func main(){
	var n, alas, tinggi int
	fmt.Scan(&n)
	for iterasi:= 1; iterasi <= n; iterasi++{
		fmt.Scan(&alas, &tinggi)
		luas := 0.5 * float64(alas * tinggi)
		fmt.Println(luas)
	}  
}