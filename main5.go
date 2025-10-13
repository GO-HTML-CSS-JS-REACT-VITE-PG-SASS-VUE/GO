package main
import "fmt"
func main() {
    a, b := 15, 4
    sum := a + b           // Сложение
    diff := a - b          // Вычитание
    product := a * b       // Умножение
    quotient := a / b      // Деление
    remainder := a % b     // Остаток от деления
    fmt.Println(sum, diff, product, quotient, remainder)
}