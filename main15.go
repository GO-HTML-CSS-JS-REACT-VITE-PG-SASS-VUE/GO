package main
import "fmt"

func main() {
    ages := make(map[string]int)  // Создание карты
    ages["Alice"] = 25            // Добавление элемента
    ages["Bob"] = 30
    fmt.Println("Возраст Alice:", ages["Alice"])  // Доступ к элементу
}
//Задание: Работа с ассоциативными массивами (map)