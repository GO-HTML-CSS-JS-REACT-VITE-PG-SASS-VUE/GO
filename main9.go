package main
import "fmt"

func main() {
    day := "Monday"
    switch day {                    // Оператор выбора
    case "Monday":
        fmt.Println("Понедельник")
    case "Tuesday":
        fmt.Println("Вторник")
    default:
        fmt.Println("Другой день")
    }
}