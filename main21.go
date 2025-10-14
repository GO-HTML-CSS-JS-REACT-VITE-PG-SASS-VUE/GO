package main
import "fmt"
type Person struct {            // Определение структуры
    Name string
    Age  int
}
func main() {
    person := Person{           // Создание экземпляра структуры
        Name: "Alice",
        Age:  25,
    }
    fmt.Println(person.Name, person.Age)
}
//Задание: Создание и использование структур