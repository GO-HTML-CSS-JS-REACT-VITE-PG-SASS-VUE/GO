package main
import "fmt"

func add(a int, b int) int {     // Объявление функции
    return a + b                 // Возврат значения
}

func main() {
    result := add(5, 3)          // Вызов функции
    fmt.Println("Результат:", result)
}
//Задание: Создание и вызов функций