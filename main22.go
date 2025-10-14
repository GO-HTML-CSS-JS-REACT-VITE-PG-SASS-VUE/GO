package main
import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Метод структуры
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    area := rect.Area()                    // Вызов метода
    fmt.Println("Площадь:", area)
}
//Задание: Добавление методов к структурам
