package main
import "fmt"

func main() {
    i := 0
    for {                         // Бесконечный цикл
        fmt.Println(i)
        i++
        if i >= 3 {               // Условие выхода
            break
        }
    }
}