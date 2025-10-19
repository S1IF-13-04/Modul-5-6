package main
import "fmt"

func main(){
	var iterasi, a, b int
	fmt.Scan(&a, &b)
	for iterasi = a; iterasi <= b; iterasi++ {
		fmt.Print(iterasi, " ")
	}
}

// iterasi++
// iterasi += 1
// iterasi = iterasi + 1
// sama saja