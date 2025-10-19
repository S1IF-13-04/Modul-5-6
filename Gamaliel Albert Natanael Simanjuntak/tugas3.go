package main
import "fmt"
func main() {
    var basis, pangkat int
    fmt.Scan(&basis, &pangkat)
    var hasil int = 1
    for i := 0; i < pangkat; i++ {
        hasil *= basis
    }
    fmt.Println(hasil)
}