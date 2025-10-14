package main
import "fmt"
func main() {
    x := 10
    ptr := &x                    // Получение указателя
    fmt.Println("Значение x:", x)
    fmt.Println("Адрес x:", ptr)
    fmt.Println("Значение через указатель:", *ptr)  // Разыменование
    *ptr = 20                    // Изменение значения через указатель
    fmt.Println("Новое значение x:", x)
}
//Задание: Основы работы с указателями