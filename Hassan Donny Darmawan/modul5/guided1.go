package main
import "fmt"
 
func main () {
var x, y, j int
fmt.Scan(&x, &y)

for j = x; j <= y; j+=1 {
	fmt.Print(j," ")
}
}