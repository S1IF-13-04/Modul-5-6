package main
import "fmt"

func main(){
	var bilangan, faktorial int
	fmt.Scan(&bilangan)
	faktorial = 1
	for i := 1; i <= bilangan; i++{
		faktorial *= i
	}
	fmt.Println(faktorial)
}