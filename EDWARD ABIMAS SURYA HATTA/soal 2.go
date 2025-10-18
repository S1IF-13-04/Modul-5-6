package main
import (
    "fmt"
    "math"
)
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var r, h float64
        fmt.Scan(&r, &h)
        volume := (1.0 / 3.0) * math.Pi * r * r * h
        fmt.Println(volume)
    }
}