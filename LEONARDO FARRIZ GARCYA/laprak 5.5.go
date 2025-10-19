package main

import "fmt"

func main() {
    const pi = 3.141592653589793
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var r, t float64
        fmt.Scan(&r, &t)
        fmt.Println((1.0 / 3.0) * pi * r * r * t)
    }
}
