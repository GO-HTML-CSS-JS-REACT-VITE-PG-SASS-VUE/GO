package main
import "fmt"

func main() {
    count := 1
    for count <= 5 {               // Цикл while-style
        fmt.Println("Счетчик:", count)
        count++
    }
}