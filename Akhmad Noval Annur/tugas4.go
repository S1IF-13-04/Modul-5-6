    package main
    import "fmt"

    func main() {
        var n, i int
        var faktorial int = 1
        fmt.Scan(&n)

        for i = 1; i <= n; i++ {
            faktorial *= i
        }

        fmt.Println(faktorial)
    }
