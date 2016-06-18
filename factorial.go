package main
import "fmt"

func factorial(n int) int {
    if n ==0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
	var num int
    fmt.Scanf("%v", &num)
    fmt.Printf("%v!=%d\n", num,factorial(num))
}