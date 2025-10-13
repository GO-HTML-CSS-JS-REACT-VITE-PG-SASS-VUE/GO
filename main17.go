package main
import "fmt"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil           // Возврат двух значений
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Результат:", result)
    }
}
//Задание: Функции с множественными возвращаемыми значениями