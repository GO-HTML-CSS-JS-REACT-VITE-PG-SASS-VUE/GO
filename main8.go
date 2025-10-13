package main
import "fmt"

func main() {
    score := 85
    if score >= 90 {
        fmt.Println("Отлично")
    } else if score >= 70 {
        fmt.Println("Хорошо")
    } else {
        fmt.Println("Удовлетворительно")
    }
}