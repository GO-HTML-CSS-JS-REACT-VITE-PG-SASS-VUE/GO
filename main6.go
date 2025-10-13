package main
import "fmt"

func main() {
    firstName := "John"
    lastName := "Doe"
    fullName := firstName + " " + lastName  // Конкатенация строк
    length := len(fullName)                 // Длина строки
    fmt.Println(fullName, "Длина:", length)
}