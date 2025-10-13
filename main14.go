package main
import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}  // Создание среза
    slice = append(slice, 6)       // Добавление элемента
    fmt.Println("Срез:", slice)
    fmt.Println("Длина:", len(slice))  // Длина среза
}
//Задание: Базовые операции со срезами