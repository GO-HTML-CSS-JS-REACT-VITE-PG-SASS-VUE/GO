package main
import "fmt"

func main() {
    // Анонимная функция
    multiply := func(x, y int) int {
        return x * y
    }
    
    result := multiply(4, 5)
    fmt.Println("Умножение:", result)
}