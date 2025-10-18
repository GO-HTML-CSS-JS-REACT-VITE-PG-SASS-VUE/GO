# Основы Go
1. Простая программа
```go
package main  // Объявление основного пакета
import "fmt"  // Импорт пакета для ввода-вывода
func main() {  // Главная функция программы
    fmt.Println("Hello, World!")  // Вывод текста в консоль
}
//Задание: Базовая структура программы Go
```
2. Объявление переменных
```go
package main
import "fmt"

func main() {
    var name string = "Alice"  // Явное объявление переменной
    age := 25                  // Короткое объявление с выводом типа
    fmt.Println(name, age)     // Вывод значений
}
//Задание: Демонстрация способов объявления переменных
```
3. Константы
```go
package main
import "fmt"
func main() {
    const Pi = 3.14159        // Объявление константы
    const AppName = "MyApp"   // Строковая константа
    fmt.Println(Pi, AppName)  // Вывод констант
}
//Задание: Работа с константами в Go
```
4. Основные типы данных
```go
package main
import "fmt"

func main() {
    var integer int = 42          // Целое число
    var floating float64 = 3.14   // Число с плавающей точкой
    var text string = "Go lang"   // Строка
    var boolean bool = true       // Логический тип
    fmt.Println(integer, floating, text, boolean)
}
//Задание: Демонстрация основных типов данных
```
5. Арифметические операции
```go
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
//Задание: Базовые арифметические операции
```
6. Строковые операции
```go
package main
import "fmt"

func main() {
    firstName := "John"
    lastName := "Doe"
    fullName := firstName + " " + lastName  // Конкатенация строк
    length := len(fullName)                 // Длина строки
    fmt.Println(fullName, "Длина:", length)
}
//Задание: Операции со строками
```
7. Условные операторы
```go
package main
import "fmt"

func main() {
    age := 18
    if age >= 18 {                    // Простое условие
        fmt.Println("Совершеннолетний")
    } else {
        fmt.Println("Несовершеннолетний")
    }
}
//Задание: Базовые условные конструкции
```
8. Множественные условия
```go
package main
import "fmt"

func main() {
    score := 85
    if score >= 90 {
        fmt.Println("Отлично")
    } else if score >= 70 {
        fmt.Println("Хорошо")
    } else {
        fmt.Println("Удовлетворительно")
    }
}
//Задание: Использование else if для множественных условий
```
9. Оператор switch
```go
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
//Задание: Использование switch для выбора
```
10. Цикл for
```go
package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {      // Классический цикл for
        fmt.Println("Итерация:", i)
    }
}
//Задание: Базовый цикл for
```
11. While-подобный цикл
```go
package main
import "fmt"

func main() {
    count := 1
    for count <= 5 {               // Цикл while-style
        fmt.Println("Счетчик:", count)
        count++
    }
}
//Задание: Цикл с условием (аналог while)
```
12. Бесконечный цикл
```go
package main
import "fmt"

func main() {
    i := 0
    for {                         // Бесконечный цикл
        fmt.Println(i)
        i++
        if i >= 3 {               // Условие выхода
            break
        }
    }
}
//Задание: Бесконечный цикл с break
```
13. Массивы
```go
package main
import "fmt"

func main() {
    var numbers [3]int           // Объявление массива
    numbers[0] = 10              // Присвоение значения
    numbers[1] = 20
    numbers[2] = 30
    fmt.Println(numbers)         // Вывод всего массива
}
//Задание: Работа с массивами
```
14. Срезы (Slices)
```go
package main
import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}  // Создание среза
    slice = append(slice, 6)       // Добавление элемента
    fmt.Println("Срез:", slice)
    fmt.Println("Длина:", len(slice))  // Длина среза
}
//Задание: Базовые операции со срезами
```
15. Карты (Maps)
```go
package main
import "fmt"

func main() {
    ages := make(map[string]int)  // Создание карты
    ages["Alice"] = 25            // Добавление элемента
    ages["Bob"] = 30
    fmt.Println("Возраст Alice:", ages["Alice"])  // Доступ к элементу
}
//Задание: Работа с ассоциативными массивами (map)
```
16. Функции
```go
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
```
17. Функция с несколькими возвращаемыми значениями
```go
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
```
18. Анонимные функции
```go
package main
import "fmt"

func main() {
    // Анонимная функция
    multiply := func(x, y int) int {
        return x * y
    }
    
    result := multiply(4, 5)
    fmt.Println("Умножение:", result)
}
//Задание: Использование анонимных функций
```
19. Замыкания
```go
package main
import "fmt"

func counter() func() int {
    count := 0
    return func() int {         // Замыкание
        count++
        return count
    }
}

func main() {
    next := counter()
    fmt.Println(next())  // 1
    fmt.Println(next())  // 2
    fmt.Println(next())  // 3
}
//Задание: Демонстрация замыканий
```
20. Указатели
```go
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
```
21. Структуры
```go
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
```
22. Методы структур
```go
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
```
23. Интерфейсы
```go
package main
import "fmt"

type Shape interface {           // Объявление интерфейса
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {  // Реализация интерфейса
    return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
    fmt.Println("Площадь:", s.Area())
}

func main() {
    circle := Circle{Radius: 5}
    printArea(circle)            // Использование интерфейса
}
//Задание: Базовое использование интерфейсов
```
24. Горутины
```go
package main
import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 3; i++ {
        fmt.Println("Число:", i)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    go printNumbers()           // Запуск горутины
    time.Sleep(time.Second * 2) // Ожидание завершения
    fmt.Println("Главная функция завершена")
}
//Задание: Базовое использование горутин
```
25. Каналы
```go
package main
import "fmt"

func sendData(ch chan string) {
    ch <- "Hello"               // Отправка данных в канал
    ch <- "World"
    close(ch)                   // Закрытие канала
}

func main() {
    ch := make(chan string)     // Создание канала
    go sendData(ch)             // Запуск горутины
    
    for msg := range ch {       // Чтение из канала
        fmt.Println("Получено:", msg)
    }
}
//Задание: Базовое использование каналов
```
26. Буферизованные каналы
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 2)     // Буферизованный канал
    ch <- 1                     // Отправка без блокировки
    ch <- 2
    fmt.Println(<-ch)           // Получение данных
    fmt.Println(<-ch)
}
//Задание: Использование буферизованных каналов
```
27. Select с каналами
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(time.Second)
        ch1 <- "из канала 1"
    }()
    
    go func() {
        time.Sleep(time.Second * 2)
        ch2 <- "из канала 2"
    }()
    
    select {                    // Ожидание нескольких каналов
    case msg1 := <-ch1:
        fmt.Println("Получено:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Получено:", msg2)
    case <-time.After(time.Second * 3):
        fmt.Println("Таймаут")
    }
}
//Задание: Использование select для работы с каналами
```
28. Обработка ошибок
```go
package main
import (
    "errors"
    "fmt"
)

func process(value int) error {
    if value < 0 {
        return errors.New("отрицательное значение недопустимо")
    }
    fmt.Println("Обработка значения:", value)
    return nil
}

func main() {
    err := process(-5)
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Успешно обработано")
    }
}
//Задание: Базовая обработка ошибок
```
29. Panic и Recover
```go
package main
import "fmt"

func safeFunction() {
    defer func() {
        if r := recover(); r != nil {  // Восстановление после panic
            fmt.Println("Восстановлено:", r)
        }
    }()
    
    panic("критическая ошибка")  // Вызов panic
}

func main() {
    safeFunction()
    fmt.Println("Программа продолжает работу")
}
//Задание: Обработка паник с recover
```
30. Defer
```go
package main
import "fmt"

func main() {
    defer fmt.Println("Это выполнится последним")  // Отложенный вызов
    defer fmt.Println("Это выполнится предпоследним")
    
    fmt.Println("Обычный вывод")
    fmt.Println("Еще один обычный вывод")
}
//Задание: Использование defer для отложенных вызовов
```
31. Работа с файлами
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // Запись в файл
    content := "Hello, File!"
    err := os.WriteFile("test.txt", []byte(content), 0644)
    if err != nil {
        fmt.Println("Ошибка записи:", err)
        return
    }
    
    // Чтение из файла
    data, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Ошибка чтения:", err)
        return
    }
    fmt.Println("Прочитано из файла:", string(data))
}
//Задание: Базовая работа с файлами
```
32. JSON маршалинг
```go
package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func main() {
    user := User{Name: "Alice", Email: "alice@example.com", Age: 25}
    
    // Преобразование в JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("JSON:", string(jsonData))
    
    // Обратное преобразование
    var newUser User
    err = json.Unmarshal(jsonData, &newUser)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Объект:", newUser)
}
//Задание: Работа с JSON
```
33. HTTP сервер
```go
package main
import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")  // Ответ клиенту
}

func main() {
    http.HandleFunc("/hello", helloHandler)  // Регистрация обработчика
    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)        // Запуск сервера
}
//Задание: Простой HTTP сервер
```
34. HTTP клиент
```go
package main
import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.github.com")  // GET запрос
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    defer resp.Body.Close()                          // Закрытие тела ответа
    
    body, err := io.ReadAll(resp.Body)               // Чтение ответа
    if err != nil {
        fmt.Println("Ошибка чтения:", err)
        return
    }
    fmt.Println("Статус код:", resp.StatusCode)
    fmt.Println("Длина ответа:", len(body))
}
//Задание: Создание HTTP клиента
```
35. Тестирование
```go
package main
import "fmt"

// Функция для тестирования
func Add(a, b int) int {
    return a + b
}

// Файл: main_test.go (обычно в отдельном файле)
// func TestAdd(t *testing.T) {
//     result := Add(2, 3)
//     expected := 5
//     if result != expected {
//         t.Errorf("Ожидалось %d, получено %d", expected, result)
//     }
// }

func main() {
    fmt.Println("Результат сложения:", Add(2, 3))
}
//Задание: Структура тестирования в Go
```
36. Пакеты и импорты
```go
package main
import (
    "fmt"
    "math"           // Стандартный пакет
    "math/rand"      // Подпакет
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())  // Инициализация генератора
    randomNum := rand.Intn(100)       // Случайное число
    sqrt := math.Sqrt(64)             // Квадратный корень
    fmt.Printf("Случайное: %d, Квадратный корень: %.2f\n", randomNum, sqrt)
}
//Задание: Использование стандартных пакетов
```
37. Строки и руны
```go
package main
import "fmt"

func main() {
    str := "Привет, мир!"
    fmt.Println("Строка:", str)
    fmt.Println("Длина в байтах:", len(str))
    fmt.Println("Длина в символах:", len([]rune(str)))
    
    // Итерация по рунам
    for i, r := range str {
        fmt.Printf("Позиция %d: %c\n", i, r)
    }
}
//Задание: Работа с Unicode строками
```
38. Ввод пользователя
```go
package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Введите ваше имя: ")
    name, _ := reader.ReadString('\n')
    
    fmt.Print("Введите ваш возраст: ")
    ageInput, _ := reader.ReadString('\n')
    age, _ := strconv.Atoi(ageInput[:len(ageInput)-1])  // Убираем \n
    
    fmt.Printf("Привет, %s! Тебе %d лет.\n", name[:len(name)-1], age)
}
//Задание: Чтение ввода от пользователя
```
39. Время и дата
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()                          // Текущее время
    fmt.Println("Текущее время:", now)
    
    formatted := now.Format("02.01.2006 15:04:05")  // Форматирование
    fmt.Println("Форматированное:", formatted)
    
    future := now.Add(24 * time.Hour)          // Добавление времени
    fmt.Println("Завтра:", future.Format("02.01.2006"))
    
    duration := future.Sub(now)                // Разница во времени
    fmt.Println("Разница:", duration)
}
//Задание: Работа с временем
```
40. Сортировка
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{5, 2, 8, 1, 9}
    fmt.Println("До сортировки:", numbers)
    
    sort.Ints(numbers)                        // Сортировка чисел
    fmt.Println("После сортировки:", numbers)
    
    strings := []string{"яблоко", "банан", "апельсин"}
    sort.Strings(strings)                     // Сортировка строк
    fmt.Println("Отсортированные строки:", strings)
}
//Задание: Сортировка данных
```
41. Работа с директориями
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    // Создание директории
    err := os.Mkdir("test_dir", 0755)
    if err != nil {
        fmt.Println("Ошибка создания:", err)
    }
    
    // Чтение директории
    files, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Ошибка чтения:", err)
        return
    }
    
    fmt.Println("Файлы в текущей директории:")
    for _, file := range files {
        fmt.Println(file.Name())
    }
}
//Задание: Работа с файловой системой
```
42. Регулярные выражения
```go
package main
import (
    "fmt"
    "regexp"
)

func main() {
    text := "Мой email: example@mail.com и другой: test@site.org"
    
    pattern := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
    re := regexp.MustCompile(pattern)
    
    emails := re.FindAllString(text, -1)
    fmt.Println("Найденные email:")
    for _, email := range emails {
        fmt.Println(email)
    }
}
//Задание: Использование регулярных выражений
```
43. Шаблоны (Templates)
```go
package main
import (
    "os"
    "text/template"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    tmpl := `Привет, {{.Name}}!
Тебе {{.Age}} лет.
{{if .Age}}Ты совершеннолетний.{{else}}Ты несовершеннолетний.{{end}}`
    
    person := Person{Name: "Алиса", Age: 25}
    
    t := template.Must(template.New("test").Parse(tmpl))
    err := t.Execute(os.Stdout, person)
    if err != nil {
        panic(err)
    }
}
//Задание: Использование текстовых шаблонов
```
44. Генерация случайных данных
```go
package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    
    // Случайное число в диапазоне
    randomInt := rand.Intn(100)
    fmt.Println("Случайное число:", randomInt)
    
    // Случайное дробное число
    randomFloat := rand.Float64() * 100
    fmt.Printf("Случайное дробное: %.2f\n", randomFloat)
    
    // Случайный выбор из слайса
    colors := []string{"красный", "зеленый", "синий", "желтый"}
    randomColor := colors[rand.Intn(len(colors))]
    fmt.Println("Случайный цвет:", randomColor)
}
//Задание: Генерация случайных данных
```
45. Работа с байтами
```go
package main
import "fmt"

func main() {
    data := []byte("Hello, World!")
    
    fmt.Println("Байтовый слайс:", data)
    fmt.Println("Как строка:", string(data))
    
    // Модификация байтов
    for i := range data {
        if data[i] == 'o' {
            data[i] = '0'  // Замена 'o' на '0'
        }
    }
    fmt.Println("После замены:", string(data))
}
//Задание: Работа с байтовыми слайсами
```
46. Интерфейс io.Reader
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    reader := strings.NewReader("Пример строки для чтения")
    
    // Чтение по 5 байт за раз
    buffer := make([]byte, 5)
    for {
        n, err := reader.Read(buffer)
        if err != nil {
            break
        }
        fmt.Printf("Прочитано %d байт: %s\n", n, string(buffer[:n]))
    }
}
//Задание: Использование интерфейса Reader
```
47. Интерфейс io.Writer
```go
package main
import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer
    
    // Запись в буфер
    buffer.WriteString("Hello, ")
    buffer.WriteString("World!")
    buffer.WriteByte('!')
    
    fmt.Println("Содержимое буфера:", buffer.String())
    fmt.Println("Длина буфера:", buffer.Len())
}
//Задание: Использование интерфейса Writer
```
48. Кастомные типы
```go
package main
import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func main() {
    tempC := Celsius(25.0)
    tempF := tempC.ToFahrenheit()
    
    fmt.Printf("%.2f°C = %.2f°F\n", tempC, tempF)
    fmt.Printf("%.2f°F = %.2f°C\n", tempF, tempF.ToCelsius())
}
//Задание: Создание кастомных типов с методами
```
49. Композиция структур
```go
package main
import "fmt"

type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Age     int
    Address // Встраивание структуры
}

func main() {
    person := Person{
        Name: "Bob",
        Age:  30,
        Address: Address{
            City:    "Moscow",
            Country: "Russia",
        },
    }
    
    fmt.Println("Имя:", person.Name)
    fmt.Println("Город:", person.City)        // Прямой доступ к полям Address
    fmt.Println("Страна:", person.Country)
}
//Задание: Композиция структур через встраивание
```
50. Рефлексия
```go
package main
import (
    "fmt"
    "reflect"
)

type Sample struct {
    Name  string
    Value int
}

func main() {
    sample := Sample{Name: "Test", Value: 42}
    
    t := reflect.TypeOf(sample)
    v := reflect.ValueOf(sample)
    
    fmt.Println("Тип:", t.Name())
    fmt.Println("Количество полей:", t.NumField())
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        fmt.Printf("Поле %d: %s = %v\n", i, field.Name, value.Interface())
    }
}
//Задание: Основы рефлексии в Go
```
Продвинутые концепции Go
51. Контекст (Context)
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():  // Проверка отмены контекста
            fmt.Println(name, "завершен")
            return
        default:
            fmt.Println(name, "работает")
            time.Sleep(time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx, "Воркер 1")
    
    time.Sleep(3 * time.Second)
    cancel()  // Отмена контекста
    time.Sleep(time.Second)
}
//Задание: Использование context для управления горутинами
```
52. Context с таймаутом
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Контекст с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Операция завершена")
    case <-ctx.Done():
        fmt.Println("Таймаут:", ctx.Err())
    }
}
//Задание: Context с автоматическим таймаутом
```
53. Sync.WaitGroup
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Уменьшение счетчика при завершении
    fmt.Printf("Воркер %d начал\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Воркер %d завершил\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)  // Увеличение счетчика
        go worker(i, &wg)
    }
    
    wg.Wait()  // Ожидание завершения всех горутин
    fmt.Println("Все воркеры завершены")
}
//Задание: Ожидание завершения группы горутин
```
54. Sync.Mutex
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // Блокировка
    defer c.mu.Unlock() // Разблокировка
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := SafeCounter{}
    
    for i := 0; i < 1000; i++ {
        go counter.Increment()
    }
    
    time.Sleep(time.Second)
    fmt.Println("Итоговое значение:", counter.Value())
}
//Задание: Синхронизация доступа к общим данным
```
55. Sync.RWMutex
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type DataStore struct {
    data map[string]string
    rw   sync.RWMutex
}

func (ds *DataStore) Set(key, value string) {
    ds.rw.Lock()  // Эксклюзивная блокировка
    defer ds.rw.Unlock()
    ds.data[key] = value
}

func (ds *DataStore) Get(key string) string {
    ds.rw.RLock()  // Блокировка для чтения
    defer ds.rw.RUnlock()
    return ds.data[key]
}

func main() {
    store := &DataStore{data: make(map[string]string)}
    store.Set("name", "Alice")
    
    // Множественные одновременные чтения
    for i := 0; i < 5; i++ {
        ```go func(id int) {
            fmt.Printf("Чтец %d: %s\n", id, store.Get("name"))
        }(i)
    }
    
    time.Sleep(time.Millisecond * 100)
}
//Задание: Оптимизированная блокировка для частого чтения
```
56. Sync.Once
```go
package main
import (
    "fmt"
    "sync"
)

type Singleton struct {
    value string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {  // Гарантирует однократное выполнение
        instance = &Singleton{value: "Единственный экземпляр"}
        fmt.Println("Создан синглтон")
    })
    return instance
}

func main() {
    for i := 0; i < 5; i++ {
        go func() {
            instance := GetInstance()
            fmt.Println(instance.value)
        }()
    }
    
    fmt.Scanln() // Ожидание завершения горутин
}
//Задание: Гарантия однократного выполнения кода
```
57. Sync.Pool
```go
package main
import (
    "fmt"
    "sync"
)

type ExpensiveObject struct {
    ID int
}

func main() {
    pool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Создан новый объект")
            return &ExpensiveObject{ID: 0}
        },
    }
    
    // Получение объекта из пула
    obj1 := pool.Get().(*ExpensiveObject)
    obj1.ID = 1
    fmt.Printf("Объект 1: %d\n", obj1.ID)
    
    // Возврат в пул
    pool.Put(obj1)
    
    // Повторное использование
    obj2 := pool.Get().(*ExpensiveObject)
    fmt.Printf("Объект 2: %d\n", obj2.ID)
}
//Задание: Использование пула объектов для оптимизации
```
58. Каналы с контекстом
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func producer(ctx context.Context, ch chan<- int) {
    for i := 0; ; i++ {
        select {
        case ch <- i:
            fmt.Printf("Отправлено: %d\n", i)
        case <-ctx.Done():
            fmt.Println("Производитель остановлен")
            close(ch)
            return
        }
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    ch := make(chan int)
    go producer(ctx, ch)
    
    for value := range ch {
        fmt.Printf("Получено: %d\n", value)
    }
}
//Задание: Управление каналами через контекст
```
59. Fan-out, Fan-in
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

// Fan-out: несколько воркеров читают из одного канала
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Воркер %d обрабатывает задание %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

// Fan-in: объединение нескольких каналов в один
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    
    wg.Add(len(channels))
    for _, c := range channels {
        go output(c)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Запуск воркеров
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Отправка заданий
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Получение результатов
    for a := 1; a <= 5; a++ {
        fmt.Println("Результат:", <-results)
    }
}
//Задание: Паттерны Fan-out и Fan-in для параллельной обработки
```
60. Атомарные операции
```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var ops atomic.Uint64  // Атомарный счетчик
    var wg sync.WaitGroup
    
    // Запуск 50 горутин
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            for c := 0; c < 1000; c++ {
                ops.Add(1)  // Атомарное увеличение
            }
            wg.Done()
        }()
    }
    
    wg.Wait()
    fmt.Println("Операций:", ops.Load())
}
//Задание: Использование атомарных операций для счетчиков
```
61. Пользовательские ошибки
```go
package main
import (
    "errors"
    "fmt"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateUser(name string, age int) error {
    if name == "" {
        return &ValidationError{Field: "name", Message: "не может быть пустым"}
    }
    if age < 0 {
        return &ValidationError{Field: "age", Message: "не может быть отрицательным"}
    }
    return nil
}

func main() {
    err := validateUser("", -1)
    if err != nil {
        var valErr *ValidationError
        if errors.As(err, &valErr) {
            fmt.Printf("Ошибка валидации в поле '%s': %s\n", valErr.Field, valErr.Message)
        }
    }
}
//Задание: Создание пользовательских типов ошибок
```
62. errors.Is и errors.As
```go
package main
import (
    "errors"
    "fmt"
    "os"
)

var ErrFileNotFound = errors.New("файл не найден")

func openFile(filename string) error {
    if filename == "missing.txt" {
        return ErrFileNotFound
    }
    _, err := os.Open(filename)
    return err
}

func main() {
    err := openFile("missing.txt")
    
    // Проверка конкретной ошибки
    if errors.Is(err, ErrFileNotFound) {
        fmt.Println("Обработка ошибки 'файл не найден'")
    }
    
    // Проверка типа ошибки
    var pathError *os.PathError
    if errors.As(err, &pathError) {
        fmt.Printf("Ошибка пути: %s\n", pathError.Path)
    }
}
//Задание: Современные методы проверки ошибок
```
63. Работа с временными зонами
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    // Текущее время в UTC
    now := time.Now().UTC()
    fmt.Println("UTC:", now)
    
    // Создание времени в конкретной зоне
    loc, _ := time.LoadLocation("America/New_York")
    nyTime := now.In(loc)
    fmt.Println("Нью-Йорк:", nyTime)
    
    // Москва
    moscowLoc, _ := time.LoadLocation("Europe/Moscow")
    moscowTime := now.In(moscowLoc)
    fmt.Println("Москва:", moscowTime)
    
    // Парсинг времени с зоной
    timeStr := "2024-01-15 14:30:00 +0300"
    parsedTime, _ := time.Parse("2006-01-02 15:04:05 -0700", timeStr)
    fmt.Println("Распарсенное время:", parsedTime)
}
//Задание: Работа с временными зонами
```
64. Таймеры и Ticker'ы
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    // Таймер - одноразовый
    timer := time.NewTimer(2 * time.Second)
    <-timer.C
    fmt.Println("Таймер сработал!")
    
    // Ticker - периодический
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
    
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Тик в", t.Format("15:04:05"))
            }
        }
    }()
    
    time.Sleep(3 * time.Second)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker остановлен")
}
//Задание: Использование таймеров и тикеров
```
65. Шаблоны HTML
```go
package main
import (
    "html/template"
    "os"
)

type PageData struct {
    Title   string
    Users   []User
    ShowFooter bool
}

type User struct {
    Name string
    Age  int
}

func main() {
    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <ul>
    {{range .Users}}
        <li>{{.Name}} - {{.Age}} лет</li>
    {{end}}
    </ul>
    {{if .ShowFooter}}
    <footer>© 2024</footer>
    {{end}}
</body>
</html>`
    
    data := PageData{
        Title: "Список пользователей",
        Users: []User{
            {"Alice", 25},
            {"Bob", 30},
            {"Charlie", 35},
        },
        ShowFooter: true,
    }
    
    t := template.Must(template.New("webpage").Parse(tmpl))
    t.Execute(os.Stdout, data)
}
//Задание: Генерация HTML с шаблонами
```
66. HTTP middleware
```go
package main
import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Middleware для аутентификации
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", helloHandler)
    
    // Применение middleware
    handler := loggingMiddleware(authMiddleware(mux))
    
    log.Println("Сервер запущен на :8080")
    http.ListenAndServe(":8080", handler)
}
//Задание: Создание HTTP middleware
```
67. WebSocket сервер
```go
package main
import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Ошибка升级:", err)
        return
    }
    defer conn.Close()
    
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Ошибка чтения:", err)
            break
        }
        log.Printf("Получено: %s", message)
        
        // Эхо-ответ
        err = conn.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Ошибка записи:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", websocketHandler)
    log.Println("WebSocket сервер запущен на :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Создание WebSocket сервера
```
68. Работа с CSV
```go
package main
import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Запись в CSV
    people := []Person{
        {"Alice", 25, "New York"},
        {"Bob", 30, "London"},
        {"Charlie", 35, "Tokyo"},
    }
    
    file, err := os.Create("people.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // Заголовок
    writer.Write([]string{"Name", "Age", "City"})
    
    // Данные
    for _, person := range people {
        writer.Write([]string{
            person.Name,
            strconv.Itoa(person.Age),
            person.City,
        })
    }
    
    // Чтение из CSV
    file, err = os.Open("people.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Прочитанные данные:")
    for _, record := range records {
        fmt.Println(record)
    }
}
//Задание: Работа с CSV файлами
```
69. Base64 кодирование
```go
package main
import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := "Hello, World! Привет, мир!"
    
    // Кодирование
    encoded := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Закодировано:", encoded)
    
    // Декодирование
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Println("Декодировано:", string(decoded))
    
    // URL-безопасное кодирование
    urlEncoded := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println("URL-закодировано:", urlEncoded)
}
//Задание: Кодирование и декодирование Base64
```
70. Хеширование
```go
package main
import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
)

func main() {
    data := "secret data"
    
    // MD5
    md5Hash := md5.Sum([]byte(data))
    fmt.Printf("MD5: %x\n", md5Hash)
    
    // SHA1
    sha1Hash := sha1.Sum([]byte(data))
    fmt.Printf("SHA1: %x\n", sha1Hash)
    
    // SHA256
    sha256Hash := sha256.Sum256([]byte(data))
    fmt.Printf("SHA256: %x\n", sha256Hash)
    
    // SHA512
    sha512Hash := sha512.Sum512([]byte(data))
    fmt.Printf("SHA512: %x\n", sha512Hash)
}
//Задание: Вычисление хешей данных
```
71. Генерация UUID
```go
package main
import (
    "fmt"
    
    "github.com/google/uuid"
)

func main() {
    // Генерация UUID v4
    id := uuid.New()
    fmt.Println("UUID v4:", id.String())
    
    // Парсинг UUID из строки
    parsedUUID, err := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    if err != nil {
        panic(err)
    }
    fmt.Println("Парсинг UUID:", parsedUUID)
    
    // UUID v1 (на основе времени)
    idv1 := uuid.NewUUID()
    fmt.Println("UUID v1:", idv1.String())
}
//Задание: Генерация и работа с UUID
```
72. Флаги командной строки
```go
package main
import (
    "flag"
    "fmt"
)

func main() {
    // Определение флагов
    var (
        host     = flag.String("host", "localhost", "Хост сервера")
        port     = flag.Int("port", 8080, "Порт сервера")
        verbose  = flag.Bool("verbose", false, "Подробный вывод")
        count    = flag.Int("count", 1, "Количество повторений")
    )
    
    // Парсинг флагов
    flag.Parse()
    
    // Использование значений
    fmt.Printf("Хост: %s\n", *host)
    fmt.Printf("Порт: %d\n", *port)
    fmt.Printf("Подробный вывод: %t\n", *verbose)
    fmt.Printf("Количество: %d\n", *count)
    
    // Аргументы без флагов
    fmt.Println("\nАргументы:")
    for i, arg := range flag.Args() {
        fmt.Printf("  %d: %s\n", i, arg)
    }
}
//Задание: Обработка аргументов командной строки
```
73. Environment variables
```go
package main
import (
    "fmt"
    "os"
    "strconv"
)

type Config struct {
    Host     string
    Port     int
    Debug    bool
    LogLevel string
}

func loadConfig() *Config {
    // Чтение переменных окружения
    host := getEnv("APP_HOST", "localhost")
    port := getEnvInt("APP_PORT", 8080)
    debug := getEnvBool("APP_DEBUG", false)
    logLevel := getEnv("APP_LOG_LEVEL", "info")
    
    return &Config{
        Host:     host,
        Port:     port,
        Debug:    debug,
        LogLevel: logLevel,
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
    if value := os.Getenv(key); value != "" {
        if boolValue, err := strconv.ParseBool(value); err == nil {
            return boolValue
        }
    }
    return defaultValue
}

func main() {
    config := loadConfig()
    fmt.Printf("Конфигурация: %+v\n", config)
}
//Задание: Работа с переменными окружения
```
74. Logging с уровнями
```go
package main
import (
    "log"
    "os"
)

type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
)

type Logger struct {
    level LogLevel
    logger *log.Logger
}

func NewLogger(level LogLevel) *Logger {
    return &Logger{
        level: level,
        logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

func (l *Logger) Debug(format string, v ...interface{}) {
    if l.level <= DEBUG {
        l.logger.Printf("[DEBUG] "+format, v...)
    }
}

func (l *Logger) Info(format string, v ...interface{}) {
    if l.level <= INFO {
        l.logger.Printf("[INFO] "+format, v...)
    }
}

func (l *Logger) Warn(format string, v ...interface{}) {
    if l.level <= WARN {
        l.logger.Printf("[WARN] "+format, v...)
    }
}

func (l *Logger) Error(format string, v ...interface{}) {
    if l.level <= ERROR {
        l.logger.Printf("[ERROR] "+format, v...)
    }
}

func main() {
    logger := NewLogger(INFO)
    
    logger.Debug("Это отладочное сообщение")
    logger.Info("Это информационное сообщение")
    logger.Warn("Это предупреждение")
    logger.Error("Это ошибка")
}
//Задание: Создание логгера с уровнями логирования
```
75. Работа с архивами
```go
package main
import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func createZip(filename string, files []string) error {
    zipFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer zipFile.Close()
    
    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()
    
    for _, file := range files {
        err := addFileToZip(zipWriter, file)
        if err != nil {
            return err
        }
    }
    
    return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    info, err := file.Stat()
    if err != nil {
        return err
    }
    
    header, err := zip.FileInfoHeader(info)
    if err != nil {
        return err
    }
    
    header.Name = filepath.Base(filename)
    header.Method = zip.Deflate
    
    writer, err := zipWriter.CreateHeader(header)
    if err != nil {
        return err
    }
    
    _, err = io.Copy(writer, file)
    return err
}

func main() {
    files := []string{"main.go"} // Добавьте существующие файлы
    err := createZip("example.zip", files)
    if err != nil {
        fmt.Println("Ошибка создания архива:", err)
    } else {
        fmt.Println("Архив создан успешно")
    }
}
//Задание: Создание ZIP архивов
```
76. Парсинг XML
```go
package main
import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
    Email   string   `xml:"email,omitempty"`
}

type People struct {
    XMLName xml.Name `xml:"people"`
    People  []Person `xml:"person"`
}

func main() {
    // Маршалинг в XML
    people := People{
        People: []Person{
            {Name: "Alice", Age: 25, Email: "alice@example.com"},
            {Name: "Bob", Age: 30},
        },
    }
    
    xmlData, err := xml.MarshalIndent(people, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println("XML:\n", string(xmlData))
    
    // Демаршалинг из XML
    var parsedPeople People
    err = xml.Unmarshal(xmlData, &parsedPeople)
    if err != nil {
        panic(err)
    }
    fmt.Printf("\nПарсинг: %+v\n", parsedPeople)
}
//Задание: Работа с XML данными
```
77. Парсинг YAML
```go
package main
import (
    "fmt"
    "gopkg.in/yaml.v3"
    "log"
)

type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Name     string `yaml:"name"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
    } `yaml:"database"`
}

func main() {
    yamlData := `
server:
  host: localhost
  port: 8080
database:
  name: myapp
  user: admin
  password: secret
`
    
    var config Config
    err := yaml.Unmarshal([]byte(yamlData), &config)
    if err != nil {
        log.Fatalf("Ошибка парсинга YAML: %v", err)
    }
    
    fmt.Printf("Конфигурация: %+v\n", config)
    
    // Обратно в YAML
    outputYAML, err := yaml.Marshal(&config)
    if err != nil {
        log.Fatalf("Ошибка маршалинга YAML: %v", err)
    }
    fmt.Println("YAML вывод:")
    fmt.Println(string(outputYAML))
}
//Задание: Работа с YAML конфигурациями
```
78. SQL база данных
```go
package main
import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func main() {
    // Открытие базы данных
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Создание таблицы
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE,
        age INTEGER
    );`
    
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    
    // Вставка данных
    insertSQL := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
    result, err := db.Exec(insertSQL, "Alice", "alice@example.com", 25)
    if err != nil {
        log.Fatal(err)
    }
    
    id, _ := result.LastInsertId()
    fmt.Printf("Добавлен пользователь с ID: %d\n", id)
    
    // Чтение данных
    rows, err := db.Query("SELECT id, name, email, age FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var u User
        err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, u)
    }
    
    fmt.Println("Пользователи:")
    for _, user := range users {
        fmt.Printf("  %d: %s (%s) - %d лет\n", user.ID, user.Name, user.Email, user.Age)
    }
}
//Задание: Работа с SQL базой данных
```
79. Миграции базы данных
```go
package main
import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/mattn/go-sqlite3"
)

type Migration struct {
    Name string
    SQL  string
}

var migrations = []Migration{
    {
        Name: "create_users_table",
        SQL: `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
    },
    {
        Name: "add_age_to_users",
        SQL: "ALTER TABLE users ADD COLUMN age INTEGER;",
    },
    {
        Name: "create_posts_table",
        SQL: `
        CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            title TEXT NOT NULL,
            content TEXT,
            FOREIGN KEY(user_id) REFERENCES users(id)
        );`,
    },
}

func main() {
    db, err := sql.Open("sqlite3", "./app.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Таблица для отслеживания миграций
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT UNIQUE,
            applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `)
    if err != nil {
        log.Fatal(err)
    }
    
    // Применение миграций
    for _, migration := range migrations {
        var count int
        err = db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ?", migration.Name).Scan(&count)
        if err != nil {
            log.Fatal(err)
        }
        
        if count == 0 {
            fmt.Printf("Применение миграции: %s\n", migration.Name)
            _, err = db.Exec(migration.SQL)
            if err != nil {
                log.Fatalf("Ошибка миграции %s: %v", migration.Name, err)
            }
            
            _, err = db.Exec("INSERT INTO migrations (name) VALUES (?)", migration.Name)
            if err != nil {
                log.Fatal(err)
            }
        }
    }
    
    fmt.Println("Все миграции применены успешно")
}
//Задание: Система миграций базы данных
```
80. REST API сервер
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    
    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}
var nextID = 3

func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    user.ID = nextID
    nextID++
    users = append(users, user)
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    var updatedUser User
    json.NewDecoder(r.Body).Decode(&updatedUser)
    
    for i, user := range users {
        if user.ID == id {
            updatedUser.ID = id
            users[i] = updatedUser
            json.NewEncoder(w).Encode(updatedUser)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/users", getUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
    
    fmt.Println("REST API сервер запущен на :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: Создание REST API сервера
```
81. JWT аутентификация
```go
package main
import (
    "fmt"
    "net/http"
    "time"
    
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Простая проверка (в реальном приложении использовать базу данных)
    if credentials.Username != "admin" || credentials.Password != "password" {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: credentials.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Error creating token", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        http.Error(w, "Missing token", http.StatusUnauthorized)
        return
    }
    
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    
    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    w.Write([]byte(fmt.Sprintf("Hello %s!", claims.Username)))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login", login).Methods("POST")
    r.HandleFunc("/protected", protectedHandler).Methods("GET")
    
    fmt.Println("JWT сервер запущен на :8080")
    http.ListenAndServe(":8080", r)
}
//Задание: JWT аутентификация в REST API
```
82. Graceful shutdown
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(5 * time.Second) // Долгая операция
        fmt.Fprintf(w, "Hello World!")
    })
    
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    
    // Запуск сервера в горутине
    go func() {
        fmt.Println("Сервер запущен на :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Ошибка сервера: %v", err)
        }
    }()
    
    // Ожидание сигналов завершения
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    fmt.Println("\nЗавершение сервера...")
    
    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Принудительное завершение: %v", err)
    }
    
    fmt.Println("Сервер остановлен")
}
//Задание: Graceful shutdown сервера
```
83. Пул соединений
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type Connection struct {
    ID   int
    Open bool
}

type ConnectionPool struct {
    mu          sync.Mutex
    connections []*Connection
    maxSize     int
}

func NewConnectionPool(maxSize int) *ConnectionPool {
    return &ConnectionPool{
        connections: make([]*Connection, 0, maxSize),
        maxSize:     maxSize,
    }
}

func (p *ConnectionPool) Get() (*Connection, error) {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    // Поиск свободного соединения
    for _, conn := range p.connections {
        if conn.Open {
            conn.Open = false
            return conn, nil
        }
    }
    
    // Создание нового соединения
    if len(p.connections) < p.maxSize {
        conn := &Connection{
            ID:   len(p.connections) + 1,
            Open: false,
        }
        p.connections = append(p.connections, conn)
        fmt.Printf("Создано новое соединение %d\n", conn.ID)
        return conn, nil
    }
    
    return nil, fmt.Errorf("пул переполнен")
}

func (p *ConnectionPool) Release(conn *Connection) {
    p.mu.Lock()
    defer p.mu.Unlock()
    conn.Open = true
}

func (p *ConnectionPool) Stats() {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    openCount := 0
    for _, conn := range p.connections {
        if conn.Open {
            openCount++
        }
    }
    
    fmt.Printf("Всего: %d, Свободных: %d, Занятых: %d\n", 
        len(p.connections), openCount, len(p.connections)-openCount)
}

func main() {
    pool := NewConnectionPool(3)
    
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            conn, err := pool.Get()
            if err != nil {
                fmt.Printf("Горутина %d: %v\n", id, err)
                return
            }
            
            fmt.Printf("Горутина %d использует соединение %d\n", id, conn.ID)
            time.Sleep(time.Second)
            pool.Release(conn)
            fmt.Printf("Горутина %d освободила соединение %d\n", id, conn.ID)
        }(i)
    }
    
    wg.Wait()
    pool.Stats()
}
//Задание: Реализация пула соединений
```
84. Кэширование
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type CacheItem struct {
    Value      interface{}
    Expiration int64
}

type Cache struct {
    items map[string]CacheItem
    mu    sync.RWMutex
}

func NewCache() *Cache {
    cache := &Cache{
        items: make(map[string]CacheItem),
    }
    go cache.cleanup()
    return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    expiration := time.Now().Add(duration).UnixNano()
    c.items[key] = CacheItem{
        Value:      value,
        Expiration: expiration,
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, found := c.items[key]
    if !found {
        return nil, false
    }
    
    if time.Now().UnixNano() > item.Expiration {
        return nil, false
    }
    
    return item.Value, true
}

func (c *Cache) cleanup() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mu.Lock()
        now := time.Now().UnixNano()
        for key, item := range c.items {
            if now > item.Expiration {
                delete(c.items, key)
            }
        }
        c.mu.Unlock()
    }
}

func main() {
    cache := NewCache()
    
    // Сохранение данных с TTL
    cache.Set("user:1", "Alice", 5*time.Second)
    cache.Set("config:timeout", 30, 10*time.Second)
    
    // Получение данных
    if value, found := cache.Get("user:1"); found {
        fmt.Println("Найден пользователь:", value)
    }
    
    // Ожидание истечения TTL
    time.Sleep(6 * time.Second)
    if value, found := cache.Get("user:1"); !found {
        fmt.Println("Данные истекли")
    } else {
        fmt.Println("Пользователь:", value)
    }
}
//Задание: Реализация TTL кэша
```
85. Rate limiting
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    mu       sync.Mutex
    requests map[string][]time.Time
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (rl *RateLimiter) Allow(identifier string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    
    // Очистка старых запросов
    if _, exists := rl.requests[identifier]; !exists {
        rl.requests[identifier] = make([]time.Time, 0)
    }
    
    validRequests := make([]time.Time, 0)
    for _, t := range rl.requests[identifier] {
        if now.Sub(t) <= rl.window {
            validRequests = append(validRequests, t)
        }
    }
    
    rl.requests[identifier] = validRequests
    
    // Проверка лимита
    if len(rl.requests[identifier]) >= rl.limit {
        return false
    }
    
    rl.requests[identifier] = append(rl.requests[identifier], now)
    return true
}

func main() {
    limiter := NewRateLimiter(5, time.Minute) // 5 запросов в минуту
    
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            if limiter.Allow("user123") {
                fmt.Printf("Запрос %d: Разрешено\n", id)
            } else {
                fmt.Printf("Запрос %d: Отклонено (превышен лимит)\n", id)
            }
        }(i)
        time.Sleep(100 * time.Millisecond)
    }
    wg.Wait()
}
//Задание: Реализация rate limiting
```
86. Конфигурационные файлы
```go
package main
import (
    "encoding/json"
    "fmt"
    "io"
    "os"
)

type Config struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Logging  LoggingConfig  `json:"logging"`
}

type ServerConfig struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

type DatabaseConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Name     string `json:"name"`
    User     string `json:"user"`
    Password string `json:"password"`
}

type LoggingConfig struct {
    Level string `json:"level"`
    File  string `json:"file"`
}

func LoadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }
    
    var config Config
    err = json.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    
    return &config, nil
}

func SaveConfig(filename string, config *Config) error {
    data, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return err
    }
    
    return os.WriteFile(filename, data, 0644)
}

func main() {
    // Создание конфигурации по умолчанию
    defaultConfig := &Config{
        Server: ServerConfig{
            Host: "localhost",
            Port: 8080,
        },
        Database: DatabaseConfig{
            Host:     "localhost",
            Port:     5432,
            Name:     "myapp",
            User:     "admin",
            Password: "secret",
        },
        Logging: LoggingConfig{
            Level: "info",
            File:  "app.log",
        },
    }
    
    // Сохранение конфигурации
    err := SaveConfig("config.json", defaultConfig)
    if err != nil {
        fmt.Println("Ошибка сохранения:", err)
        return
    }
    
    // Загрузка конфигурации
    config, err := LoadConfig("config.json")
    if err != nil {
        fmt.Println("Ошибка загрузки:", err)
        return
    }
    
    fmt.Printf("Загруженная конфигурация: %+v\n", config)
}
//Задание: Работа с конфигурационными файлами
```
87. Health check
```go
package main
import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "time"
)

type HealthStatus struct {
    Status    string            `json:"status"`
    Timestamp time.Time         `json:"timestamp"`
    Checks    map[string]string `json:"checks"`
}

type HealthChecker struct {
    mu      sync.RWMutex
    checks  map[string]func() error
    status  HealthStatus
}

func NewHealthChecker() *HealthChecker {
    hc := &HealthChecker{
        checks: make(map[string]func() error),
        status: HealthStatus{
            Checks: make(map[string]string),
        },
    }
    
    go hc.monitor()
    return hc
}

func (hc *HealthChecker) AddCheck(name string, check func() error) {
    hc.mu.Lock()
    defer hc.mu.Unlock()
    hc.checks[name] = check
}

func (hc *HealthChecker) monitor() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        hc.performChecks()
    }
}

func (hc *HealthChecker) performChecks() {
    hc.mu.Lock()
    defer hc.mu.Unlock()
    
    overallStatus := "healthy"
    hc.status.Timestamp = time.Now()
    
    for name, check := range hc.checks {
        if err := check(); err != nil {
            hc.status.Checks[name] = err.Error()
            overallStatus = "unhealthy"
        } else {
            hc.status.Checks[name] = "ok"
        }
    }
    
    hc.status.Status = overallStatus
}

func (hc *HealthChecker) HealthHandler(w http.ResponseWriter, r *http.Request) {
    hc.mu.RLock()
    defer hc.mu.RUnlock()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(hc.status)
}

func main() {
    healthChecker := NewHealthChecker()
    
    // Добавление проверок
    healthChecker.AddCheck("database", func() error {
        // Здесь должна быть реальная проверка БД
        return nil
    })
    
    healthChecker.AddCheck("external_api", func() error {
        // Проверка внешнего API
        return nil
    })
    
    http.HandleFunc("/health", healthChecker.HealthHandler)
    fmt.Println("Health check сервер запущен на :8080")
    http.ListenAndServe(":8080", nil)
}
//Задание: Система health check для мониторинга
```
88. Метрики Prometheus
```go
package main
import (
    "fmt"
    "net/http"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    requestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "path"},
    )
    
    activeRequests = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "http_active_requests",
            Help: "Number of active HTTP requests",
        },
    )
)

func metricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        activeRequests.Inc()
        
        // Создание ResponseWriter для перехвата статуса
        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        
        next.ServeHTTP(rw, r)
        
        duration := time.Since(start).Seconds()
        requestsTotal.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", rw.statusCode)).Inc()
        requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
        activeRequests.Dec()
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(100 * time.Millisecond) // Имитация работы
    w.Write([]byte("Hello World!"))
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/metrics", promhttp.Handler())
    mux.Handle("/hello", metricsMiddleware(http.HandlerFunc(helloHandler)))
    
    fmt.Println("Сервер метрик запущен на :8080")
    http.ListenAndServe(":8080", mux)
}
//Задание: Интеграция с Prometheus для сбора метрик
```
89. Трассировка (Tracing)
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func initTracer() (*sdktrace.TracerProvider, error) {
    exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        return nil, err
    }
    
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            "service.name", "my-service",
        )),
    )
    
    otel.SetTracerProvider(tp)
    tracer = tp.Tracer("example-tracer")
    return tp, nil
}

func processOrder(ctx context.Context, orderID string) {
    ctx, span := tracer.Start(ctx, "processOrder")
    defer span.End()
    
    // Имитация работы
    time.Sleep(100 * time.Millisecond)
    validatePayment(ctx, orderID)
    updateInventory(ctx, orderID)
}

func validatePayment(ctx context.Context, orderID string) {
    ctx, span := tracer.Start(ctx, "validatePayment")
    defer span.End()
    
    time.Sleep(50 * time.Millisecond)
}

func updateInventory(ctx context.Context, orderID string) {
    ctx, span := tracer.Start(ctx, "updateInventory")
    defer span.End()
    
    time.Sleep(75 * time.Millisecond)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    // Создание корневого span
    ctx, span := tracer.Start(ctx, "orderHandler")
    defer span.End()
    
    orderID := "12345"
    processOrder(ctx, orderID)
    
    w.Write([]byte(fmt.Sprintf("Order %s processed", orderID)))
}

func main() {
    tp, err := initTracer()
    if err != nil {
        log.Fatal(err)
    }
    defer tp.Shutdown(context.Background())
    
    http.HandleFunc("/order", orderHandler)
    fmt.Println("Сервер с трассировкой запущен на :8080")
    http.ListenAndServe(":8080", nil)
}
//Задание: Инструментирование приложения с трассировкой
```
90. Feature flags
```go
package main
import (
    "encoding/json"
    "fmt"
    "sync"
    "time"
)

type FeatureFlag struct {
    Name        string    `json:"name"`
    Enabled     bool      `json:"enabled"`
    Description string    `json:"description"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type FeatureManager struct {
    mu     sync.RWMutex
    flags  map[string]*FeatureFlag
    config string
}

func NewFeatureManager() *FeatureManager {
    fm := &FeatureManager{
        flags: make(map[string]*FeatureFlag),
    }
    
    // Загрузка флагов по умолчанию
    fm.flags["new_ui"] = &FeatureFlag{
        Name:        "new_ui",
        Enabled:     false,
        Description: "Новый пользовательский интерфейс",
        UpdatedAt:   time.Now(),
    }
    
    fm.flags["beta_features"] = &FeatureFlag{
        Name:        "beta_features",
        Enabled:     true,
        Description: "Бета-функциональность",
        UpdatedAt:   time.Now(),
    }
    
    return fm
}

func (fm *FeatureManager) IsEnabled(flagName string) bool {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    if flag, exists := fm.flags[flagName]; exists {
        return flag.Enabled
    }
    return false
}

func (fm *FeatureManager) SetFlag(flagName string, enabled bool) {
    fm.mu.Lock()
    defer fm.mu.Unlock()
    
    if flag, exists := fm.flags[flagName]; exists {
        flag.Enabled = enabled
        flag.UpdatedAt = time.Now()
    } else {
        fm.flags[flagName] = &FeatureFlag{
            Name:      flagName,
            Enabled:   enabled,
            UpdatedAt: time.Now(),
        }
    }
}

func (fm *FeatureManager) GetFlags() map[string]*FeatureFlag {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    flags := make(map[string]*FeatureFlag)
    for k, v := range fm.flags {
        flags[k] = v
    }
    return flags
}

func (fm *FeatureManager) SaveToFile(filename string) error {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    data, err := json.MarshalIndent(fm.flags, "", "  ")
    if err != nil {
        return err
    }
    
    // В реальном приложении сохранить в файл
    fm.config = string(data)
    return nil
}

func main() {
    featureManager := NewFeatureManager()
    
    // Проверка флагов
    fmt.Println("New UI enabled:", featureManager.IsEnabled("new_ui"))
    fmt.Println("Beta features enabled:", featureManager.IsEnabled("beta_features"))
    
    // Включение флага
    featureManager.SetFlag("new_ui", true)
    fmt.Println("New UI after enable:", featureManager.IsEnabled("new_ui"))
    
    // Получение всех флагов
    flags := featureManager.GetFlags()
    for name, flag := range flags {
        fmt.Printf("%s: %t\n", name, flag.Enabled)
    }
}
//Задание: Система feature flags для управления функциональностью
```
91. Работа с большими файлами
```go
package main
import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
    "sync"
)

type WordCount struct {
    Word  string
    Count int
}

func countWords(filename string) (map[string]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    wordCounts := make(map[string]int)
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        word = strings.Trim(word, ".,!?;:\"()[]{}")
        if word != "" {
            wordCounts[word]++
        }
    }
    
    return wordCounts, scanner.Err()
}

func processLargeFile(filename string, batchSize int) ([]WordCount, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var mu sync.Mutex
    var wg sync.WaitGroup
    wordCounts := make(map[string]int)
    
    scanner := bufio.NewScanner(file)
    lines := make([]string, 0, batchSize)
    
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
        
        if len(lines) >= batchSize {
            wg.Add(1)
            batch := make([]string, len(lines))
            copy(batch, lines)
            lines = lines[:0]
            
            go func(batch []string) {
                defer wg.Done()
                batchCounts := make(map[string]int)
                
                for _, line := range batch {
                    words := strings.Fields(line)
                    for _, word := range words {
                        cleanWord := strings.ToLower(strings.Trim(word, ".,!?;:\"()[]{}"))
                        if cleanWord != "" {
                            batchCounts[cleanWord]++
                        }
                    }
                }
                
                mu.Lock()
                for word, count := range batchCounts {
                    wordCounts[word] += count
                }
                mu.Unlock()
            }(batch)
        }
    }
    
    // Обработка оставшихся строк
    if len(lines) > 0 {
        batchCounts := make(map[string]int)
        for _, line := range lines {
            words := strings.Fields(line)
            for _, word := range words {
                cleanWord := strings.ToLower(strings.Trim(word, ".,!?;:\"()[]{}"))
                if cleanWord != "" {
                    batchCounts[cleanWord]++
                }
            }
        }
        
        mu.Lock()
        for word, count := range batchCounts {
            wordCounts[word] += count
        }
        mu.Unlock()
    }
    
    wg.Wait()
    
    // Преобразование в слайс и сортировка
    var result []WordCount
    for word, count := range wordCounts {
        result = append(result, WordCount{Word: word, Count: count})
    }
    
    sort.Slice(result, func(i, j int) bool {
        return result[i].Count > result[j].Count
    })
    
    return result, nil
}

func main() {
    // Создание тестового файла
    content := "Hello world hello Go world Go programming language Go is awesome"
    os.WriteFile("test.txt", []byte(content), 0644)
    
    // Простой подсчет
    counts, err := countWords("test.txt")
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    
    fmt.Println("Подсчет слов:")
    for word, count := range counts {
        fmt.Printf("%s: %d\n", word, count)
    }
    
    // Параллельная обработка (для больших файлов)
    fmt.Println("\nТоп слов:")
    topWords, err := processLargeFile("test.txt", 2)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    
    for i, wc := range topWords {
        if i >= 5 { // Показать топ-5
            break
        }
        fmt.Printf("%s: %d\n", wc.Word, wc.Count)
    }
}
//Задание: Эффективная обработка больших файлов
```
92. Парсинг аргументов с cobra
```go
package main
import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "Мое приложение",
    Long:  "Длинное описание моего приложения",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Запуск основного приложения")
    },
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Версия приложения",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("MyApp v1.0.0")
    },
}

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Запуск сервера",
    Run: func(cmd *cobra.Command, args []string) {
        host, _ := cmd.Flags().GetString("host")
        port, _ := cmd.Flags().GetInt("port")
        fmt.Printf("Запуск сервера на %s:%d\n", host, port)
    },
}

func init() {
    serveCmd.Flags().StringP("host", "H", "localhost", "Хост сервера")
    serveCmd.Flags().IntP("port", "p", 8080, "Порт сервера")
    
    rootCmd.AddCommand(versionCmd)
    rootCmd.AddCommand(serveCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
//Задание: Создание CLI приложения с cobra
```
93. Генерация документации
```go
package main
import "fmt"

// Calculator предоставляет базовые математические операции
type Calculator struct{}

// NewCalculator создает новый экземпляр калькулятора
func NewCalculator() *Calculator {
    return &Calculator{}
}

// Add возвращает сумму двух чисел
// Пример: Add(2, 3) возвращает 5
func (c *Calculator) Add(a, b int) int {
    return a + b
}

// Subtract возвращает разность двух чисел
// Пример: Subtract(5, 3) возвращает 2
func (c *Calculator) Subtract(a, b int) int {
    return a - b
}

// Multiply возвращает произведение двух чисел
// Пример: Multiply(2, 3) возвращает 6
func (c *Calculator) Multiply(a, b int) int {
    return a * b
}

// Divide возвращает результат деления a на b
// Если b == 0, возвращает ошибку
// Пример: Divide(6, 3) возвращает 2
func (c *Calculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Power возвращает a в степени b
// Пример: Power(2, 3) возвращает 8
func (c *Calculator) Power(a, b int) int {
    result := 1
    for i := 0; i < b; i++ {
        result *= a
    }
    return result
}

// IsEven проверяет, является ли число четным
// Пример: IsEven(4) возвращает true
func (c *Calculator) IsEven(n int) bool {
    return n%2 == 0
}

// Fibonacci возвращает n-ное число Фибоначчи
// Пример: Fibonacci(6) возвращает 8
func (c *Calculator) Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

func main() {
    calc := NewCalculator()
    
    fmt.Println("Сложение:", calc.Add(5, 3))
    fmt.Println("Вычитание:", calc.Subtract(5, 3))
    fmt.Println("Умножение:", calc.Multiply(5, 3))
    
    result, err := calc.Divide(6, 3)
    if err != nil {
        fmt.Println("Ошибка деления:", err)
    } else {
        fmt.Println("Деление:", result)
    }
    
    fmt.Println("Степень:", calc.Power(2, 3))
    fmt.Println("Четное:", calc.IsEven(4))
    fmt.Println("Фибоначчи:", calc.Fibonacci(6))
}
//Задание: Документирование кода с примерами
```
94. Benchmark тесты
```go
package main
import (
    "crypto/sha256"
    "testing"
)

// Функция для бенчмарка
func calculateSHA256(data []byte) [32]byte {
    return sha256.Sum256(data)
}

// Функция для бенчмарка - конкатенация строк
func concatenateStrings(strings []string) string {
    var result string
    for _, s := range strings {
        result += s
    }
    return result
}

// Функция для бенчмарка - эффективная конкатенация
func concatenateStringsBuilder(strings []string) string {
    var builder strings.Builder
    for _, s := range strings {
        builder.WriteString(s)
    }
    return builder.String()
}

// Бенчмарк тесты
func BenchmarkSHA256(b *testing.B) {
    data := []byte("test data for benchmarking")
    for i := 0; i < b.N; i++ {
        calculateSHA256(data)
    }
}

func BenchmarkConcatenateStrings(b *testing.B) {
    testStrings := []string{"hello", "world", "golang", "benchmark", "testing"}
    for i := 0; i < b.N; i++ {
        concatenateStrings(testStrings)
    }
}

func BenchmarkConcatenateStringsBuilder(b *testing.B) {
    testStrings := []string{"hello", "world", "golang", "benchmark", "testing"}
    for i := 0; i < b.N; i++ {
        concatenateStringsBuilder(testStrings)
    }
}

func BenchmarkMapAccess(b *testing.B) {
    m := make(map[int]string)
    for i := 0; i < 1000; i++ {
        m[i] = fmt.Sprintf("value%d", i)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = m[i%1000]
    }
}

func BenchmarkSliceAccess(b *testing.B) {
    slice := make([]string, 1000)
    for i := 0; i < 1000; i++ {
        slice[i] = fmt.Sprintf("value%d", i)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = slice[i%1000]
    }
}

// Пример запуска: go test -bench=. -benchmem
//Задание: Написание benchmark тестов для измерения производительности
```
95. Table-driven тесты
```go
package main
import (
    "testing"
)

// Функции для тестирования
func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Table-driven тесты для Add
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed signs", -2, 3, 1},
        {"zero", 0, 5, 5},
        {"both zero", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Table-driven тесты для Subtract
func TestSubtract(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive result", 5, 3, 2},
        {"negative result", 3, 5, -2},
        {"zero result", 5, 5, 0},
        {"negative numbers", -2, -3, 1},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Subtract(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Subtract(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Table-driven тесты для Multiply
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 2, 3, 6},
        {"with zero", 5, 0, 0},
        {"negative numbers", -2, 3, -6},
        {"both negative", -2, -3, 6},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Table-driven тесты для Divide
func TestDivide(t *testing.T) {
    tests := []struct {
        name        string
        a           int
        b           int
        expected    int
        expectError bool
    }{
        {"normal division", 6, 3, 2, false},
        {"division by zero", 5, 0, 0, true},
        {"fraction result", 5, 2, 2, false},
        {"negative division", -6, 3, -2, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Divide(tt.a, tt.b)
            
            if tt.expectError {
                if err == nil {
                    t.Errorf("Divide(%d, %d) expected error, but got none", tt.a, tt.b)
                }
            } else {
                if err != nil {
                    t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
                }
                if result != tt.expected {
                    t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
                }
            }
        })
    }
}

// Пример запуска: go test -v
//Задание: Table-driven тестирование
```
96. Mock тестирование
```go
package main
import (
    "errors"
    "testing"
)

// Интерфейс для зависимости
type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
}

type User struct {
    ID   int
    Name string
}

// Реальная реализация
type RealUserRepository struct {
    users map[int]*User
}

func NewRealUserRepository() *RealUserRepository {
    return &RealUserRepository{
        users: map[int]*User{
            1: {ID: 1, Name: "Alice"},
            2: {ID: 2, Name: "Bob"},
        },
    }
}

func (r *RealUserRepository) FindByID(id int) (*User, error) {
    user, exists := r.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (r *RealUserRepository) Save(user *User) error {
    r.users[user.ID] = user
    return nil
}

// Mock реализация для тестов
type MockUserRepository struct {
    FindByIDFunc func(id int) (*User, error)
    SaveFunc     func(user *User) error
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    return m.FindByIDFunc(id)
}

func (m *MockUserRepository) Save(user *User) error {
    return m.SaveFunc(user)
}

// Сервис, который использует репозиторий
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}

func (s *UserService) CreateUser(id int, name string) error {
    user := &User{ID: id, Name: name}
    return s.repo.Save(user)
}

// Тесты с mock
func TestUserService_GetUserName(t *testing.T) {
    mockRepo := &MockUserRepository{
        FindByIDFunc: func(id int) (*User, error) {
            if id == 1 {
                return &User{ID: 1, Name: "Test User"}, nil
            }
            return nil, errors.New("not found")
        },
    }
    
    service := NewUserService(mockRepo)
    
    // Тест успешного случая
    name, err := service.GetUserName(1)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if name != "Test User" {
        t.Errorf("Expected 'Test User', got '%s'", name)
    }
    
    // Тест случая с ошибкой
    _, err = service.GetUserName(2)
    if err == nil {
        t.Error("Expected error, got nil")
    }
}

func TestUserService_CreateUser(t *testing.T) {
    var savedUser *User
    mockRepo := &MockUserRepository{
        SaveFunc: func(user *User) error {
            savedUser = user
            return nil
        },
    }
    
    service := NewUserService(mockRepo)
    
    err := service.CreateUser(3, "Charlie")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if savedUser == nil {
        t.Error("User was not saved")
    }
    if savedUser.ID != 3 || savedUser.Name != "Charlie" {
        t.Errorf("Saved user mismatch: %+v", savedUser)
    }
}
//Задание: Mock тестирование с зависимостями
```
97. Интеграционные тесты
```go
package main
import (
    "database/sql"
    "testing"
    
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID   int
    Name string
    Age  int
}

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateTable() error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER
    )`
    _, err := r.db.Exec(query)
    return err
}

func (r *UserRepository) CreateUser(user *User) error {
    query := "INSERT INTO users (name, age) VALUES (?, ?)"
    result, err := r.db.Exec(query, user.Name, user.Age)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    user.ID = int(id)
    return nil
}

func (r *UserRepository) GetUser(id int) (*User, error) {
    query := "SELECT id, name, age FROM users WHERE id = ?"
    row := r.db.QueryRow(query, id)
    
    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]*User, error) {
    query := "SELECT id, name, age FROM users"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []*User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}

// Интеграционный тест
func TestUserRepository_Integration(t *testing.T) {
    // Использование in-memory базы данных для тестов
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()
    
    repo := NewUserRepository(db)
    
    // Создание таблицы
    err = repo.CreateTable()
    if err != nil {
        t.Fatalf("Failed to create table: %v", err)
    }
    
    // Тест создания пользователя
    user := &User{Name: "Test User", Age: 25}
    err = repo.CreateUser(user)
    if err != nil {
        t.Fatalf("Failed to create user: %v", err)
    }
    
    if user.ID == 0 {
        t.Error("User ID should be set after creation")
    }
    
    // Тест получения пользователя
    retrievedUser, err := repo.GetUser(user.ID)
    if err != nil {
        t.Fatalf("Failed to get user: %v", err)
    }
    
    if retrievedUser.Name != user.Name || retrievedUser.Age != user.Age {
        t.Errorf("Retrieved user mismatch: got %+v, expected %+v", retrievedUser, user)
    }
    
    // Тест получения всех пользователей
    users, err := repo.GetAllUsers()
    if err != nil {
        t.Fatalf("Failed to get all users: %v", err)
    }
    
    if len(users) != 1 {
        t.Errorf("Expected 1 user, got %d", len(users))
    }
}

// Пример запуска: go test -tags=integration
//Задание: Интеграционное тестирование с базой данных
```
98. Property-based тестирование
```go
package main
import (
    "testing"
    "github.com/leanovate/gopter"
    "github.com/leanovate/gopter/gen"
    "github.com/leanovate/gopter/prop"
)

// Функции для тестирования
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

// Property-based тесты
func TestReverseStringProperties(t *testing.T) {
    parameters := gopter.DefaultTestParameters()
    parameters.MinSuccessfulTests = 1000
    properties := gopter.NewProperties(parameters)
    
    properties.Property("reverse(reverse(s)) == s", prop.ForAll(
        func(s string) bool {
            return ReverseString(ReverseString(s)) == s
        },
        gen.AnyString(),
    ))
    
    properties.Property("reverse preserves length", prop.ForAll(
        func(s string) bool {
            return len(ReverseString(s)) == len(s)
        },
        gen.AnyString(),
    ))
    
    properties.TestingRun(t)
}

func TestAddProperties(t *testing.T) {
    parameters := gopter.DefaultTestParameters()
    properties := gopter.NewProperties(parameters)
    
    properties.Property("commutative", prop.ForAll(
        func(a, b int) bool {
            return Add(a, b) == Add(b, a)
        },
        gen.Int(),
        gen.Int(),
    ))
    
    properties.Property("associative", prop.ForAll(
        func(a, b, c int) bool {
            return Add(Add(a, b), c) == Add(a, Add(b, c))
        },
        gen.Int(),
        gen.Int(),
        gen.Int(),
    ))
    
    properties.Property("identity", prop.ForAll(
        func(a int) bool {
            return Add(a, 0) == a
        },
        gen.Int(),
    ))
    
    properties.TestingRun(t)
}

func TestMultiplyProperties(t *testing.T) {
    parameters := gopter.DefaultTestParameters()
    properties := gopter.NewProperties(parameters)
    
    properties.Property("commutative", prop.ForAll(
        func(a, b int) bool {
            return Multiply(a, b) == Multiply(b, a)
        },
        gen.Int(),
        gen.Int(),
    ))
    
    properties.Property("associative", prop.ForAll(
        func(a, b, c int) bool {
            return Multiply(Multiply(a, b), c) == Multiply(a, Multiply(b, c))
        },
        gen.Int(),
        gen.Int(),
        gen.Int(),
    ))
    
    properties.Property("identity", prop.ForAll(
        func(a int) bool {
            return Multiply(a, 1) == a
        },
        gen.Int(),
    ))
    
    properties.Property("zero", prop.ForAll(
        func(a int) bool {
            return Multiply(a, 0) == 0
        },
        gen.Int(),
    ))
    
    properties.Property("distributive", prop.ForAll(
        func(a, b, c int) bool {
            return Multiply(a, Add(b, c)) == Add(Multiply(a, b), Multiply(a, c))
        },
        gen.Int(),
        gen.Int(),
        gen.Int(),
    ))
    
    properties.TestingRun(t)
}
//Задание: Property-based тестирование
```
99. Fuzz тестирование
```go
package main
import (
    "testing"
)

// Функции для fuzz тестирования
func ParseInteger(s string) (int, error) {
    // Упрощенная реализация парсинга
    if s == "" {
        return 0, fmt.Errorf("empty string")
    }
    
    result := 0
    sign := 1
    start := 0
    
    if s[0] == '-' {
        sign = -1
        start = 1
    } else if s[0] == '+' {
        start = 1
    }
    
    for i := start; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            return 0, fmt.Errorf("invalid character")
        }
        result = result*10 + int(s[i]-'0')
    }
    
    return result * sign, nil
}

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ValidateEmail(email string) bool {
    if len(email) < 3 || len(email) > 254 {
        return false
    }
    
    at := -1
    dot := -1
    
    for i, c := range email {
        if c == '@' {
            if at != -1 {
                return false // Множественные @
            }
            at = i
        } else if c == '.' {
            dot = i
        }
    }
    
    return at > 0 && dot > at+1 && dot < len(email)-1
}

// Fuzz тесты
func FuzzParseInteger(f *testing.F) {
    // Добавление seed корпусов
    f.Add("123")
    f.Add("-456")
    f.Add("0")
    f.Add("+789")
    
    f.Fuzz(func(t *testing.T, input string) {
        result, err := ParseInteger(input)
        
        // Проверка, что функция не паникует
        if err == nil {
            // Если парсинг успешен, проверяем базовые свойства
            if result == 0 && input != "0" && input != "+0" && input != "-0" {
                t.Errorf("Zero result for non-zero input: %s", input)
            }
        }
    })
}

func FuzzReverseString(f *testing.F) {
    f.Add("hello")
    f.Add("world")
    f.Add("")
    f.Add("a")
    f.Add("абвгд")
    
    f.Fuzz(func(t *testing.T, input string) {
        reversed := ReverseString(input)
        doubleReversed := ReverseString(reversed)
        
        if doubleReversed != input {
            t.Errorf("Reverse(Reverse(%q)) = %q, want %q", input, doubleReversed, input)
        }
        
        // Дополнительные проверки
        if len(reversed) != len(input) {
            t.Errorf("Length changed: input %d, reversed %d", len(input), len(reversed))
        }
    })
}

func FuzzValidateEmail(f *testing.F) {
    f.Add("test@example.com")
    f.Add("user.name@domain.co.uk")
    f.Add("invalid")
    f.Add("")
    f.Add("a@b.c")
    
    f.Fuzz(func(t *testing.T, email string) {
        // Основная проверка - функция не должна паниковать
        result := ValidateEmail(email)
        
        // Дополнительные проверки логики
        if result {
            // Если email валиден, проверяем базовые требования
            if len(email) < 3 || len(email) > 254 {
                t.Errorf("Valid email with invalid length: %s", email)
            }
        }
    })
}

// Пример запуска: go test -fuzz=FuzzParseInteger
//Задание: Fuzz тестирование для поиска краевых случаев
```
100. E2E тестирование
```go
package main
import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var users = []User{
    {ID: 1, Name: "Alice", Age: 25},
    {ID: 2, Name: "Bob", Age: 30},
}
var nextID = 3

func setupAPI() *http.ServeMux {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        switch r.Method {
        case "GET":
            json.NewEncoder(w).Encode(users)
        case "POST":
            var user User
            if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            user.ID = nextID
            nextID++
            users = append(users, user)
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(user)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        id := 1 // Упрощенное извлечение ID
        switch r.Method {
        case "GET":
            for _, user := range users {
                if user.ID == id {
                    json.NewEncoder(w).Encode(user)
                    return
                }
            }
            http.Error(w, "User not found", http.StatusNotFound)
        case "PUT":
            var updatedUser User
            if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            for i, user := range users {
                if user.ID == id {
                    updatedUser.ID = id
                    users[i] = updatedUser
                    json.NewEncoder(w).Encode(updatedUser)
                    return
                }
            }
            http.Error(w, "User not found", http.StatusNotFound)
        case "DELETE":
            for i, user := range users {
                if user.ID == id {
                    users = append(users[:i], users[i+1:]...)
                    w.WriteHeader(http.StatusNoContent)
                    return
                }
            }
            http.Error(w, "User not found", http.StatusNotFound)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    return mux
}

// E2E тесты
func TestE2EUserAPI(t *testing.T) {
    server := httptest.NewServer(setupAPI())
    defer server.Close()
    
    client := server.Client()
    
    t.Run("GET /users returns all users", func(t *testing.T) {
        resp, err := client.Get(server.URL + "/users")
        if err != nil {
            t.Fatalf("Failed to make request: %v", err)
        }
        defer resp.Body.Close()
        
        if resp.StatusCode != http.StatusOK {
            t.Errorf("Expected status 200, got %d", resp.StatusCode)
        }
        
        var users []User
        if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
            t.Fatalf("Failed to decode response: %v", err)
        }
        
        if len(users) != 2 {
            t.Errorf("Expected 2 users, got %d", len(users))
        }
    })
    
    t.Run("POST /users creates new user", func(t *testing.T) {
        newUser := User{Name: "Charlie", Age: 35}
        jsonData, _ := json.Marshal(newUser)
        
        resp, err := client.Post(server.URL+"/users", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            t.Fatalf("Failed to make request: %v", err)
        }
        defer resp.Body.Close()
        
        if resp.StatusCode != http.StatusCreated {
            t.Errorf("Expected status 201, got %d", resp.StatusCode)
        }
        
        var createdUser User
        if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
            t.Fatalf("Failed to decode response: %v", err)
        }
        
        if createdUser.Name != "Charlie" || createdUser.Age != 35 {
            t.Errorf("User data mismatch: %+v", createdUser)
        }
    })
    
    t.Run("GET /users/{id} returns specific user", func(t *testing.T) {
        resp, err := client.Get(server.URL + "/users/1")
        if err != nil {
            t.Fatalf("Failed to make request: %v", err)
        }
        defer resp.Body.Close()
        
        if resp.StatusCode != http.StatusOK {
            t.Errorf("Expected status 200, got %d", resp.StatusCode)
        }
        
        var user User
        if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
            t.Fatalf("Failed to decode response: %v", err)
        }
        
        if user.ID != 1 {
            t.Errorf("Expected user ID 1, got %d", user.ID)
        }
    })
}

// Пример запуска: go test -v
//Задание: End-to-end тестирование API
```
Продвинутые техники и реальные сценарии
101. Graceful shutdown HTTP сервера
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(2 * time.Second) // Имитация долгой обработки
        fmt.Fprintf(w, "Hello World!")
    })
    
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    
    // Канал для получения ошибок сервера
    serverErrors := make(chan error, 1)
    
    // Запуск сервера в горутине
    go func() {
        log.Println("Server listening on :8080")
        serverErrors <- server.ListenAndServe()
    }()
    
    // Канал для сигналов ОС
    osSignals := make(chan os.Signal, 1)
    signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
    
    // Ожидание сигнала завершения или ошибки сервера
    select {
    case err := <-serverErrors:
        log.Fatalf("Error starting server: %v", err)
        
    case <-osSignals:
        log.Println("Received shutdown signal")
        
        // Даем 30 секунд на завершение активных запросов
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()
        
        if err := server.Shutdown(ctx); err != nil {
            log.Printf("Graceful shutdown failed: %v", err)
            server.Close()
        }
    }
    
    log.Println("Server stopped")
}
//Задание: Полная реализация graceful shutdown для HTTP сервера
```
102. Конфигурация с Viper
```go
package main
import (
    "fmt"
    "log"
    
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
    Host string `mapstructure:"host"`
    Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

func LoadConfig(configPath string) (*Config, error) {
    viper.SetConfigFile(configPath)
    viper.SetConfigType("yaml")
    
    // Значения по умолчанию
    viper.SetDefault("server.host", "localhost")
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("redis.port", 6379)
    viper.SetDefault("redis.db", 0)
    
    // Чтение конфигурационного файла
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading config file: %w", err)
    }
    
    // Автоматическое чтение переменных окружения
    viper.AutomaticEnv()
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("error unmarshaling config: %w", err)
    }
    
    return &config, nil
}

func main() {
    config, err := LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }
    
    fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("Database: %s@%s:%d/%s\n", 
        config.Database.Username, 
        config.Database.Host, 
        config.Database.Port, 
        config.Database.Name)
    fmt.Printf("Redis: %s:%d/%d\n", 
        config.Redis.Host, 
        config.Redis.Port, 
        config.Redis.DB)
}
//Задание: Управление конфигурацией с помощью Viper
```
103. Миграции с Goose
```go
package main
import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
    "github.com/pressly/goose"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "password"
    dbname   = "mydb"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    
    db, err := sql.Open("postgresql", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Установка директории с миграциями
    goose.SetDialect("postgres")
    
    // Команды для управления миграциями:
    // goose create add_users_table sql
    // goose up
    // goose down
    // goose status
    
    // Пример создания миграции:
    // Файл: migrations/001_create_users_table.sql
    /*
    -- +goose Up
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    CREATE INDEX idx_users_username ON users(username);
    CREATE INDEX idx_users_email ON users(email);
    
    -- +goose Down
    DROP TABLE users;
    */
    
    // Запуск миграций
    if err := goose.Up(db, "migrations"); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Migrations applied successfully")
}
//Задание: Управление миграциями базы данных с Goose
```
104. gRPC сервер
```go
package main
import (
    "context"
    "log"
    "net"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "github.com/yourname/yourpackage/pb"
)

type userService struct {
    pb.UnimplementedUserServiceServer
    users map[int32]*pb.User
}

func NewUserService() *userService {
    return &userService{
        users: map[int32]*pb.User{
            1: {Id: 1, Name: "Alice", Email: "alice@example.com"},
            2: {Id: 2, Name: "Bob", Email: "bob@example.com"},
        },
    }
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    user, exists := s.users[req.Id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
    id := int32(len(s.users) + 1)
    user := &pb.User{
        Id:    id,
        Name:  req.Name,
        Email: req.Email,
    }
    s.users[id] = user
    return user, nil
}

func (s *userService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
    users := make([]*pb.User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return &pb.ListUsersResponse{Users: users}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, NewUserService())
    
    // Включение reflection для тестирования
    reflection.Register(s)
    
    log.Println("gRPC server listening on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
//Задание: Создание gRPC сервера
```
105. gRPC клиент
```go
package main
import (
    "context"
    "log"
    "time"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "github.com/yourname/yourpackage/pb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    
    client := pb.NewUserServiceClient(conn)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    // Создание пользователя
    createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
        Name:  "Charlie",
        Email: "charlie@example.com",
    })
    if err != nil {
        log.Fatalf("could not create user: %v", err)
    }
    log.Printf("Created user: %v", createResp)
    
    // Получение пользователя
    getResp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
    if err != nil {
        log.Fatalf("could not get user: %v", err)
    }
    log.Printf("User: %v", getResp)
    
    // Список пользователей
    listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
    if err != nil {
        log.Fatalf("could not list users: %v", err)
    }
    log.Printf("Users: %v", listResp.Users)
}
//Задание: Создание gRPC клиента
```
106. GraphQL сервер с gqlgen
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/yourname/yourpackage/graph"
    "github.com/yourname/yourpackage/graph/generated"
)

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []*User{
    {ID: "1", Name: "Alice", Email: "alice@example.com"},
    {ID: "2", Name: "Bob", Email: "bob@example.com"},
}

type resolver struct{}

func (r *resolver) Users(ctx context.Context) ([]*User, error) {
    return users, nil
}

func (r *resolver) User(ctx context.Context, id string) (*User, error) {
    for _, user := range users {
        if user.ID == id {
            return user, nil
        }
    }
    return nil, fmt.Errorf("user not found")
}

func (r *resolver) CreateUser(ctx context.Context, input graph.CreateUserInput) (*User, error) {
    user := &User{
        ID:    fmt.Sprintf("%d", len(users)+1),
        Name:  input.Name,
        Email: input.Email,
    }
    users = append(users, user)
    return user, nil
}

func main() {
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(
        generated.Config{Resolvers: &resolver{}}))
    
    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)
    
    log.Printf("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Создание GraphQL сервера с gqlgen
```
107. WebSocket чат
```go
package main
import (
    "log"
    "net/http"
    "sync"
    
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
    conn *websocket.Conn
    send chan []byte
}

type ChatRoom struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
    mu         sync.RWMutex
}

func NewChatRoom() *ChatRoom {
    return &ChatRoom{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (cr *ChatRoom) run() {
    for {
        select {
        case client := <-cr.register:
            cr.mu.Lock()
            cr.clients[client] = true
            cr.mu.Unlock()
            
        case client := <-cr.unregister:
            cr.mu.Lock()
            if _, ok := cr.clients[client]; ok {
                delete(cr.clients, client)
                close(client.send)
            }
            cr.mu.Unlock()
            
        case message := <-cr.broadcast:
            cr.mu.RLock()
            for client := range cr.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(cr.clients, client)
                }
            }
            cr.mu.RUnlock()
        }
    }
}

func (c *Client) readPump(chatRoom *ChatRoom) {
    defer func() {
        chatRoom.unregister <- c
        c.conn.Close()
    }()
    
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            break
        }
        chatRoom.broadcast <- message
    }
}

func (c *Client) writePump() {
    defer c.conn.Close()
    
    for message := range c.send {
        err := c.conn.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            break
        }
    }
}

func serveChat(chatRoom *ChatRoom) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println("Upgrade error:", err)
            return
        }
        
        client := &Client{
            conn: conn,
            send: make(chan []byte, 256),
        }
        
        chatRoom.register <- client
        
        go client.writePump()
        go client.readPump(chatRoom)
    }
}

func main() {
    chatRoom := NewChatRoom()
    go chatRoom.run()
    
    http.HandleFunc("/ws", serveChat(chatRoom))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "chat.html")
    })
    
    log.Println("Chat server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Реализация многопользовательского WebSocket чата
```
108. OAuth2 аутентификация
```go
package main
import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var oauthConfig = &oauth2.Config{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    RedirectURL:  "http://localhost:8080/auth/callback",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
    Endpoint:     google.Endpoint,
}

type UserInfo struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/auth/callback", authCallbackHandler)
    http.HandleFunc("/profile", profileHandler)
    
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    html := `
    <html>
        <body>
            <h1>OAuth2 Example</h1>
            <a href="/login">Login with Google</a>
        </body>
    </html>`
    fmt.Fprintf(w, html)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    
    token, err := oauthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Сохранение токена в сессии (в реальном приложении используйте secure cookies)
    http.SetCookie(w, &http.Cookie{
        Name:  "oauth_token",
        Value: token.AccessToken,
        Path:  "/",
    })
    
    http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("oauth_token")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
    
    client := oauthConfig.Client(context.Background(), &oauth2.Token{AccessToken: cookie.Value})
    
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        http.Error(w, "Failed to get user info", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    
    var userInfo UserInfo
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, `
    <html>
        <body>
            <h1>Profile</h1>
            <p>ID: %s</p>
            <p>Name: %s</p>
            <p>Email: %s</p>
            <a href="/">Home</a>
        </body>
    </html>`, userInfo.ID, userInfo.Name, userInfo.Email)
}
//Задание: Реализация OAuth2 аутентификации с Google
```
109. JWT с refresh токенами
```go
package main
import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
)

var (
    jwtKey = []byte("your-secret-key")
    refreshTokens = make(map[string]string)
)

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type AuthTokens struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

func generateRefreshToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}

func generateTokens(username string) (*AuthTokens, error) {
    // Access Token (15 минут)
    accessExpirationTime := time.Now().Add(15 * time.Minute)
    accessClaims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: accessExpirationTime.Unix(),
        },
    }
    
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessTokenString, err := accessToken.SignedString(jwtKey)
    if err != nil {
        return nil, err
    }
    
    // Refresh Token (7 дней)
    refreshToken, err := generateRefreshToken()
    if err != nil {
        return nil, err
    }
    
    refreshTokens[refreshToken] = username
    
    return &AuthTokens{
        AccessToken:  accessTokenString,
        RefreshToken: refreshToken,
    }, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Простая проверка (в реальном приложении проверяйте в БД)
    if credentials.Username != "admin" || credentials.Password != "password" {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    
    tokens, err := generateTokens(credentials.Username)
    if err != nil {
        http.Error(w, "Error generating tokens", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tokens)
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
    var request struct {
        RefreshToken string `json:"refresh_token"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    username, exists := refreshTokens[request.RefreshToken]
    if !exists {
        http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        return
    }
    
    // Удаляем использованный refresh token
    delete(refreshTokens, request.RefreshToken)
    
    tokens, err := generateTokens(username)
    if err != nil {
        http.Error(w, "Error generating tokens", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tokens)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        http.Error(w, "Missing token", http.StatusUnauthorized)
        return
    }
    
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    
    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
    
    w.Write([]byte(fmt.Sprintf("Hello %s!", claims.Username)))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login", loginHandler).Methods("POST")
    r.HandleFunc("/refresh", refreshHandler).Methods("POST")
    r.HandleFunc("/protected", protectedHandler).Methods("GET")
    
    log.Println("JWT server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: JWT аутентификация с refresh токенами
```
110. Rate limiting с Redis
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "time"
    
    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func initRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    
    if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}

func rateLimitMiddleware(next http.HandlerFunc, limit int, window time.Duration) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        key := fmt.Sprintf("rate_limit:%s", ip)
        
        // Используем Redis для подсчета запросов
        pipe := rdb.Pipeline()
        incr := pipe.Incr(ctx, key)
        pipe.Expire(ctx, key, window)
        _, err := pipe.Exec(ctx)
        
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        
        currentCount, _ := incr.Result()
        
        // Устанавливаем заголовки с информацией о лимите
        w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
        w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(int64(limit)-currentCount, 10))
        w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(window).Unix(), 10))
        
        if currentCount > int64(limit) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    }
}

func slidingWindowRateLimit(next http.HandlerFunc, limit int, window time.Duration) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        now := time.Now().UnixNano()
        windowSize := window.Nanoseconds()
        
        key := fmt.Sprintf("sliding_rate_limit:%s", ip)
        
        // Удаляем старые записи
        minScore := strconv.FormatInt(now-windowSize, 10)
        rdb.ZRemRangeByScore(ctx, key, "0", minScore)
        
        // Добавляем текущий запрос
        member := strconv.FormatInt(now, 10)
        pipe := rdb.Pipeline()
        pipe.ZAdd(ctx, key, &redis.Z{
            Score:  float64(now),
            Member: member,
        })
        pipe.Expire(ctx, key, window)
        count := pipe.ZCard(ctx, key)
        _, err := pipe.Exec(ctx)
        
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        
        currentCount, _ := count.Result()
        
        w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
        w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(int64(limit)-currentCount, 10))
        
        if currentCount > int64(limit) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    }
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API response"))
}

func main() {
    initRedis()
    
    http.HandleFunc("/api", rateLimitMiddleware(apiHandler, 10, time.Minute))
    http.HandleFunc("/api/sliding", slidingWindowRateLimit(apiHandler, 10, time.Minute))
    
    log.Println("Rate limiting server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Rate limiting с использованием Redis
```
111. Кэширование с Redis
```go
package main
import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func initRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
}

func getUserFromDB(id int) (*User, error) {
    // Имитация запроса к базе данных
    time.Sleep(100 * time.Millisecond)
    
    // В реальном приложении здесь был бы запрос к БД
    return &User{
        ID:    id,
        Name:  fmt.Sprintf("User %d", id),
        Email: fmt.Sprintf("user%d@example.com", id),
    }, nil
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Попытка получить данные из кэша
    cachedData, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        // Данные найдены в кэше
        var user User
        json.Unmarshal([]byte(cachedData), &user)
        
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("X-Cache", "HIT")
        json.NewEncoder(w).Encode(user)
        return
    }
    
    // Данных нет в кэше, получаем из БД
    user, err := getUserFromDB(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    // Сохраняем в кэш
    userJSON, _ := json.Marshal(user)
    rdb.SetEX(ctx, cacheKey, userJSON, 5*time.Minute)
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Cache", "MISS")
    json.NewEncoder(w).Encode(user)
}

// Паттерн "Cache Aside"
func getUserWithCacheAside(id int) (*User, error) {
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Чтение из кэша
    cachedData, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var user User
        json.Unmarshal([]byte(cachedData), &user)
        return &user, nil
    }
    
    // Чтение из БД
    user, err := getUserFromDB(id)
    if err != nil {
        return nil, err
    }
    
    // Запись в кэш
    userJSON, _ := json.Marshal(user)
    rdb.SetEX(ctx, cacheKey, userJSON, 5*time.Minute)
    
    return user, nil
}

// Паттерн "Write Through"
func updateUserWithWriteThrough(user *User) error {
    // Обновление в БД
    // db.UpdateUser(user)
    
    // Обновление в кэше
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    userJSON, _ := json.Marshal(user)
    return rdb.SetEX(ctx, cacheKey, userJSON, 5*time.Minute).Err()
}

func main() {
    initRedis()
    
    http.HandleFunc("/user", getUserHandler)
    
    log.Println("Caching server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Кэширование данных с Redis
```
112. Message Queue с RabbitMQ
```go
package main
import (
    "context"
    "encoding/json"
    "log"
    "time"
    
    "github.com/streadway/amqp"
)

type EmailTask struct {
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func connectRabbitMQ() (*amqp.Connection, error) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return nil, err
    }
    return conn, nil
}

func setupQueue(ch *amqp.Channel) error {
    _, err := ch.QueueDeclare(
        "email_tasks", // name
        true,          // durable
        false,         // delete when unused
        false,         // exclusive
        false,         // no-wait
        nil,           // arguments
    )
    return err
}

func publishEmailTask(ch *amqp.Channel, task EmailTask) error {
    body, err := json.Marshal(task)
    if err != nil {
        return err
    }
    
    return ch.Publish(
        "",             // exchange
        "email_tasks",  // routing key
        false,          // mandatory
        false,          // immediate
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "application/json",
            Body:         body,
        })
}

func startEmailWorker(ch *amqp.Channel) error {
    msgs, err := ch.Consume(
        "email_tasks", // queue
        "",            // consumer
        false,         // auto-ack
        false,         // exclusive
        false,         // no-local
        false,         // no-wait
        nil,           // args
    )
    if err != nil {
        return err
    }
    
    go func() {
        for msg := range msgs {
            var task EmailTask
            if err := json.Unmarshal(msg.Body, &task); err != nil {
                log.Printf("Error decoding message: %v", err)
                msg.Nack(false, false) // Не пере-поставлять сообщение
                continue
            }
            
            // Имитация отправки email
            log.Printf("Sending email to: %s, Subject: %s", task.To, task.Subject)
            time.Sleep(2 * time.Second) // Имитация работы
            log.Printf("Email sent to: %s", task.To)
            
            msg.Ack(false) // Подтверждение обработки
        }
    }()
    
    return nil
}

func main() {
    conn, err := connectRabbitMQ()
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()
    
    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()
    
    if err := setupQueue(ch); err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }
    
    // Запуск воркеров
    for i := 0; i < 3; i++ {
        if err := startEmailWorker(ch); err != nil {
            log.Fatalf("Failed to start worker %d: %v", i, err)
        }
    }
    
    // Публикация тестовых задач
    tasks := []EmailTask{
        {To: "user1@example.com", Subject: "Welcome", Body: "Welcome to our service!"},
        {To: "user2@example.com", Subject: "Notification", Body: "You have a new notification."},
        {To: "user3@example.com", Subject: "Reminder", Body: "Don't forget about our meeting."},
    }
    
    for _, task := range tasks {
        if err := publishEmailTask(ch, task); err != nil {
            log.Printf("Failed to publish task: %v", err)
        } else {
            log.Printf("Published task for: %s", task.To)
        }
    }
    
    // Ожидание завершения работы
    log.Println("Press Ctrl+C to exit")
    select {}
}
//Задание: Асинхронная обработка задач с RabbitMQ
```
113. Микросервисная архитектура
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/gorilla/mux"
)

// Сервис пользователей
type UserService struct {
    users map[string]*User
}

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func NewUserService() *UserService {
    return &UserService{
        users: map[string]*User{
            "1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
            "2": {ID: "2", Name: "Bob", Email: "bob@example.com"},
        },
    }
}

func (s *UserService) GetUser(id string) (*User, error) {
    user, exists := s.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

// Сервис заказов
type OrderService struct {
    orders map[string]*Order
}

type Order struct {
    ID     string  `json:"id"`
    UserID string  `json:"user_id"`
    Amount float64 `json:"amount"`
    Status string  `json:"status"`
}

func NewOrderService() *OrderService {
    return &OrderService{
        orders: map[string]*Order{
            "1": {ID: "1", UserID: "1", Amount: 100.50, Status: "completed"},
            "2": {ID: "2", UserID: "2", Amount: 200.75, Status: "pending"},
        },
    }
}

func (s *OrderService) GetUserOrders(userID string) ([]*Order, error) {
    var userOrders []*Order
    for _, order := range s.orders {
        if order.UserID == userID {
            userOrders = append(userOrders, order)
        }
    }
    return userOrders, nil
}

// API Gateway
type APIGateway struct {
    userService  *UserService
    orderService *OrderService
}

func NewAPIGateway() *APIGateway {
    return &APIGateway{
        userService:  NewUserService(),
        orderService: NewOrderService(),
    }
}

type UserWithOrders struct {
    User   *User    `json:"user"`
    Orders []*Order `json:"orders"`
}

func (g *APIGateway) getUserWithOrders(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    
    // Получаем пользователя
    user, err := g.userService.GetUser(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    // Получаем заказы пользователя
    orders, err := g.orderService.GetUserOrders(userID)
    if err != nil {
        http.Error(w, "Error fetching orders", http.StatusInternalServerError)
        return
    }
    
    response := UserWithOrders{
        User:   user,
        Orders: orders,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Health check endpoint
func (g *APIGateway) healthCheck(w http.ResponseWriter, r *http.Request) {
    health := map[string]string{
        "status":    "healthy",
        "timestamp": time.Now().Format(time.RFC3339),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(health)
}

func main() {
    gateway := NewAPIGateway()
    
    r := mux.NewRouter()
    r.HandleFunc("/health", gateway.healthCheck).Methods("GET")
    r.HandleFunc("/users/{id}/orders", gateway.getUserWithOrders).Methods("GET")
    
    log.Println("API Gateway started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: Простая микросервисная архитектура с API Gateway
```
114. Docker-совместимое приложение
```go
package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
)

type Config struct {
    Port        int
    Environment string
    DatabaseURL string
    RedisURL    string
}

func loadConfig() *Config {
    port, _ := strconv.Atoi(getEnv("PORT", "8080"))
    
    return &Config{
        Port:        port,
        Environment: getEnv("ENVIRONMENT", "development"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/db"),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"status": "ok", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

func main() {
    config := loadConfig()
    
    log.Printf("Starting server in %s environment", config.Environment)
    log.Printf("Server listening on port %d", config.Port)
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Dockerized Go App!")
    })
    
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Environment: %s\n", config.Environment)
        fmt.Fprintf(w, "Database: %s\n", config.DatabaseURL)
        fmt.Fprintf(w, "Redis: %s\n", config.RedisURL)
    })
    
    addr := fmt.Sprintf(":%d", config.Port)
    log.Fatal(http.ListenAndServe(addr, nil))
}

// Dockerfile для этого приложения:
/*
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
*/

// docker-compose.yml:
/*
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ENVIRONMENT=production
      - DATABASE_URL=postgres://user:pass@db:5432/app
      - REDIS_URL=redis://redis:6379
    depends_on:
      - db
      - redis

  db:
    image: postgres:13
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
      - POSTGRES_DB=app

  redis:
    image: redis:6-alpine
*/
//Задание: Создание Docker-совместимого приложения
```
115. Kubernetes-ready приложение
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type App struct {
    server *http.Server
}

func NewApp(port string) *App {
    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/health", healthHandler)
    mux.HandleFunc("/ready", readyHandler)
    mux.HandleFunc("/metrics", metricsHandler)
    
    server := &http.Server{
        Addr:         ":" + port,
        Handler:      mux,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    return &App{server: server}
}

func (a *App) Start() error {
    log.Printf("Server starting on %s", a.server.Addr)
    return a.server.ListenAndServe()
}

func (a *App) Stop() error {
    log.Println("Server shutting down...")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    return a.server.Shutdown(ctx)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(w, "Hello from Kubernetes! Host: %s\n", hostname)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"status": "healthy", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
    // Проверка зависимостей (БД, Redis, etc.)
    if checkDependencies() {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintf(w, `{"status": "ready", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
    } else {
        http.Error(w, `{"status": "not ready"}`, http.StatusServiceUnavailable)
    }
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    // Простые метрики для Prometheus
    metrics := `
# HELP go_app_requests_total Total number of requests
# TYPE go_app_requests_total counter
go_app_requests_total 100
# HELP go_app_requests_duration_seconds Duration of requests
# TYPE go_app_requests_duration_seconds histogram
go_app_requests_duration_seconds_bucket{le="0.1"} 50
go_app_requests_duration_seconds_bucket{le="0.5"} 80
go_app_requests_duration_seconds_bucket{le="1.0"} 95
go_app_requests_duration_seconds_bucket{le="+Inf"} 100
`
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(metrics))
}

func checkDependencies() bool {
    // Проверка подключения к БД, Redis и другим зависимостям
    return true
}

func main() {
    port := getEnv("PORT", "8080")
    app := NewApp(port)
    
    // Graceful shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        if err := app.Start(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()
    
    <-stop
    app.Stop()
    log.Println("Server stopped gracefully")
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

// Kubernetes deployment.yaml:
/*
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
      - name: go-app
        image: your-registry/go-app:latest
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        - name: ENVIRONMENT
          value: "production"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "64Mi"
            cpu: "50m"
          limits:
            memory: "128Mi"
            cpu: "100m"
*/
//Задание: Создание Kubernetes-ready приложения
```
116. Приложение с мониторингом
```go
package main
import (
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    // Метрики
    httpRequestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "path"},
    )
    
    activeConnections = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of active connections",
        },
    )
    
    businessTransactions = promauto.NewCounter(
        prometheus.CounterOpts{
            Name: "business_transactions_total",
            Help: "Total number of business transactions",
        },
    )
    
    customMetrics = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "custom_metric",
            Help: "Custom business metric",
        },
        []string{"type"},
    )
)

type monitoringResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (mrw *monitoringResponseWriter) WriteHeader(code int) {
    mrw.statusCode = code
    mrw.ResponseWriter.WriteHeader(code)
}

func monitoringMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        activeConnections.Inc()
        
        mrw := &monitoringResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        
        next.ServeHTTP(mrw, r)
        
        duration := time.Since(start).Seconds()
        httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", mrw.statusCode)).Inc()
        httpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
        activeConnections.Dec()
    })
}

func businessHandler(w http.ResponseWriter, r *http.Request) {
    // Имитация бизнес-логики
    time.Sleep(time.Duration(100+time.Now().UnixNano()%100) * time.Millisecond)
    
    businessTransactions.Inc()
    customMetrics.WithLabelValues("success").Inc()
    
    w.Write([]byte("Business transaction completed"))
}

func main() {
    // Установка начальных значений метрик
    customMetrics.WithLabelValues("success").Set(0)
    customMetrics.WithLabelValues("error").Set(0)
    
    mux := http.NewServeMux()
    mux.Handle("/metrics", promhttp.Handler())
    mux.Handle("/business", monitoringMiddleware(http.HandlerFunc(businessHandler)))
    mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    }))
    
    // Запуск фоновой задачи для обновления метрик
    go updateCustomMetrics()
    
    log.Println("Monitoring server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func updateCustomMetrics() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        // Обновление кастомных метрик
        customMetrics.WithLabelValues("background").Set(float64(time.Now().Unix() % 100))
    }
}
//Задание: Приложение с комплексным мониторингом Prometheus
```
117. Distributed tracing с Jaeger
```go
package main
import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
    "go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func initTracer() (*sdktrace.TracerProvider, error) {
    // Создание Jaeger exporter
    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
    if err != nil {
        return nil, err
    }
    
    // Настройка провайдера трассировки
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("my-service"),
            attribute.String("environment", "development"),
        )),
    )
    
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{}, propagation.Baggage{}))
    
    tracer = tp.Tracer("example-tracer")
    return tp, nil
}

func tracingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
        
        ctx, span := tracer.Start(ctx, "http-request",
            trace.WithAttributes(
                attribute.String("http.method", r.Method),
                attribute.String("http.url", r.URL.String()),
                attribute.String("http.user_agent", r.UserAgent()),
            ))
        defer span.End()
        
        // Передаем контекст с трассировкой дальше
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
        
        span.SetAttributes(attribute.Int("http.status_code", http.StatusOK))
    })
}

func orderProcessing(ctx context.Context, orderID string) error {
    ctx, span := tracer.Start(ctx, "order-processing")
    defer span.End()
    
    span.SetAttributes(
        attribute.String("order.id", orderID),
        attribute.String("processing.stage", "started"),
    )
    
    // Имитация обработки заказа
    time.Sleep(50 * time.Millisecond)
    
    if err := validateOrder(ctx, orderID); err != nil {
        span.RecordError(err)
        return err
    }
    
    if err := processPayment(ctx, orderID); err != nil {
        span.RecordError(err)
        return err
    }
    
    span.SetAttributes(attribute.String("processing.stage", "completed"))
    return nil
}

func validateOrder(ctx context.Context, orderID string) error {
    ctx, span := tracer.Start(ctx, "validate-order")
    defer span.End()
    
    time.Sleep(20 * time.Millisecond)
    
    // Имитация проверки
    if orderID == "invalid" {
        return fmt.Errorf("invalid order")
    }
    
    span.SetAttributes(attribute.Bool("validation.success", true))
    return nil
}

func processPayment(ctx context.Context, orderID string) error {
    ctx, span := tracer.Start(ctx, "process-payment")
    defer span.End()
    
    time.Sleep(30 * time.Millisecond)
    
    span.SetAttributes(attribute.Bool("payment.success", true))
    return nil
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    orderID := r.URL.Query().Get("order_id")
    if orderID == "" {
        http.Error(w, "order_id is required", http.StatusBadRequest)
        return
    }
    
    if err := orderProcessing(ctx, orderID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Write([]byte(fmt.Sprintf("Order %s processed successfully", orderID)))
}

func main() {
    tp, err := initTracer()
    if err != nil {
        log.Fatal(err)
    }
    defer tp.Shutdown(context.Background())
    
    mux := http.NewServeMux()
    mux.Handle("/order", tracingMiddleware(http.HandlerFunc(orderHandler)))
    
    log.Println("Tracing server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
//Задание: Distributed tracing с Jaeger
```
118. Feature flags с запуском в продакшн
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
    
    "github.com/gorilla/mux"
)

type FeatureFlag struct {
    Name          string    `json:"name"`
    Description   string    `json:"description"`
    Enabled       bool      `json:"enabled"`
    Percentage    int       `json:"percentage"` // 0-100
    Users         []string  `json:"users"`      // Список пользователей
    UpdatedAt     time.Time `json:"updated_at"`
    EnableAfter   time.Time `json:"enable_after,omitempty"`
    DisableAfter  time.Time `json:"disable_after,omitempty"`
}

type FeatureManager struct {
    mu     sync.RWMutex
    flags  map[string]*FeatureFlag
    config string
}

func NewFeatureManager() *FeatureManager {
    fm := &FeatureManager{
        flags: make(map[string]*FeatureFlag),
    }
    
    // Флаги по умолчанию
    fm.flags["new_ui"] = &FeatureFlag{
        Name:        "new_ui",
        Description: "Новый пользовательский интерфейс",
        Enabled:     false,
        Percentage:  0,
        UpdatedAt:   time.Now(),
    }
    
    fm.flags["beta_features"] = &FeatureFlag{
        Name:        "beta_features",
        Description: "Бета-функциональность",
        Enabled:     true,
        Percentage:  100,
        UpdatedAt:   time.Now(),
    }
    
    fm.flags["gradual_rollout"] = &FeatureFlag{
        Name:        "gradual_rollout",
        Description: "Постепенный rollout",
        Enabled:     true,
        Percentage:  50, // 50% пользователей
        UpdatedAt:   time.Now(),
    }
    
    return fm
}

func (fm *FeatureManager) IsEnabled(flagName string) bool {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    flag, exists := fm.flags[flagName]
    if !exists {
        return false
    }
    
    // Проверка временных ограничений
    now := time.Now()
    if !flag.EnableAfter.IsZero() && now.Before(flag.EnableAfter) {
        return false
    }
    if !flag.DisableAfter.IsZero() && now.After(flag.DisableAfter) {
        return false
    }
    
    return flag.Enabled
}

func (fm *FeatureManager) IsEnabledForUser(flagName, userID string) bool {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    flag, exists := fm.flags[flagName]
    if !exists {
        return false
    }
    
    // Проверка конкретного пользователя
    for _, user := range flag.Users {
        if user == userID {
            return true
        }
    }
    
    // Проверка процентного распределения
    if flag.Percentage > 0 {
        // Простой хеш для детерминированного распределения
        hash := simpleHash(userID + flagName)
        return hash%100 < flag.Percentage
    }
    
    return flag.Enabled
}

func simpleHash(s string) int {
    hash := 0
    for _, char := range s {
        hash = (hash << 5) - hash + int(char)
    }
    return hash & 0x7FFFFFFF // Положительное число
}

func (fm *FeatureManager) SetFlag(flagName string, flag *FeatureFlag) {
    fm.mu.Lock()
    defer fm.mu.Unlock()
    
    flag.UpdatedAt = time.Now()
    fm.flags[flagName] = flag
}

func (fm *FeatureManager) GetFlags() map[string]*FeatureFlag {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    
    flags := make(map[string]*FeatureFlag)
    for k, v := range fm.flags {
        flags[k] = v
    }
    return flags
}

func main() {
    featureManager := NewFeatureManager()
    
    r := mux.NewRouter()
    
    // API для управления флагами
    r.HandleFunc("/api/flags", func(w http.ResponseWriter, r *http.Request) {
        flags := featureManager.GetFlags()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(flags)
    }).Methods("GET")
    
    r.HandleFunc("/api/flags/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        flagName := vars["name"]
        
        var flag FeatureFlag
        if err := json.NewDecoder(r.Body).Decode(&flag); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        featureManager.SetFlag(flagName, &flag)
        w.WriteHeader(http.StatusOK)
    }).Methods("PUT")
    
    // Endpoint для проверки флагов
    r.HandleFunc("/api/check/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        flagName := vars["name"]
        userID := r.URL.Query().Get("user_id")
        
        var enabled bool
        if userID != "" {
            enabled = featureManager.IsEnabledForUser(flagName, userID)
        } else {
            enabled = featureManager.IsEnabled(flagName)
        }
        
        response := map[string]interface{}{
            "enabled": enabled,
            "flag":    flagName,
            "user_id": userID,
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }).Methods("GET")
    
    // Демонстрационная страница
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        userID := "user123" // В реальном приложении из сессии или JWT
        
        html := fmt.Sprintf(`
        <html>
        <body>
            <h1>Feature Flags Demo</h1>
            <p>User ID: %s</p>
            <ul>
                <li>New UI: %v</li>
                <li>Beta Features: %v</li>
                <li>Gradual Rollout: %v</li>
            </ul>
            <a href="/api/flags">View All Flags</a>
        </body>
        </html>`, userID,
            featureManager.IsEnabledForUser("new_ui", userID),
            featureManager.IsEnabledForUser("beta_features", userID),
            featureManager.IsEnabledForUser("gradual_rollout", userID))
        
        w.Write([]byte(html))
    })
    
    log.Println("Feature flag server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: Продвинутая система feature flags
```
119. A/B тестирование
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
    
    "github.com/gorilla/mux"
)

type Experiment struct {
    Name        string            `json:"name"`
    Variants    []Variant         `json:"variants"`
    StartDate   time.Time         `json:"start_date"`
    EndDate     time.Time         `json:"end_date"`
    Enabled     bool              `json:"enabled"`
    TargetUsers []string          `json:"target_users,omitempty"`
    Metadata    map[string]string `json:"metadata,omitempty"`
}

type Variant struct {
    Name        string                 `json:"name"`
    Weight      int                    `json:"weight"` // Относительный вес (1-100)
    Parameters  map[string]interface{} `json:"parameters"`
    IsControl   bool                   `json:"is_control"`
}

type ExperimentResult struct {
    ExperimentName string               `json:"experiment_name"`
    VariantName    string               `json:"variant_name"`
    UserID         string               `json:"user_id"`
    Timestamp      time.Time            `json:"timestamp"`
    Events         []Event              `json:"events"`
    Parameters     map[string]interface{} `json:"parameters"`
}

type Event struct {
    Type      string                 `json:"type"`
    Timestamp time.Time              `json:"timestamp"`
    Data      map[string]interface{} `json:"data"`
}

type ABTestManager struct {
    mu         sync.RWMutex
    experiments map[string]*Experiment
    results    []ExperimentResult
}

func NewABTestManager() *ABTestManager {
    return &ABTestManager{
        experiments: make(map[string]*Experiment),
        results:     make([]ExperimentResult, 0),
    }
}

func (m *ABTestManager) AddExperiment(exp *Experiment) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.experiments[exp.Name] = exp
}

func (m *ABTestManager) GetVariant(experimentName, userID string) (*Variant, error) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    
    exp, exists := m.experiments[experimentName]
    if !exists || !exp.Enabled {
        return nil, fmt.Errorf("experiment not found or disabled")
    }
    
    // Проверка времени эксперимента
    now := time.Now()
    if now.Before(exp.StartDate) || now.After(exp.EndDate) {
        return nil, fmt.Errorf("experiment not active")
    }
    
    // Проверка целевых пользователей
    if len(exp.TargetUsers) > 0 {
        found := false
        for _, user := range exp.TargetUsers {
            if user == userID {
                found = true
                break
            }
        }
        if !found {
            return nil, fmt.Errorf("user not in target audience")
        }
    }
    
    // Детерминированное распределение по вариантам
    totalWeight := 0
    for _, variant := range exp.Variants {
        totalWeight += variant.Weight
    }
    
    if totalWeight == 0 {
        return nil, fmt.Errorf("invalid variant weights")
    }
    
    // Детерминированный выбор на основе userID
    hash := simpleHash(userID + experimentName)
    choice := hash % totalWeight
    
    currentWeight := 0
    for _, variant := range exp.Variants {
        currentWeight += variant.Weight
        if choice < currentWeight {
            return &variant, nil
        }
    }
    
    return &exp.Variants[0], nil // fallback
}

func (m *ABTestManager) RecordEvent(experimentName, variantName, userID, eventType string, data map[string]interface{}) {
    m.mu.Lock()
    defer m.mu.Unlock()
    
    result := ExperimentResult{
        ExperimentName: experimentName,
        VariantName:    variantName,
        UserID:         userID,
        Timestamp:      time.Now(),
        Events: []Event{
            {
                Type:      eventType,
                Timestamp: time.Now(),
                Data:      data,
            },
        },
    }
    
    m.results = append(m.results, result)
}

func (m *ABTestManager) GetResults() []ExperimentResult {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return m.results
}

func main() {
    abTestManager := NewABTestManager()
    
    // Создание тестового эксперимента
    buttonColorExperiment := &Experiment{
        Name:      "button_color",
        Enabled:   true,
        StartDate: time.Now(),
        EndDate:   time.Now().Add(30 * 24 * time.Hour),
        Variants: []Variant{
            {
                Name:      "control",
                Weight:    50,
                IsControl: true,
                Parameters: map[string]interface{}{
                    "button_color": "blue",
                    "button_text":  "Buy Now",
                },
            },
            {
                Name:   "variant_a",
                Weight: 25,
                Parameters: map[string]interface{}{
                    "button_color": "green",
                    "button_text":  "Purchase",
                },
            },
            {
                Name:   "variant_b",
                Weight: 25,
                Parameters: map[string]interface{}{
                    "button_color": "red",
                    "button_text":  "Get It Now",
                },
            },
        },
    }
    
    abTestManager.AddExperiment(buttonColorExperiment)
    
    r := mux.NewRouter()
    
    // API для получения варианта эксперимента
    r.HandleFunc("/api/experiment/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        experimentName := vars["name"]
        userID := r.URL.Query().Get("user_id")
        
        if userID == "" {
            http.Error(w, "user_id is required", http.StatusBadRequest)
            return
        }
        
        variant, err := abTestManager.GetVariant(experimentName, userID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        
        response := map[string]interface{}{
            "experiment": experimentName,
            "variant":    variant.Name,
            "parameters": variant.Parameters,
            "user_id":    userID,
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }).Methods("GET")
    
    // API для записи событий
    r.HandleFunc("/api/event", func(w http.ResponseWriter, r *http.Request) {
        var event struct {
            Experiment string                 `json:"experiment"`
            Variant    string                 `json:"variant"`
            UserID     string                 `json:"user_id"`
            EventType  string                 `json:"event_type"`
            Data       map[string]interface{} `json:"data"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        abTestManager.RecordEvent(event.Experiment, event.Variant, event.UserID, event.EventType, event.Data)
        w.WriteHeader(http.StatusOK)
    }).Methods("POST")
    
    // API для получения результатов
    r.HandleFunc("/api/results", func(w http.ResponseWriter, r *http.Request) {
        results := abTestManager.GetResults()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(results)
    }).Methods("GET")
    
    log.Println("A/B Testing server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: Система A/B тестирования
```
120. Real-time дашборд с SSE
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
    
    "github.com/gorilla/mux"
)

type Metric struct {
    Name      string  `json:"name"`
    Value     float64 `json:"value"`
    Timestamp int64   `json:"timestamp"`
}

type Dashboard struct {
    mu          sync.RWMutex
    clients     map[chan []byte]bool
    metrics     map[string]*Metric
}

func NewDashboard() *Dashboard {
    return &Dashboard{
        clients: make(map[chan []byte]bool),
        metrics: make(map[string]*Metric),
    }
}

func (d *Dashboard) AddClient(ch chan []byte) {
    d.mu.Lock()
    defer d.mu.Unlock()
    d.clients[ch] = true
}

func (d *Dashboard) RemoveClient(ch chan []byte) {
    d.mu.Lock()
    defer d.mu.Unlock()
    delete(d.clients, ch)
    close(ch)
}

func (d *Dashboard) Broadcast(data []byte) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    
    for client := range d.clients {
        select {
        case client <- data:
        default:
            // Клиент не успевает читать, пропускаем
        }
    }
}

func (d *Dashboard) UpdateMetric(name string, value float64) {
    d.mu.Lock()
    defer d.mu.Unlock()
    
    d.metrics[name] = &Metric{
        Name:      name,
        Value:     value,
        Timestamp: time.Now().Unix(),
    }
    
    // Рассылаем обновление всем клиентам
    metricJSON, _ := json.Marshal(d.metrics[name])
    d.Broadcast(metricJSON)
}

func (d *Dashboard) GetMetrics() map[string]*Metric {
    d.mu.RLock()
    defer d.mu.RUnlock()
    
    metrics := make(map[string]*Metric)
    for k, v := range d.metrics {
        metrics[k] = v
    }
    return metrics
}

func (d *Dashboard) StartMetricGenerator() {
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        // Генерация случайных метрик
        d.UpdateMetric("cpu_usage", rand.Float64()*100)
        d.UpdateMetric("memory_usage", 30+rand.Float64()*50)
        d.UpdateMetric("active_users", float64(rand.Intn(1000)))
        d.UpdateMetric("response_time", rand.Float64()*500)
        d.UpdateMetric("error_rate", rand.Float64()*5)
    }
}

func sseHandler(dashboard *Dashboard) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Устанавливаем заголовки для SSE
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        
        // Создаем канал для клиента
        ch := make(chan []byte, 10)
        dashboard.AddClient(ch)
        
        // Удаляем клиента при закрытии соединения
        defer dashboard.RemoveClient(ch)
        
        // Отправляем начальные данные
        metrics := dashboard.GetMetrics()
        initialData, _ := json.Marshal(metrics)
        fmt.Fprintf(w, "data: %s\n\n", initialData)
        w.(http.Flusher).Flush()
        
        // Отправляем обновления
        for {
            select {
            case data := <-ch:
                fmt.Fprintf(w, "data: %s\n\n", data)
                w.(http.Flusher).Flush()
            case <-r.Context().Done():
                return
            }
        }
    }
}

func metricsHandler(dashboard *Dashboard) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        metrics := dashboard.GetMetrics()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(metrics)
    }
}

func main() {
    dashboard := NewDashboard()
    
    // Запуск генератора метрик
    go dashboard.StartMetricGenerator()
    
    r := mux.NewRouter()
    
    // SSE endpoint для реального времени
    r.HandleFunc("/events", sseHandler(dashboard))
    
    // REST API для получения текущих метрик
    r.HandleFunc("/metrics", metricsHandler(dashboard))
    
    // HTML страница с дашбордом
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Real-time Dashboard</title>
            <script>
                const eventSource = new EventSource('/events');
                const metrics = {};
                
                eventSource.onmessage = function(event) {
                    const data = JSON.parse(event.data);
                    metrics[data.name] = data;
                    updateDashboard();
                };
                
                function updateDashboard() {
                    const container = document.getElementById('metrics');
                    container.innerHTML = '';
                    
                    for (const name in metrics) {
                        const metric = metrics[name];
                        const div = document.createElement('div');
                        div.className = 'metric';
                        div.innerHTML = \`
                            <h3>\${metric.name}</h3>
                            <div class="value">\${metric.value.toFixed(2)}</div>
                            <div class="timestamp">\${new Date(metric.timestamp * 1000).toLocaleTimeString()}</div>
                        \`;
                        container.appendChild(div);
                    }
                }
                
                // Стилизация
                const style = document.createElement('style');
                style.textContent = \`
                    .metric {
                        border: 1px solid #ddd;
                        border-radius: 5px;
                        padding: 10px;
                        margin: 5px;
                        display: inline-block;
                        width: 200px;
                    }
                    .value {
                        font-size: 24px;
                        font-weight: bold;
                    }
                    .timestamp {
                        color: #666;
                        font-size: 12px;
                    }
                \`;
                document.head.appendChild(style);
            </script>
        </head>
        <body>
            <h1>Real-time Metrics Dashboard</h1>
            <div id="metrics"></div>
        </body>
        </html>`
        w.Write([]byte(html))
    })
    
    log.Println("Real-time dashboard started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
//Задание: Real-time дашборд с Server-Sent Events
```
Специализированные техники и продвинутые сценарии
121. Работа с бинарными данными
```go
package main
import (
    "encoding/binary"
    "fmt"
    "log"
)

type Person struct {
    Name   [32]byte
    Age    uint32
    Height float32
    Active bool
}

func main() {
    // Создание структуры
    person := Person{
        Name:   [32]byte{},
        Age:    30,
        Height: 175.5,
        Active: true,
    }
    copy(person.Name[:], "John Doe")
    
    // Сериализация в бинарный формат
    data := make([]byte, 32+4+4+1) // Name[32] + Age[4] + Height[4] + Active[1]
    
    copy(data[0:32], person.Name[:])
    binary.LittleEndian.PutUint32(data[32:36], person.Age)
    binary.LittleEndian.PutUint32(data[36:40], math.Float32bits(person.Height))
    if person.Active {
        data[40] = 1
    } else {
        data[40] = 0
    }
    
    fmt.Printf("Binary data: %x\n", data)
    
    // Десериализация обратно
    var decodedPerson Person
    copy(decodedPerson.Name[:], data[0:32])
    decodedPerson.Age = binary.LittleEndian.Uint32(data[32:36])
    decodedPerson.Height = math.Float32frombits(binary.LittleEndian.Uint32(data[36:40]))
    decodedPerson.Active = data[40] == 1
    
    fmt.Printf("Decoded: Name=%s, Age=%d, Height=%.1f, Active=%t\n",
        string(decodedPerson.Name[:]), decodedPerson.Age, decodedPerson.Height, decodedPerson.Active)
}
//Задание: Работа с бинарными данными и сериализация структур
```
122. Bit manipulation
```go
package main
import (
    "fmt"
)

type BitSet struct {
    data []uint64
}

func NewBitSet(size int) *BitSet {
    words := (size + 63) / 64
    return &BitSet{
        data: make([]uint64, words),
    }
}

func (bs *BitSet) Set(bit int) {
    word := bit / 64
    pos := uint(bit % 64)
    bs.data[word] |= (1 << pos)
}

func (bs *BitSet) Clear(bit int) {
    word := bit / 64
    pos := uint(bit % 64)
    bs.data[word] &^= (1 << pos)
}

func (bs *BitSet) Get(bit int) bool {
    word := bit / 64
    pos := uint(bit % 64)
    return (bs.data[word] & (1 << pos)) != 0
}

func (bs *BitSet) Toggle(bit int) {
    word := bit / 64
    pos := uint(bit % 64)
    bs.data[word] ^= (1 << pos)
}

func (bs *BitSet) Count() int {
    count := 0
    for _, word := range bs.data {
        count += countBits(word)
    }
    return count
}

func countBits(x uint64) int {
    count := 0
    for x != 0 {
        count++
        x &= x - 1
    }
    return count
}

func main() {
    bs := NewBitSet(100)
    
    // Установка битов
    bs.Set(5)
    bs.Set(10)
    bs.Set(15)
    
    fmt.Printf("Bit 5: %t\n", bs.Get(5))   // true
    fmt.Printf("Bit 6: %t\n", bs.Get(6))   // false
    
    // Переключение бита
    bs.Toggle(5)
    fmt.Printf("Bit 5 after toggle: %t\n", bs.Get(5)) // false
    
    // Подсчет установленных битов
    fmt.Printf("Bits set: %d\n", bs.Count()) // 2
    
    // Битовая арифметика
    a := uint8(0b10101010)
    b := uint8(0b11001100)
    
    fmt.Printf("a: %08b\n", a)
    fmt.Printf("b: %08b\n", b)
    fmt.Printf("a & b: %08b\n", a&b)  // AND
    fmt.Printf("a | b: %08b\n", a|b)  // OR
    fmt.Printf("a ^ b: %08b\n", a^b)  // XOR
    fmt.Printf("a << 2: %08b\n", a<<2) // Left shift
    fmt.Printf("a >> 2: %08b\n", a>>2) // Right shift
}
//Задание: Работа с битовыми операциями и битовыми множествами
```
123. Memory pool для объектов
```go
package main
import (
    "fmt"
    "sync"
    "time"
)

type Object struct {
    ID    int
    Data  string
    Timestamp time.Time
}

type ObjectPool struct {
    pool    chan *Object
    created int
    maxSize int
    mu      sync.Mutex
}

func NewObjectPool(maxSize int) *ObjectPool {
    return &ObjectPool{
        pool:    make(chan *Object, maxSize),
        maxSize: maxSize,
    }
}

func (p *ObjectPool) Get() *Object {
    select {
    case obj := <-p.pool:
        return obj
    default:
        p.mu.Lock()
        defer p.mu.Unlock()
        
        if p.created < p.maxSize {
            p.created++
            return &Object{
                Timestamp: time.Now(),
            }
        }
        
        // Если пул полон, создаем новый объект
        return &Object{
            Timestamp: time.Now(),
        }
    }
}

func (p *ObjectPool) Put(obj *Object) {
    // Сбрасываем состояние объекта
    obj.ID = 0
    obj.Data = ""
    
    select {
    case p.pool <- obj:
        // Успешно вернули в пул
    default:
        // Пул полон, объект будет собран GC
    }
}

func (p *ObjectPool) Stats() {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    fmt.Printf("Pool stats: created=%d, in pool=%d/%d\n", 
        p.created, len(p.pool), p.maxSize)
}

func main() {
    pool := NewObjectPool(5)
    
    var wg sync.WaitGroup
    
    // Использование пула из нескольких горутин
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Получаем объект из пула
            obj := pool.Get()
            obj.ID = id
            obj.Data = fmt.Sprintf("Data from goroutine %d", id)
            
            fmt.Printf("Goroutine %d using object: %+v\n", id, obj)
            
            // Имитация работы
            time.Sleep(100 * time.Millisecond)
            
            // Возвращаем объект в пул
            pool.Put(obj)
        }(i)
    }
    
    wg.Wait()
    pool.Stats()
}
//Задание: Реализация пула объектов для уменьшения нагрузки на GC
```
124. Lock-free структуры данных
```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
    "unsafe"
)

type Node struct {
    value interface{}
    next  unsafe.Pointer
}

type LockFreeQueue struct {
    head unsafe.Pointer
    tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
    dummy := &Node{}
    return &LockFreeQueue{
        head: unsafe.Pointer(dummy),
        tail: unsafe.Pointer(dummy),
    }
}

func (q *LockFreeQueue) Enqueue(value interface{}) {
    newNode := &Node{value: value}
    
    for {
        tail := atomic.LoadPointer(&q.tail)
        tailNode := (*Node)(tail)
        next := atomic.LoadPointer(&tailNode.next)
        
        if next == nil {
            if atomic.CompareAndSwapPointer(&tailNode.next, next, unsafe.Pointer(newNode)) {
                atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newNode))
                return
            }
        } else {
            atomic.CompareAndSwapPointer(&q.tail, tail, next)
        }
    }
}

func (q *LockFreeQueue) Dequeue() interface{} {
    for {
        head := atomic.LoadPointer(&q.head)
        headNode := (*Node)(head)
        tail := atomic.LoadPointer(&q.tail)
        next := atomic.LoadPointer(&headNode.next)
        
        if head == tail {
            if next == nil {
                return nil // Очередь пуста
            }
            atomic.CompareAndSwapPointer(&q.tail, tail, next)
        } else {
            if atomic.CompareAndSwapPointer(&q.head, head, next) {
                nextNode := (*Node)(next)
                return nextNode.value
            }
        }
    }
}

func main() {
    queue := NewLockFreeQueue()
    var wg sync.WaitGroup
    
    // Писатели
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 5; j++ {
                value := fmt.Sprintf("value-%d-%d", id, j)
                queue.Enqueue(value)
                fmt.Printf("Producer %d enqueued: %s\n", id, value)
            }
        }(i)
    }
    
    // Читатели
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 7; j++ {
                value := queue.Dequeue()
                if value != nil {
                    fmt.Printf("Consumer %d dequeued: %s\n", id, value)
                }
            }
        }(i)
    }
    
    wg.Wait()
}
//Задание: Lock-free очередь на основе атомарных операций
```
125. SIMD-подобные операции
```go
package main
import (
    "fmt"
    "time"
)

// Векторные операции без SIMD (для демонстрации концепции)
type Vector struct {
    data []float64
}

func NewVector(size int) *Vector {
    return &Vector{
        data: make([]float64, size),
    }
}

func (v *Vector) Add(other *Vector) *Vector {
    result := NewVector(len(v.data))
    for i := range v.data {
        result.data[i] = v.data[i] + other.data[i]
    }
    return result
}

func (v *Vector) Multiply(other *Vector) *Vector {
    result := NewVector(len(v.data))
    for i := range v.data {
        result.data[i] = v.data[i] * other.data[i]
    }
    return result
}

func (v *Vector) Dot(other *Vector) float64 {
    result := 0.0
    for i := range v.data {
        result += v.data[i] * other.data[i]
    }
    return result
}

func (v *Vector) Scale(scalar float64) *Vector {
    result := NewVector(len(v.data))
    for i := range v.data {
        result.data[i] = v.data[i] * scalar
    }
    return result
}

// Оптимизированная версия с развертыванием цикла
func (v *Vector) AddOptimized(other *Vector) *Vector {
    result := NewVector(len(v.data))
    n := len(v.data)
    
    // Развертывание цикла для лучшей производительности
    for i := 0; i < n-3; i += 4 {
        result.data[i] = v.data[i] + other.data[i]
        result.data[i+1] = v.data[i+1] + other.data[i+1]
        result.data[i+2] = v.data[i+2] + other.data[i+2]
        result.data[i+3] = v.data[i+3] + other.data[i+3]
    }
    
    // Обработка оставшихся элементов
    for i := n - (n % 4); i < n; i++ {
        result.data[i] = v.data[i] + other.data[i]
    }
    
    return result
}

func main() {
    size := 1000000
    a := NewVector(size)
    b := NewVector(size)
    
    // Инициализация векторов
    for i := range a.data {
        a.data[i] = float64(i)
        b.data[i] = float64(i * 2)
    }
    
    // Тестирование производительности
    start := time.Now()
    result1 := a.Add(b)
    elapsed1 := time.Since(start)
    
    start = time.Now()
    result2 := a.AddOptimized(b)
    elapsed2 := time.Since(start)
    
    fmt.Printf("Standard add: %v\n", elapsed1)
    fmt.Printf("Optimized add: %v\n", elapsed2)
    fmt.Printf("Speedup: %.2fx\n", float64(elapsed1)/float64(elapsed2))
    
    // Проверка корректности
    dot := a.Dot(b)
    fmt.Printf("Dot product: %.2f\n", dot)
    
    scaled := a.Scale(2.5)
    fmt.Printf("First 5 scaled values: %v\n", scaled.data[:5])
}
//Задание: Векторные операции и оптимизации циклов
```
126. Genetic algorithm
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
    "sort"
    "time"
)

type Individual struct {
    genes   []float64
    fitness float64
}

type GeneticAlgorithm struct {
    populationSize int
    mutationRate  float64
    crossoverRate float64
    elitismCount  int
    target        []float64
}

func NewGeneticAlgorithm(popSize int, mutation, crossover float64, elitism int, target []float64) *GeneticAlgorithm {
    return &GeneticAlgorithm{
        populationSize: popSize,
        mutationRate:  mutation,
        crossoverRate: crossover,
        elitismCount:  elitism,
        target:        target,
    }
}

func (ga *GeneticAlgorithm) createIndividual() *Individual {
    ind := &Individual{
        genes: make([]float64, len(ga.target)),
    }
    for i := range ind.genes {
        ind.genes[i] = rand.Float64() * 10 // Случайные значения от 0 до 10
    }
    return ind
}

func (ga *GeneticAlgorithm) calculateFitness(individual *Individual) float64 {
    // Фитнес функция: обратная величина от суммы квадратов разностей
    sumSquares := 0.0
    for i, gene := range individual.genes {
        diff := gene - ga.target[i]
        sumSquares += diff * diff
    }
    return 1.0 / (1.0 + sumSquares)
}

func (ga *GeneticAlgorithm) crossover(parent1, parent2 *Individual) *Individual {
    child := &Individual{
        genes: make([]float64, len(parent1.genes)),
    }
    
    // Одноточечный кроссовер
    if rand.Float64() < ga.crossoverRate {
        crossoverPoint := rand.Intn(len(parent1.genes))
        for i := 0; i < crossoverPoint; i++ {
            child.genes[i] = parent1.genes[i]
        }
        for i := crossoverPoint; i < len(parent1.genes); i++ {
            child.genes[i] = parent2.genes[i]
        }
    } else {
        copy(child.genes, parent1.genes)
    }
    
    return child
}

func (ga *GeneticAlgorithm) mutate(individual *Individual) {
    for i := range individual.genes {
        if rand.Float64() < ga.mutationRate {
            // Случайная мутация
            individual.genes[i] += (rand.Float64() - 0.5) * 2
            // Ограничение значений
            individual.genes[i] = math.Max(0, math.Min(10, individual.genes[i]))
        }
    }
}

func (ga *GeneticAlgorithm) evolvePopulation(population []*Individual) []*Individual {
    newPopulation := make([]*Individual, ga.populationSize)
    
    // Элитизм: сохраняем лучших особей
    sort.Slice(population, func(i, j int) bool {
        return population[i].fitness > population[j].fitness
    })
    
    for i := 0; i < ga.elitismCount; i++ {
        newPopulation[i] = &Individual{
            genes:   make([]float64, len(population[i].genes)),
            fitness: population[i].fitness,
        }
        copy(newPopulation[i].genes, population[i].genes)
    }
    
    // Создание нового поколения
    for i := ga.elitismCount; i < ga.populationSize; i++ {
        // Селекция (турнирная)
        parent1 := ga.tournamentSelection(population, 3)
        parent2 := ga.tournamentSelection(population, 3)
        
        // Кроссовер
        child := ga.crossover(parent1, parent2)
        
        // Мутация
        ga.mutate(child)
        
        newPopulation[i] = child
    }
    
    return newPopulation
}

func (ga *GeneticAlgorithm) tournamentSelection(population []*Individual, tournamentSize int) *Individual {
    best := population[rand.Intn(len(population))]
    
    for i := 1; i < tournamentSize; i++ {
        contender := population[rand.Intn(len(population))]
        if contender.fitness > best.fitness {
            best = contender
        }
    }
    
    return best
}

func (ga *GeneticAlgorithm) Run(generations int) *Individual {
    rand.Seed(time.Now().UnixNano())
    
    // Создание начальной популяции
    population := make([]*Individual, ga.populationSize)
    for i := range population {
        population[i] = ga.createIndividual()
        population[i].fitness = ga.calculateFitness(population[i])
    }
    
    var bestIndividual *Individual
    
    for generation := 0; generation < generations; generation++ {
        // Эволюция популяции
        population = ga.evolvePopulation(population)
        
        // Пересчет фитнеса
        for _, individual := range population {
            individual.fitness = ga.calculateFitness(individual)
        }
        
        // Нахождение лучшей особи
        sort.Slice(population, func(i, j int) bool {
            return population[i].fitness > population[j].fitness
        })
        
        bestIndividual = population[0]
        
        if generation%100 == 0 {
            fmt.Printf("Generation %d: Best fitness = %.6f\n", generation, bestIndividual.fitness)
            fmt.Printf("Best genes: %v\n", bestIndividual.genes)
        }
        
        // Критерий остановки
        if bestIndividual.fitness > 0.95 {
            fmt.Printf("Solution found at generation %d\n", generation)
            break
        }
    }
    
    return bestIndividual
}

func main() {
    // Целевой вектор, который мы хотим найти
    target := []float64{2.5, 3.8, 1.2, 4.5, 2.1}
    
    ga := NewGeneticAlgorithm(
        100,  // Размер популяции
        0.01, // Вероятность мутации
        0.9,  // Вероятность кроссовера
        2,    // Количество элитных особей
        target,
    )
    
    best := ga.Run(1000)
    
    fmt.Printf("\nFinal result:\n")
    fmt.Printf("Target: %v\n", target)
    fmt.Printf("Found:  %v\n", best.genes)
    fmt.Printf("Fitness: %.6f\n", best.fitness)
}
//Задание: Реализация генетического алгоритма для оптимизации
```
127. Neural network (простая реализация)
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type NeuralNetwork struct {
    inputSize   int
    hiddenSize  int
    outputSize  int
    weightsIH   [][]float64 // Веса вход->скрытый слой
    weightsHO   [][]float64 // Веса скрытый->выходной слой
    biasH       []float64   // Смещения скрытого слоя
    biasO       []float64   // Смещения выходного слоя
    learningRate float64
}

func NewNeuralNetwork(input, hidden, output int, learningRate float64) *NeuralNetwork {
    nn := &NeuralNetwork{
        inputSize:    input,
        hiddenSize:   hidden,
        outputSize:   output,
        learningRate: learningRate,
    }
    
    // Инициализация весов случайными значениями
    nn.weightsIH = make([][]float64, hidden)
    for i := range nn.weightsIH {
        nn.weightsIH[i] = make([]float64, input)
        for j := range nn.weightsIH[i] {
            nn.weightsIH[i][j] = rand.Float64()*2 - 1 // [-1, 1]
        }
    }
    
    nn.weightsHO = make([][]float64, output)
    for i := range nn.weightsHO {
        nn.weightsHO[i] = make([]float64, hidden)
        for j := range nn.weightsHO[i] {
            nn.weightsHO[i][j] = rand.Float64()*2 - 1
        }
    }
    
    nn.biasH = make([]float64, hidden)
    for i := range nn.biasH {
        nn.biasH[i] = rand.Float64()*2 - 1
    }
    
    nn.biasO = make([]float64, output)
    for i := range nn.biasO {
        nn.biasO[i] = rand.Float64()*2 - 1
    }
    
    return nn
}

func sigmoid(x float64) float64 {
    return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
    return x * (1.0 - x)
}

func (nn *NeuralNetwork) Forward(input []float64) ([]float64, []float64) {
    // Вход -> скрытый слой
    hidden := make([]float64, nn.hiddenSize)
    for i := 0; i < nn.hiddenSize; i++ {
        sum := nn.biasH[i]
        for j := 0; j < nn.inputSize; j++ {
            sum += input[j] * nn.weightsIH[i][j]
        }
        hidden[i] = sigmoid(sum)
    }
    
    // Скрытый -> выходной слой
    output := make([]float64, nn.outputSize)
    for i := 0; i < nn.outputSize; i++ {
        sum := nn.biasO[i]
        for j := 0; j < nn.hiddenSize; j++ {
            sum += hidden[j] * nn.weightsHO[i][j]
        }
        output[i] = sigmoid(sum)
    }
    
    return output, hidden
}

func (nn *NeuralNetwork) Train(input, target []float64) {
    // Прямое распространение
    output, hidden := nn.Forward(input)
    
    // Ошибка выходного слоя
    outputErrors := make([]float64, nn.outputSize)
    outputDeltas := make([]float64, nn.outputSize)
    for i := 0; i < nn.outputSize; i++ {
        outputErrors[i] = target[i] - output[i]
        outputDeltas[i] = outputErrors[i] * sigmoidDerivative(output[i])
    }
    
    // Ошибка скрытого слоя
    hiddenErrors := make([]float64, nn.hiddenSize)
    hiddenDeltas := make([]float64, nn.hiddenSize)
    for i := 0; i < nn.hiddenSize; i++ {
        hiddenErrors[i] = 0.0
        for j := 0; j < nn.outputSize; j++ {
            hiddenErrors[i] += outputDeltas[j] * nn.weightsHO[j][i]
        }
        hiddenDeltas[i] = hiddenErrors[i] * sigmoidDerivative(hidden[i])
    }
    
    // Обновление весов выходного слоя
    for i := 0; i < nn.outputSize; i++ {
        for j := 0; j < nn.hiddenSize; j++ {
            nn.weightsHO[i][j] += nn.learningRate * outputDeltas[i] * hidden[j]
        }
        nn.biasO[i] += nn.learningRate * outputDeltas[i]
    }
    
    // Обновление весов скрытого слоя
    for i := 0; i < nn.hiddenSize; i++ {
        for j := 0; j < nn.inputSize; j++ {
            nn.weightsIH[i][j] += nn.learningRate * hiddenDeltas[i] * input[j]
        }
        nn.biasH[i] += nn.learningRate * hiddenDeltas[i]
    }
}

func (nn *NeuralNetwork) Predict(input []float64) []float64 {
    output, _ := nn.Forward(input)
    return output
}

func main() {
    rand.Seed(42)
    
    // Создание нейронной сети для задачи XOR
    nn := NewNeuralNetwork(2, 4, 1, 0.5)
    
    // Данные для обучения (XOR)
    trainingData := [][][]float64{
        {{0, 0}, {0}},
        {{0, 1}, {1}},
        {{1, 0}, {1}},
        {{1, 1}, {0}},
    }
    
    // Обучение
    fmt.Println("Training neural network...")
    for epoch := 0; epoch < 10000; epoch++ {
        for _, data := range trainingData {
            nn.Train(data[0], data[1])
        }
        
        if epoch%1000 == 0 {
            totalError := 0.0
            for _, data := range trainingData {
                output := nn.Predict(data[0])
                error := math.Abs(data[1][0] - output[0])
                totalError += error
            }
            fmt.Printf("Epoch %d, Error: %.4f\n", epoch, totalError)
        }
    }
    
    // Тестирование
    fmt.Println("\nTesting neural network:")
    for _, data := range trainingData {
        output := nn.Predict(data[0])
        fmt.Printf("Input: %v, Target: %v, Output: [%.3f]\n", 
            data[0], data[1], output[0])
    }
}
//Задание: Простая реализация нейронной сети с обратным распространением ошибки
```
128. Blockchain simulation
```go
package main
import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Hash      string
    Nonce     int
}

type Blockchain struct {
    blocks []*Block
}

func NewBlockchain() *Blockchain {
    genesisBlock := createGenesisBlock()
    return &Blockchain{
        blocks: []*Block{genesisBlock},
    }
}

func createGenesisBlock() *Block {
    return &Block{
        Index:     0,
        Timestamp: time.Now().String(),
        Data:      "Genesis Block",
        PrevHash:  "0",
        Hash:      "",
        Nonce:     0,
    }
}

func (b *Block) calculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func (b *Block) mineBlock(difficulty int) {
    target := ""
    for i := 0; i < difficulty; i++ {
        target += "0"
    }
    
    for {
        b.Hash = b.calculateHash()
        if b.Hash[:difficulty] == target {
            break
        }
        b.Nonce++
    }
    
    fmt.Printf("Block mined: %s\n", b.Hash)
}

func (bc *Blockchain) addBlock(data string, difficulty int) {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := &Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevBlock.Hash,
        Nonce:     0,
    }
    
    newBlock.mineBlock(difficulty)
    bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) isValid() bool {
    for i := 1; i < len(bc.blocks); i++ {
        currentBlock := bc.blocks[i]
        prevBlock := bc.blocks[i-1]
        
        // Проверка хеша текущего блока
        if currentBlock.Hash != currentBlock.calculateHash() {
            return false
        }
        
        // Проверка связи с предыдущим блоком
        if currentBlock.PrevHash != prevBlock.Hash {
            return false
        }
    }
    return true
}

func (bc *Blockchain) printBlocks() {
    for _, block := range bc.blocks {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Println("------------------------")
    }
}

func main() {
    // Создание блокчейна
    blockchain := NewBlockchain()
    
    // Майнинг новых блоков
    fmt.Println("Mining block 1...")
    blockchain.addBlock("First block data", 2)
    
    fmt.Println("Mining block 2...")
    blockchain.addBlock("Second block data", 2)
    
    fmt.Println("Mining block 3...")
    blockchain.addBlock("Third block data", 2)
    
    // Вывод всех блоков
    fmt.Println("\nBlockchain:")
    blockchain.printBlocks()
    
    // Проверка валидности блокчейна
    fmt.Printf("Blockchain valid: %t\n", blockchain.isValid())
    
    // Попытка подделки блока
    fmt.Println("\nAttempting to tamper with block 2...")
    blockchain.blocks[1].Data = "Tampered data"
    
    fmt.Printf("Blockchain valid after tampering: %t\n", blockchain.isValid())
}
//Задание: Симуляция простого блокчейна с proof-of-work
```
129. Game server (WebSocket)
```go
package main
import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
    
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

type Player struct {
    ID        string  `json:"id"`
    X         float64 `json:"x"`
    Y         float64 `json:"y"`
    Color     string  `json:"color"`
    LastSeen  time.Time
    conn      *websocket.Conn
}

type GameState struct {
    Players map[string]*Player `json:"players"`
}

type GameServer struct {
    mu      sync.RWMutex
    players map[string]*Player
    games   map[string]*GameState
}

func NewGameServer() *GameServer {
    return &GameServer{
        players: make(map[string]*Player),
        games:   make(map[string]*GameState),
    }
}

func (gs *GameServer) generatePlayerID() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, 8)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func (gs *GameServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    
    // Создание нового игрока
    player := &Player{
        ID:       gs.generatePlayerID(),
        X:        rand.Float64() * 800,
        Y:        rand.Float64() * 600,
        Color:    fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF)),
        LastSeen: time.Now(),
        conn:     conn,
    }
    
    gs.mu.Lock()
    gs.players[player.ID] = player
    gs.mu.Unlock()
    
    log.Printf("Player %s connected", player.ID)
    
    // Отправка начального состояния
    gs.sendGameState(player)
    
    // Обработка сообщений от клиента
    go gs.handlePlayerMessages(player)
    
    // Оповещение других игроков о новом игроке
    gs.broadcastPlayerJoined(player)
}

func (gs *GameServer) handlePlayerMessages(player *Player) {
    defer func() {
        gs.removePlayer(player.ID)
        player.conn.Close()
    }()
    
    for {
        _, message, err := player.conn.ReadMessage()
        if err != nil {
            log.Printf("Read error for player %s: %v", player.ID, err)
            break
        }
        
        var msg map[string]interface{}
        if err := json.Unmarshal(message, &msg); err != nil {
            log.Printf("Error parsing message: %v", err)
            continue
        }
        
        switch msg["type"] {
        case "move":
            gs.handlePlayerMove(player, msg)
        case "chat":
            gs.handleChatMessage(player, msg)
        }
        
        player.LastSeen = time.Now()
    }
}

func (gs *GameServer) handlePlayerMove(player *Player, msg map[string]interface{}) {
    if x, ok := msg["x"].(float64); ok {
        player.X = x
    }
    if y, ok := msg["y"].(float64); ok {
        player.Y = y
    }
    
    // Рассылка обновления позиции всем игрокам
    gs.broadcastPlayerUpdate(player)
}

func (gs *GameServer) handleChatMessage(player *Player, msg map[string]interface{}) {
    text, ok := msg["text"].(string)
    if !ok {
        return
    }
    
    chatMsg := map[string]interface{}{
        "type":   "chat",
        "player": player.ID,
        "text":   text,
        "color":  player.Color,
    }
    
    gs.broadcastToAll(chatMsg)
}

func (gs *GameServer) sendGameState(player *Player) {
    gs.mu.RLock()
    defer gs.mu.RUnlock()
    
    players := make([]*Player, 0, len(gs.players))
    for _, p := range gs.players {
        if p.ID != player.ID {
            players = append(players, p)
        }
    }
    
    state := map[string]interface{}{
        "type":    "game_state",
        "players": players,
        "you":     player,
    }
    
    player.conn.WriteJSON(state)
}

func (gs *GameServer) broadcastPlayerUpdate(player *Player) {
    update := map[string]interface{}{
        "type": "player_update",
        "player": map[string]interface{}{
            "id": player.ID,
            "x":  player.X,
            "y":  player.Y,
        },
    }
    
    gs.broadcastToOthers(player.ID, update)
}

func (gs *GameServer) broadcastPlayerJoined(player *Player) {
    joinMsg := map[string]interface{}{
        "type": "player_joined",
        "player": player,
    }
    
    gs.broadcastToOthers(player.ID, joinMsg)
}

func (gs *GameServer) broadcastToAll(message interface{}) {
    gs.mu.RLock()
    defer gs.mu.RUnlock()
    
    for _, player := range gs.players {
        player.conn.WriteJSON(message)
    }
}

func (gs *GameServer) broadcastToOthers(excludeID string, message interface{}) {
    gs.mu.RLock()
    defer gs.mu.RUnlock()
    
    for _, player := range gs.players {
        if player.ID != excludeID {
            player.conn.WriteJSON(message)
        }
    }
}

func (gs *GameServer) removePlayer(playerID string) {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    
    delete(gs.players, playerID)
    
    // Оповещение об уходе игрока
    leaveMsg := map[string]interface{}{
        "type": "player_left",
        "id":   playerID,
    }
    
    for _, player := range gs.players {
        player.conn.WriteJSON(leaveMsg)
    }
    
    log.Printf("Player %s disconnected", playerID)
}

func (gs *GameServer) startCleanupRoutine() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        gs.mu.Lock()
        now := time.Now()
        for id, player := range gs.players {
            if now.Sub(player.LastSeen) > time.Minute {
                log.Printf("Removing inactive player %s", id)
                delete(gs.players, id)
                player.conn.Close()
            }
        }
        gs.mu.Unlock()
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    gameServer := NewGameServer()
    
    // Запуск очистки неактивных игроков
    go gameServer.startCleanupRoutine()
    
    http.HandleFunc("/ws", gameServer.handleWebSocket)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
    log.Println("Game server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Многопользовательский игровой сервер на WebSocket
```
130. Image processing
```go
package main
import (
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"
    "math"
    "os"
)

type ImageProcessor struct{}

func NewImageProcessor() *ImageProcessor {
    return &ImageProcessor{}
}

func (ip *ImageProcessor) LoadImage(filename string) (image.Image, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    
    return img, nil
}

func (ip *ImageProcessor) SaveImage(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    switch {
    case len(filename) > 4 && filename[len(filename)-4:] == ".png":
        return png.Encode(file, img)
    case len(filename) > 4 && filename[len(filename)-4:] == ".jpg":
        return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
    default:
        return png.Encode(file, img)
    }
}

func (ip *ImageProcessor) Grayscale(img image.Image) image.Image {
    bounds := img.Bounds()
    gray := image.NewGray(bounds)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            originalColor := img.At(x, y)
            grayColor := color.GrayModel.Convert(originalColor)
            gray.Set(x, y, grayColor)
        }
    }
    
    return gray
}

func (ip *ImageProcessor) Resize(img image.Image, newWidth, newHeight int) image.Image {
    bounds := img.Bounds()
    oldWidth := bounds.Dx()
    oldHeight := bounds.Dy()
    
    resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
    
    xRatio := float64(oldWidth) / float64(newWidth)
    yRatio := float64(oldHeight) / float64(newHeight)
    
    for y := 0; y < newHeight; y++ {
        for x := 0; x < newWidth; x++ {
            srcX := int(float64(x) * xRatio)
            srcY := int(float64(y) * yRatio)
            
            if srcX >= oldWidth {
                srcX = oldWidth - 1
            }
            if srcY >= oldHeight {
                srcY = oldHeight - 1
            }
            
            resized.Set(x, y, img.At(srcX, srcY))
        }
    }
    
    return resized
}

func (ip *ImageProcessor) Blur(img image.Image, radius int) image.Image {
    bounds := img.Bounds()
    blurred := image.NewRGBA(bounds)
    
    kernel := ip.createGaussianKernel(radius)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            var r, g, b, a float64
            weightSum := 0.0
            
            for ky := -radius; ky <= radius; ky++ {
                for kx := -radius; kx <= radius; kx++ {
                    nx := x + kx
                    ny := y + ky
                    
                    if nx >= bounds.Min.X && nx < bounds.Max.X && 
                       ny >= bounds.Min.Y && ny < bounds.Max.Y {
                        weight := kernel[ky+radius][kx+radius]
                        pixel := img.At(nx, ny)
                        pr, pg, pb, pa := pixel.RGBA()
                        
                        r += float64(pr>>8) * weight
                        g += float64(pg>>8) * weight
                        b += float64(pb>>8) * weight
                        a += float64(pa>>8) * weight
                        weightSum += weight
                    }
                }
            }
            
            if weightSum > 0 {
                r /= weightSum
                g /= weightSum
                b /= weightSum
                a /= weightSum
            }
            
            blurred.Set(x, y, color.RGBA{
                R: uint8(r),
                G: uint8(g),
                B: uint8(b),
                A: uint8(a),
            })
        }
    }
    
    return blurred
}

func (ip *ImageProcessor) createGaussianKernel(radius int) [][]float64 {
    size := 2*radius + 1
    kernel := make([][]float64, size)
    sigma := float64(radius) / 2.0
    sum := 0.0
    
    for y := -radius; y <= radius; y++ {
        kernel[y+radius] = make([]float64, size)
        for x := -radius; x <= radius; x++ {
            exponent := -(float64(x*x + y*y) / (2 * sigma * sigma))
            kernel[y+radius][x+radius] = math.Exp(exponent)
            sum += kernel[y+radius][x+radius]
        }
    }
    
    // Нормализация
    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            kernel[y][x] /= sum
        }
    }
    
    return kernel
}

func (ip *ImageProcessor) DrawText(img image.Image, text string, x, y int, fontColor color.Color) image.Image {
    bounds := img.Bounds()
    result := image.NewRGBA(bounds)
    draw.Draw(result, bounds, img, bounds.Min, draw.Src)
    
    // Простая реализация рисования текста (в реальном приложении используйте библиотеку)
    for i, char := range text {
        ip.drawChar(result, char, x+i*8, y, fontColor)
    }
    
    return result
}

func (ip *ImageProcessor) drawChar(img *image.RGBA, char rune, x, y int, fontColor color.Color) {
    // Простой 8x8 шрифт для демонстрации
    font := map[rune][]byte{
        'H': {0x81, 0x81, 0x81, 0xFF, 0x81, 0x81, 0x81, 0x00},
        'e': {0x00, 0x00, 0x7C, 0x82, 0xFE, 0x80, 0x7C, 0x00},
        'l': {0x00, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x00},
        'o': {0x00, 0x00, 0x7C, 0x82, 0x82, 0x82, 0x7C, 0x00},
        'W': {0x41, 0x41, 0x41, 0x49, 0x49, 0x55, 0x63, 0x00},
        'r': {0x00, 0x00, 0xBC, 0xC2, 0x80, 0x80, 0x80, 0x00},
        'd': {0x02, 0x02, 0x7A, 0x86, 0x82, 0x82, 0x7E, 0x00},
        '!': {0x80, 0x80, 0x80, 0x80, 0x80, 0x00, 0x80, 0x00},
    }
    
    if pixels, exists := font[char]; exists {
        for py := 0; py < 8; py++ {
            for px := 0; px < 8; px++ {
                if (pixels[py]>>(7-px))&1 == 1 {
                    img.Set(x+px, y+py, fontColor)
                }
            }
        }
    }
}

func main() {
    processor := NewImageProcessor()
    
    // Создание тестового изображения
    img := image.NewRGBA(image.Rect(0, 0, 200, 200))
    
    // Заливка градиентом
    for y := 0; y < 200; y++ {
        for x := 0; x < 200; x++ {
            r := uint8(x)
            g := uint8(y)
            b := uint8((x + y) / 2)
            img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
        }
    }
    
    // Добавление текста
    imgWithText := processor.DrawText(img, "Hello World!", 50, 100, color.White)
    
    // Сохранение оригинального изображения
    processor.SaveImage(imgWithText, "original.png")
    
    // Применение фильтров
    gray := processor.Grayscale(imgWithText)
    processor.SaveImage(gray, "grayscale.png")
    
    resized := processor.Resize(imgWithText, 100, 100)
    processor.SaveImage(resized, "resized.png")
    
    blurred := processor.Blur(imgWithText, 2)
    processor.SaveImage(blurred, "blurred.png")
    
    fmt.Println("Image processing completed!")
}
//Задание: Обработка изображений: фильтры, преобразования, рисование
```
Криптография и безопасность
131. AES шифрование
```go
package main
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

type AESEncryptor struct {
    key []byte
}

func NewAESEncryptor(key string) *AESEncryptor {
    // Ключ должен быть 16, 24 или 32 байта
    keyBytes := []byte(key)
    if len(keyBytes) < 16 {
        padded := make([]byte, 16)
        copy(padded, keyBytes)
        keyBytes = padded
    } else if len(keyBytes) > 32 {
        keyBytes = keyBytes[:32]
    }
    
    return &AESEncryptor{key: keyBytes}
}

func (ae *AESEncryptor) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(ae.key)
    if err != nil {
        return "", err
    }
    
    // Создаем GCM режим
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    // Создаем nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    
    // Шифруем
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (ae *AESEncryptor) Decrypt(encrypted string) (string, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return "", err
    }
    
    block, err := aes.NewCipher(ae.key)
    if err != nil {
        return "", err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return "", fmt.Errorf("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    
    return string(plaintext), nil
}

func main() {
    encryptor := NewAESEncryptor("my-secret-key-12345")
    
    originalText := "Hello, World! This is a secret message."
    fmt.Printf("Original: %s\n", originalText)
    
    // Шифрование
    encrypted, err := encryptor.Encrypt(originalText)
    if err != nil {
        fmt.Printf("Encryption error: %v\n", err)
        return
    }
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Дешифрование
    decrypted, err := encryptor.Decrypt(encrypted)
    if err != nil {
        fmt.Printf("Decryption error: %v\n", err)
        return
    }
    fmt.Printf("Decrypted: %s\n", decrypted)
    
    // Проверка
    if originalText == decrypted {
        fmt.Println("✅ Encryption/Decryption successful!")
    } else {
        fmt.Println("❌ Encryption/Decryption failed!")
    }
}
//Задание: AES шифрование с использованием GCM режима
```
132. RSA шифрование и подписи
```go
package main
import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "fmt"
    "os"
)

type RSAHelper struct {
    privateKey *rsa.PrivateKey
    publicKey  *rsa.PublicKey
}

func NewRSAHelper(bits int) (*RSAHelper, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
    }
    
    return &RSAHelper{
        privateKey: privateKey,
        publicKey:  &privateKey.PublicKey,
    }, nil
}

func (rh *RSAHelper) Encrypt(plaintext []byte) ([]byte, error) {
    return rsa.EncryptOAEP(
        sha256.New(),
        rand.Reader,
        rh.publicKey,
        plaintext,
        nil,
    )
}

func (rh *RSAHelper) Decrypt(ciphertext []byte) ([]byte, error) {
    return rsa.DecryptOAEP(
        sha256.New(),
        rand.Reader,
        rh.privateKey,
        ciphertext,
        nil,
    )
}

func (rh *RSAHelper) Sign(data []byte) ([]byte, error) {
    hashed := sha256.Sum256(data)
    return rsa.SignPSS(
        rand.Reader,
        rh.privateKey,
        crypto.SHA256,
        hashed[:],
        nil,
    )
}

func (rh *RSAHelper) Verify(data []byte, signature []byte) error {
    hashed := sha256.Sum256(data)
    return rsa.VerifyPSS(
        rh.publicKey,
        crypto.SHA256,
        hashed[:],
        signature,
        nil,
    )
}

func (rh *RSAHelper) SavePrivateKey(filename string) error {
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(rh.privateKey)
    privateKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privateKeyBytes,
    })
    return os.WriteFile(filename, privateKeyPEM, 0600)
}

func (rh *RSAHelper) SavePublicKey(filename string) error {
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(rh.publicKey)
    if err != nil {
        return err
    }
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: publicKeyBytes,
    })
    return os.WriteFile(filename, publicKeyPEM, 0644)
}

func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    block, _ := pem.Decode(data)
    if block == nil {
        return nil, fmt.Errorf("failed to parse PEM block")
    }
    
    return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    block, _ := pem.Decode(data)
    if block == nil {
        return nil, fmt.Errorf("failed to parse PEM block")
    }
    
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    
    return pub.(*rsa.PublicKey), nil
}

func main() {
    // Генерация ключей
    rsaHelper, err := NewRSAHelper(2048)
    if err != nil {
        fmt.Printf("Error generating keys: %v\n", err)
        return
    }
    
    // Сохранение ключей
    rsaHelper.SavePrivateKey("private.pem")
    rsaHelper.SavePublicKey("public.pem")
    fmt.Println("✅ RSA keys generated and saved")
    
    // Шифрование и дешифрование
    message := "Hello, RSA Encryption!"
    fmt.Printf("Original message: %s\n", message)
    
    encrypted, err := rsaHelper.Encrypt([]byte(message))
    if err != nil {
        fmt.Printf("Encryption error: %v\n", err)
        return
    }
    fmt.Printf("Encrypted (base64): %s\n", base64.StdEncoding.EncodeToString(encrypted))
    
    decrypted, err := rsaHelper.Decrypt(encrypted)
    if err != nil {
        fmt.Printf("Decryption error: %v\n", err)
        return
    }
    fmt.Printf("Decrypted: %s\n", string(decrypted))
    
    // Цифровая подпись
    dataToSign := []byte("Important document data")
    signature, err := rsaHelper.Sign(dataToSign)
    if err != nil {
        fmt.Printf("Signing error: %v\n", err)
        return
    }
    fmt.Printf("Signature (base64): %s\n", base64.StdEncoding.EncodeToString(signature))
    
    // Проверка подписи
    err = rsaHelper.Verify(dataToSign, signature)
    if err != nil {
        fmt.Printf("Signature verification failed: %v\n", err)
    } else {
        fmt.Println("✅ Signature verified successfully!")
    }
    
    // Загрузка ключей из файлов
    loadedPrivateKey, err := LoadPrivateKey("private.pem")
    if err != nil {
        fmt.Printf("Error loading private key: %v\n", err)
        return
    }
    
    loadedPublicKey, err := LoadPublicKey("public.pem")
    if err != nil {
        fmt.Printf("Error loading public key: %v\n", err)
        return
    }
    
    fmt.Printf("✅ Keys loaded successfully: private=%t, public=%t\n", 
        loadedPrivateKey != nil, loadedPublicKey != nil)
}
//Задание: RSA шифрование, цифровые подписи и работа с PEM файлами
```
133. Password hashing с argon2
```go
package main
import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "fmt"
    "strings"
    
    "golang.org/x/crypto/argon2"
)

type Argon2Params struct {
    memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

type PasswordHasher struct {
    params *Argon2Params
}

func NewPasswordHasher() *PasswordHasher {
    return &PasswordHasher{
        params: &Argon2Params{
            memory:      64 * 1024, // 64 MB
            iterations:  3,
            parallelism: 2,
            saltLength:  16,
            keyLength:   32,
        },
    }
}

func (ph *PasswordHasher) GenerateHash(password string) (string, error) {
    // Генерация соли
    salt := make([]byte, ph.params.saltLength)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    // Хеширование пароля
    hash := argon2.IDKey(
        []byte(password),
        salt,
        ph.params.iterations,
        ph.params.memory,
        ph.params.parallelism,
        ph.params.keyLength,
    )
    
    // Кодирование в строку
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)
    
    // Формат: argon2id$v=19$m=65536,t=3,p=2$salt$hash
    encodedHash := fmt.Sprintf("argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
        ph.params.memory,
        ph.params.iterations,
        ph.params.parallelism,
        b64Salt,
        b64Hash,
    )
    
    return encodedHash, nil
}

func (ph *PasswordHasher) VerifyPassword(password, encodedHash string) (bool, error) {
    // Парсинг encodedHash
    parts := strings.Split(encodedHash, "$")
    if len(parts) != 6 {
        return false, fmt.Errorf("invalid hash format")
    }
    
    if parts[1] != "argon2id" {
        return false, fmt.Errorf("unsupported algorithm")
    }
    
    // Парсинг параметров
    var version int
    _, err := fmt.Sscanf(parts[2], "v=%d", &version)
    if err != nil {
        return false, err
    }
    if version != 19 {
        return false, fmt.Errorf("unsupported version")
    }
    
    var memory, iterations uint32
    var parallelism uint8
    _, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
    if err != nil {
        return false, err
    }
    
    // Декодирование соли и хеша
    salt, err := base64.RawStdEncoding.DecodeString(parts[4])
    if err != nil {
        return false, err
    }
    
    expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
    if err != nil {
        return false, err
    }
    
    // Вычисление хеша для проверяемого пароля
    actualHash := argon2.IDKey(
        []byte(password),
        salt,
        iterations,
        memory,
        parallelism,
        uint32(len(expectedHash)),
    )
    
    // Сравнение с постоянным временем выполнения
    return subtle.ConstantTimeCompare(actualHash, expectedHash) == 1, nil
}

func main() {
    hasher := NewPasswordHasher()
    
    // Хеширование паролей
    passwords := []string{
        "mySecurePassword123",
        "anotherPassword!",
        "test123",
    }
    
    var hashes []string
    
    for i, password := range passwords {
        hash, err := hasher.GenerateHash(password)
        if err != nil {
            fmt.Printf("Error hashing password %d: %v\n", i, err)
            continue
        }
        hashes = append(hashes, hash)
        fmt.Printf("Password %d: %s\n", i+1, password)
        fmt.Printf("Hash: %s\n\n", hash)
    }
    
    // Проверка паролей
    fmt.Println("Verifying passwords:")
    for i, password := range passwords {
        if i >= len(hashes) {
            break
        }
        
        valid, err := hasher.VerifyPassword(password, hashes[i])
        if err != nil {
            fmt.Printf("Error verifying password %d: %v\n", i, err)
            continue
        }
        
        if valid {
            fmt.Printf("✅ Password %d: CORRECT\n", i+1)
        } else {
            fmt.Printf("❌ Password %d: INCORRECT\n", i+1)
        }
        
        // Проверка с неправильным паролем
        wrongValid, _ := hasher.VerifyPassword("wrongpassword", hashes[i])
        if !wrongValid {
            fmt.Printf("✅ Wrong password correctly rejected\n")
        }
    }
    
    // Бенчмарк производительности
    fmt.Println("\nPerformance test:")
    testPassword := "benchmarkPassword"
    
    start := time.Now()
    hash, _ := hasher.GenerateHash(testPassword)
    elapsed := time.Since(start)
    
    fmt.Printf("Time to hash: %v\n", elapsed)
    fmt.Printf("Hash length: %d bytes\n", len(hash))
}
//Задание: Безопасное хеширование паролей с Argon2
```
134. TLS сервер и клиент
```go
package main
import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "os"
)

// Генерация сертификатов (для тестирования):
// openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes

func startTLSServer() {
    // Загрузка сертификата
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Fatalf("Error loading certificate: %v", err)
    }
    
    // Настройка TLS конфигурации
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        MinVersion:   tls.VersionTLS12,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }
    
    // Создание listener'а
    listener, err := tls.Listen("tcp", ":8443", tlsConfig)
    if err != nil {
        log.Fatalf("Error creating listener: %v", err)
    }
    defer listener.Close()
    
    fmt.Println("TLS Server listening on :8443")
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %v", err)
            continue
        }
        
        go handleTLSConnection(conn)
    }
}

func handleTLSConnection(conn net.Conn) {
    defer conn.Close()
    
    tlsConn, ok := conn.(*tls.Conn)
    if !ok {
        log.Println("Not a TLS connection")
        return
    }
    
    // Handshake
    if err := tlsConn.Handshake(); err != nil {
        log.Printf("TLS handshake error: %v", err)
        return
    }
    
    // Информация о соединении
    state := tlsConn.ConnectionState()
    fmt.Printf("TLS Version: %s\n", tlsVersionToString(state.Version))
    fmt.Printf("Cipher Suite: %s\n", tls.CipherSuiteName(state.CipherSuite))
    fmt.Printf("Server Name: %s\n", state.ServerName)
    
    // Обработка данных
    buf := make([]byte, 1024)
    n, err := tlsConn.Read(buf)
    if err != nil {
        if err != io.EOF {
            log.Printf("Read error: %v", err)
        }
        return
    }
    
    message := string(buf[:n])
    fmt.Printf("Received: %s\n", message)
    
    // Ответ
    response := fmt.Sprintf("Hello from TLS Server! Your message: %s", message)
    tlsConn.Write([]byte(response))
}

func tlsVersionToString(version uint16) string {
    switch version {
    case tls.VersionTLS10:
        return "TLS 1.0"
    case tls.VersionTLS11:
        return "TLS 1.1"
    case tls.VersionTLS12:
        return "TLS 1.2"
    case tls.VersionTLS13:
        return "TLS 1.3"
    default:
        return "Unknown"
    }
}

func startTLSClient() {
    // Загрузка корневого сертификата
    caCert, err := os.ReadFile("server.crt")
    if err != nil {
        log.Fatalf("Error reading CA certificate: %v", err)
    }
    
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)
    
    // Настройка TLS конфигурации клиента
    tlsConfig := &tls.Config{
        RootCAs:    caCertPool,
        ServerName: "localhost",
    }
    
    // Подключение к серверу
    conn, err := tls.Dial("tcp", "localhost:8443", tlsConfig)
    if err != nil {
        log.Fatalf("Error connecting to server: %v", err)
    }
    defer conn.Close()
    
    // Отправка сообщения
    message := "Hello from TLS Client!"
    _, err = conn.Write([]byte(message))
    if err != nil {
        log.Fatalf("Error writing: %v", err)
    }
    
    // Чтение ответа
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        log.Fatalf("Error reading: %v", err)
    }
    
    fmt.Printf("Server response: %s\n", string(buf[:n]))
}

func startHTTPServer() {
    // HTTP сервер с TLS
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from HTTPS Server!"))
    })
    
    http.HandleFunc("/tls-info", func(w http.ResponseWriter, r *http.Request) {
        if r.TLS != nil {
            info := fmt.Sprintf(
                "TLS Version: %s\nCipher Suite: %s\nServer Name: %s\n",
                tlsVersionToString(r.TLS.Version),
                tls.CipherSuiteName(r.TLS.CipherSuite),
                r.TLS.ServerName,
            )
            w.Write([]byte(info))
        } else {
            w.Write([]byte("Not a TLS connection"))
        }
    })
    
    fmt.Println("HTTPS Server listening on :8444")
    err := http.ListenAndServeTLS(":8444", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatalf("Error starting HTTPS server: %v", err)
    }
}

func main() {
    // Запуск в разных горутинах для демонстрации
    go startTLSServer()
    go startHTTPServer()
    
    // Даем серверам время запуститься
    time.Sleep(2 * time.Second)
    
    // Тестирование TLS клиента
    fmt.Println("Testing TLS client...")
    startTLSClient()
    
    // Тестирование HTTPS клиента
    fmt.Println("\nTesting HTTPS client...")
    testHTTPSClient()
    
    // Ожидание
    select {}
}

func testHTTPSClient() {
    // Настройка HTTPS клиента
    caCert, err := os.ReadFile("server.crt")
    if err != nil {
        log.Fatalf("Error reading CA certificate: %v", err)
    }
    
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)
    
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs: caCertPool,
            },
        },
    }
    
    resp, err := client.Get("https://localhost:8444/tls-info")
    if err != nil {
        log.Fatalf("Error making request: %v", err)
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response: %v", err)
    }
    
    fmt.Printf("HTTPS Response:\n%s\n", string(body))
}
//Задание: TLS сервер и клиент с взаимной аутентификацией
```
Системное программирование
135. Работа с процессами
```go
package main
import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strconv"
    "syscall"
    "time"
)

type ProcessManager struct{}

func NewProcessManager() *ProcessManager {
    return &ProcessManager{}
}

func (pm *ProcessManager) GetProcessInfo() {
    // Информация о текущем процессе
    pid := os.Getpid()
    ppid := os.Getppid()
    
    fmt.Printf("Process ID: %d\n", pid)
    fmt.Printf("Parent Process ID: %d\n", ppid)
    fmt.Printf("Working Directory: %s\n", pm.getWorkingDir())
    fmt.Printf("Environment Variables: %d\n", len(os.Environ()))
    
    // Информация о системе
    fmt.Printf("GOOS: %s\n", runtime.GOOS)
    fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
    fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

func (pm *ProcessManager) getWorkingDir() string {
    wd, err := os.Getwd()
    if err != nil {
        return "unknown"
    }
    return wd
}

func (pm *ProcessManager) ExecuteCommand(name string, args ...string) error {
    cmd := exec.Command(name, args...)
    
    // Настройка ввода/вывода
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    
    fmt.Printf("Executing: %s %v\n", name, args)
    
    // Запуск команды
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("command failed: %v", err)
    }
    
    return nil
}

func (pm *ProcessManager) ExecuteCommandWithOutput(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("command failed: %v", err)
    }
    
    return string(output), nil
}

func (pm *ProcessManager) StartBackgroundProcess(name string, args ...string) (*exec.Cmd, error) {
    cmd := exec.Command(name, args...)
    
    // Запуск в фоновом режиме
    if err := cmd.Start(); err != nil {
        return nil, err
    }
    
    fmt.Printf("Started background process PID: %d\n", cmd.Process.Pid)
    return cmd, nil
}

func (pm *ProcessManager) KillProcess(pid int) error {
    process, err := os.FindProcess(pid)
    if err != nil {
        return err
    }
    
    // Отправка сигнала SIGTERM
    err = process.Signal(syscall.SIGTERM)
    if err != nil {
        // Если SIGTERM не сработал, отправляем SIGKILL
        err = process.Signal(syscall.SIGKILL)
        if err != nil {
            return err
        }
    }
    
    fmt.Printf("Sent termination signal to PID: %d\n", pid)
    return nil
}

func (pm *ProcessManager) MonitorProcess(pid int, timeout time.Duration) error {
    process, err := os.FindProcess(pid)
    if err != nil {
        return err
    }
    
    // Канал для ожидания завершения
    done := make(chan error, 1)
    go func() {
        _, err := process.Wait()
        done <- err
    }()
    
    // Таймаут для ожидания
    select {
    case err := <-done:
        if err != nil {
            return fmt.Errorf("process failed: %v", err)
        }
        fmt.Printf("Process %d completed successfully\n", pid)
        return nil
    case <-time.After(timeout):
        return fmt.Errorf("process %d timed out after %v", pid, timeout)
    }
}

func (pm *ProcessManager) CreateChildProcess() (int, error) {
    // Аргументы для нового процесса
    args := []string{
        "child",
        strconv.Itoa(os.Getpid()),
    }
    
    // Создание команды
    cmd := exec.Command(os.Args[0], args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    // Запуск
    if err := cmd.Start(); err != nil {
        return 0, err
    }
    
    return cmd.Process.Pid, nil
}

func childProcess() {
    if len(os.Args) < 3 {
        fmt.Println("This is a child process")
        return
    }
    
    if os.Args[1] == "child" {
        parentPid, _ := strconv.Atoi(os.Args[2])
        fmt.Printf("Child process PID: %d, Parent PID: %d\n", 
            os.Getpid(), parentPid)
        
        // Демонстрационная работа
        time.Sleep(5 * time.Second)
        fmt.Println("Child process completed")
        os.Exit(0)
    }
}

func main() {
    // Проверка, не является ли процесс дочерним
    if len(os.Args) > 1 && os.Args[1] == "child" {
        childProcess()
        return
    }
    
    pm := NewProcessManager()
    
    fmt.Println("=== Process Information ===")
    pm.GetProcessInfo()
    
    fmt.Println("\n=== Executing Commands ===")
    
    // Выполнение простой команды
    if runtime.GOOS == "windows" {
        pm.ExecuteCommand("cmd", "/c", "echo", "Hello from Windows!")
    } else {
        pm.ExecuteCommand("echo", "Hello from Unix!")
    }
    
    // Получение вывода команды
    output, err := pm.ExecuteCommandWithOutput("go", "version")
    if err == nil {
        fmt.Printf("Go version: %s", output)
    }
    
    fmt.Println("\n=== Background Process ===")
    
    // Запуск фонового процесса
    if runtime.GOOS != "windows" {
        cmd, err := pm.StartBackgroundProcess("sleep", "10")
        if err == nil {
            // Мониторинг процесса
            go func() {
                time.Sleep(3 * time.Second)
                pm.KillProcess(cmd.Process.Pid)
            }()
            
            pm.MonitorProcess(cmd.Process.Pid, 15*time.Second)
        }
    }
    
    fmt.Println("\n=== Child Process ===")
    childPid, err := pm.CreateChildProcess()
    if err == nil {
        fmt.Printf("Created child process with PID: %d\n", childPid)
        pm.MonitorProcess(childPid, 10*time.Second)
    }
    
    fmt.Println("\n=== Process Signals ===")
    // Демонстрация обработки сигналов
    setupSignalHandlers()
    
    fmt.Println("Main process waiting... Press Ctrl+C to exit")
    select {}
}

func setupSignalHandlers() {
    // Канал для получения сигналов
    sigChan := make(chan os.Signal, 1)
    // signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        for {
            sig := <-sigChan
            fmt.Printf("\nReceived signal: %v\n", sig)
            
            switch sig {
            case syscall.SIGINT:
                fmt.Println("Handling SIGINT...")
                os.Exit(0)
            case syscall.SIGTERM:
                fmt.Println("Handling SIGTERM...")
                os.Exit(0)
            }
        }
    }()
}
//Задание: Управление процессами: запуск, мониторинг, сигналы
```
136. Работа с файловой системой (low-level)
```go
package main
import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "syscall"
    "time"
)

type FileSystemManager struct{}

func NewFileSystemManager() *FileSystemManager {
    return &FileSystemManager{}
}

func (fsm *FileSystemManager) GetFileInfo(path string) error {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return err
    }
    
    fmt.Printf("File: %s\n", path)
    fmt.Printf("Size: %d bytes\n", fileInfo.Size())
    fmt.Printf("Permissions: %s\n", fileInfo.Mode())
    fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())
    fmt.Printf("Is Directory: %t\n", fileInfo.IsDir())
    
    // Дополнительная системная информация
    if sysStat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
        fmt.Printf("Inode: %d\n", sysStat.Ino)
        fmt.Printf("Device: %d\n", sysStat.Dev)
        fmt.Printf("Links: %d\n", sysStat.Nlink)
        fmt.Printf("UID: %d\n", sysStat.Uid)
        fmt.Printf("GID: %d\n", sysStat.Gid)
    }
    
    return nil
}

func (fsm *FileSystemManager) CreateHardLink(src, dst string) error {
    return os.Link(src, dst)
}

func (fsm *FileSystemManager) CreateSymlink(src, dst string) error {
    return os.Symlink(src, dst)
}

func (fsm *FileSystemManager) ReadSymlink(path string) (string, error) {
    return os.Readlink(path)
}

func (fsm *FileSystemManager) ChangeOwner(path string, uid, gid int) error {
    return os.Chown(path, uid, gid)
}

func (fsm *FileSystemManager) ChangePermissions(path string, mode os.FileMode) error {
    return os.Chmod(path, mode)
}

func (fsm *FileSystemManager) SetFileTimes(path string, atime, mtime time.Time) error {
    return os.Chtimes(path, atime, mtime)
}

func (fsm *FileSystemManager) MemoryMapFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    fileInfo, err := file.Stat()
    if err != nil {
        return nil, err
    }
    
    size := fileInfo.Size()
    if size == 0 {
        return []byte{}, nil
    }
    
    // Для memory mapping в Go нужно использовать syscall.Mmap
    // Это упрощенная версия с чтением всего файла
    data := make([]byte, size)
    _, err = io.ReadFull(file, data)
    if err != nil {
        return nil, err
    }
    
    return data, nil
}

func (fsm *FileSystemManager) WatchFileChanges(path string) error {
    initialStat, err := os.Stat(path)
    if err != nil {
        return err
    }
    
    initialModTime := initialStat.ModTime()
    
    fmt.Printf("Watching file: %s\n", path)
    
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        stat, err := os.Stat(path)
        if err != nil {
            fmt.Printf("Error watching file: %v\n", err)
            continue
        }
        
        if stat.ModTime() != initialModTime {
            fmt.Printf("File changed at: %s\n", stat.ModTime())
            initialModTime = stat.ModTime()
        }
    }
    
    return nil
}

func (fsm *FileSystemManager) GetDiskUsage(path string) error {
    var totalSize int64
    var fileCount int
    
    err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if !info.IsDir() {
            totalSize += info.Size()
            fileCount++
        }
        
        return nil
    })
    
    if err != nil {
        return err
    }
    
    fmt.Printf("Path: %s\n", path)
    fmt.Printf("Total Size: %d bytes (%.2f MB)\n", totalSize, float64(totalSize)/(1024*1024))
    fmt.Printf("File Count: %d\n", fileCount)
    
    return nil
}

func (fsm *FileSystemManager) CreateSparseFile(filename string, size int64) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Установка размера файла
    err = file.Truncate(size)
    if err != nil {
        return err
    }
    
    fmt.Printf("Created sparse file: %s, size: %d bytes\n", filename, size)
    return nil
}

func (fsm *FileSystemManager) LockFile(path string) error {
    file, err := os.OpenFile(path, os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Попытка заблокировать файл
    err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
    if err != nil {
        return fmt.Errorf("file is locked: %v", err)
    }
    
    fmt.Printf("File locked: %s\n", path)
    
    // Держим блокировку некоторое время
    time.Sleep(5 * time.Second)
    
    // Разблокировка
    syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
    fmt.Printf("File unlocked: %s\n", path)
    
    return nil
}

func main() {
    fsm := NewFileSystemManager()
    
    // Создание тестового файла
    testFile := "test_file.txt"
    os.WriteFile(testFile, []byte("Hello, File System!"), 0644)
    
    fmt.Println("=== File Information ===")
    fsm.GetFileInfo(testFile)
    
    fmt.Println("\n=== File Links ===")
    hardLink := "test_file_hard.txt"
    fsm.CreateHardLink(testFile, hardLink)
    fmt.Printf("Created hard link: %s\n", hardLink)
    
    symLink := "test_file_sym.txt"
    fsm.CreateSymlink(testFile, symLink)
    fmt.Printf("Created symlink: %s\n", symLink)
    
    target, _ := fsm.ReadSymlink(symLink)
    fmt.Printf("Symlink target: %s\n", target)
    
    fmt.Println("\n=== File Operations ===")
    fsm.ChangePermissions(testFile, 0755)
    fmt.Printf("Changed permissions of %s\n", testFile)
    
    // Memory mapping (упрощенное)
    data, err := fsm.MemoryMapFile(testFile)
    if err == nil {
        fmt.Printf("File content: %s\n", string(data))
    }
    
    fmt.Println("\n=== Disk Usage ===")
    fsm.GetDiskUsage(".")
    
    fmt.Println("\n=== Sparse File ===")
    sparseFile := "sparse_file.bin"
    fsm.CreateSparseFile(sparseFile, 1024*1024) // 1MB
    
    fmt.Println("\n=== File Locking ===")
    go func() {
        time.Sleep(1 * time.Second)
        fsm.LockFile(testFile)
    }()
    
    // Ожидание завершения
    time.Sleep(10 * time.Second)
    
    // Очистка
    os.Remove(testFile)
    os.Remove(hardLink)
    os.Remove(symLink)
    os.Remove(sparseFile)
}
//Задание: Низкоуровневая работа с файловой системой
```
137. System calls и raw sockets
```go
package main
import (
    "encoding/binary"
    "fmt"
    "net"
    "os"
    "syscall"
    "time"
)

type RawSocketManager struct {
    socketFd int
}

func NewRawSocketManager(protocol int) (*RawSocketManager, error) {
    // Создание raw socket
    fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, protocol)
    if err != nil {
        return nil, err
    }
    
    return &RawSocketManager{socketFd: fd}, nil
}

func (rsm *RawSocketManager) Close() {
    syscall.Close(rsm.socketFd)
}

func (rsm *RawSocketManager) SendPacket(destIP string, data []byte) error {
    // Преобразование IP адреса
    ip := net.ParseIP(destIP)
    if ip == nil {
        return fmt.Errorf("invalid IP address")
    }
    
    ip4 := ip.To4()
    if ip4 == nil {
        return fmt.Errorf("IPv6 not supported")
    }
    
    // Создание адреса назначения
    destAddr := syscall.SockaddrInet4{
        Port: 0,
        Addr: [4]byte{ip4[0], ip4[1], ip4[2], ip4[3]},
    }
    
    // Отправка пакета
    err := syscall.Sendto(rsm.socketFd, data, 0, &destAddr)
    if err != nil {
        return err
    }
    
    fmt.Printf("Sent %d bytes to %s\n", len(data), destIP)
    return nil
}

func (rsm *RawSocketManager) ReceivePacket() ([]byte, syscall.Sockaddr, error) {
    buffer := make([]byte, 4096)
    
    // Чтение пакета
    n, addr, err := syscall.Recvfrom(rsm.socketFd, buffer, 0)
    if err != nil {
        return nil, nil, err
    }
    
    return buffer[:n], addr, nil
}

// IP заголовок (упрощенный)
type IPHeader struct {
    VersionIHL     byte
    TOS            byte
    TotalLength    uint16
    Identification uint16
    FlagsFragOffset uint16
    TTL            byte
    Protocol       byte
    Checksum       uint16
    SourceIP       [4]byte
    DestIP         [4]byte
}

func createIPHeader(srcIP, dstIP string, protocol byte, dataLen int) *IPHeader {
    src := net.ParseIP(srcIP).To4()
    dst := net.ParseIP(dstIP).To4()
    
    header := &IPHeader{
        VersionIHL:     0x45, // IPv4, header length 5 words (20 bytes)
        TOS:            0,
        TotalLength:    uint16(20 + dataLen),
        Identification: uint16(time.Now().UnixNano() & 0xFFFF),
        FlagsFragOffset: 0x4000, // Don't fragment
        TTL:            64,
        Protocol:       protocol,
        Checksum:       0,
    }
    
    copy(header.SourceIP[:], src)
    copy(header.DestIP[:], dst)
    
    // Расчет checksum
    header.Checksum = calculateChecksum(header)
    
    return header
}

func calculateChecksum(header *IPHeader) uint16 {
    // Преобразование заголовка в байты для расчета checksum
    var data []byte
    data = append(data, header.VersionIHL)
    data = append(data, header.TOS)
    data = append(data, byte(header.TotalLength>>8), byte(header.TotalLength))
    data = append(data, byte(header.Identification>>8), byte(header.Identification))
    data = append(data, byte(header.FlagsFragOffset>>8), byte(header.FlagsFragOffset))
    data = append(data, header.TTL)
    data = append(data, header.Protocol)
    data = append(data, 0, 0) // Checksum будет рассчитан
    data = append(data, header.SourceIP[:]...)
    data = append(data, header.DestIP[:]...)
    
    return internetChecksum(data)
}

func internetChecksum(data []byte) uint16 {
    var sum uint32
    
    for i := 0; i < len(data); i += 2 {
        if i+1 < len(data) {
            sum += uint32(data[i])<<8 | uint32(data[i+1])
        } else {
            sum += uint32(data[i]) << 8
        }
    }
    
    for sum>>16 != 0 {
        sum = (sum & 0xFFFF) + (sum >> 16)
    }
    
    return uint16(^sum)
}

func (rsm *RawSocketManager) SendICMPEchoRequest(destIP string) error {
    // Создание ICMP Echo Request
    icmpData := []byte("Hello Raw Socket!")
    
    // ICMP заголовок
    icmpHeader := []byte{
        8,  // Type: Echo Request
        0,  // Code
        0,  // Checksum (будет рассчитан)
        0,  // Checksum продолжение
        0,  // Identifier
        1,  // Identifier продолжение  
        0,  // Sequence Number
        1,  // Sequence Number продолжение
    }
    
    // Расчет ICMP checksum
    icmpPacket := append(icmpHeader, icmpData...)
    checksum := internetChecksum(icmpPacket)
    icmpPacket[2] = byte(checksum >> 8)
    icmpPacket[3] = byte(checksum)
    
    // Создание IP пакета
    ipHeader := createIPHeader("127.0.0.1", destIP, 1, len(icmpPacket))
    
    // Сериализация IP заголовка
    packet := make([]byte, 20)
    packet[0] = ipHeader.VersionIHL
    packet[1] = ipHeader.TOS
    binary.BigEndian.PutUint16(packet[2:4], ipHeader.TotalLength)
    binary.BigEndian.PutUint16(packet[4:6], ipHeader.Identification)
    binary.BigEndian.PutUint16(packet[6:8], ipHeader.FlagsFragOffset)
    packet[8] = ipHeader.TTL
    packet[9] = ipHeader.Protocol
    binary.BigEndian.PutUint16(packet[10:12], ipHeader.Checksum)
    copy(packet[12:16], ipHeader.SourceIP[:])
    copy(packet[16:20], ipHeader.DestIP[:])
    
    // Добавление ICMP пакета
    packet = append(packet, icmpPacket...)
    
    // Отправка
    return rsm.SendPacket(destIP, packet)
}

func main() {
    // Требуются права root для raw sockets
    
    fmt.Println("=== Raw Socket Demo ===")
    
    // Создание raw socket для ICMP
    rsm, err := NewRawSocketManager(syscall.IPPROTO_ICMP)
    if err != nil {
        fmt.Printf("Error creating raw socket: %v\n", err)
        fmt.Println("Note: This requires root privileges on most systems")
        os.Exit(1)
    }
    defer rsm.Close()
    
    fmt.Println("Raw socket created successfully")
    
    // Отправка ICMP Echo Request
    fmt.Println("\n=== Sending ICMP Echo Request ===")
    err = rsm.SendICMPEchoRequest("127.0.0.1")
    if err != nil {
        fmt.Printf("Error sending ICMP: %v\n", err)
    }
    
    // Прием пакетов
    fmt.Println("\n=== Receiving Packets ===")
    fmt.Println("Listening for packets (timeout: 5 seconds)...")
    
    go func() {
        time.Sleep(5 * time.Second)
        fmt.Println("Timeout reached")
        os.Exit(0)
    }()
    
    for {
        data, addr, err := rsm.ReceivePacket()
        if err != nil {
            fmt.Printf("Error receiving packet: %v\n", err)
            continue
        }
        
        if sockAddr, ok := addr.(*syscall.SockaddrInet4); ok {
            srcIP := fmt.Sprintf("%d.%d.%d.%d", 
                sockAddr.Addr[0], sockAddr.Addr[1], 
                sockAddr.Addr[2], sockAddr.Addr[3])
            
            fmt.Printf("Received %d bytes from %s\n", len(data), srcIP)
            
            // Анализ IP заголовка
            if len(data) >= 20 {
                version := data[0] >> 4
                ihl := (data[0] & 0x0F) * 4
                protocol := data[9]
                
                fmt.Printf("IP Version: %d, Header Length: %d, Protocol: %d\n", 
                    version, ihl, protocol)
                
                // Если это ICMP
                if protocol == 1 && len(data) >= int(ihl+8) {
                    icmpType := data[ihl]
                    icmpCode := data[ihl+1]
                    
                    fmt.Printf("ICMP Type: %d, Code: %d\n", icmpType, icmpCode)
                    
                    if icmpType == 0 {
                        fmt.Println("ICMP Echo Reply received!")
                    }
                }
            }
        }
    }
}
//Задание: Работа с raw sockets и системными вызовами
```
138. Memory management и указатели
```go
package main
import (
    "fmt"
    "reflect"
    "runtime"
    "unsafe"
)

type MemoryManager struct{}

func NewMemoryManager() *MemoryManager {
    return &MemoryManager{}
}

func (mm *MemoryManager) ShowMemoryStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
    fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
    fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func (mm *MemoryManager) AnalyzeSliceMemory(slice []int) {
    sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
    
    fmt.Printf("Slice analysis:\n")
    fmt.Printf("  Data pointer: %p\n", unsafe.Pointer(sliceHeader.Data))
    fmt.Printf("  Length: %d\n", sliceHeader.Len)
    fmt.Printf("  Capacity: %d\n", sliceHeader.Cap)
    fmt.Printf("  Total memory: %d bytes\n", sliceHeader.Cap*int(unsafe.Sizeof(slice[0])))
}

func (mm *MemoryManager) ManualMemoryAllocation(size int) unsafe.Pointer {
    // Выделение памяти вручную
    memory := make([]byte, size)
    
    // Получение указателя на данные
    sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&memory))
    return unsafe.Pointer(sliceHeader.Data)
}

func (mm *MemoryManager) PointerArithmetic() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    // Получение указателя на первый элемент
    ptr := unsafe.Pointer(&arr[0])
    
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Base pointer: %p\n", ptr)
    
    // Арифметика указателей через unsafe
    for i := 0; i < len(arr); i++ {
        // Вычисление адреса элемента
        elemPtr := unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(arr[0]))
        value := *(*int)(elemPtr)
        
        fmt.Printf("Element %d: address=%p, value=%d\n", i, elemPtr, value)
    }
}

func (mm *MemoryManager) StringInternals() {
    str := "Hello, World!"
    
    // Анализ внутренней структуры строки
    stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
    
    fmt.Printf("String: \"%s\"\n", str)
    fmt.Printf("Data pointer: %p\n", unsafe.Pointer(stringHeader.Data))
    fmt.Printf("Length: %d\n", stringHeader.Len)
    
    // Конвертация в байты без копирования
    bytes := (*[1 << 30]byte)(unsafe.Pointer(stringHeader.Data))[:stringHeader.Len:stringHeader.Len]
    fmt.Printf("As bytes: %v\n", bytes[:5]) // Первые 5 байт
}

func (mm *MemoryManager) StructMemoryLayout() {
    type ExampleStruct struct {
        flag    bool    // 1 byte
        number  int32   // 4 bytes  
        decimal float64 // 8 bytes
        active  bool    // 1 byte
    }
    
    instance := ExampleStruct{
        flag:    true,
        number:  42,
        decimal: 3.14,
        active:  false,
    }
    
    fmt.Printf("Struct size: %d bytes\n", unsafe.Sizeof(instance))
    fmt.Printf("Field offsets:\n")
    fmt.Printf("  flag: %d\n", unsafe.Offsetof(instance.flag))
    fmt.Printf("  number: %d\n", unsafe.Offsetof(instance.number))
    fmt.Printf("  decimal: %d\n", unsafe.Offsetof(instance.decimal))
    fmt.Printf("  active: %d\n", unsafe.Offsetof(instance.active))
    
    // Выравнивание
    fmt.Printf("Alignment: %d\n", unsafe.Alignof(instance))
}

func (mm *MemoryManager) MemoryPoolDemo() {
    type Object struct {
        data [1024]byte // 1KB объект
        id   int
    }
    
    pool := make([]*Object, 0, 100)
    
    fmt.Println("Memory pool demonstration:")
    
    // Выделение объектов
    for i := 0; i < 10; i++ {
        obj := &Object{id: i}
        pool = append(pool, obj)
    }
    
    mm.ShowMemoryStats()
    
    // "Освобождение" памяти (в Go это делает GC)
    pool = nil
    runtime.GC()
    
    fmt.Println("After GC:")
    mm.ShowMemoryStats()
}

func (mm *MemoryManager) EscapeAnalysisDemo() {
    // Эта функция демонстрирует escape analysis
    
    // Локальная переменная - может остаться в stack
    localVar := 42
    fmt.Printf("Local variable: %d (address: %p)\n", localVar, &localVar)
    
    // Переменная, которая escape в heap
    escapedVar := mm.createObject()
    fmt.Printf("Escaped variable: %v (address: %p)\n", escapedVar, escapedVar)
}

func (mm *MemoryManager) createObject() *map[string]int {
    // Этот объект escape в heap, потому что возвращается из функции
    obj := make(map[string]int)
    obj["result"] = 100
    return &obj
}

func (mm *MemoryManager) UnsafeConversion() {
    // Конвертация между типами через unsafe
    
    // float64 -> int64
    floatValue := 3.14159
    floatBits := *(*uint64)(unsafe.Pointer(&floatValue))
    fmt.Printf("Float: %f, Bits: 0x%016x\n", floatValue, floatBits)
    
    // int64 -> float64
    intValue := int64(42)
    floatFromInt := *(*float64)(unsafe.Pointer(&intValue))
    fmt.Printf("Int: %d, As float: %f\n", intValue, floatFromInt)
    
    // Важно: это опасные операции!
}

func main() {
    mm := NewMemoryManager()
    
    fmt.Println("=== Memory Statistics ===")
    mm.ShowMemoryStats()
    
    fmt.Println("\n=== Slice Memory Analysis ===")
    slice := make([]int, 10, 20)
    for i := range slice {
        slice[i] = i * i
    }
    mm.AnalyzeSliceMemory(slice)
    
    fmt.Println("\n=== Pointer Arithmetic ===")
    mm.PointerArithmetic()
    
    fmt.Println("\n=== String Internals ===")
    mm.StringInternals()
    
    fmt.Println("\n=== Struct Memory Layout ===")
    mm.StructMemoryLayout()
    
    fmt.Println("\n=== Memory Pool ===")
    mm.MemoryPoolDemo()
    
    fmt.Println("\n=== Escape Analysis ===")
    mm.EscapeAnalysisDemo()
    
    fmt.Println("\n=== Unsafe Conversions ===")
    mm.UnsafeConversion()
    
    fmt.Println("\n=== Final Memory Stats ===")
    mm.ShowMemoryStats()
}
//Задание: Управление памятью, указатели и unsafe операции
```
139. Profiling и оптимизация
```go
package main
import (
    "fmt"
    "log"
    "os"
    "runtime"
    "runtime/pprof"
    "sort"
    "time"
)

type ProfilingManager struct {
    cpuProfile *os.File
    memProfile *os.File
}

func NewProfilingManager() *ProfilingManager {
    return &ProfilingManager{}
}

func (pm *ProfilingManager) StartCPUProfile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    
    pm.cpuProfile = file
    pprof.StartCPUProfile(file)
    fmt.Printf("CPU profiling started: %s\n", filename)
    return nil
}

func (pm *ProfilingManager) StopCPUProfile() {
    if pm.cpuProfile != nil {
        pprof.StopCPUProfile()
        pm.cpuProfile.Close()
        fmt.Println("CPU profiling stopped")
    }
}

func (pm *ProfilingManager) WriteMemProfile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    pm.memProfile = file
    runtime.GC() // Принудительный GC перед профилированием памяти
    return pprof.WriteHeapProfile(file)
}

func (pm *ProfilingManager) MeasureExecutionTime(fn func(), name string) time.Duration {
    start := time.Now()
    fn()
    duration := time.Since(start)
    fmt.Printf("%s executed in: %v\n", name, duration)
    return duration
}

func (pm *ProfilingManager) MemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Memory Usage:\n")
    fmt.Printf("  Alloc:      %v MB\n", m.Alloc/1024/1024)
    fmt.Printf("  TotalAlloc: %v MB\n", m.TotalAlloc/1024/1024)
    fmt.Printf("  Sys:        %v MB\n", m.Sys/1024/1024)
    fmt.Printf("  NumGC:      %v\n", m.NumGC)
    fmt.Printf("  GCSys:      %v MB\n", m.GCSys/1024/1024)
}

// Функции для профилирования
func inefficientFunction() {
    // Неэффективная реализация
    var result []int
    for i := 0; i < 100000; i++ {
        result = append(result, i)
    }
    
    // Ненужная сортировка
    sort.Ints(result)
    
    // Лишние операции
    for i := 0; i < len(result); i++ {
        result[i] = result[i] * 2
    }
}

func optimizedFunction() {
    // Оптимизированная реализация
    result := make([]int, 100000)
    for i := 0; i < 100000; i++ {
        result[i] = i * 2
    }
}

func memoryIntensiveFunction() {
    // Функция, интенсивно использующая память
    data := make([][]byte, 1000)
    for i := 0; i < 1000; i++ {
        data[i] = make([]byte, 1024*1024) // 1MB каждый
    }
    
    // Симуляция работы
    time.Sleep(100 * time.Millisecond)
    
    // "Освобождение" памяти
    data = nil
}

func (pm *ProfilingManager) RunProfilingDemo() {
    fmt.Println("=== Profiling Demo ===")
    
    // Запуск CPU профилирования
    pm.StartCPUProfile("cpu.prof")
    defer pm.StopCPUProfile()
    
    // Сравнение производительности
    fmt.Println("\n=== Performance Comparison ===")
    inefficientTime := pm.MeasureExecutionTime(inefficientFunction, "Inefficient function")
    optimizedTime := pm.MeasureExecutionTime(optimizedFunction, "Optimized function")
    
    improvement := float64(inefficientTime-optimizedTime) / float64(inefficientTime) * 100
    fmt.Printf("Performance improvement: %.2f%%\n", improvement)
    
    // Мониторинг памяти
    fmt.Println("\n=== Memory Monitoring ===")
    pm.MemoryUsage()
    
    // Запуск memory-intensive функции
    fmt.Println("\n=== Memory Intensive Operation ===")
    pm.MeasureExecutionTime(memoryIntensiveFunction, "Memory intensive function")
    
    // Сбор мусора
    runtime.GC()
    
    // Финальная статистика памяти
    fmt.Println("\n=== Final Memory Stats ===")
    pm.MemoryUsage()
    
    // Сохранение memory profile
    pm.WriteMemProfile("mem.prof")
    fmt.Println("Memory profile saved: mem.prof")
}

func (pm *ProfilingManager) GoroutineProfile() {
    fmt.Println("\n=== Goroutine Profile ===")
    
    // Создание нескольких горутин
    for i := 0; i < 10; i++ {
        go func(id int) {
            time.Sleep(time.Duration(id) * time.Second)
        }(i)
    }
    
    // Получение профиля горутин
    profile := pprof.Lookup("goroutine")
    profile.WriteTo(os.Stdout, 1)
    
    fmt.Printf("Current goroutines: %d\n", runtime.NumGoroutine())
}

func (pm *ProfilingManager) BlockProfile() {
    fmt.Println("\n=== Block Profile ===")
    
    // Включение блокирующего профилирования
    runtime.SetBlockProfileRate(1)
    
    // Функция с блокировками
    ch := make(chan int)
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch <- 42
    }()
    
    // Блокирующая операция
    <-ch
    
    // Получение профиля блокировок
    profile := pprof.Lookup("block")
    profile.WriteTo(os.Stdout, 1)
}

func (pm *ProfilingManager) CustomProfile() {
    fmt.Println("\n=== Custom Profile ===")
    
    // Создание кастомного профиля
    customProf := pprof.NewProfile("custom.operations")
    
    // Регистрация операций
    startOperation := func(name string) func() {
        start := time.Now()
        return func() {
            duration := time.Since(start)
            customProf.Add(nil, int64(duration))
            fmt.Printf("Operation '%s' took: %v\n", name, duration)
        }
    }
    
    // Измерение операций
    defer startOperation("database_query")()
    time.Sleep(50 * time.Millisecond)
    
    defer startOperation("file_operation")()
    time.Sleep(30 * time.Millisecond)
    
    // Сохранение кастомного профиля
    file, _ := os.Create("custom.prof")
    defer file.Close()
    customProf.WriteTo(file, 0)
    fmt.Println("Custom profile saved: custom.prof")
}

func main() {
    pm := NewProfilingManager()
    
    // Запуск демонстрации профилирования
    pm.RunProfilingDemo()
    
    // Дополнительные профили
    pm.GoroutineProfile()
    pm.BlockProfile()
    pm.CustomProfile()
    
    fmt.Println("\n=== Profiling Completed ===")
    fmt.Println("To analyze profiles, run:")
    fmt.Println("  go tool pprof cpu.prof")
    fmt.Println("  go tool pprof mem.prof")
    fmt.Println("  go tool pprof custom.prof")
}

// Дополнительная утилита для анализа
func analyzePerformance() {
    // Анализ аллокаций
    type allocation struct {
        size int
        desc string
    }
    
    allocations := []allocation{
        {1024, "small buffer"},
        {1024 * 1024, "large buffer"},
        {512, "medium buffer"},
    }
    
    // Сортировка по размеру
    sort.Slice(allocations, func(i, j int) bool {
        return allocations[i].size < allocations[j].size
    })
    
    fmt.Println("Allocations sorted by size:")
    for _, alloc := range allocations {
        fmt.Printf("  %s: %d bytes\n", alloc.desc, alloc.size)
    }
}
//Задание: Профилирование производительности и оптимизация кода
```
140. Cross-platform компиляция
```go
package main
import (
    "fmt"
    "runtime"
    "syscall"
)

type PlatformInfo struct {
    OS      string
    Arch    string
    Version string
}

func GetPlatformInfo() *PlatformInfo {
    info := &PlatformInfo{
        OS:   runtime.GOOS,
        Arch: runtime.GOARCH,
    }
    
    // Получение информации о версии ОС
    switch runtime.GOOS {
    case "windows":
        info.Version = getWindowsVersion()
    case "linux":
        info.Version = getLinuxVersion()
    case "darwin":
        info.Version = getMacVersion()
    default:
        info.Version = "unknown"
    }
    
    return info
}

func getWindowsVersion() string {
    // Используем syscall для получения версии Windows
    kernel32 := syscall.NewLazyDLL("kernel32.dll")
    getVersion := kernel32.NewProc("GetVersion")
    
    ret, _, _ := getVersion.Call()
    version := byte(ret)
    major := byte(ret >> 8)
    
    return fmt.Sprintf("%d.%d", version, major)
}

func getLinuxVersion() string {
    // Чтение информации о версии Linux
    // В реальном приложении нужно читать /etc/os-release
    return "Linux"
}

func getMacVersion() string {
    // Получение версии macOS
    // В реальном приложении используйте sysctl или другие методы
    return "macOS"
}

func ShowCompilationInstructions() {
    fmt.Println("=== Cross-Platform Compilation ===")
    fmt.Println("To compile for different platforms, use:")
    fmt.Println()
    
    targets := map[string]string{
        "Windows 64-bit": "GOOS=windows GOARCH=amd64 go build -o app.exe",
        "Windows 32-bit": "GOOS=windows GOARCH=386 go build -o app.exe",
        "Linux 64-bit":   "GOOS=linux GOARCH=amd64 go build -o app",
        "Linux 32-bit":   "GOOS=linux GOARCH=386 go build -o app",
        "macOS 64-bit":   "GOOS=darwin GOARCH=amd64 go build -o app",
        "macOS ARM":      "GOOS=darwin GOARCH=arm64 go build -o app",
        "FreeBSD":        "GOOS=freebsd GOARCH=amd64 go build -o app",
        "Android":        "GOOS=android GOARCH=arm64 go build -o app",
    }
    
    for platform, command := range targets {
        fmt.Printf("%s:\n  %s\n\n", platform, command)
    }
}

func PlatformSpecificCode() {
    fmt.Println("=== Platform-Specific Features ===")
    
    switch runtime.GOOS {
    case "windows":
        fmt.Println("Windows-specific features:")
        fmt.Println("  - Registry access")
        fmt.Println("  - COM components")
        fmt.Println("  - Windows API calls")
        
    case "linux":
        fmt.Println("Linux-specific features:")
        fmt.Println("  - Systemd integration")
        fmt.Println("  - cgroups support")
        fmt.Println("  - Linux namespaces")
        
    case "darwin":
        fmt.Println("macOS-specific features:")
        fmt.Println("  - Cocoa integration")
        fmt.Println("  - Launchd services")
        fmt.Println("  - Apple Script")
        
    default:
        fmt.Println("Unknown platform features")
    }
}

func FilePathExample() {
    fmt.Println("=== Cross-Platform File Paths ===")
    
    // Использование filepath для кроссплатформенных путей
    path := "directory" + string(filepath.Separator) + "file.txt"
    fmt.Printf("Platform-specific path: %s\n", path)
    
    // Примеры путей
    paths := []string{
        filepath.Join("usr", "local", "bin"),
        filepath.Join("C:", "Program Files", "App"),
        filepath.Join("/home", "user", "documents"),
    }
    
    for _, p := range paths {
        fmt.Printf("Normalized path: %s\n", filepath.Clean(p))
    }
}

func EnvironmentVariables() {
    fmt.Println("=== Platform-Specific Environment ===")
    
    // Переменные окружения, специфичные для платформ
    envVars := map[string]string{
        "windows": "PATH",
        "linux":   "PATH",
        "darwin":  "PATH",
    }
    
    if varName, exists := envVars[runtime.GOOS]; exists {
        value := os.Getenv(varName)
        if len(value) > 100 {
            value = value[:100] + "..."
        }
        fmt.Printf("%s: %s\n", varName, value)
    }
    
    // Дополнительные переменные
    switch runtime.GOOS {
    case "windows":
        fmt.Printf("OS: %s\n", os.Getenv("OS"))
        fmt.Printf("COMPUTERNAME: %s\n", os.Getenv("COMPUTERNAME"))
    case "linux", "darwin":
        fmt.Printf("USER: %s\n", os.Getenv("USER"))
        fmt.Printf("HOME: %s\n", os.Getenv("HOME"))
    }
}

func BuildConstraintsDemo() {
    fmt.Println("=== Build Constraints ===")
    fmt.Println("Add these comments to conditionally compile code:")
    fmt.Println()
    
    constraints := []string{
        "//go:build windows",
        "// +build windows",
        "",
        "//go:build linux || darwin",
        "// +build linux darwin",
        "",
        "//go:build !windows",
        "// +build !windows",
    }
    
    for _, constraint := range constraints {
        fmt.Println(constraint)
    }
}

func main() {
    // Информация о платформе
    info := GetPlatformInfo()
    fmt.Printf("Current Platform: %s/%s\n", info.OS, info.Arch)
    fmt.Printf("OS Version: %s\n", info.Version)
    fmt.Println()
    
    // Инструкции по компиляции
    ShowCompilationInstructions()
    
    // Платформо-специфичные особенности
    PlatformSpecificCode()
    fmt.Println()
    
    // Работа с путями
    FilePathExample()
    fmt.Println()
    
    // Переменные окружения
    EnvironmentVariables()
    fmt.Println()
    
    // Build constraints
    BuildConstraintsDemo()
    
    // Демонстрация кроссплатформенного кода
    fmt.Println("\n=== Cross-Platform Demo ===")
    demoCrossPlatformFeatures()
}

func demoCrossPlatformFeatures() {
    // Кроссплатформенные операции
    
    // 1. Работа с файлами
    tempFile, err := os.CreateTemp("", "crossplatform_*.txt")
    if err == nil {
        defer os.Remove(tempFile.Name())
        tempFile.WriteString("Cross-platform file content")
        fmt.Printf("Created temp file: %s\n", tempFile.Name())
    }
    
    // 2. Работа с временем
    now := time.Now()
    fmt.Printf("Current time: %s\n", now.Format(time.RFC1123))
    
    // 3. Работа с сетью
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Printf("Hostname: %s\n", hostname)
    }
    
    // 4. Параллелизм
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d on %s\n", id, runtime.GOOS)
        }(i)
    }
    wg.Wait()
}
//Задание: Кроссплатформенная компиляция и платформо-специфичный код
```
Сетевые протоколы и распределенные системы
141. DNS клиент и сервер
```go
package main
import (
    "fmt"
    "net"
    "strings"
)

type DNSClient struct{}

func NewDNSClient() *DNSClient {
    return &DNSClient{}
}

func (dc *DNSClient) LookupHost(hostname string) ([]string, error) {
    return net.LookupHost(hostname)
}

func (dc *DNSClient) LookupIP(hostname string) ([]net.IP, error) {
    return net.LookupIP(hostname)
}

func (dc *DNSClient) LookupCNAME(hostname string) (string, error) {
    return net.LookupCNAME(hostname)
}

func (dc *DNSClient) LookupMX(hostname string) ([]*net.MX, error) {
    return net.LookupMX(hostname)
}

func (dc *DNSClient) LookupTXT(hostname string) ([]string, error) {
    return net.LookupTXT(hostname)
}

func (dc *DNSClient) LookupNS(hostname string) ([]*net.NS, error) {
    return net.LookupNS(hostname)
}

func (dc *DNSClient) CustomDNSLookup(hostname string, dnsServer string) ([]string, error) {
    resolver := &net.Resolver{
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
            d := net.Dialer{}
            return d.DialContext(ctx, "udp", dnsServer+":53")
        },
    }
    
    return resolver.LookupHost(context.Background(), hostname)
}

func (dc *DNSClient) ReverseDNSLookup(ip string) ([]string, error) {
    return net.LookupAddr(ip)
}

func main() {
    dnsClient := NewDNSClient()
    
    hostname := "google.com"
    
    fmt.Printf("=== DNS Lookup for %s ===\n", hostname)
    
    // A records
    ips, err := dnsClient.LookupHost(hostname)
    if err == nil {
        fmt.Printf("A Records: %v\n", ips)
    }
    
    // MX records
    mxRecords, err := dnsClient.LookupMX(hostname)
    if err == nil {
        fmt.Printf("MX Records:\n")
        for _, mx := range mxRecords {
            fmt.Printf("  %s (pref: %d)\n", mx.Host, mx.Pref)
        }
    }
    
    // TXT records
    txtRecords, err := dnsClient.LookupTXT(hostname)
    if err == nil {
        fmt.Printf("TXT Records: %v\n", txtRecords)
    }
    
    // NS records
    nsRecords, err := dnsClient.LookupNS(hostname)
    if err == nil {
        fmt.Printf("NS Records:\n")
        for _, ns := range nsRecords {
            fmt.Printf("  %s\n", ns.Host)
        }
    }
    
    // CNAME
    cname, err := dnsClient.LookupCNAME("www." + hostname)
    if err == nil {
        fmt.Printf("CNAME: %s\n", cname)
    }
    
    // Custom DNS server
    customIPs, err := dnsClient.CustomDNSLookup(hostname, "8.8.8.8")
    if err == nil {
        fmt.Printf("Using Google DNS: %v\n", customIPs)
    }
    
    // Reverse DNS
    if len(ips) > 0 {
        reverse, err := dnsClient.ReverseDNSLookup(ips[0])
        if err == nil {
            fmt.Printf("Reverse DNS for %s: %v\n", ips[0], reverse)
        }
    }
}

// Простой DNS сервер
type DNSServer struct {
    records map[string]string
}

func NewDNSServer() *DNSServer {
    return &DNSServer{
        records: map[string]string{
            "test.local.": "127.0.0.1",
            "api.local.":  "192.168.1.100",
            "db.local.":   "192.168.1.101",
        },
    }
}

func (ds *DNSServer) Start() error {
    conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 5353})
    if err != nil {
        return err
    }
    defer conn.Close()
    
    fmt.Println("DNS Server listening on :5353")
    
    buffer := make([]byte, 512)
    for {
        n, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            continue
        }
        
        go ds.handleQuery(conn, addr, buffer[:n])
    }
}

func (ds *DNSServer) handleQuery(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
    // Упрощенная обработка DNS запроса
    query := string(data)
    if strings.Contains(query, "test.local") {
        response := []byte("test.local 127.0.0.1")
        conn.WriteToUDP(response, addr)
    }
}
//Задание: DNS клиент с различными типами запросов и простой DNS сервер
```
142. FTP клиент
```go
package main
import (
    "bufio"
    "fmt"
    "io"
    "net"
    "net/textproto"
    "strconv"
    "strings"
)

type FTPClient struct {
    conn     net.Conn
    reader   *textproto.Reader
    writer   *textproto.Writer
    host     string
    username string
    password string
}

func NewFTPClient(host string) *FTPClient {
    return &FTPClient{
        host: host,
    }
}

func (fc *FTPClient) Connect() error {
    conn, err := net.Dial("tcp", fc.host+":21")
    if err != nil {
        return err
    }
    
    fc.conn = conn
    fc.reader = textproto.NewReader(bufio.NewReader(conn))
    fc.writer = textproto.NewWriter(bufio.NewWriter(conn))
    
    // Чтение приветственного сообщения
    message, err := fc.reader.ReadLine()
    if err != nil {
        return err
    }
    fmt.Printf("Server: %s\n", message)
    
    return nil
}

func (fc *FTPClient) Login(username, password string) error {
    fc.username = username
    fc.password = password
    
    // USER command
    if err := fc.sendCommand("USER " + username); err != nil {
        return err
    }
    
    // PASS command
    if err := fc.sendCommand("PASS " + password); err != nil {
        return err
    }
    
    return nil
}

func (fc *FTPClient) sendCommand(command string) error {
    fmt.Printf("Client: %s\n", command)
    if err := fc.writer.PrintfLine(command); err != nil {
        return err
    }
    
    response, err := fc.reader.ReadLine()
    if err != nil {
        return err
    }
    fmt.Printf("Server: %s\n", response)
    
    if !strings.HasPrefix(response, "2") && !strings.HasPrefix(response, "3") {
        return fmt.Errorf("command failed: %s", response)
    }
    
    return nil
}

func (fc *FTPClient) ListFiles() error {
    // Переход в пассивный режим
    if err := fc.sendCommand("PASV"); err != nil {
        return err
    }
    
    response, _ := fc.reader.ReadLine()
    
    // Парсинг PASV response для получения адреса данных
    start := strings.Index(response, "(")
    end := strings.Index(response, ")")
    if start == -1 || end == -1 {
        return fmt.Errorf("invalid PASV response")
    }
    
    pasvData := strings.Split(response[start+1:end], ",")
    if len(pasvData) != 6 {
        return fmt.Errorf("invalid PASV data")
    }
    
    ip := strings.Join(pasvData[0:4], ".")
    port1, _ := strconv.Atoi(pasvData[4])
    port2, _ := strconv.Atoi(pasvData[5])
    port := port1*256 + port2
    
    // Подключение к data connection
    dataConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
    if err != nil {
        return err
    }
    defer dataConn.Close()
    
    // Отправка LIST команды
    if err := fc.sendCommand("LIST"); err != nil {
        return err
    }
    
    // Чтение данных
    data, err := io.ReadAll(dataConn)
    if err != nil {
        return err
    }
    
    fmt.Printf("Directory listing:\n%s\n", string(data))
    
    // Чтение завершающего ответа
    response, _ = fc.reader.ReadLine()
    fmt.Printf("Server: %s\n", response)
    
    return nil
}

func (fc *FTPClient) DownloadFile(filename string) error {
    // Аналогично ListFiles, но для RETR команды
    if err := fc.sendCommand("RETR " + filename); err != nil {
        return err
    }
    
    // Здесь должна быть реализация загрузки файла
    // Аналогично ListFiles с созданием data connection
    
    return nil
}

func (fc *FTPClient) UploadFile(filename string, data []byte) error {
    // Аналогично ListFiles, но для STOR команды
    if err := fc.sendCommand("STOR " + filename); err != nil {
        return err
    }
    
    // Здесь должна быть реализация загрузки файла
    // Аналогично ListFiles с созданием data connection
    
    return nil
}

func (fc *FTPClient) ChangeDirectory(path string) error {
    return fc.sendCommand("CWD " + path)
}

func (fc *FTPClient) PrintWorkingDirectory() error {
    return fc.sendCommand("PWD")
}

func (fc *FTPClient) Quit() error {
    if err := fc.sendCommand("QUIT"); err != nil {
        return err
    }
    return fc.conn.Close()
}

func main() {
    client := NewFTPClient("localhost")
    
    // Подключение
    if err := client.Connect(); err != nil {
        fmt.Printf("Connection error: %v\n", err)
        return
    }
    
    // Аутентификация (анонимный доступ)
    if err := client.Login("anonymous", "guest"); err != nil {
        fmt.Printf("Login error: %v\n", err)
        return
    }
    
    // Получение списка файлов
    if err := client.ListFiles(); err != nil {
        fmt.Printf("List error: %v\n", err)
    }
    
    // Текущая директория
    client.PrintWorkingDirectory()
    
    // Завершение сессии
    client.Quit()
}
//Задание: FTP клиент с поддержкой основных команд
```
143. SMTP клиент для отправки email
```go
package main
import (
    "crypto/tls"
    "fmt"
    "net"
    "net/smtp"
    "strings"
)

type Email struct {
    From    string
    To      []string
    Subject string
    Body    string
}

type SMTPClient struct {
    host     string
    port     int
    username string
    password string
    tls      bool
}

func NewSMTPClient(host string, port int, username, password string, useTLS bool) *SMTPClient {
    return &SMTPClient{
        host:     host,
        port:     port,
        username: username,
        password: password,
        tls:      useTLS,
    }
}

func (sc *SMTPClient) SendEmail(email *Email) error {
    // Подготовка сообщения
    message := sc.buildMessage(email)
    
    // Адрес сервера
    addr := fmt.Sprintf("%s:%d", sc.host, sc.port)
    
    if sc.tls {
        return sc.sendWithTLS(addr, email, message)
    } else {
        return sc.sendPlain(addr, email, message)
    }
}

func (sc *SMTPClient) sendPlain(addr string, email *Email, message string) error {
    // Аутентификация
    auth := smtp.PlainAuth("", sc.username, sc.password, sc.host)
    
    // Отправка email
    return smtp.SendMail(addr, auth, email.From, email.To, []byte(message))
}

func (sc *SMTPClient) sendWithTLS(addr string, email *Email, message string) error {
    // Подключение к серверу
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }
    defer conn.Close()
    
    // Создание SMTP клиента
    client, err := smtp.NewClient(conn, sc.host)
    if err != nil {
        return err
    }
    defer client.Close()
    
    // STARTTLS
    if err = client.StartTLS(&tls.Config{ServerName: sc.host}); err != nil {
        return err
    }
    
    // Аутентификация
    auth := smtp.PlainAuth("", sc.username, sc.password, sc.host)
    if err = client.Auth(auth); err != nil {
        return err
    }
    
    // Отправитель
    if err = client.Mail(email.From); err != nil {
        return err
    }
    
    // Получатели
    for _, to := range email.To {
        if err = client.Rcpt(to); err != nil {
            return err
        }
    }
    
    // Данные
    w, err := client.Data()
    if err != nil {
        return err
    }
    
    // Запись сообщения
    _, err = w.Write([]byte(message))
    if err != nil {
        return err
    }
    
    // Закрытие writer'а
    err = w.Close()
    if err != nil {
        return err
    }
    
    return client.Quit()
}

func (sc *SMTPClient) buildMessage(email *Email) string {
    var message strings.Builder
    
    // Заголовки
    message.WriteString(fmt.Sprintf("From: %s\r\n", email.From))
    message.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ", ")))
    message.WriteString(fmt.Sprintf("Subject: %s\r\n", email.Subject))
    message.WriteString("MIME-Version: 1.0\r\n")
    message.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
    message.WriteString("\r\n")
    
    // Тело сообщения
    message.WriteString(email.Body)
    message.WriteString("\r\n")
    
    return message.String()
}

func (sc *SMTPClient) SendHTMLEmail(email *Email, htmlBody string) error {
    // Создание HTML сообщения
    var message strings.Builder
    
    message.WriteString(fmt.Sprintf("From: %s\r\n", email.From))
    message.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ", ")))
    message.WriteString(fmt.Sprintf("Subject: %s\r\n", email.Subject))
    message.WriteString("MIME-Version: 1.0\r\n")
    message.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n")
    message.WriteString("\r\n")
    message.WriteString(htmlBody)
    message.WriteString("\r\n")
    
    // Отправка
    return sc.SendEmail(&Email{
        From:    email.From,
        To:      email.To,
        Subject: email.Subject,
        Body:    message.String(),
    })
}

func (sc *SMTPClient) VerifyConnection() error {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", sc.host, sc.port))
    if err != nil {
        return err
    }
    defer conn.Close()
    
    // Чтение приветственного сообщения
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        return err
    }
    
    fmt.Printf("SMTP Server: %s", string(buffer[:n]))
    return nil
}

func main() {
    // Настройки SMTP (пример для Gmail)
    client := NewSMTPClient("smtp.gmail.com", 587, "your-email@gmail.com", "your-password", true)
    
    // Проверка подключения
    if err := client.VerifyConnection(); err != nil {
        fmt.Printf("Connection test failed: %v\n", err)
        return
    }
    
    // Создание email
    email := &Email{
        From:    "your-email@gmail.com",
        To:      []string{"recipient@example.com"},
        Subject: "Test Email from Go",
        Body:    "This is a test email sent from Go program using SMTP.",
    }
    
    // Отправка plain text email
    if err := client.SendEmail(email); err != nil {
        fmt.Printf("Error sending email: %v\n", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
    
    // Отправка HTML email
    htmlEmail := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Test Email</title>
    </head>
    <body>
        <h1>Hello from Go!</h1>
        <p>This is a <strong>HTML email</strong> sent from Go program.</p>
        <p>Current time: ` + time.Now().Format(time.RFC1123) + `</p>
    </body>
    </html>`
    
    if err := client.SendHTMLEmail(email, htmlEmail); err != nil {
        fmt.Printf("Error sending HTML email: %v\n", err)
    } else {
        fmt.Println("HTML email sent successfully!")
    }
}
//Задание: SMTP клиент для отправки email с поддержкой TLS и HTML
```
144. Web crawler
```go
package main
import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "net/url"
    "sync"
    "time"
)

type Crawler struct {
    visited   sync.Map
    maxDepth  int
    delay     time.Duration
    userAgent string
    mu        sync.Mutex
}

type Page struct {
    URL   string
    Title string
    Links []string
}

func NewCrawler(maxDepth int, delay time.Duration) *Crawler {
    return &Crawler{
        maxDepth:  maxDepth,
        delay:     delay,
        userAgent: "GoWebCrawler/1.0",
        visited:   sync.Map{},
    }
}

func (c *Crawler) Crawl(startURL string) ([]*Page, error) {
    var pages []*Page
    var wg sync.WaitGroup
    var mu sync.Mutex
    
    // Канал для ограничения параллелизма
    semaphore := make(chan struct{}, 10)
    
    // Функция для рекурсивного crawling'а
    var crawl func(string, int)
    crawl = func(currentURL string, depth int) {
        if depth > c.maxDepth {
            return
        }
        
        // Проверка, не посещали ли уже эту страницу
        if _, visited := c.visited.LoadOrStore(currentURL, true); visited {
            return
        }
        
        // Ограничение параллелизма
        semaphore <- struct{}{}
        defer func() { <-semaphore }()
        
        // Задержка для соблюдения robots.txt
        time.Sleep(c.delay)
        
        // Получение страницы
        page, err := c.fetchPage(currentURL)
        if err != nil {
            fmt.Printf("Error fetching %s: %v\n", currentURL, err)
            return
        }
        
        mu.Lock()
        pages = append(pages, page)
        mu.Unlock()
        
        fmt.Printf("Crawled: %s (depth: %d, links: %d)\n", currentURL, depth, len(page.Links))
        
        // Рекурсивный обход ссылок
        for _, link := range page.Links {
            wg.Add(1)
            go func(l string, d int) {
                defer wg.Done()
                crawl(l, d+1)
            }(link, depth)
        }
    }
    
    wg.Add(1)
    go func() {
        defer wg.Done()
        crawl(startURL, 0)
    }()
    
    wg.Wait()
    return pages, nil
}

func (c *Crawler) fetchPage(pageURL string) (*Page, error) {
    // Создание HTTP клиента
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    // Создание запроса
    req, err := http.NewRequest("GET", pageURL, nil)
    if err != nil {
        return nil, err
    }
    
    // Установка User-Agent
    req.Header.Set("User-Agent", c.userAgent)
    
    // Выполнение запроса
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    // Парсинг HTML
    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }
    
    page := &Page{
        URL: pageURL,
    }
    
    // Извлечение данных
    c.extractData(doc, page)
    
    return page, nil
}

func (c *Crawler) extractData(n *html.Node, page *Page) {
    if n.Type == html.ElementNode {
        switch n.Data {
        case "title":
            if n.FirstChild != nil {
                page.Title = n.FirstChild.Data
            }
        case "a":
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    absoluteURL := c.resolveURL(page.URL, attr.Val)
                    if absoluteURL != "" {
                        page.Links = append(page.Links, absoluteURL)
                    }
                    break
                }
            }
        }
    }
    
    for child := n.FirstChild; child != nil; child = child.NextSibling {
        c.extractData(child, page)
    }
}

func (c *Crawler) resolveURL(base, relative string) string {
    baseURL, err := url.Parse(base)
    if err != nil {
        return ""
    }
    
    relativeURL, err := url.Parse(relative)
    if err != nil {
        return ""
    }
    
    resolvedURL := baseURL.ResolveReference(relativeURL)
    
    // Фильтрация URL (только HTTP/HTTPS)
    if resolvedURL.Scheme != "http" && resolvedURL.Scheme != "https" {
        return ""
    }
    
    return resolvedURL.String()
}

func (c *Crawler) GenerateSitemap(pages []*Page) string {
    var sitemap strings.Builder
    
    sitemap.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
    sitemap.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
    
    for _, page := range pages {
        sitemap.WriteString("<url>")
        sitemap.WriteString(fmt.Sprintf("<loc>%s</loc>", page.URL))
        if page.Title != "" {
            sitemap.WriteString(fmt.Sprintf("<title>%s</title>", page.Title))
        }
        sitemap.WriteString("</url>")
    }
    
    sitemap.WriteString("</urlset>")
    return sitemap.String()
}

func (c *Crawler) SaveResults(pages []*Page, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    for _, page := range pages {
        fmt.Fprintf(file, "URL: %s\n", page.URL)
        fmt.Fprintf(file, "Title: %s\n", page.Title)
        fmt.Fprintf(file, "Links: %d\n", len(page.Links))
        fmt.Fprintf(file, "---\n")
    }
    
    return nil
}

func main() {
    crawler := NewCrawler(2, 1*time.Second)
    
    startURL := "https://example.com"
    
    fmt.Printf("Starting crawl from: %s\n", startURL)
    fmt.Printf("Max depth: %d, Delay: %v\n", crawler.maxDepth, crawler.delay)
    
    startTime := time.Now()
    pages, err := crawler.Crawl(startURL)
    elapsed := time.Since(startTime)
    
    if err != nil {
        fmt.Printf("Crawl error: %v\n", err)
        return
    }
    
    fmt.Printf("\nCrawl completed in %v\n", elapsed)
    fmt.Printf("Pages crawled: %d\n", len(pages))
    
    // Генерация sitemap
    sitemap := crawler.GenerateSitemap(pages)
    os.WriteFile("sitemap.xml", []byte(sitemap), 0644)
    fmt.Println("Sitemap saved: sitemap.xml")
    
    // Сохранение результатов
    crawler.SaveResults(pages, "crawl_results.txt")
    fmt.Println("Results saved: crawl_results.txt")
    
    // Статистика
    totalLinks := 0
    for _, page := range pages {
        totalLinks += len(page.Links)
    }
    
    fmt.Printf("Total links found: %d\n", totalLinks)
    fmt.Printf("Average links per page: %.2f\n", float64(totalLinks)/float64(len(pages)))
}
//Задание: Веб-краулер с рекурсивным обходом и генерацией sitemap
```
145. Distributed key-value store
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net"
    "net/http"
    "net/rpc"
    "sync"
    "time"
)

type KeyValue struct {
    Key   string      `json:"key"`
    Value interface{} `json:"value"`
    Time  time.Time   `json:"timestamp"`
}

type Store struct {
    data  map[string]*KeyValue
    mutex sync.RWMutex
}

type StoreService struct {
    store *Store
}

type PutArgs struct {
    Key   string      `json:"key"`
    Value interface{} `json:"value"`
}

type GetArgs struct {
    Key string `json:"key"`
}

type GetReply struct {
    Value *KeyValue `json:"value"`
    Error string    `json:"error,omitempty"`
}

func NewStore() *Store {
    return &Store{
        data: make(map[string]*KeyValue),
    }
}

func (s *Store) Put(key string, value interface{}) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    s.data[key] = &KeyValue{
        Key:   key,
        Value: value,
        Time:  time.Now(),
    }
}

func (s *Store) Get(key string) (*KeyValue, bool) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    value, exists := s.data[key]
    return value, exists
}

func (s *Store) Delete(key string) bool {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    if _, exists := s.data[key]; exists {
        delete(s.data, key)
        return true
    }
    return false
}

func (s *Store) Keys() []string {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    keys := make([]string, 0, len(s.data))
    for key := range s.data {
        keys = append(keys, key)
    }
    return keys
}

func (s *Store) Size() int {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    return len(s.data)
}

func (ss *StoreService) Put(args *PutArgs, reply *bool) error {
    ss.store.Put(args.Key, args.Value)
    *reply = true
    return nil
}

func (ss *StoreService) Get(args *GetArgs, reply *GetReply) error {
    value, exists := ss.store.Get(args.Key)
    if !exists {
        reply.Error = "key not found"
        return nil
    }
    
    reply.Value = value
    return nil
}

func (ss *StoreService) Delete(args *GetArgs, reply *bool) error {
    *reply = ss.store.Delete(args.Key)
    return nil
}

func (ss *StoreService) Keys(args struct{}, reply *[]string) error {
    *reply = ss.store.Keys()
    return nil
}

type Node struct {
    ID      string
    Address string
    Store   *Store
    Server  *rpc.Server
}

func NewNode(id, address string) *Node {
    node := &Node{
        ID:      id,
        Address: address,
        Store:   NewStore(),
        Server:  rpc.NewServer(),
    }
    
    // Регистрация RPC сервиса
    storeService := &StoreService{store: node.Store}
    node.Server.Register(storeService)
    
    return node
}

func (n *Node) Start() error {
    listener, err := net.Listen("tcp", n.Address)
    if err != nil {
        return err
    }
    
    log.Printf("Node %s listening on %s", n.ID, n.Address)
    
    // HTTP handler для RPC
    http.Handle("/rpc", n.Server)
    
    return http.Serve(listener, nil)
}

func (n *Node) ConnectToNode(address string) (*rpc.Client, error) {
    return rpc.DialHTTP("tcp", address)
}

type Cluster struct {
    Nodes map[string]*Node
    mutex sync.RWMutex
}

func NewCluster() *Cluster {
    return &Cluster{
        Nodes: make(map[string]*Node),
    }
}

func (c *Cluster) AddNode(node *Node) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.Nodes[node.ID] = node
}

func (c *Cluster) RemoveNode(nodeID string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    delete(c.Nodes, nodeID)
}

func (c *Cluster) ReplicatePut(key string, value interface{}) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    for _, node := range c.Nodes {
        go func(n *Node) {
            client, err := n.ConnectToNode(n.Address)
            if err != nil {
                log.Printf("Failed to connect to node %s: %v", n.ID, err)
                return
            }
            defer client.Close()
            
            args := &PutArgs{Key: key, Value: value}
            var reply bool
            err = client.Call("StoreService.Put", args, &reply)
            if err != nil {
                log.Printf("Failed to replicate to node %s: %v", n.ID, err)
            }
        }(node)
    }
}

func (c *Cluster) GetFromAnyNode(key string) (*KeyValue, error) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    for _, node := range c.Nodes {
        client, err := node.ConnectToNode(node.Address)
        if err != nil {
            continue
        }
        defer client.Close()
        
        args := &GetArgs{Key: key}
        var reply GetReply
        err = client.Call("StoreService.Get", args, &reply)
        if err == nil && reply.Error == "" {
            return reply.Value, nil
        }
    }
    
    return nil, fmt.Errorf("key not found in any node")
}

func main() {
    // Создание кластера
    cluster := NewCluster()
    
    // Создание узлов
    node1 := NewNode("node1", ":8081")
    node2 := NewNode("node2", ":8082")
    node3 := NewNode("node3", ":8083")
    
    // Добавление узлов в кластер
    cluster.AddNode(node1)
    cluster.AddNode(node2)
    cluster.AddNode(node3)
    
    // Запуск узлов
    for _, node := range cluster.Nodes {
        go func(n *Node) {
            if err := n.Start(); err != nil {
                log.Printf("Node %s failed: %v", n.ID, err)
            }
        }(node)
    }
    
    // Даем узлам время запуститься
    time.Sleep(2 * time.Second)
    
    // Демонстрация работы
    fmt.Println("=== Distributed Key-Value Store Demo ===")
    
    // Запись данных с репликацией
    cluster.ReplicatePut("name", "Alice")
    cluster.ReplicatePut("age", 30)
    cluster.ReplicatePut("city", "New York")
    
    fmt.Println("Data replicated across cluster")
    
    // Чтение данных
    value, err := cluster.GetFromAnyNode("name")
    if err == nil {
        fmt.Printf("Retrieved: %s = %v\n", value.Key, value.Value)
    }
    
    value, err = cluster.GetFromAnyNode("age")
    if err == nil {
        fmt.Printf("Retrieved: %s = %v\n", value.Key, value.Value)
    }
    
    // Статистика
    fmt.Printf("Cluster size: %d nodes\n", len(cluster.Nodes))
    
    // Ожидание
    select {}
}
//Задание: Распределенное key-value хранилище с RPC
```
146. Real-time collaboration сервер
```go
package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
    
    "github.com/gorilla/websocket"
)

type Document struct {
    ID      string    `json:"id"`
    Content string    `json:"content"`
    Version int       `json:"version"`
    Created time.Time `json:"created"`
    Updated time.Time `json:"updated"`
}

type Operation struct {
    Type      string `json:"type"` // "insert", "delete", "cursor"
    Position  int    `json:"position"`
    Text      string `json:"text"`
    ClientID  string `json:"clientId"`
    Timestamp int64  `json:"timestamp"`
}

type Client struct {
    ID       string
    Document string
    Conn     *websocket.Conn
    Send     chan []byte
}

type CollaborationServer struct {
    documents map[string]*Document
    clients   map[string]*Client
    mutex     sync.RWMutex
    upgrader  websocket.Upgrader
}

func NewCollaborationServer() *CollaborationServer {
    return &CollaborationServer{
        documents: make(map[string]*Document),
        clients:   make(map[string]*Client),
        upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool { return true },
        },
    }
}

func (cs *CollaborationServer) CreateDocument(id string) *Document {
    cs.mutex.Lock()
    defer cs.mutex.Unlock()
    
    doc := &Document{
        ID:      id,
        Content: "",
        Version: 0,
        Created: time.Now(),
        Updated: time.Now(),
    }
    
    cs.documents[id] = doc
    return doc
}

func (cs *CollaborationServer) GetDocument(id string) (*Document, bool) {
    cs.mutex.RLock()
    defer cs.mutex.RUnlock()
    
    doc, exists := cs.documents[id]
    return doc, exists
}

func (cs *CollaborationServer) ApplyOperation(docID string, op *Operation) error {
    cs.mutex.Lock()
    defer cs.mutex.Unlock()
    
    doc, exists := cs.documents[docID]
    if !exists {
        return fmt.Errorf("document not found")
    }
    
    // Применение операции к документу
    switch op.Type {
    case "insert":
        if op.Position >= 0 && op.Position <= len(doc.Content) {
            doc.Content = doc.Content[:op.Position] + op.Text + doc.Content[op.Position:]
        }
    case "delete":
        if op.Position >= 0 && op.Position+len(op.Text) <= len(doc.Content) {
            doc.Content = doc.Content[:op.Position] + doc.Content[op.Position+len(op.Text):]
        }
    }
    
    doc.Version++
    doc.Updated = time.Now()
    
    // Рассылка операции всем клиентам
    cs.broadcastToDocumentClients(docID, op)
    
    return nil
}

func (cs *CollaborationServer) RegisterClient(client *Client) {
    cs.mutex.Lock()
    defer cs.mutex.Unlock()
    
    cs.clients[client.ID] = client
    log.Printf("Client %s connected to document %s", client.ID, client.Document)
}

func (cs *CollaborationServer) UnregisterClient(clientID string) {
    cs.mutex.Lock()
    defer cs.mutex.Unlock()
    
    if client, exists := cs.clients[clientID]; exists {
        close(client.Send)
        delete(cs.clients, clientID)
        log.Printf("Client %s disconnected", clientID)
    }
}

func (cs *CollaborationServer) broadcastToDocumentClients(docID string, op *Operation) {
    message, err := json.Marshal(op)
    if err != nil {
        log.Printf("Error marshaling operation: %v", err)
        return
    }
    
    for _, client := range cs.clients {
        if client.Document == docID {
            select {
            case client.Send <- message:
            default:
                close(client.Send)
                delete(cs.clients, client.ID)
            }
        }
    }
}

func (cs *CollaborationServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := cs.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("WebSocket upgrade error: %v", err)
        return
    }
    
    // Параметры из query string
    query := r.URL.Query()
    docID := query.Get("doc")
    clientID := query.Get("client")
    
    if docID == "" || clientID == "" {
        conn.Close()
        return
    }
    
    // Создание или получение документа
    doc, exists := cs.GetDocument(docID)
    if !exists {
        doc = cs.CreateDocument(docID)
    }
    
    // Создание клиента
    client := &Client{
        ID:       clientID,
        Document: docID,
        Conn:     conn,
        Send:     make(chan []byte, 256),
    }
    
    cs.RegisterClient(client)
    
    // Запуск горутин для чтения и записи
    go cs.writePump(client)
    go cs.readPump(client)
    
    // Отправка текущего состояния документа
    initialMessage := map[string]interface{}{
        "type":    "document_state",
        "content": doc.Content,
        "version": doc.Version,
    }
    
    message, _ := json.Marshal(initialMessage)
    client.Send <- message
}

func (cs *CollaborationServer) writePump(client *Client) {
    defer func() {
        client.Conn.Close()
        cs.UnregisterClient(client.ID)
    }()
    
    for {
        select {
        case message, ok := <-client.Send:
            if !ok {
                client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            
            err := client.Conn.WriteMessage(websocket.TextMessage, message)
            if err != nil {
                return
            }
        }
    }
}

func (cs *CollaborationServer) readPump(client *Client) {
    defer func() {
        client.Conn.Close()
        cs.UnregisterClient(client.ID)
    }()
    
    for {
        _, message, err := client.Conn.ReadMessage()
        if err != nil {
            break
        }
        
        var op Operation
        if err := json.Unmarshal(message, &op); err != nil {
            log.Printf("Error parsing operation: %v", err)
            continue
        }
        
        op.ClientID = client.ID
        op.Timestamp = time.Now().UnixNano()
        
        // Применение операции
        if err := cs.ApplyOperation(client.Document, &op); err != nil {
            log.Printf("Error applying operation: %v", err)
        }
    }
}

func (cs *CollaborationServer) GetStats() map[string]interface{} {
    cs.mutex.RLock()
    defer cs.mutex.RUnlock()
    
    return map[string]interface{}{
        "documents": len(cs.documents),
        "clients":   len(cs.clients),
    }
}

func main() {
    server := NewCollaborationServer()
    
    // HTTP handlers
    http.HandleFunc("/ws", server.HandleWebSocket)
    http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
        stats := server.GetStats()
        json.NewEncoder(w).Encode(stats)
    })
    
    http.HandleFunc("/documents/", func(w http.ResponseWriter, r *http.Request) {
        docID := r.URL.Path[len("/documents/"):]
        
        switch r.Method {
        case "GET":
            if doc, exists := server.GetDocument(docID); exists {
                json.NewEncoder(w).Encode(doc)
            } else {
                http.NotFound(w, r)
            }
        case "POST":
            var doc Document
            if err := json.NewDecoder(r.Body).Decode(&doc); err == nil {
                server.CreateDocument(docID)
                w.WriteHeader(http.StatusCreated)
            }
        }
    })
    
    // Статический файл с демо
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
    log.Println("Collaboration server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
//Задание: Real-time collaboration сервер для совместного редактирования документов
```
Машинное обучение и AI
147. Линейная регрессия
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type LinearRegression struct {
    weights []float64
    bias    float64
    lr      float64 // learning rate
}

func NewLinearRegression(features int, learningRate float64) *LinearRegression {
    weights := make([]float64, features)
    for i := range weights {
        weights[i] = rand.Float64()*2 - 1 // Инициализация случайными значениями
    }
    
    return &LinearRegression{
        weights: weights,
        bias:    rand.Float64()*2 - 1,
        lr:      learningRate,
    }
}

func (lr *LinearRegression) Predict(features []float64) float64 {
    if len(features) != len(lr.weights) {
        panic("Features count doesn't match weights count")
    }
    
    prediction := lr.bias
    for i, feature := range features {
        prediction += feature * lr.weights[i]
    }
    return prediction
}

func (lr *LinearRegression) Train(features [][]float64, targets []float64, epochs int) {
    if len(features) != len(targets) {
        panic("Features and targets must have same length")
    }
    
    for epoch := 0; epoch < epochs; epoch++ {
        totalError := 0.0
        
        for i, sample := range features {
            prediction := lr.Predict(sample)
            error := prediction - targets[i]
            totalError += error * error
            
            // Градиентный спуск
            lr.bias -= lr.lr * error
            for j := range lr.weights {
                lr.weights[j] -= lr.lr * error * sample[j]
            }
        }
        
        if epoch%100 == 0 {
            mse := totalError / float64(len(features))
            fmt.Printf("Epoch %d, MSE: %.4f\n", epoch, mse)
        }
    }
}

func (lr *LinearRegression) RMSE(features [][]float64, targets []float64) float64 {
    totalError := 0.0
    for i, sample := range features {
        prediction := lr.Predict(sample)
        error := prediction - targets[i]
        totalError += error * error
    }
    return math.Sqrt(totalError / float64(len(features)))
}

func (lr *LinearRegression) R2Score(features [][]float64, targets []float64) float64 {
    meanTarget := 0.0
    for _, target := range targets {
        meanTarget += target
    }
    meanTarget /= float64(len(targets))
    
    totalSumSquares := 0.0
    residualSumSquares := 0.0
    
    for i, sample := range features {
        prediction := lr.Predict(sample)
        totalSumSquares += (targets[i] - meanTarget) * (targets[i] - meanTarget)
        residualSumSquares += (targets[i] - prediction) * (targets[i] - prediction)
    }
    
    return 1 - (residualSumSquares / totalSumSquares)
}

func main() {
    // Генерация синтетических данных
    rand.Seed(42)
    nSamples := 1000
    nFeatures := 3
    
    features := make([][]float64, nSamples)
    targets := make([]float64, nSamples)
    
    // Истинные веса для генерации данных
    trueWeights := []float64{2.5, -1.8, 0.7}
    trueBias := 3.2
    
    for i := 0; i < nSamples; i++ {
        sample := make([]float64, nFeatures)
        for j := 0; j < nFeatures; j++ {
            sample[j] = rand.Float64()*10 - 5 // Значения от -5 до 5
        }
        features[i] = sample
        
        // Генерация целевой переменной с небольшим шумом
        target := trueBias
        for j := 0; j < nFeatures; j++ {
            target += sample[j] * trueWeights[j]
        }
        target += rand.NormFloat64() * 0.5 // Добавление шума
        targets[i] = target
    }
    
    // Разделение на обучающую и тестовую выборки
    splitIdx := int(0.8 * float64(nSamples))
    trainFeatures := features[:splitIdx]
    trainTargets := targets[:splitIdx]
    testFeatures := features[splitIdx:]
    testTargets := targets[splitIdx:]
    
    // Создание и обучение модели
    model := NewLinearRegression(nFeatures, 0.01)
    fmt.Println("Training linear regression...")
    model.Train(trainFeatures, trainTargets, 1000)
    
    // Оценка модели
    trainRMSE := model.RMSE(trainFeatures, trainTargets)
    testRMSE := model.RMSE(testFeatures, testTargets)
    r2 := model.R2Score(testFeatures, testTargets)
    
    fmt.Printf("\nModel Evaluation:\n")
    fmt.Printf("Train RMSE: %.4f\n", trainRMSE)
    fmt.Printf("Test RMSE: %.4f\n", testRMSE)
    fmt.Printf("R² Score: %.4f\n", r2)
    
    fmt.Printf("\nTrue weights: %v, bias: %.2f\n", trueWeights, trueBias)
    fmt.Printf("Learned weights: %v, bias: %.2f\n", model.weights, model.bias)
    
    // Прогнозирование на новых данных
    newSample := []float64{1.0, -2.0, 0.5}
    prediction := model.Predict(newSample)
    fmt.Printf("\nPrediction for %v: %.2f\n", newSample, prediction)
}
//Задание: Реализация линейной регрессии с градиентным спуском
```
148. K-ближайших соседей (KNN)
```go
package main
import (
    "fmt"
    "math"
    "sort"
)

type Point struct {
    Features []float64
    Label    string
    Distance float64
}

type KNN struct {
    k      int
    points []Point
}

func NewKNN(k int) *KNN {
    return &KNN{
        k: k,
    }
}

func (knn *KNN) Fit(points []Point) {
    knn.points = make([]Point, len(points))
    copy(knn.points, points)
}

func (knn *KNN) euclideanDistance(a, b []float64) float64 {
    if len(a) != len(b) {
        panic("Feature dimensions must match")
    }
    
    sum := 0.0
    for i := range a {
        diff := a[i] - b[i]
        sum += diff * diff
    }
    return math.Sqrt(sum)
}

func (knn *KNN) Predict(features []float64) string {
    // Вычисление расстояний до всех точек
    for i := range knn.points {
        knn.points[i].Distance = knn.euclideanDistance(knn.points[i].Features, features)
    }
    
    // Сортировка по расстоянию
    sortedPoints := make([]Point, len(knn.points))
    copy(sortedPoints, knn.points)
    sort.Slice(sortedPoints, func(i, j int) bool {
        return sortedPoints[i].Distance < sortedPoints[j].Distance
    })
    
    // Выбор k ближайших соседей
    kNeighbors := sortedPoints[:knn.k]
    
    // Голосование
    votes := make(map[string]int)
    for _, neighbor := range kNeighbors {
        votes[neighbor.Label]++
    }
    
    // Нахождение наиболее частого класса
    maxVotes := 0
    predictedLabel := ""
    for label, count := range votes {
        if count > maxVotes {
            maxVotes = count
            predictedLabel = label
        }
    }
    
    return predictedLabel
}

func (knn *KNN) PredictWithConfidence(features []float64) (string, float64) {
    for i := range knn.points {
        knn.points[i].Distance = knn.euclideanDistance(knn.points[i].Features, features)
    }
    
    sortedPoints := make([]Point, len(knn.points))
    copy(sortedPoints, knn.points)
    sort.Slice(sortedPoints, func(i, j int) bool {
        return sortedPoints[i].Distance < sortedPoints[j].Distance
    })
    
    kNeighbors := sortedPoints[:knn.k]
    
    votes := make(map[string]int)
    for _, neighbor := range kNeighbors {
        votes[neighbor.Label]++
    }
    
    maxVotes := 0
    predictedLabel := ""
    for label, count := range votes {
        if count > maxVotes {
            maxVotes = count
            predictedLabel = label
        }
    }
    
    confidence := float64(maxVotes) / float64(knn.k)
    return predictedLabel, confidence
}

func (knn *KNN) Accuracy(testPoints []Point) float64 {
    correct := 0
    for _, point := range testPoints {
        prediction := knn.Predict(point.Features)
        if prediction == point.Label {
            correct++
        }
    }
    return float64(correct) / float64(len(testPoints))
}

func main() {
    // Создание синтетических данных (классификация ирисов)
    points := []Point{
        {Features: []float64{5.1, 3.5, 1.4, 0.2}, Label: "setosa"},
        {Features: []float64{4.9, 3.0, 1.4, 0.2}, Label: "setosa"},
        {Features: []float64{4.7, 3.2, 1.3, 0.2}, Label: "setosa"},
        {Features: []float64{7.0, 3.2, 4.7, 1.4}, Label: "versicolor"},
        {Features: []float64{6.4, 3.2, 4.5, 1.5}, Label: "versicolor"},
        {Features: []float64{6.9, 3.1, 4.9, 1.5}, Label: "versicolor"},
        {Features: []float64{6.3, 3.3, 6.0, 2.5}, Label: "virginica"},
        {Features: []float64{5.8, 2.7, 5.1, 1.9}, Label: "virginica"},
        {Features: []float64{7.1, 3.0, 5.9, 2.1}, Label: "virginica"},
    }
    
    // Тестовые данные
    testPoints := []Point{
        {Features: []float64{5.0, 3.6, 1.4, 0.3}, Label: "setosa"},
        {Features: []float64{6.5, 3.0, 4.6, 1.5}, Label: "versicolor"},
        {Features: []float64{6.0, 2.7, 5.1, 1.6}, Label: "virginica"},
    }
    
    // Создание и обучение KNN
    knn := NewKNN(3)
    knn.Fit(points)
    
    fmt.Println("K-Nearest Neighbors Classification")
    fmt.Printf("Training points: %d\n", len(points))
    fmt.Printf("K: %d\n", knn.k)
    
    // Прогнозирование
    fmt.Println("\nPredictions:")
    for i, testPoint := range testPoints {
        prediction, confidence := knn.PredictWithConfidence(testPoint.Features)
        fmt.Printf("Test %d: Predicted: %s, Actual: %s, Confidence: %.2f\n", 
            i+1, prediction, testPoint.Label, confidence)
    }
    
    // Оценка точности
    accuracy := knn.Accuracy(testPoints)
    fmt.Printf("\nAccuracy: %.2f%%\n", accuracy*100)
    
    // Поиск оптимального K
    fmt.Println("\nFinding optimal K:")
    for k := 1; k <= 5; k++ {
        testKNN := NewKNN(k)
        testKNN.Fit(points)
        acc := testKNN.Accuracy(testPoints)
        fmt.Printf("K=%d, Accuracy=%.2f%%\n", k, acc*100)
    }
}
//Задание: Алгоритм K-ближайших соседей для классификации
```
149. Дерево решений
```go
package main
import (
    "fmt"
    "math"
)

type DecisionTreeNode struct {
    IsLeaf     bool
    Class      string
    FeatureIdx int
    Threshold  float64
    Left       *DecisionTreeNode
    Right      *DecisionTreeNode
}

type DecisionTree struct {
    Root       *DecisionTreeNode
    MaxDepth   int
    MinSamples int
}

type Dataset struct {
    Features [][]float64
    Labels   []string
}

func NewDecisionTree(maxDepth, minSamples int) *DecisionTree {
    return &DecisionTree{
        MaxDepth:   maxDepth,
        MinSamples: minSamples,
    }
}

func (dt *DecisionTree) Fit(features [][]float64, labels []string) {
    dataset := Dataset{Features: features, Labels: labels}
    dt.Root = dt.buildTree(dataset, 0)
}

func (dt *DecisionTree) buildTree(dataset Dataset, depth int) *DecisionTreeNode {
    // Критерии остановки
    if depth >= dt.MaxDepth || len(dataset.Labels) <= dt.MinSamples || dt.isPure(dataset.Labels) {
        return &DecisionTreeNode{
            IsLeaf: true,
            Class:  dt.majorityClass(dataset.Labels),
        }
    }
    
    // Поиск лучшего разделения
    bestFeature, bestThreshold, bestGain := dt.findBestSplit(dataset)
    if bestGain == 0 {
        return &DecisionTreeNode{
            IsLeaf: true,
            Class:  dt.majorityClass(dataset.Labels),
        }
    }
    
    // Разделение данных
    leftData, rightData := dt.splitDataset(dataset, bestFeature, bestThreshold)
    
    // Рекурсивное построение дерева
    node := &DecisionTreeNode{
        IsLeaf:     false,
        FeatureIdx: bestFeature,
        Threshold:  bestThreshold,
    }
    
    node.Left = dt.buildTree(leftData, depth+1)
    node.Right = dt.buildTree(rightData, depth+1)
    
    return node
}

func (dt *DecisionTree) isPure(labels []string) bool {
    if len(labels) == 0 {
        return true
    }
    first := labels[0]
    for _, label := range labels[1:] {
        if label != first {
            return false
        }
    }
    return true
}

func (dt *DecisionTree) majorityClass(labels []string) string {
    counts := make(map[string]int)
    for _, label := range labels {
        counts[label]++
    }
    
    maxCount := 0
    majority := ""
    for label, count := range counts {
        if count > maxCount {
            maxCount = count
            majority = label
        }
    }
    return majority
}

func (dt *DecisionTree) entropy(labels []string) float64 {
    if len(labels) == 0 {
        return 0
    }
    
    counts := make(map[string]int)
    for _, label := range labels {
        counts[label]++
    }
    
    entropy := 0.0
    for _, count := range counts {
        probability := float64(count) / float64(len(labels))
        entropy -= probability * math.Log2(probability)
    }
    return entropy
}

func (dt *DecisionTree) informationGain(parent []string, left []string, right []string) float64 {
    parentEntropy := dt.entropy(parent)
    
    weightLeft := float64(len(left)) / float64(len(parent))
    weightRight := float64(len(right)) / float64(len(parent))
    
    childrenEntropy := weightLeft*dt.entropy(left) + weightRight*dt.entropy(right)
    return parentEntropy - childrenEntropy
}

func (dt *DecisionTree) findBestSplit(dataset Dataset) (int, float64, float64) {
    bestGain := 0.0
    bestFeature := -1
    bestThreshold := 0.0
    
    nFeatures := len(dataset.Features[0])
    
    for featureIdx := 0; featureIdx < nFeatures; featureIdx++ {
        // Получение уникальных значений для признака
        values := make([]float64, len(dataset.Features))
        for i, sample := range dataset.Features {
            values[i] = sample[featureIdx]
        }
        
        // Попробовать различные пороги
        for _, threshold := range values {
            leftLabels, rightLabels := dt.splitLabels(dataset, featureIdx, threshold)
            
            gain := dt.informationGain(dataset.Labels, leftLabels, rightLabels)
            if gain > bestGain {
                bestGain = gain
                bestFeature = featureIdx
                bestThreshold = threshold
            }
        }
    }
    
    return bestFeature, bestThreshold, bestGain
}

func (dt *DecisionTree) splitLabels(dataset Dataset, featureIdx int, threshold float64) ([]string, []string) {
    var leftLabels, rightLabels []string
    
    for i, sample := range dataset.Features {
        if sample[featureIdx] <= threshold {
            leftLabels = append(leftLabels, dataset.Labels[i])
        } else {
            rightLabels = append(rightLabels, dataset.Labels[i])
        }
    }
    
    return leftLabels, rightLabels
}

func (dt *DecisionTree) splitDataset(dataset Dataset, featureIdx int, threshold float64) (Dataset, Dataset) {
    var leftData, rightData Dataset
    
    for i, sample := range dataset.Features {
        newSample := make([]float64, len(sample))
        copy(newSample, sample)
        
        if sample[featureIdx] <= threshold {
            leftData.Features = append(leftData.Features, newSample)
            leftData.Labels = append(leftData.Labels, dataset.Labels[i])
        } else {
            rightData.Features = append(rightData.Features, newSample)
            rightData.Labels = append(rightData.Labels, dataset.Labels[i])
        }
    }
    
    return leftData, rightData
}

func (dt *DecisionTree) Predict(features []float64) string {
    return dt.traverseTree(dt.Root, features)
}

func (dt *DecisionTree) traverseTree(node *DecisionTreeNode, features []float64) string {
    if node.IsLeaf {
        return node.Class
    }
    
    if features[node.FeatureIdx] <= node.Threshold {
        return dt.traverseTree(node.Left, features)
    } else {
        return dt.traverseTree(node.Right, features)
    }
}

func (dt *DecisionTree) Print() {
    dt.printNode(dt.Root, 0)
}

func (dt *DecisionTree) printNode(node *DecisionTreeNode, depth int) {
    indent := ""
    for i := 0; i < depth; i++ {
        indent += "  "
    }
    
    if node.IsLeaf {
        fmt.Printf("%sLeaf: %s\n", indent, node.Class)
    } else {
        fmt.Printf("%sFeature[%d] <= %.2f\n", indent, node.FeatureIdx, node.Threshold)
        dt.printNode(node.Left, depth+1)
        dt.printNode(node.Right, depth+1)
    }
}

func main() {
    // Данные для классификации (погода -> играть в теннис)
    features := [][]float64{
        {1, 1, 1, 1}, // солнечно, жарко, высокая, слабый
        {1, 1, 1, 2}, // солнечно, жарко, высокая, сильный
        {2, 1, 1, 1}, // пасмурно, жарко, высокая, слабый
        {3, 2, 1, 1}, // дождь, умеренно, высокая, слабый
        {3, 3, 2, 1}, // дождь, холодно, нормальная, слабый
        {3, 3, 2, 2}, // дождь, холодно, нормальная, сильный
        {2, 3, 2, 2}, // пасмурно, холодно, нормальная, сильный
        {1, 2, 1, 1}, // солнечно, умеренно, высокая, слабый
        {1, 3, 2, 1}, // солнечно, холодно, нормальная, слабый
        {3, 2, 2, 1}, // дождь, умеренно, нормальная, слабый
        {1, 2, 2, 2}, // солнечно, умеренно, нормальная, сильный
        {2, 2, 1, 2}, // пасмурно, умеренно, высокая, сильный
        {2, 1, 2, 1}, // пасмурно, жарко, нормальная, слабый
        {3, 2, 1, 2}, // дождь, умеренно, высокая, сильный
    }
    
    labels := []string{
        "no", "no", "yes", "yes", "yes", "no", "yes", "no", "yes", "yes", "yes", "yes", "yes", "no",
    }
    
    // Создание и обучение дерева решений
    tree := NewDecisionTree(5, 2)
    tree.Fit(features, labels)
    
    fmt.Println("Decision Tree Structure:")
    tree.Print()
    
    // Прогнозирование
    testCases := [][]float64{
        {1, 1, 1, 1}, // солнечно, жарко, высокая, слабый -> no
        {3, 2, 2, 1}, // дождь, умеренно, нормальная, слабый -> yes
        {2, 2, 2, 2}, // пасмурно, умеренно, нормальная, сильный -> ?
    }
    
    fmt.Println("\nPredictions:")
    for i, testCase := range testCases {
        prediction := tree.Predict(testCase)
        fmt.Printf("Test case %d: %v -> %s\n", i+1, testCase, prediction)
    }
    
    // Оценка точности на обучающих данных
    correct := 0
    for i, sample := range features {
        prediction := tree.Predict(sample)
        if prediction == labels[i] {
            correct++
        }
    }
    accuracy := float64(correct) / float64(len(features))
    fmt.Printf("\nTraining accuracy: %.2f%%\n", accuracy*100)
}
//Задание: Дерево решений для классификации с использованием Information Gain
```
150. K-средних (K-means) кластеризация
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type Point struct {
    X, Y float64
    Cluster int
}

type KMeans struct {
    K        int
    Points   []Point
    Centroids []Point
    MaxIter  int
}

func NewKMeans(k int, maxIter int) *KMeans {
    return &KMeans{
        K:       k,
        MaxIter: maxIter,
    }
}

func (km *KMeans) Initialize(points []Point) {
    km.Points = make([]Point, len(points))
    copy(km.Points, points)
    
    // Инициализация центроидов случайными точками
    km.Centroids = make([]Point, km.K)
    used := make(map[int]bool)
    
    for i := 0; i < km.K; i++ {
        for {
            idx := rand.Intn(len(km.Points))
            if !used[idx] {
                km.Centroids[i] = Point{
                    X: km.Points[idx].X,
                    Y: km.Points[idx].Y,
                    Cluster: i,
                }
                used[idx] = true
                break
            }
        }
    }
}

func (km *KMeans) EuclideanDistance(p1, p2 Point) float64 {
    dx := p1.X - p2.X
    dy := p1.Y - p2.Y
    return math.Sqrt(dx*dx + dy*dy)
}

func (km *KMeans) AssignClusters() bool {
    changed := false
    
    for i := range km.Points {
        minDist := math.MaxFloat64
        closestCluster := -1
        
        for j, centroid := range km.Centroids {
            dist := km.EuclideanDistance(km.Points[i], centroid)
            if dist < minDist {
                minDist = dist
                closestCluster = j
            }
        }
        
        if km.Points[i].Cluster != closestCluster {
            km.Points[i].Cluster = closestCluster
            changed = true
        }
    }
    
    return changed
}

func (km *KMeans) UpdateCentroids() {
    // Сброс центроидов
    for i := range km.Centroids {
        km.Centroids[i].X = 0
        km.Centroids[i].Y = 0
    }
    
    // Суммирование координат по кластерам
    counts := make([]int, km.K)
    for _, point := range km.Points {
        cluster := point.Cluster
        km.Centroids[cluster].X += point.X
        km.Centroids[cluster].Y += point.Y
        counts[cluster]++
    }
    
    // Вычисление новых центроидов
    for i := range km.Centroids {
        if counts[i] > 0 {
            km.Centroids[i].X /= float64(counts[i])
            km.Centroids[i].Y /= float64(counts[i])
        }
    }
}

func (km *KMeans) Fit() {
    for iter := 0; iter < km.MaxIter; iter++ {
        changed := km.AssignClusters()
        km.UpdateCentroids()
        
        if !changed {
            fmt.Printf("Converged after %d iterations\n", iter+1)
            break
        }
        
        if iter == km.MaxIter-1 {
            fmt.Printf("Reached maximum iterations: %d\n", km.MaxIter)
        }
    }
}

func (km *KMeans) WCSS() float64 {
    total := 0.0
    for _, point := range km.Points {
        centroid := km.Centroids[point.Cluster]
        dist := km.EuclideanDistance(point, centroid)
        total += dist * dist
    }
    return total
}

func (km *KMeans) SilhouetteScore() float64 {
    if km.K <= 1 {
        return 0
    }
    
    totalScore := 0.0
    
    for i, point := range km.Points {
        // Среднее расстояние до точек в своем кластере
        a := 0.0
        countA := 0
        
        for j, other := range km.Points {
            if i != j && point.Cluster == other.Cluster {
                dist := km.EuclideanDistance(point, other)
                a += dist
                countA++
            }
        }
        
        if countA > 0 {
            a /= float64(countA)
        }
        
        // Среднее расстояние до точек в ближайшем соседнем кластере
        b := math.MaxFloat64
        
        for cluster := 0; cluster < km.K; cluster++ {
            if cluster == point.Cluster {
                continue
            }
            
            clusterDist := 0.0
            countB := 0
            
            for _, other := range km.Points {
                if other.Cluster == cluster {
                    dist := km.EuclideanDistance(point, other)
                    clusterDist += dist
                    countB++
                }
            }
            
            if countB > 0 {
                clusterDist /= float64(countB)
                if clusterDist < b {
                    b = clusterDist
                }
            }
        }
        
        // Silhouette score для точки
        if a < b {
            totalScore += 1 - a/b
        } else if a > b {
            totalScore += b/a - 1
        }
        // если a == b, score = 0
    }
    
    return totalScore / float64(len(km.Points))
}

func (km *KMeans) PrintResults() {
    fmt.Printf("\nK-Means Clustering Results (K=%d)\n", km.K)
    fmt.Printf("Within-Cluster Sum of Squares: %.4f\n", km.WCSS())
    fmt.Printf("Silhouette Score: %.4f\n", km.SilhouetteScore())
    
    fmt.Println("\nCentroids:")
    for i, centroid := range km.Centroids {
        fmt.Printf("Cluster %d: (%.2f, %.2f)\n", i, centroid.X, centroid.Y)
    }
    
    fmt.Println("\nCluster sizes:")
    counts := make([]int, km.K)
    for _, point := range km.Points {
        counts[point.Cluster]++
    }
    for i, count := range counts {
        fmt.Printf("Cluster %d: %d points\n", i, count)
    }
}

func main() {
    // Генерация синтетических данных
    rand.Seed(42)
    nPoints := 300
    points := make([]Point, nPoints)
    
    // Генерация данных из трех кластеров
    for i := 0; i < nPoints; i++ {
        cluster := i % 3
        var x, y float64
        
        switch cluster {
        case 0:
            x = rand.NormFloat64()*0.5 + 2.0
            y = rand.NormFloat64()*0.5 + 2.0
        case 1:
            x = rand.NormFloat64()*0.5 + 5.0
            y = rand.NormFloat64()*0.5 + 5.0
        case 2:
            x = rand.NormFloat64()*0.5 + 8.0
            y = rand.NormFloat64()*0.5 + 2.0
        }
        
        points[i] = Point{X: x, Y: y, Cluster: -1}
    }
    
    // Поиск оптимального K с использованием метода локтя
    fmt.Println("Finding optimal K using elbow method:")
    bestK := 1
    bestScore := math.MaxFloat64
    
    for k := 1; k <= 6; k++ {
        kmeans := NewKMeans(k, 100)
        kmeans.Initialize(points)
        kmeans.Fit()
        wcss := kmeans.WCSS()
        
        fmt.Printf("K=%d, WCSS=%.4f\n", k, wcss)
        
        if k > 1 {
            // Простой метод локтя: ищем "изгиб" в графике WCSS
            improvement := bestScore - wcss
            if improvement < bestScore*0.1 { // Если улучшение менее 10%
                fmt.Printf("Elbow found at K=%d\n", k-1)
                break
            }
        }
        
        bestK = k
        bestScore = wcss
    }
    
    // Запуск с оптимальным K
    fmt.Printf("\nRunning K-Means with K=%d\n", bestK)
    kmeans := NewKMeans(bestK, 100)
    kmeans.Initialize(points)
    kmeans.Fit()
    kmeans.PrintResults()
    
    // Визуализация (псевдографика)
    fmt.Println("\nCluster visualization (approximate):")
    km.Visualize(10, 10)
}

func (km *KMeans) Visualize(width, height int) {
    // Нахождение границ данных
    minX, maxX := math.MaxFloat64, -math.MaxFloat64
    minY, maxY := math.MaxFloat64, -math.MaxFloat64
    
    for _, point := range km.Points {
        if point.X < minX {
            minX = point.X
        }
        if point.X > maxX {
            maxX = point.X
        }
        if point.Y < minY {
            minY = point.Y
        }
        if point.Y > maxY {
            maxY = point.Y
        }
    }
    
    // Создание сетки
    grid := make([][]int, height)
    for i := range grid {
        grid[i] = make([]int, width)
        for j := range grid[i] {
            grid[i][j] = -1
        }
    }
    
    // Отображение точек на сетку
    for _, point := range km.Points {
        x := int((point.X - minX) / (maxX - minX) * float64(width-1))
        y := int((point.Y - minY) / (maxY - minY) * float64(height-1))
        
        if x >= 0 && x < width && y >= 0 && y < height {
            grid[y][x] = point.Cluster
        }
    }
    
    // Вывод сетки
    symbols := []string{"·", "○", "●", "▲", "■", "★"}
    
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            cluster := grid[y][x]
            if cluster == -1 {
                fmt.Print(" ")
            } else {
                symbol := symbols[cluster%len(symbols)]
                fmt.Print(symbol)
            }
        }
        fmt.Println()
    }
}
//Задание: Алгоритм K-средних для кластеризации с методом локтя
```
151. Обработка естественного языка (NLP) - базовые методы
```go
package main
import (
    "fmt"
    "regexp"
    "sort"
    "strings"
    "unicode"
)

type TextProcessor struct {
    stopWords map[string]bool
}

func NewTextProcessor() *TextProcessor {
    stopWords := map[string]bool{
        "the": true, "a": true, "an": true, "and": true, "or": true, "but": true,
        "in": true, "on": true, "at": true, "to": true, "for": true, "of": true,
        "with": true, "by": true, "is": true, "are": true, "was": true, "were": true,
        "be": true, "been": true, "being": true, "have": true, "has": true, "had": true,
        "do": true, "does": true, "did": true, "will": true, "would": true, "could": true,
        "should": true, "may": true, "might": true, "must": true, "can": true,
    }
    
    return &TextProcessor{stopWords: stopWords}
}

func (tp *TextProcessor) CleanText(text string) string {
    // Приведение к нижнему регистру
    text = strings.ToLower(text)
    
    // Удаление специальных символов, оставляем только буквы и пробелы
    reg := regexp.MustCompile(`[^a-zA-Z\s]`)
    text = reg.ReplaceAllString(text, " ")
    
    // Удаление лишних пробелов
    reg = regexp.MustCompile(`\s+`)
    text = reg.ReplaceAllString(text, " ")
    
    return strings.TrimSpace(text)
}

func (tp *TextProcessor) Tokenize(text string) []string {
    return strings.Fields(text)
}

func (tp *TextProcessor) RemoveStopWords(tokens []string) []string {
    var filtered []string
    for _, token := range tokens {
        if !tp.stopWords[token] {
            filtered = append(filtered, token)
        }
    }
    return filtered
}

func (tp *TextProcessor) Stem(word string) string {
    // Простой стеммер (Porter stemmer упрощенный)
    if len(word) < 3 {
        return word
    }
    
    // Удаление окончаний
    suffixes := []string{"ing", "ed", "ly", "es", "s"}
    for _, suffix := range suffixes {
        if strings.HasSuffix(word, suffix) {
            word = word[:len(word)-len(suffix)]
            break
        }
    }
    
    return word
}

func (tp *TextProcessor) StemTokens(tokens []string) []string {
    stemmed := make([]string, len(tokens))
    for i, token := range tokens {
        stemmed[i] = tp.Stem(token)
    }
    return stemmed
}

type TFIDFVectorizer struct {
    vocabulary map[string]int
    idf        map[string]float64
    documents  [][]string
}

func NewTFIDFVectorizer() *TFIDFVectorizer {
    return &TFIDFVectorizer{
        vocabulary: make(map[string]int),
        idf:        make(map[string]float64),
    }
}

func (tv *TFIDFVectorizer) Fit(documents []string) {
    // Очистка и токенизация документов
    processor := NewTextProcessor()
    tv.documents = make([][]string, len(documents))
    
    for i, doc := range documents {
        cleaned := processor.CleanText(doc)
        tokens := processor.Tokenize(cleaned)
        tokens = processor.RemoveStopWords(tokens)
        tokens = processor.StemTokens(tokens)
        tv.documents[i] = tokens
    }
    
    // Построение словаря
    tv.buildVocabulary()
    
    // Вычисление IDF
    tv.calculateIDF()
}

func (tv *TFIDFVectorizer) buildVocabulary() {
    index := 0
    for _, tokens := range tv.documents {
        for _, token := range tokens {
            if _, exists := tv.vocabulary[token]; !exists {
                tv.vocabulary[token] = index
                index++
            }
        }
    }
}

func (tv *TFIDFVectorizer) calculateIDF() {
    n := float64(len(tv.documents))
    
    for term := range tv.vocabulary {
        docCount := 0
        for _, tokens := range tv.documents {
            for _, token := range tokens {
                if token == term {
                    docCount++
                    break
                }
            }
        }
        
        tv.idf[term] = math.Log(n / float64(docCount))
    }
}

func (tv *TFIDFVectorizer) Transform(documents []string) [][]float64 {
    processor := NewTextProcessor()
    vectors := make([][]float64, len(documents))
    
    for i, doc := range documents {
        // Предобработка документа
        cleaned := processor.CleanText(doc)
        tokens := processor.Tokenize(cleaned)
        tokens = processor.RemoveStopWords(tokens)
        tokens = processor.StemTokens(tokens)
        
        // Вычисление TF
        tf := make(map[string]float64)
        for _, token := range tokens {
            tf[token]++
        }
        
        // Нормализация TF
        totalTerms := float64(len(tokens))
        for term := range tf {
            tf[term] /= totalTerms
        }
        
        // Создание вектора TF-IDF
        vector := make([]float64, len(tv.vocabulary))
        for term, tfValue := range tf {
            if idfValue, exists := tv.idf[term]; exists {
                index := tv.vocabulary[term]
                vector[index] = tfValue * idfValue
            }
        }
        
        vectors[i] = vector
    }
    
    return vectors
}

func (tv *TFIDFVectorizer) CosineSimilarity(vec1, vec2 []float64) float64 {
    dotProduct := 0.0
    norm1 := 0.0
    norm2 := 0.0
    
    for i := range vec1 {
        dotProduct += vec1[i] * vec2[i]
        norm1 += vec1[i] * vec1[i]
        norm2 += vec2[i] * vec2[i]
    }
    
    if norm1 == 0 || norm2 == 0 {
        return 0
    }
    
    return dotProduct / (math.Sqrt(norm1) * math.Sqrt(norm2))
}

type NaiveBayesClassifier struct {
    classProbabilities map[string]float64
    featureProbabilities map[string]map[string]float64
    vocabulary map[string]bool
}

func NewNaiveBayesClassifier() *NaiveBayesClassifier {
    return &NaiveBayesClassifier{
        classProbabilities: make(map[string]float64),
        featureProbabilities: make(map[string]map[string]float64),
        vocabulary: make(map[string]bool),
    }
}

func (nb *NaiveBayesClassifier) Fit(documents []string, labels []string) {
    processor := NewTextProcessor()
    
    // Подсчет количества документов в каждом классе
    classCounts := make(map[string]int)
    for _, label := range labels {
        classCounts[label]++
    }
    
    // Вычисление априорных вероятностей
    totalDocs := float64(len(documents))
    for class, count := range classCounts {
        nb.classProbabilities[class] = float64(count) / totalDocs
    }
    
    // Подсчет частот слов для каждого класса
    wordCounts := make(map[string]map[string]int)
    for class := range nb.classProbabilities {
        wordCounts[class] = make(map[string]int)
    }
    
    totalWordsPerClass := make(map[string]int)
    
    for i, doc := range documents {
        class := labels[i]
        cleaned := processor.CleanText(doc)
        tokens := processor.Tokenize(cleaned)
        tokens = processor.RemoveStopWords(tokens)
        tokens = processor.StemTokens(tokens)
        
        for _, token := range tokens {
            wordCounts[class][token]++
            totalWordsPerClass[class]++
            nb.vocabulary[token] = true
        }
    }
    
    // Вычисление условных вероятностей (с лапласовским сглаживанием)
    vocabularySize := len(nb.vocabulary)
    
    for class := range nb.classProbabilities {
        nb.featureProbabilities[class] = make(map[string]float64)
        totalWords := totalWordsPerClass[class]
        
        for word := range nb.vocabulary {
            count := wordCounts[class][word]
            // Лапласовское сглаживание
            nb.featureProbabilities[class][word] = float64(count + 1) / float64(totalWords + vocabularySize)
        }
    }
}

func (nb *NaiveBayesClassifier) Predict(document string) string {
    processor := NewTextProcessor()
    cleaned := processor.CleanText(document)
    tokens := processor.Tokenize(cleaned)
    tokens = processor.RemoveStopWords(tokens)
    tokens = processor.StemTokens(tokens)
    
    bestClass := ""
    bestScore := math.Inf(-1)
    
    for class, prior := range nb.classProbabilities {
        score := math.Log(prior) // Логарифм для избежания underflow
        for _, token := range tokens {
            if prob, exists := nb.featureProbabilities[class][token]; exists {
                score += math.Log(prob)
            } else {
                // Если слово не встречалось в обучении, используем маленькую вероятность
                score += math.Log(1e-10)
            }
        }
        
        if score > bestScore {
            bestScore = score
            bestClass = class
        }
    }
    
    return bestClass
}

func main() {
    fmt.Println("=== Natural Language Processing Demo ===")
    
    // Демонстрация текстовой обработки
    processor := NewTextProcessor()
    text := "The quick brown fox jumps over the lazy dog. Running quickly, he escaped!"
    
    fmt.Println("Original text:", text)
    cleaned := processor.CleanText(text)
    fmt.Println("Cleaned text:", cleaned)
    
    tokens := processor.Tokenize(cleaned)
    fmt.Println("Tokens:", tokens)
    
    filtered := processor.RemoveStopWords(tokens)
    fmt.Println("Without stop words:", filtered)
    
    stemmed := processor.StemTokens(filtered)
    fmt.Println("Stemmed:", stemmed)
    
    // Демонстрация TF-IDF
    fmt.Println("\n=== TF-IDF Vectorization ===")
    documents := []string{
        "the cat sat on the mat",
        "the dog sat on the log",
        "cats and dogs are great pets",
        "the mat is on the floor",
    }
    
    vectorizer := NewTFIDFVectorizer()
    vectorizer.Fit(documents)
    
    testDoc := "the cat and dog play"
    vectors := vectorizer.Transform([]string{testDoc})
    fmt.Printf("TF-IDF vector for '%s': %v\n", testDoc, vectors[0])
    
    // Демонстрация классификатора
    fmt.Println("\n=== Naive Bayes Text Classification ===")
    
    trainingDocs := []string{
        "love this movie its great",
        "amazing film wonderful acting",
        "hate this movie terrible acting",
        "awful film boring story",
        "great movie fantastic plot",
        "terrible movie waste of time",
        "wonderful film amazing cinematography",
        "boring movie fell asleep",
    }
    
    trainingLabels := []string{
        "positive", "positive", "negative", "negative",
        "positive", "negative", "positive", "negative",
    }
    
    classifier := NewNaiveBayesClassifier()
    classifier.Fit(trainingDocs, trainingLabels)
    
    testTexts := []string{
        "great amazing wonderful film",
        "terrible awful boring movie",
        "not bad actually quite good",
    }
    
    for _, testText := range testTexts {
        prediction := classifier.Predict(testText)
        fmt.Printf("Text: '%s' -> %s\n", testText, prediction)
    }
}
//Задание: Базовые методы NLP: обработка текста, TF-IDF, наивный Байес
```
152. Компьютерное зрение - базовые фильтры
```go
package main
import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"
    "math"
    "os"
)

type ImageProcessor struct{}

func NewImageProcessor() *ImageProcessor {
    return &ImageProcessor{}
}

func (ip *ImageProcessor) LoadImage(filename string) (image.Image, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    
    return img, nil
}

func (ip *ImageProcessor) SaveImage(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    switch {
    case len(filename) > 4 && filename[len(filename)-4:] == ".png":
        return png.Encode(file, img)
    case len(filename) > 4 && filename[len(filename)-4:] == ".jpg":
        return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
    default:
        return png.Encode(file, img)
    }
}

func (ip *ImageProcessor) Grayscale(img image.Image) *image.Gray {
    bounds := img.Bounds()
    gray := image.NewGray(bounds)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            originalColor := img.At(x, y)
            grayColor := color.GrayModel.Convert(originalColor)
            gray.Set(x, y, grayColor)
        }
    }
    
    return gray
}

func (ip *ImageProcessor) ApplyConvolution(img image.Image, kernel [][]float64) *image.Gray {
    bounds := img.Bounds()
    gray := ip.Grayscale(img)
    result := image.NewGray(bounds)
    
    kernelSize := len(kernel)
    offset := kernelSize / 2
    
    for y := bounds.Min.Y + offset; y < bounds.Max.Y - offset; y++ {
        for x := bounds.Min.X + offset; x < bounds.Max.X - offset; x++ {
            var sum float64
            
            for ky := -offset; ky <= offset; ky++ {
                for kx := -offset; kx <= offset; kx++ {
                    pixel := gray.GrayAt(x + kx, y + ky)
                    kernelValue := kernel[ky + offset][kx + offset]
                    sum += float64(pixel.Y) * kernelValue
                }
            }
            
            // Нормализация
            sum = math.Max(0, math.Min(255, sum))
            result.SetGray(x, y, color.Gray{Y: uint8(sum)})
        }
    }
    
    return result
}

func (ip *ImageProcessor) GaussianBlur(img image.Image, radius int, sigma float64) *image.Gray {
    size := 2*radius + 1
    kernel := make([][]float64, size)
    
    // Создание гауссова ядра
    sum := 0.0
    for y := -radius; y <= radius; y++ {
        kernel[y + radius] = make([]float64, size)
        for x := -radius; x <= radius; x++ {
            exponent := -(float64(x*x + y*y) / (2 * sigma * sigma))
            value := math.Exp(exponent)
            kernel[y + radius][x + radius] = value
            sum += value
        }
    }
    
    // Нормализация ядра
    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            kernel[y][x] /= sum
        }
    }
    
    return ip.ApplyConvolution(img, kernel)
}

func (ip *ImageProcessor) SobelEdgeDetection(img image.Image) *image.Gray {
    // Ядра Собеля
    sobelX := [][]float64{
        {-1, 0, 1},
        {-2, 0, 2},
        {-1, 0, 1},
    }
    
    sobelY := [][]float64{
        {-1, -2, -1},
        {0, 0, 0},
        {1, 2, 1},
    }
    
    bounds := img.Bounds()
    gray := ip.Grayscale(img)
    result := image.NewGray(bounds)
    
    for y := 1; y < bounds.Max.Y - 1; y++ {
        for x := 1; x < bounds.Max.X - 1; x++ {
            var gx, gy float64
            
            // Применение ядер
            for ky := -1; ky <= 1; ky++ {
                for kx := -1; kx <= 1; kx++ {
                    pixel := gray.GrayAt(x + kx, y + ky)
                    value := float64(pixel.Y)
                    
                    gx += value * sobelX[ky + 1][kx + 1]
                    gy += value * sobelY[ky + 1][kx + 1]
                }
            }
            
            // Вычисление градиента
            gradient := math.Sqrt(gx*gx + gy*gy)
            gradient = math.Min(255, gradient)
            
            result.SetGray(x, y, color.Gray{Y: uint8(gradient)})
        }
    }
    
    return result
}

func (ip *ImageProcessor) CannyEdgeDetection(img image.Image, lowThreshold, highThreshold float64) *image.Gray {
    // 1. Размытие по Гауссу
    blurred := ip.GaussianBlur(img, 2, 1.4)
    
    // 2. Градиенты Собеля
    bounds := blurred.Bounds()
    gradient := image.NewGray(bounds)
    direction := make([][]float64, bounds.Dy())
    for i := range direction {
        direction[i] = make([]float64, bounds.Dx())
    }
    
    sobelX := [][]float64{{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1}}
    sobelY := [][]float64{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}
    
    for y := 1; y < bounds.Max.Y - 1; y++ {
        for x := 1; x < bounds.Max.X - 1; x++ {
            var gx, gy float64
            
            for ky := -1; ky <= 1; ky++ {
                for kx := -1; kx <= 1; kx++ {
                    pixel := blurred.GrayAt(x + kx, y + ky)
                    value := float64(pixel.Y)
                    
                    gx += value * sobelX[ky + 1][kx + 1]
                    gy += value * sobelY[ky + 1][kx + 1]
                }
            }
            
            grad := math.Sqrt(gx*gx + gy*gy)
            gradient.SetGray(x, y, color.Gray{Y: uint8(math.Min(255, grad))})
            
            // Направление градиента
            angle := math.Atan2(gy, gx) * 180 / math.Pi
            if angle < 0 {
                angle += 180
            }
            direction[y][x] = angle
        }
    }
    
    // 3. Подавление немаксимумов
    suppressed := image.NewGray(bounds)
    
    for y := 2; y < bounds.Max.Y - 2; y++ {
        for x := 2; x < bounds.Max.X - 2; x++ {
            angle := direction[y][x]
            var p1, p2 float64
            
            // Определение соседей в направлении градиента
            if (0 <= angle && angle < 22.5) || (157.5 <= angle && angle <= 180) {
                // Горизонтальное направление
                p1 = float64(gradient.GrayAt(x + 1, y).Y)
                p2 = float64(gradient.GrayAt(x - 1, y).Y)
            } else if 22.5 <= angle && angle < 67.5 {
                // Диагональ 45°
                p1 = float64(gradient.GrayAt(x + 1, y - 1).Y)
                p2 = float64(gradient.GrayAt(x - 1, y + 1).Y)
            } else if 67.5 <= angle && angle < 112.5 {
                // Вертикальное направление
                p1 = float64(gradient.GrayAt(x, y - 1).Y)
                p2 = float64(gradient.GrayAt(x, y + 1).Y)
            } else if 112.5 <= angle && angle < 157.5 {
                // Диагональ 135°
                p1 = float64(gradient.GrayAt(x - 1, y - 1).Y)
                p2 = float64(gradient.GrayAt(x + 1, y + 1).Y)
            }
            
            current := float64(gradient.GrayAt(x, y).Y)
            if current >= p1 && current >= p2 {
                suppressed.SetGray(x, y, color.Gray{Y: uint8(current)})
            } else {
                suppressed.SetGray(x, y, color.Gray{Y: 0})
            }
        }
    }
    
    // 4. Двойная пороговая фильтрация
    result := image.NewGray(bounds)
    
    for y := 0; y < bounds.Max.Y; y++ {
        for x := 0; x < bounds.Max.X; x++ {
            value := float64(suppressed.GrayAt(x, y).Y)
            
            if value >= highThreshold {
                result.SetGray(x, y, color.Gray{Y: 255}) // Сильный край
            } else if value >= lowThreshold {
                result.SetGray(x, y, color.Gray{Y: 128}) // Слабый край
            } else {
                result.SetGray(x, y, color.Gray{Y: 0}) // Не край
            }
        }
    }
    
    return result
}

func (ip *ImageProcessor) HistogramEqualization(img image.Image) *image.Gray {
    bounds := img.Bounds()
    gray := ip.Grayscale(img)
    
    // Вычисление гистограммы
    histogram := [256]int{}
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            pixel := gray.GrayAt(x, y)
            histogram[pixel.Y]++
        }
    }
    
    // Вычисление кумулятивной гистограммы
    cumulative := [256]int{}
    cumulative[0] = histogram[0]
    for i := 1; i < 256; i++ {
        cumulative[i] = cumulative[i-1] + histogram[i]
    }
    
    // Нормализация
    totalPixels := bounds.Dx() * bounds.Dy()
    result := image.NewGray(bounds)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            pixel := gray.GrayAt(x, y)
            newValue := cumulative[pixel.Y] * 255 / totalPixels
            result.SetGray(x, y, color.Gray{Y: uint8(newValue)})
        }
    }
    
    return result
}

func main() {
    processor := NewImageProcessor()
    
    // Создание тестового изображения
    img := image.NewRGBA(image.Rect(0, 0, 400, 400))
    
    // Рисование тестовых фигур
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
    
    // Круг
    for y := 100; y < 300; y++ {
        for x := 100; x < 300; x++ {
            dx, dy := x-200, y-200
            if dx*dx+dy*dy <= 100*100 {
                img.Set(x, y, color.Black)
            }
        }
    }
    
    // Прямоугольник
    for y := 50; y < 150; y++ {
        for x := 50; x < 150; x++ {
            img.Set(x, y, color.Gray{Y: 100})
        }
    }
    
    // Сохранение оригинального изображения
    processor.SaveImage(img, "original.png")
    fmt.Println("Created test image: original.png")
    
    // Применение фильтров
    fmt.Println("Applying image filters...")
    
    // Grayscale
    gray := processor.Grayscale(img)
    processor.SaveImage(gray, "grayscale.png")
    fmt.Println("Grayscale saved: grayscale.png")
    
    // Gaussian blur
    blurred := processor.GaussianBlur(img, 3, 2.0)
    processor.SaveImage(blurred, "blurred.png")
    fmt.Println("Gaussian blur saved: blurred.png")
    
    // Sobel edge detection
    sobel := processor.SobelEdgeDetection(img)
    processor.SaveImage(sobel, "sobel_edges.png")
    fmt.Println("Sobel edges saved: sobel_edges.png")
    
    // Canny edge detection
    canny := processor.CannyEdgeDetection(img, 50, 150)
    processor.SaveImage(canny, "canny_edges.png")
    fmt.Println("Canny edges saved: canny_edges.png")
    
    // Histogram equalization
    equalized := processor.HistogramEqualization(img)
    processor.SaveImage(equalized, "equalized.png")
    fmt.Println("Histogram equalization saved: equalized.png")
    
    fmt.Println("\nAll image processing operations completed!")
}
//Задание: Базовые операции компьютерного зрения: фильтры, детекция краев
```
Глубокое обучение и нейронные сети
153. Полносвязная нейронная сеть
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type Layer struct {
    Weights [][]float64
    Biases  []float64
    Outputs []float64
    Deltas  []float64
}

type NeuralNetwork struct {
    Layers []*Layer
    LR     float64 // learning rate
}

func NewNeuralNetwork(layers []int, learningRate float64) *NeuralNetwork {
    nn := &NeuralNetwork{LR: learningRate}
    
    for i := 1; i < len(layers); i++ {
        layer := &Layer{
            Weights: make([][]float64, layers[i]),
            Biases:  make([]float64, layers[i]),
            Outputs: make([]float64, layers[i]),
            Deltas:  make([]float64, layers[i]),
        }
        
        // Инициализация весов (Xavier initialization)
        for j := range layer.Weights {
            layer.Weights[j] = make([]float64, layers[i-1])
            for k := range layer.Weights[j] {
                layer.Weights[j][k] = rand.NormFloat64() * math.Sqrt(2.0/float64(layers[i-1]))
            }
            layer.Biases[j] = 0.1
        }
        
        nn.Layers = append(nn.Layers, layer)
    }
    
    return nn
}

func sigmoid(x float64) float64 {
    return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
    return x * (1.0 - x)
}

func relu(x float64) float64 {
    if x > 0 {
        return x
    }
    return 0
}

func reluDerivative(x float64) float64 {
    if x > 0 {
        return 1
    }
    return 0
}

func (nn *NeuralNetwork) Forward(input []float64) []float64 {
    current := input
    
    for _, layer := range nn.Layers {
        next := make([]float64, len(layer.Weights))
        
        for i, neuron := range layer.Weights {
            sum := layer.Biases[i]
            for j, weight := range neuron {
                sum += weight * current[j]
            }
            // ReLU для скрытых слоев, sigmoid для выходного
            if layer != nn.Layers[len(nn.Layers)-1] {
                next[i] = relu(sum)
            } else {
                next[i] = sigmoid(sum)
            }
            layer.Outputs[i] = next[i]
        }
        
        current = next
    }
    
    return current
}

func (nn *NeuralNetwork) Backward(input, target []float64) {
    // Output layer deltas
    outputLayer := nn.Layers[len(nn.Layers)-1]
    for i := range outputLayer.Deltas {
        error := target[i] - outputLayer.Outputs[i]
        outputLayer.Deltas[i] = error * sigmoidDerivative(outputLayer.Outputs[i])
    }
    
    // Hidden layers deltas
    for l := len(nn.Layers) - 2; l >= 0; l-- {
        currentLayer := nn.Layers[l]
        nextLayer := nn.Layers[l+1]
        
        for i := range currentLayer.Deltas {
            var sum float64
            for j := range nextLayer.Deltas {
                sum += nextLayer.Weights[j][i] * nextLayer.Deltas[j]
            }
            currentLayer.Deltas[i] = sum * reluDerivative(currentLayer.Outputs[i])
        }
    }
    
    // Update weights and biases
    prevOutput := input
    for _, layer := range nn.Layers {
        for i := range layer.Weights {
            for j := range layer.Weights[i] {
                layer.Weights[i][j] += nn.LR * layer.Deltas[i] * prevOutput[j]
            }
            layer.Biases[i] += nn.LR * layer.Deltas[i]
        }
        prevOutput = layer.Outputs
    }
}

func (nn *NeuralNetwork) Train(inputs [][]float64, targets [][]float64, epochs int) {
    for epoch := 0; epoch < epochs; epoch++ {
        totalLoss := 0.0
        
        for i, input := range inputs {
            output := nn.Forward(input)
            nn.Backward(input, targets[i])
            
            // Calculate loss
            for j := range output {
                diff := targets[i][j] - output[j]
                totalLoss += diff * diff
            }
        }
        
        if epoch%100 == 0 {
            avgLoss := totalLoss / float64(len(inputs))
            fmt.Printf("Epoch %d, Loss: %.4f\n", epoch, avgLoss)
        }
    }
}

func (nn *NeuralNetwork) Predict(input []float64) []float64 {
    return nn.Forward(input)
}

func main() {
    rand.Seed(42)
    
    // Создание нейронной сети для XOR
    layers := []int{2, 4, 1} // 2 входа, 4 нейрона в скрытом слое, 1 выход
    nn := NewNeuralNetwork(layers, 0.1)
    
    // Данные XOR
    inputs := [][]float64{
        {0, 0},
        {0, 1},
        {1, 0},
        {1, 1},
    }
    targets := [][]float64{
        {0},
        {1},
        {1},
        {0},
    }
    
    fmt.Println("Training Neural Network for XOR problem...")
    nn.Train(inputs, targets, 10000)
    
    fmt.Println("\nTesting trained network:")
    for i, input := range inputs {
        prediction := nn.Predict(input)
        fmt.Printf("Input: %v, Target: %.0f, Prediction: %.4f\n", 
            input, targets[i][0], prediction[0])
    }
    
    // Тестирование на новых данных
    testInputs := [][]float64{
        {0.5, 0.5},
        {0.2, 0.8},
        {0.8, 0.2},
    }
    
    fmt.Println("\nTesting on new data:")
    for _, input := range testInputs {
        prediction := nn.Predict(input)
        fmt.Printf("Input: %v, Prediction: %.4f\n", input, prediction[0])
    }
}
//Задание: Полносвязная нейронная сеть с обратным распространением ошибки
```
154. Сверточная нейронная сеть (CNN)
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type ConvLayer struct {
    Filters [][][][]float64 // [filter][channel][row][col]
    Biases  []float64
    Stride  int
    Padding int
}

type PoolingLayer struct {
    Size    int
    Stride  int
}

type FlattenLayer struct {
    InputSize  []int
    OutputSize int
}

type CNN struct {
    Conv     *ConvLayer
    Pool     *PoolingLayer
    Flatten  *FlattenLayer
    Dense    *NeuralNetwork
}

func NewConvLayer(numFilters, filterSize, numChannels, stride, padding int) *ConvLayer {
    filters := make([][][][]float64, numFilters)
    biases := make([]float64, numFilters)
    
    for f := 0; f < numFilters; f++ {
        filters[f] = make([][][]float64, numChannels)
        for c := 0; c < numChannels; c++ {
            filters[f][c] = make([][]float64, filterSize)
            for i := 0; i < filterSize; i++ {
                filters[f][c][i] = make([]float64, filterSize)
                for j := 0; j < filterSize; j++ {
                    filters[f][c][i][j] = rand.NormFloat64() * 0.1
                }
            }
        }
        biases[f] = 0.1
    }
    
    return &ConvLayer{
        Filters: filters,
        Biases:  biases,
        Stride:  stride,
        Padding: padding,
    }
}

func (cl *ConvLayer) Forward(input [][][]float64) [][][]float64 {
    channels := len(input)
    inputHeight := len(input[0])
    inputWidth := len(input[0][0])
    
    outputHeight := (inputHeight + 2*cl.Padding - len(cl.Filters[0][0]))/cl.Stride + 1
    outputWidth := (inputWidth + 2*cl.Padding - len(cl.Filters[0][0][0]))/cl.Stride + 1
    
    output := make([][][]float64, len(cl.Filters))
    for f := range output {
        output[f] = make([][]float64, outputHeight)
        for i := range output[f] {
            output[f][i] = make([]float64, outputWidth)
        }
    }
    
    for f := range cl.Filters {
        for i := 0; i < outputHeight; i++ {
            for j := 0; j < outputWidth; j++ {
                sum := cl.Biases[f]
                
                for c := range cl.Filters[f] {
                    for di := 0; di < len(cl.Filters[f][c]); di++ {
                        for dj := 0; dj < len(cl.Filters[f][c][di]); dj++ {
                            inputI := i*cl.Stride + di - cl.Padding
                            inputJ := j*cl.Stride + dj - cl.Padding
                            
                            if inputI >= 0 && inputI < inputHeight && inputJ >= 0 && inputJ < inputWidth {
                                sum += input[c][inputI][inputJ] * cl.Filters[f][c][di][dj]
                            }
                        }
                    }
                }
                
                output[f][i][j] = relu(sum)
            }
        }
    }
    
    return output
}

func NewPoolingLayer(size, stride int) *PoolingLayer {
    return &PoolingLayer{
        Size:   size,
        Stride: stride,
    }
}

func (pl *PoolingLayer) Forward(input [][][]float64) [][][]float64 {
    channels := len(input)
    inputHeight := len(input[0])
    inputWidth := len(input[0][0])
    
    outputHeight := (inputHeight - pl.Size)/pl.Stride + 1
    outputWidth := (inputWidth - pl.Size)/pl.Stride + 1
    
    output := make([][][]float64, channels)
    for c := range output {
        output[c] = make([][]float64, outputHeight)
        for i := range output[c] {
            output[c][i] = make([]float64, outputWidth)
        }
    }
    
    for c := range input {
        for i := 0; i < outputHeight; i++ {
            for j := 0; j < outputWidth; j++ {
                maxVal := math.Inf(-1)
                
                for di := 0; di < pl.Size; di++ {
                    for dj := 0; dj < pl.Size; dj++ {
                        inputI := i*pl.Stride + di
                        inputJ := j*pl.Stride + dj
                        if inputI < inputHeight && inputJ < inputWidth {
                            if input[c][inputI][inputJ] > maxVal {
                                maxVal = input[c][inputI][inputJ]
                            }
                        }
                    }
                }
                
                output[c][i][j] = maxVal
            }
        }
    }
    
    return output
}

func NewFlattenLayer() *FlattenLayer {
    return &FlattenLayer{}
}

func (fl *FlattenLayer) Forward(input [][][]float64) []float64 {
    channels := len(input)
    height := len(input[0])
    width := len(input[0][0])
    
    fl.InputSize = []int{channels, height, width}
    fl.OutputSize = channels * height * width
    
    output := make([]float64, fl.OutputSize)
    idx := 0
    
    for c := range input {
        for i := range input[c] {
            for j := range input[c][i] {
                output[idx] = input[c][i][j]
                idx++
            }
        }
    }
    
    return output
}

func NewCNN() *CNN {
    conv := NewConvLayer(8, 3, 1, 1, 1) // 8 фильтров 3x3
    pool := NewPoolingLayer(2, 2)       // Max pooling 2x2
    flatten := NewFlattenLayer()
    dense := NewNeuralNetwork([]int{32, 16, 10}, 0.01) // Полносвязные слои
    
    return &CNN{
        Conv:    conv,
        Pool:    pool,
        Flatten: flatten,
        Dense:   dense,
    }
}

func (cnn *CNN) Forward(input [][][]float64) []float64 {
    // Сверточный слой
    convOut := cnn.Conv.Forward(input)
    
    // Pooling слой
    poolOut := cnn.Pool.Forward(convOut)
    
    // Выравнивание
    flattenOut := cnn.Flatten.Forward(poolOut)
    
    // Полносвязные слои
    finalOut := cnn.Dense.Forward(flattenOut)
    
    return finalOut
}

func (cnn *CNN) Train(inputs [][][][]float64, targets [][]float64, epochs int) {
    for epoch := 0; epoch < epochs; epoch++ {
        totalLoss := 0.0
        
        for i, input := range inputs {
            output := cnn.Forward(input)
            
            // Упрощенное обучение (без обратного распространения для сверточных слоев)
            // В реальной реализации нужно реализовать полное обратное распространение
            
            // Расчет потерь
            for j := range output {
                diff := targets[i][j] - output[j]
                totalLoss += diff * diff
            }
        }
        
        if epoch%10 == 0 {
            avgLoss := totalLoss / float64(len(inputs))
            fmt.Printf("Epoch %d, Loss: %.4f\n", epoch, avgLoss)
        }
    }
}

func main() {
    rand.Seed(42)
    
    // Создание простой CNN
    cnn := NewCNN()
    
    // Создание синтетических данных (упрощенные изображения 28x28)
    fmt.Println("Generating synthetic image data...")
    
    numSamples := 100
    inputs := make([][][][]float64, numSamples)
    targets := make([][]float64, numSamples)
    
    for i := 0; i < numSamples; i++ {
        // Создание "изображения" 28x28 с одним каналом
        image := make([][]float64, 28)
        for row := range image {
            image[row] = make([]float64, 28)
            for col := range image[row] {
                // Случайные пиксели с некоторой структурой
                if row < 14 && col < 14 {
                    image[row][col] = rand.Float64()
                } else {
                    image[row][col] = rand.Float64() * 0.3
                }
            }
        }
        
        inputs[i] = [][][]float64{image} // Один канал (grayscale)
        
        // Случайная метка класса (10 классов)
        target := make([]float64, 10)
        class := rand.Intn(10)
        target[class] = 1.0
        targets[i] = target
    }
    
    fmt.Println("Training CNN...")
    cnn.Train(inputs, targets, 100)
    
    // Тестирование
    fmt.Println("\nTesting CNN:")
    testImage := inputs[0]
    prediction := cnn.Forward(testImage)
    
    fmt.Printf("Prediction: %v\n", prediction)
    
    // Нахождение предсказанного класса
    maxProb := 0.0
    predictedClass := -1
    for i, prob := range prediction {
        if prob > maxProb {
            maxProb = prob
            predictedClass = i
        }
    }
    
    fmt.Printf("Predicted class: %d (confidence: %.2f%%)\n", 
        predictedClass, maxProb*100)
}
//Задание: Сверточная нейронная сеть для классификации изображений
```
155. Рекуррентная нейронная сеть (RNN)
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type RNNLayer struct {
    Wxh [][]float64 // Веса вход->скрытый
    Whh [][]float64 // Веса скрытый->скрытый
    Why [][]float64 // Веса скрытый->выход
    Bh  []float64   // Смещения скрытого слоя
    By  []float64   // Смещения выходного слоя
    H   []float64   // Состояние скрытого слоя
}

type RNN struct {
    Layer    *RNNLayer
    LR       float64
    HSize    int // Размер скрытого состояния
    InSize   int // Размер входа
    OutSize  int // Размер выхода
}

func NewRNN(inputSize, hiddenSize, outputSize int, learningRate float64) *RNN {
    rnn := &RNN{
        HSize:   hiddenSize,
        InSize:  inputSize,
        OutSize: outputSize,
        LR:      learningRate,
    }
    
    // Инициализация весов
    layer := &RNNLayer{
        Wxh: make([][]float64, hiddenSize),
        Whh: make([][]float64, hiddenSize),
        Why: make([][]float64, outputSize),
        Bh:  make([]float64, hiddenSize),
        By:  make([]float64, outputSize),
        H:   make([]float64, hiddenSize),
    }
    
    // Инициализация Wxh (вход->скрытый)
    for i := range layer.Wxh {
        layer.Wxh[i] = make([]float64, inputSize)
        for j := range layer.Wxh[i] {
            layer.Wxh[i][j] = rand.NormFloat64() * 0.01
        }
    }
    
    // Инициализация Whh (скрытый->скрытый)
    for i := range layer.Whh {
        layer.Whh[i] = make([]float64, hiddenSize)
        for j := range layer.Whh[i] {
            layer.Whh[i][j] = rand.NormFloat64() * 0.01
        }
    }
    
    // Инициализация Why (скрытый->выход)
    for i := range layer.Why {
        layer.Why[i] = make([]float64, hiddenSize)
        for j := range layer.Why[i] {
            layer.Why[i][j] = rand.NormFloat64() * 0.01
        }
    }
    
    rnn.Layer = layer
    return rnn
}

func (rnn *RNN) Forward(inputs [][]float64) ([][]float64, [][]float64) {
    seqLength := len(inputs)
    hiddenStates := make([][]float64, seqLength+1)
    outputs := make([][]float64, seqLength)
    
    // Инициализация начального скрытого состояния
    hiddenStates[0] = make([]float64, rnn.HSize)
    copy(hiddenStates[0], rnn.Layer.H)
    
    for t := 0; t < seqLength; t++ {
        // Новое скрытое состояние
        hiddenStates[t+1] = make([]float64, rnn.HSize)
        
        for i := 0; i < rnn.HSize; i++ {
            sum := rnn.Layer.Bh[i]
            
            // Вход->скрытый
            for j := 0; j < rnn.InSize; j++ {
                sum += inputs[t][j] * rnn.Layer.Wxh[i][j]
            }
            
            // Скрытый->скрытый
            for j := 0; j < rnn.HSize; j++ {
                sum += hiddenStates[t][j] * rnn.Layer.Whh[i][j]
            }
            
            hiddenStates[t+1][i] = math.Tanh(sum)
        }
        
        // Выходной слой
        outputs[t] = make([]float64, rnn.OutSize)
        for i := 0; i < rnn.OutSize; i++ {
            sum := rnn.Layer.By[i]
            for j := 0; j < rnn.HSize; j++ {
                sum += hiddenStates[t+1][j] * rnn.Layer.Why[i][j]
            }
            outputs[t][i] = sum // Linear activation for output
        }
    }
    
    // Сохранение последнего скрытого состояния
    copy(rnn.Layer.H, hiddenStates[seqLength])
    
    return outputs, hiddenStates
}

func (rnn *RNN) Backward(inputs [][]float64, targets [][]float64, outputs [][]float64, hiddenStates [][]float64) {
    seqLength := len(inputs)
    
    // Градиенты
    dWxh := make([][]float64, rnn.HSize)
    dWhh := make([][]float64, rnn.HSize)
    dWhy := make([][]float64, rnn.OutSize)
    dBh := make([]float64, rnn.HSize)
    dBy := make([]float64, rnn.OutSize)
    
    for i := range dWxh {
        dWxh[i] = make([]float64, rnn.InSize)
    }
    for i := range dWhh {
        dWhh[i] = make([]float64, rnn.HSize)
    }
    for i := range dWhy {
        dWhy[i] = make([]float64, rnn.HSize)
    }
    
    // Обратное распространение через время (BPTT)
    dhNext := make([]float64, rnn.HSize)
    
    for t := seqLength - 1; t >= 0; t-- {
        // Градиенты выходного слоя
        dy := make([]float64, rnn.OutSize)
        for i := 0; i < rnn.OutSize; i++ {
            dy[i] = outputs[t][i] - targets[t][i]
            
            // Градиенты Why и By
            for j := 0; j < rnn.HSize; j++ {
                dWhy[i][j] += dy[i] * hiddenStates[t+1][j]
            }
            dBy[i] += dy[i]
        }
        
        // Градиенты скрытого состояния
        dh := make([]float64, rnn.HSize)
        for i := 0; i < rnn.HSize; i++ {
            // От выхода
            for j := 0; j < rnn.OutSize; j++ {
                dh[i] += dy[j] * rnn.Layer.Why[j][i]
            }
            
            // От следующего скрытого состояния
            for j := 0; j < rnn.HSize; j++ {
                dh[i] += dhNext[j] * rnn.Layer.Whh[j][i]
            }
            
            // Градиент tanh
            dh[i] *= (1 - hiddenStates[t+1][i]*hiddenStates[t+1][i])
        }
        
        // Градиенты весов
        for i := 0; i < rnn.HSize; i++ {
            for j := 0; j < rnn.InSize; j++ {
                dWxh[i][j] += dh[i] * inputs[t][j]
            }
            for j := 0; j < rnn.HSize; j++ {
                dWhh[i][j] += dh[i] * hiddenStates[t][j]
            }
            dBh[i] += dh[i]
        }
        
        dhNext = dh
    }
    
    // Обновление весов
    for i := 0; i < rnn.HSize; i++ {
        for j := 0; j < rnn.InSize; j++ {
            rnn.Layer.Wxh[i][j] -= rnn.LR * dWxh[i][j]
        }
        for j := 0; j < rnn.HSize; j++ {
            rnn.Layer.Whh[i][j] -= rnn.LR * dWhh[i][j]
        }
        rnn.Layer.Bh[i] -= rnn.LR * dBh[i]
    }
    
    for i := 0; i < rnn.OutSize; i++ {
        for j := 0; j < rnn.HSize; j++ {
            rnn.Layer.Why[i][j] -= rnn.LR * dWhy[i][j]
        }
        rnn.Layer.By[i] -= rnn.LR * dBy[i]
    }
}

func (rnn *RNN) Train(inputs [][]float64, targets [][]float64, epochs int) {
    for epoch := 0; epoch < epochs; epoch++ {
        totalLoss := 0.0
        
        outputs, hiddenStates := rnn.Forward(inputs)
        rnn.Backward(inputs, targets, outputs, hiddenStates)
        
        // Расчет потерь
        for t := range outputs {
            for i := range outputs[t] {
                diff := outputs[t][i] - targets[t][i]
                totalLoss += diff * diff
            }
        }
        
        if epoch%100 == 0 {
            avgLoss := totalLoss / float64(len(outputs))
            fmt.Printf("Epoch %d, Loss: %.4f\n", epoch, avgLoss)
        }
    }
}

func (rnn *RNN) Predict(seed []float64, length int) [][]float64 {
    predictions := make([][]float64, length)
    currentInput := seed
    currentHidden := make([]float64, rnn.HSize)
    copy(currentHidden, rnn.Layer.H)
    
    for i := 0; i < length; i++ {
        // Прямой проход
        nextHidden := make([]float64, rnn.HSize)
        for j := 0; j < rnn.HSize; j++ {
            sum := rnn.Layer.Bh[j]
            for k := 0; k < rnn.InSize; k++ {
                sum += currentInput[k] * rnn.Layer.Wxh[j][k]
            }
            for k := 0; k < rnn.HSize; k++ {
                sum += currentHidden[k] * rnn.Layer.Whh[j][k]
            }
            nextHidden[j] = math.Tanh(sum)
        }
        
        // Выход
        output := make([]float64, rnn.OutSize)
        for j := 0; j < rnn.OutSize; j++ {
            sum := rnn.Layer.By[j]
            for k := 0; k < rnn.HSize; k++ {
                sum += nextHidden[k] * rnn.Layer.Why[j][k]
            }
            output[j] = sum
        }
        
        predictions[i] = output
        
        // Использование предсказания как следующего входа
        currentInput = output
        currentHidden = nextHidden
    }
    
    return predictions
}

func main() {
    rand.Seed(42)
    
    // Создание RNN для предсказания последовательностей
    rnn := NewRNN(1, 32, 1, 0.01)
    
    // Генерация синтетических данных (синусоида)
    fmt.Println("Generating sine wave data...")
    
    seqLength := 50
    inputs := make([][]float64, seqLength)
    targets := make([][]float64, seqLength)
    
    for i := 0; i < seqLength; i++ {
        x := float64(i) * 0.1
        inputs[i] = []float64{math.Sin(x)}
        targets[i] = []float64{math.Sin(x + 0.1)} // Предсказание следующего значения
    }
    
    fmt.Println("Training RNN...")
    rnn.Train(inputs, targets, 1000)
    
    // Тестирование предсказания
    fmt.Println("\nTesting sequence prediction:")
    seed := []float64{math.Sin(0)}
    predictions := rnn.Predict(seed, 20)
    
    fmt.Println("Time\tPredicted\tActual")
    for i := 0; i < len(predictions); i++ {
        actual := math.Sin(float64(i+1) * 0.1)
        fmt.Printf("%d\t%.4f\t\t%.4f\n", i+1, predictions[i][0], actual)
    }
    
    // Сброс состояния для нового предсказания
    rnn.Layer.H = make([]float64, rnn.HSize)
    
    // Предсказание более длинной последовательности
    fmt.Println("\nLong-term prediction:")
    longPredictions := rnn.Predict(seed, 50)
    
    fmt.Println("Time\tPredicted\tActual")
    for i := 0; i < len(longPredictions); i++ {
        actual := math.Sin(float64(i+1) * 0.1)
        fmt.Printf("%d\t%.4f\t\t%.4f\n", i+1, longPredictions[i][0], actual)
    }
}
//Задание: Рекуррентная нейронная сеть для предсказания временных рядов
```
156. Генеративно-состязательная сеть (GAN)
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type Generator struct {
    Layers []*Layer
    LR     float64
}

type Discriminator struct {
    Layers []*Layer
    LR     float64
}

type GAN struct {
    Generator     *Generator
    Discriminator *Discriminator
    NoiseSize     int
    DataSize      int
}

func NewGenerator(noiseSize, hiddenSize, outputSize int, learningRate float64) *Generator {
    layers := []int{noiseSize, hiddenSize, hiddenSize, outputSize}
    nn := NewNeuralNetwork(layers, learningRate)
    
    return &Generator{
        Layers: nn.Layers,
        LR:     learningRate,
    }
}

func NewDiscriminator(inputSize, hiddenSize int, learningRate float64) *Discriminator {
    layers := []int{inputSize, hiddenSize, hiddenSize, 1}
    nn := NewNeuralNetwork(layers, learningRate)
    
    return &Discriminator{
        Layers: nn.Layers,
        LR:     learningRate,
    }
}

func NewGAN(noiseSize, dataSize, hiddenSize int, learningRate float64) *GAN {
    generator := NewGenerator(noiseSize, hiddenSize, dataSize, learningRate)
    discriminator := NewDiscriminator(dataSize, hiddenSize, learningRate)
    
    return &GAN{
        Generator:     generator,
        Discriminator: discriminator,
        NoiseSize:     noiseSize,
        DataSize:      dataSize,
    }
}

func (g *Generator) Forward(noise []float64) []float64 {
    current := noise
    
    for _, layer := range g.Layers {
        next := make([]float64, len(layer.Weights))
        
        for i, neuron := range layer.Weights {
            sum := layer.Biases[i]
            for j, weight := range neuron {
                sum += weight * current[j]
            }
            
            // Tanh для генератора (выход в диапазоне [-1, 1])
            if layer == g.Layers[len(g.Layers)-1] {
                next[i] = math.Tanh(sum)
            } else {
                next[i] = relu(sum)
            }
            layer.Outputs[i] = next[i]
        }
        
        current = next
    }
    
    return current
}

func (d *Discriminator) Forward(data []float64) float64 {
    current := data
    
    for _, layer := range d.Layers {
        next := make([]float64, len(layer.Weights))
        
        for i, neuron := range layer.Weights {
            sum := layer.Biases[i]
            for j, weight := range neuron {
                sum += weight * current[j]
            }
            
            // Sigmoid для дискриминатора (вероятность)
            if layer == d.Layers[len(d.Layers)-1] {
                next[i] = sigmoid(sum)
            } else {
                next[i] = relu(sum)
            }
            layer.Outputs[i] = next[i]
        }
        
        current = next
    }
    
    return current[0] // Один выход - вероятность реальности
}

func (gan *GAN) Train(realData [][]float64, epochs, batchSize int) {
    for epoch := 0; epoch < epochs; epoch++ {
        // Обучение дискриминатора
        dLoss := gan.trainDiscriminator(realData, batchSize)
        
        // Обучение генератора
        gLoss := gan.trainGenerator(batchSize)
        
        if epoch%100 == 0 {
            fmt.Printf("Epoch %d, D Loss: %.4f, G Loss: %.4f\n", epoch, dLoss, gLoss)
        }
    }
}

func (gan *GAN) trainDiscriminator(realData [][]float64, batchSize int) float64 {
    totalLoss := 0.0
    batches := 0
    
    for i := 0; i < len(realData); i += batchSize {
        end := i + batchSize
        if end > len(realData) {
            end = len(realData)
        }
        
        batchReal := realData[i:end]
        batchSize := len(batchReal)
        
        // Реальные данные
        realLoss := 0.0
        for _, data := range batchReal {
            prediction := gan.Discriminator.Forward(data)
            // Дискриминатор должен предсказать 1 для реальных данных
            realLoss += -math.Log(prediction)
        }
        realLoss /= float64(batchSize)
        
        // Сгенерированные данные
        fakeLoss := 0.0
        for j := 0; j < batchSize; j++ {
            noise := make([]float64, gan.NoiseSize)
            for k := range noise {
                noise[k] = rand.NormFloat64()
            }
            
            fakeData := gan.Generator.Forward(noise)
            prediction := gan.Discriminator.Forward(fakeData)
            // Дискриминатор должен предсказать 0 для сгенерированных данных
            fakeLoss += -math.Log(1 - prediction)
        }
        fakeLoss /= float64(batchSize)
        
        totalLoss += (realLoss + fakeLoss) / 2
        batches++
    }
    
    return totalLoss / float64(batches)
}

func (gan *GAN) trainGenerator(batchSize int) float64 {
    totalLoss := 0.0
    batches := 0
    
    for i := 0; i < 5; i++ { // Обучаем генератор несколько раз за эпоху
        batchLoss := 0.0
        
        for j := 0; j < batchSize; j++ {
            noise := make([]float64, gan.NoiseSize)
            for k := range noise {
                noise[k] = rand.NormFloat64()
            }
            
            fakeData := gan.Generator.Forward(noise)
            prediction := gan.Discriminator.Forward(fakeData)
            
            // Генератор хочет, чтобы дискриминатор предсказал 1 для сгенерированных данных
            batchLoss += -math.Log(prediction)
        }
        
        totalLoss += batchLoss / float64(batchSize)
        batches++
    }
    
    return totalLoss / float64(batches)
}

func (gan *GAN) Generate(numSamples int) [][]float64 {
    samples := make([][]float64, numSamples)
    
    for i := 0; i < numSamples; i++ {
        noise := make([]float64, gan.NoiseSize)
        for j := range noise {
            noise[j] = rand.NormFloat64()
        }
        
        sample := gan.Generator.Forward(noise)
        samples[i] = sample
    }
    
    return samples
}

func main() {
    rand.Seed(42)
    
    // Создание GAN для генерации 2D данных
    noiseSize := 10
    dataSize := 2
    hiddenSize := 64
    
    gan := NewGAN(noiseSize, dataSize, hiddenSize, 0.001)
    
    // Генерация реальных данных (круг)
    fmt.Println("Generating real data (circle)...")
    numRealSamples := 1000
    realData := make([][]float64, numRealSamples)
    
    for i := 0; i < numRealSamples; i++ {
        angle := rand.Float64() * 2 * math.Pi
        radius := 2.0
        
        x := radius * math.Cos(angle) + rand.NormFloat64()*0.1
        y := radius * math.Sin(angle) + rand.NormFloat64()*0.1
        
        realData[i] = []float64{x, y}
    }
    
    fmt.Println("Training GAN...")
    gan.Train(realData, 1000, 32)
    
    // Генерация новых данных
    fmt.Println("\nGenerating new samples...")
    generatedSamples := gan.Generate(10)
    
    fmt.Println("Generated samples:")
    for i, sample := range generatedSamples {
        fmt.Printf("Sample %d: (%.3f, %.3f)\n", i+1, sample[0], sample[1])
    }
    
    // Оценка качества генерации
    fmt.Println("\nQuality evaluation:")
    
    // Проверка, лежат ли сгенерированные точки на круге
    correct := 0
    for _, sample := range generatedSamples {
        distance := math.Sqrt(sample[0]*sample[0] + sample[1]*sample[1])
        if math.Abs(distance-2.0) < 0.5 {
            correct++
        }
    }
    
    fmt.Printf("Samples on circle: %d/%d (%.1f%%)\n", 
        correct, len(generatedSamples), float64(correct)/float64(len(generatedSamples))*100)
    
    // Генерация большего количества samples для визуализации
    manySamples := gan.Generate(100)
    
    fmt.Println("\nStatistics of 100 generated samples:")
    var avgX, avgY, stdX, stdY float64
    
    for _, sample := range manySamples {
        avgX += sample[0]
        avgY += sample[1]
    }
    avgX /= 100
    avgY /= 100
    
    for _, sample := range manySamples {
        stdX += (sample[0] - avgX) * (sample[0] - avgX)
        stdY += (sample[1] - avgY) * (sample[1] - avgY)
    }
    stdX = math.Sqrt(stdX / 100)
    stdY = math.Sqrt(stdY / 100)
    
    fmt.Printf("Mean: (%.3f, %.3f)\n", avgX, avgY)
    fmt.Printf("Std:  (%.3f, %.3f)\n", stdX, stdY)
}
//Задание: Генеративно-состязательная сеть для создания синтетических данных
```
157. Reinforcement Learning - Q-learning
```go
package main
import (
    "fmt"
    "math"
    "math/rand"
)

type QLearning struct {
    QTable      map[string]map[int]float64
    LearningRate float64
    Discount     float64
    Exploration  float64
    MinExploration float64
    ExplorationDecay float64
}

type Environment struct {
    States  []string
    Actions []int
    Rewards map[string]float64
    CurrentState string
}

func NewQLearning(learningRate, discount, exploration float64) *QLearning {
    return &QLearning{
        QTable:          make(map[string]map[int]float64),
        LearningRate:    learningRate,
        Discount:        discount,
        Exploration:     exploration,
        MinExploration:  0.01,
        ExplorationDecay: 0.995,
    }
}

func NewEnvironment() *Environment {
    states := []string{
        "start", "state1", "state2", "state3", "goal",
        "trap1", "trap2",
    }
    
    actions := []int{0, 1, 2, 3} // up, right, down, left
    
    rewards := map[string]float64{
        "start": 0,
        "state1": 0,
        "state2": 0,
        "state3": 0,
        "goal":  100,
        "trap1": -50,
        "trap2": -50,
    }
    
    return &Environment{
        States:  states,
        Actions: actions,
        Rewards: rewards,
        CurrentState: "start",
    }
}

func (env *Environment) Reset() string {
    env.CurrentState = "start"
    return env.CurrentState
}

func (env *Environment) Step(action int) (string, float64, bool) {
    nextState := env.getNextState(env.CurrentState, action)
    reward := env.Rewards[nextState]
    done := (nextState == "goal" || nextState == "trap1" || nextState == "trap2")
    
    env.CurrentState = nextState
    return nextState, reward, done
}

func (env *Environment) getNextState(state string, action int) string {
    transitions := map[string]map[int]string{
        "start": {
            0: "trap1",
            1: "state1",
            2: "start",
            3: "trap2",
        },
        "state1": {
            0: "state1",
            1: "state2",
            2: "start",
            3: "state1",
        },
        "state2": {
            0: "state1",
            1: "state3",
            2: "state2",
            3: "state2",
        },
        "state3": {
            0: "state2",
            1: "goal",
            2: "state3",
            3: "trap1",
        },
        "trap1": {
            0: "trap1", 1: "trap1", 2: "trap1", 3: "trap1",
        },
        "trap2": {
            0: "trap2", 1: "trap2", 2: "trap2", 3: "trap2",
        },
        "goal": {
            0: "goal", 1: "goal", 2: "goal", 3: "goal",
        },
    }
    
    return transitions[state][action]
}

func (ql *QLearning) GetQValue(state string, action int) float64 {
    if _, exists := ql.QTable[state]; !exists {
        ql.QTable[state] = make(map[int]float64)
        for _, a := range []int{0, 1, 2, 3} {
            ql.QTable[state][a] = 0
        }
    }
    return ql.QTable[state][action]
}

func (ql *QLearning) SetQValue(state string, action int, value float64) {
    if _, exists := ql.QTable[state]; !exists {
        ql.QTable[state] = make(map[int]float64)
    }
    ql.QTable[state][action] = value
}

func (ql *QLearning) ChooseAction(state string) int {
    // ε-greedy стратегия
    if rand.Float64() < ql.Exploration {
        // Случайное действие (exploration)
        return rand.Intn(4)
    } else {
        // Лучшее действие (exploitation)
        return ql.GetBestAction(state)
    }
}

func (ql *QLearning) GetBestAction(state string) int {
    bestAction := 0
    bestValue := math.Inf(-1)
    
    for action := 0; action < 4; action++ {
        value := ql.GetQValue(state, action)
        if value > bestValue {
            bestValue = value
            bestAction = action
        }
    }
    
    return bestAction
}

func (ql *QLearning) Update(state string, action int, reward float64, nextState string) {
    currentQ := ql.GetQValue(state, action)
    
    // Максимальное Q-value для следующего состояния
    maxNextQ := math.Inf(-1)
    for a := 0; a < 4; a++ {
        nextQ := ql.GetQValue(nextState, a)
        if nextQ > maxNextQ {
            maxNextQ = nextQ
        }
    }
    
    // Q-learning формула
    newQ := currentQ + ql.LearningRate * (reward + ql.Discount * maxNextQ - currentQ)
    ql.SetQValue(state, action, newQ)
}

func (ql *QLearning) Train(env *Environment, episodes int) {
    for episode := 0; episode < episodes; episode++ {
        state := env.Reset()
        totalReward := 0.0
        steps := 0
        
        for {
            action := ql.ChooseAction(state)
            nextState, reward, done := env.Step(action)
            
            ql.Update(state, action, reward, nextState)
            
            state = nextState
            totalReward += reward
            steps++
            
            if done {
                break
            }
            
            if steps > 100 { // Защита от бесконечных циклов
                break
            }
        }
        
        // Decay exploration rate
        ql.Exploration = math.Max(ql.MinExploration, ql.Exploration * ql.ExplorationDecay)
        
        if episode%100 == 0 {
            fmt.Printf("Episode %d: steps=%d, reward=%.1f, exploration=%.3f\n", 
                episode, steps, totalReward, ql.Exploration)
        }
    }
}

func (ql *QLearning) Test(env *Environment, episodes int) {
    totalRewards := 0.0
    successes := 0
    
    for episode := 0; episode < episodes; episode++ {
        state := env.Reset()
        totalReward := 0.0
        steps := 0
        
        for {
            action := ql.GetBestAction(state) // Всегда лучшее действие
            nextState, reward, done := env.Step(action)
            
            state = nextState
            totalReward += reward
            steps++
            
            if done {
                if state == "goal" {
                    successes++
                }
                break
            }
            
            if steps > 100 {
                break
            }
        }
        
        totalRewards += totalReward
    }
    
    avgReward := totalRewards / float64(episodes)
    successRate := float64(successes) / float64(episodes) * 100
    
    fmt.Printf("\nTest Results (%d episodes):\n", episodes)
    fmt.Printf("Average Reward: %.2f\n", avgReward)
    fmt.Printf("Success Rate: %.1f%%\n", successRate)
}

func (ql *QLearning) PrintPolicy() {
    fmt.Println("\nLearned Policy:")
    
    states := []string{"start", "state1", "state2", "state3"}
    actionNames := map[int]string{0: "↑", 1: "→", 2: "↓", 3: "←"}
    
    for _, state := range states {
        bestAction := ql.GetBestAction(state)
        bestValue := ql.GetQValue(state, bestAction)
        
        fmt.Printf("%s: %s (Q=%.2f) [", state, actionNames[bestAction], bestValue)
        for action := 0; action < 4; action++ {
            fmt.Printf("%s:%.2f ", actionNames[action], ql.GetQValue(state, action))
        }
        fmt.Printf("]\n")
    }
}

func main() {
    rand.Seed(42)
    
    // Создание среды и агента
    env := NewEnvironment()
    ql := NewQLearning(0.1, 0.9, 1.0)
    
    fmt.Println("Training Q-learning agent...")
    ql.Train(env, 1000)
    
    // Тестирование обученного агента
    ql.Test(env, 100)
    
    // Вывод изученной политики
    ql.PrintPolicy()
    
    // Демонстрация одного эпизода
    fmt.Println("\nDemonstration episode:")
    state := env.Reset()
    fmt.Printf("Start state: %s\n", state)
    
    steps := 0
    for {
        action := ql.GetBestAction(state)
        nextState, reward, done := env.Step(action)
        
        actionNames := map[int]string{0: "↑", 1: "→", 2: "↓", 3: "←"}
        fmt.Printf("Step %d: %s -> %s (%s), Reward: %.1f\n", 
            steps, state, nextState, actionNames[action], reward)
        
        state = nextState
        steps++
        
        if done {
            if state == "goal" {
                fmt.Println("🎉 Reached goal!")
            } else {
                fmt.Println("💥 Fell into trap!")
            }
            break
        }
        
        if steps > 20 {
            fmt.Println("⏰ Timeout!")
            break
        }
    }
}
//Задание: Q-learning для reinforcement learning с ε-greedy стратегией
```
Обработка аудио и мультимедиа
158. Аудио обработка - генерация и анализ звука
```go
package main
import (
    "encoding/binary"
    "fmt"
    "math"
    "os"
)

type AudioProcessor struct {
    SampleRate float64
    BitDepth   int
}

type WaveHeader struct {
    ChunkID       [4]byte
    ChunkSize     uint32
    Format        [4]byte
    Subchunk1ID   [4]byte
    Subchunk1Size uint32
    AudioFormat   uint16
    NumChannels   uint16
    SampleRate    uint32
    ByteRate      uint32
    BlockAlign    uint16
    BitsPerSample uint16
    Subchunk2ID   [4]byte
    Subchunk2Size uint32
}

func NewAudioProcessor(sampleRate float64, bitDepth int) *AudioProcessor {
    return &AudioProcessor{
        SampleRate: sampleRate,
        BitDepth:   bitDepth,
    }
}

func (ap *AudioProcessor) GenerateSineWave(frequency, duration float64) []float64 {
    numSamples := int(ap.SampleRate * duration)
    samples := make([]float64, numSamples)
    
    for i := 0; i < numSamples; i++ {
        time := float64(i) / ap.SampleRate
        samples[i] = math.Sin(2 * math.Pi * frequency * time)
    }
    
    return samples
}

func (ap *AudioProcessor) GenerateSquareWave(frequency, duration float64) []float64 {
    numSamples := int(ap.SampleRate * duration)
    samples := make([]float64, numSamples)
    
    for i := 0; i < numSamples; i++ {
        time := float64(i) / ap.SampleRate
        phase := 2 * math.Pi * frequency * time
        if math.Sin(phase) >= 0 {
            samples[i] = 1.0
        } else {
            samples[i] = -1.0
        }
    }
    
    return samples
}

func (ap *AudioProcessor) GenerateSawtoothWave(frequency, duration float64) []float64 {
    numSamples := int(ap.SampleRate * duration)
    samples := make([]float64, numSamples)
    
    for i := 0; i < numSamples; i++ {
        time := float64(i) / ap.SampleRate
        phase := math.Mod(time*frequency, 1.0)
        samples[i] = 2*phase - 1
    }
    
    return samples
}

func (ap *AudioProcessor) MixWaves(waves ...[]float64) []float64 {
    if len(waves) == 0 {
        return nil
    }
    
    length := len(waves[0])
    mixed := make([]float64, length)
    
    for _, wave := range waves {
        for i := range wave {
            if i < length {
                mixed[i] += wave[i]
            }
        }
    }
    
    // Нормализация
    maxAmplitude := 0.0
    for i := range mixed {
        if math.Abs(mixed[i]) > maxAmplitude {
            maxAmplitude = math.Abs(mixed[i])
        }
    }
    
    if maxAmplitude > 0 {
        for i := range mixed {
            mixed[i] /= maxAmplitude
        }
    }
    
    return mixed
}

func (ap *AudioProcessor) ApplyFade(samples []float64, fadeIn, fadeOut float64) []float64 {
    result := make([]float64, len(samples))
    copy(result, samples)
    
    fadeInSamples := int(ap.SampleRate * fadeIn)
    fadeOutSamples := int(ap.SampleRate * fadeOut)
    
    // Fade in
    for i := 0; i < fadeInSamples && i < len(result); i++ {
        factor := float64(i) / float64(fadeInSamples)
        result[i] *= factor
    }
    
    // Fade out
    for i := 0; i < fadeOutSamples && i < len(result); i++ {
        idx := len(result) - 1 - i
        factor := float64(i) / float64(fadeOutSamples)
        result[idx] *= factor
    }
    
    return result
}

func (ap *AudioProcessor) SaveWAV(filename string, samples []float64) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Подготовка заголовка WAV
    header := WaveHeader{
        ChunkID:       [4]byte{'R', 'I', 'F', 'F'},
        Format:        [4]byte{'W', 'A', 'V', 'E'},
        Subchunk1ID:   [4]byte{'f', 'm', 't', ' '},
        Subchunk1Size: 16,
        AudioFormat:   1, // PCM
        NumChannels:   1, // Mono
        SampleRate:    uint32(ap.SampleRate),
        BitsPerSample: uint16(ap.BitDepth),
    }
    
    header.BlockAlign = header.NumChannels * header.BitsPerSample / 8
    header.ByteRate = header.SampleRate * uint32(header.BlockAlign)
    
    dataSize := uint32(len(samples) * int(header.BitsPerSample) / 8)
    header.ChunkSize = 4 + (8 + header.Subchunk1Size) + (8 + dataSize)
    header.Subchunk2ID = [4]byte{'d', 'a', 't', 'a'}
    header.Subchunk2Size = dataSize
    
    // Запись заголовка
    binary.Write(file, binary.LittleEndian, &header)
    
    // Запись данных
    for _, sample := range samples {
        // Конвертация в 16-bit PCM
        pcmValue := int16(sample * 32767)
        binary.Write(file, binary.LittleEndian, pcmValue)
    }
    
    return nil
}

func (ap *AudioProcessor) FFT(samples []float64) []complex128 {
    n := len(samples)
    if n == 1 {
        return []complex128{complex(samples[0], 0)}
    }
    
    // Разделение на четные и нечетные
    even := make([]float64, n/2)
    odd := make([]float64, n/2)
    
    for i := 0; i < n/2; i++ {
        even[i] = samples[2*i]
        odd[i] = samples[2*i+1]
    }
    
    // Рекурсивное вычисление FFT
    evenFFT := ap.FFT(even)
    oddFFT := ap.FFT(odd)
    
    // Комбинирование результатов
    result := make([]complex128, n)
    for k := 0; k < n/2; k++ {
        angle := -2 * math.Pi * float64(k) / float64(n)
        twiddle := complex(math.Cos(angle), math.Sin(angle))
        
        result[k] = evenFFT[k] + twiddle*oddFFT[k]
        result[k+n/2] = evenFFT[k] - twiddle*oddFFT[k]
    }
    
    return result
}

func (ap *AudioProcessor) CalculateSpectrum(samples []float64) []float64 {
    // Применение оконной функции Хэннинга
    windowed := make([]float64, len(samples))
    for i := range samples {
        window := 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(len(samples)-1)))
        windowed[i] = samples[i] * window
    }
    
    // FFT
    fftResult := ap.FFT(windowed)
    
    // Вычисление амплитудного спектра
    spectrum := make([]float64, len(fftResult)/2)
    for i := range spectrum {
        re := real(fftResult[i])
        im := imag(fftResult[i])
        spectrum[i] = math.Sqrt(re*re + im*im)
    }
    
    return spectrum
}

func (ap *AudioProcessor) FindFundamentalFrequency(samples []float64) float64 {
    spectrum := ap.CalculateSpectrum(samples)
    
    // Поиск пика в спектре (исключая DC компонент)
    maxAmp := 0.0
    maxBin := 0
    
    for i := 1; i < len(spectrum)/2; i++ { // Игнорируем верхнюю половину (симметрия)
        if spectrum[i] > maxAmp {
            maxAmp = spectrum[i]
            maxBin = i
        }
    }
    
    // Конвертация бина в частоту
    fundamentalFreq := float64(maxBin) * ap.SampleRate / float64(len(samples))
    return fundamentalFreq
}

func main() {
    ap := NewAudioProcessor(44100, 16)
    
    fmt.Println("=== Audio Synthesis and Analysis ===")
    
    // Генерация различных волн
    sine := ap.GenerateSineWave(440, 2.0)      // Ля 440 Гц
    square := ap.GenerateSquareWave(220, 2.0)  // Ля 220 Гц
    sawtooth := ap.GenerateSawtoothWave(330, 2.0) // Ми 330 Гц
    
    // Смешивание волн
    mixed := ap.MixWaves(sine, square, sawtooth)
    
    // Применение fade in/out
    mixedWithFade := ap.ApplyFade(mixed, 0.1, 0.1)
    
    // Сохранение в WAV файл
    err := ap.SaveWAV("output.wav", mixedWithFade)
    if err != nil {
        fmt.Printf("Error saving WAV: %v\n", err)
    } else {
        fmt.Println("Audio saved: output.wav")
    }
    
    // Анализ аудио
    fmt.Println("\n=== Audio Analysis ===")
    
    // Генерация тестового сигнала для анализа
    testSignal := ap.GenerateSineWave(1000, 1.0) // 1 кГц
    fundamental := ap.FindFundamentalFrequency(testSignal)
    
    fmt.Printf("Generated 1000 Hz signal\n")
    fmt.Printf("Detected fundamental frequency: %.2f Hz\n", fundamental)
    
    // Спектральный анализ
    spectrum := ap.CalculateSpectrum(testSignal)
    
    fmt.Println("\nFirst 10 frequency bins:")
    for i := 0; i < 10 && i < len(spectrum); i++ {
        freq := float64(i) * ap.SampleRate / float64(len(testSignal))
        fmt.Printf("Bin %d: %.1f Hz, Amplitude: %.2f\n", i, freq, spectrum[i])
    }
    
    // Создание аккорда
    fmt.Println("\n=== Chord Generation ===")
    chordFrequencies := []float64{261.63, 329.63, 392.00} // C, E, G
    var chordWaves [][]float64
    
    for _, freq := range chordFrequencies {
        wave := ap.GenerateSineWave(freq, 3.0)
        chordWaves = append(chordWaves, wave)
    }
    
    chord := ap.MixWaves(chordWaves...)
    chord = ap.ApplyFade(chord, 0.5, 0.5)
    
    ap.SaveWAV("chord.wav", chord)
    fmt.Println("Chord saved: chord.wav")
    
    // Анализ аккорда
    chordFundamental := ap.FindFundamentalFrequency(chord)
    fmt.Printf("Chord fundamental frequency: %.2f Hz\n", chordFundamental)
}
//Задание: Генерация и анализ аудио сигналов с FFT
```
159. Криптовалютный кошелек и транзакции
```go
package main
import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "math/big"
)

type Wallet struct {
    PrivateKey *ecdsa.PrivateKey
    PublicKey  *ecdsa.PublicKey
    Address    string
}

type Transaction struct {
    From   string
    To     string
    Amount float64
    Fee    float64
    Nonce  int
    Signature string
    Hash     string
}

type Block struct {
    Index        int
    Timestamp    int64
    Transactions []*Transaction
    PreviousHash string
    Hash         string
    Nonce        int
}

type Blockchain struct {
    Chain        []*Block
    PendingTransactions []*Transaction
    Difficulty   int
    MiningReward float64
}

func NewWallet() *Wallet {
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return nil
    }
    
    publicKey := &privateKey.PublicKey
    
    // Генерация адреса из публичного ключа
    publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
    hash := sha256.Sum256(publicKeyBytes)
    address := hex.EncodeToString(hash[:20]) // Берем первые 20 байт
    
    return &Wallet{
        PrivateKey: privateKey,
        PublicKey:  publicKey,
        Address:    address,
    }
}

func (w *Wallet) Sign(data []byte) (string, error) {
    hash := sha256.Sum256(data)
    r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
    if err != nil {
        return "", err
    }
    
    signature := append(r.Bytes(), s.Bytes()...)
    return hex.EncodeToString(signature), nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, data []byte, signature string) bool {
    sigBytes, err := hex.DecodeString(signature)
    if err != nil {
        return false
    }
    
    if len(sigBytes) != 64 { // 32 bytes for r + 32 bytes for s
        return false
    }
    
    r := new(big.Int).SetBytes(sigBytes[:32])
    s := new(big.Int).SetBytes(sigBytes[32:])
    
    hash := sha256.Sum256(data)
    return ecdsa.Verify(publicKey, hash[:], r, s)
}

func NewTransaction(from, to string, amount, fee float64, nonce int) *Transaction {
    tx := &Transaction{
        From:   from,
        To:     to,
        Amount: amount,
        Fee:    fee,
        Nonce:  nonce,
    }
    
    tx.Hash = tx.CalculateHash()
    return tx
}

func (tx *Transaction) CalculateHash() string {
    data := fmt.Sprintf("%s%s%.8f%.8f%d", tx.From, tx.To, tx.Amount, tx.Fee, tx.Nonce)
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

func (tx *Transaction) Sign(wallet *Wallet) error {
    if wallet.Address != tx.From {
        return fmt.Errorf("cannot sign transaction from other wallet")
    }
    
    signature, err := wallet.Sign([]byte(tx.Hash))
    if err != nil {
        return err
    }
    
    tx.Signature = signature
    return nil
}

func (tx *Transaction) IsValid() bool {
    if tx.Amount <= 0 {
        return false
    }
    
    if tx.From == "" { // Mining reward transaction
        return true
    }
    
    if tx.Signature == "" {
        return false
    }
    
    // В реальной системе здесь была бы проверка подписи
    // и баланса отправителя
    return true
}

func NewBlockchain(difficulty int, miningReward float64) *Blockchain {
    bc := &Blockchain{
        Difficulty:   difficulty,
        MiningReward: miningReward,
    }
    
    // Создание генезис-блока
    genesisBlock := &Block{
        Index:        0,
        Timestamp:    1625097600, // Пример timestamp
        Transactions: []*Transaction{},
        PreviousHash: "0",
        Nonce:        0,
    }
    
    genesisBlock.Hash = bc.CalculateHash(genesisBlock)
    bc.Chain = []*Block{genesisBlock}
    
    return bc
}

func (bc *Blockchain) CalculateHash(block *Block) string {
    data := fmt.Sprintf("%d%d%s%s%d", 
        block.Index, block.Timestamp, 
        bc.SerializeTransactions(block.Transactions),
        block.PreviousHash, block.Nonce)
    
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

func (bc *Blockchain) SerializeTransactions(txs []*Transaction) string {
    var result string
    for _, tx := range txs {
        result += tx.Hash
    }
    return result
}

func (bc *Blockchain) GetLatestBlock() *Block {
    return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddTransaction(tx *Transaction) bool {
    if !tx.IsValid() {
        return false
    }
    
    bc.PendingTransactions = append(bc.PendingTransactions, tx)
    return true
}

func (bc *Blockchain) MineBlock(minerAddress string) {
    // Награда за майнинг
    rewardTx := NewTransaction("", minerAddress, bc.MiningReward, 0, 0)
    bc.PendingTransactions = append(bc.PendingTransactions, rewardTx)
    
    block := &Block{
        Index:        len(bc.Chain),
        Timestamp:    1625097600 + int64(len(bc.Chain)*600), // +10 минут за блок
        Transactions: bc.PendingTransactions,
        PreviousHash: bc.GetLatestBlock().Hash,
        Nonce:        0,
    }
    
    // Proof of Work
    for {
        block.Hash = bc.CalculateHash(block)
        if bc.IsValidHash(block.Hash) {
            break
        }
        block.Nonce++
    }
    
    bc.Chain = append(bc.Chain, block)
    bc.PendingTransactions = []*Transaction{}
}

func (bc *Blockchain) IsValidHash(hash string) bool {
    prefix := ""
    for i := 0; i < bc.Difficulty; i++ {
        prefix += "0"
    }
    return len(hash) >= bc.Difficulty && hash[:bc.Difficulty] == prefix
}

func (bc *Blockchain) IsChainValid() bool {
    for i := 1; i < len(bc.Chain); i++ {
        currentBlock := bc.Chain[i]
        previousBlock := bc.Chain[i-1]
        
        // Проверка хеша блока
        if currentBlock.Hash != bc.CalculateHash(currentBlock) {
            return false
        }
        
        // Проверка связи с предыдущим блоком
        if currentBlock.PreviousHash != previousBlock.Hash {
            return false
        }
        
        // Проверка Proof of Work
        if !bc.IsValidHash(currentBlock.Hash) {
            return false
        }
    }
    return true
}

func (bc *Blockchain) GetBalance(address string) float64 {
    balance := 0.0
    
    for _, block := range bc.Chain {
        for _, tx := range block.Transactions {
            if tx.From == address {
                balance -= tx.Amount + tx.Fee
            }
            if tx.To == address {
                balance += tx.Amount
            }
        }
    }
    
    return balance
}

func main() {
    fmt.Println("=== Cryptocurrency Wallet and Blockchain Demo ===")
    
    // Создание кошельков
    wallet1 := NewWallet()
    wallet2 := NewWallet()
    minerWallet := NewWallet()
    
    fmt.Printf("Wallet 1: %s\n", wallet1.Address)
    fmt.Printf("Wallet 2: %s\n", wallet2.Address)
    fmt.Printf("Miner Wallet: %s\n", minerWallet.Address)
    
    // Создание блокчейна
    blockchain := NewBlockchain(2, 50.0) // Сложность 2, награда 50
    
    // Майнинг начальных блоков для создания баланса
    fmt.Println("\nMining initial blocks...")
    blockchain.MineBlock(minerWallet.Address)
    blockchain.MineBlock(minerWallet.Address)
    
    // Создание транзакций
    fmt.Println("\nCreating transactions...")
    
    // Транзакция от майнера к wallet1
    tx1 := NewTransaction(minerWallet.Address, wallet1.Address, 25.0, 0.1, 1)
    tx1.Sign(minerWallet)
    blockchain.AddTransaction(tx1)
    
    // Транзакция от wallet1 к wallet2
    tx2 := NewTransaction(wallet1.Address, wallet2.Address, 10.0, 0.1, 1)
    tx2.Sign(wallet1)
    blockchain.AddTransaction(tx2)
    
    // Майнинг блока с транзакциями
    blockchain.MineBlock(minerWallet.Address)
    
    // Проверка балансов
    fmt.Println("\n=== Balances ===")
    fmt.Printf("Miner: %.2f\n", blockchain.GetBalance(minerWallet.Address))
    fmt.Printf("Wallet 1: %.2f\n", blockchain.GetBalance(wallet1.Address))
    fmt.Printf("Wallet 2: %.2f\n", blockchain.GetBalance(wallet2.Address))
    
    // Информация о блокчейне
    fmt.Println("\n=== Blockchain Info ===")
    fmt.Printf("Blockchain length: %d\n", len(blockchain.Chain))
    fmt.Printf("Pending transactions: %d\n", len(blockchain.PendingTransactions))
    fmt.Printf("Chain valid: %t\n", blockchain.IsChainValid())
    
    // Вывод информации о блоках
    fmt.Println("\n=== Blocks ===")
    for i, block := range blockchain.Chain {
        fmt.Printf("Block %d: %s\n", i, block.Hash)
        fmt.Printf("  Transactions: %d\n", len(block.Transactions))
        fmt.Printf("  Nonce: %d\n", block.Nonce)
        
        for j, tx := range block.Transactions {
            fmt.Printf("  TX %d: %s -> %s (%.2f)\n", 
                j, tx.From, tx.To, tx.Amount)
        }
    }
    
    // Демонстрация безопасности
    fmt.Println("\n=== Security Demo ===")
    
    // Попытка подделки транзакции
    fakeTx := NewTransaction(minerWallet.Address, wallet2.Address, 100.0, 0.1, 1)
    // Не подписываем - транзакция невалидна
    if blockchain.AddTransaction(fakeTx) {
        fmt.Println("ERROR: Fake transaction accepted!")
    } else {
        fmt.Println("✓ Fake transaction correctly rejected")
    }
    
    // Проверка целостности блокчейна
    if blockchain.IsChainValid() {
        fmt.Println("✓ Blockchain integrity verified")
    } else {
        fmt.Println("ERROR: Blockchain corrupted!")
    }
}
//Задание: Криптовалютный кошелек и простой блокчейн с Proof of Work
```
160. IoT устройство - симуляция умного дома
```go
package main
import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

type SensorData struct {
    DeviceID    string  `json:"device_id"`
    Temperature float64 `json:"temperature"`
    Humidity    float64 `json:"humidity"`
    LightLevel  float64 `json:"light_level"`
    Motion      bool    `json:"motion"`
    Timestamp   int64   `json:"timestamp"`
}

type Device struct {
    ID       string
    Name     string
    Type     string
    Location string
    State    map[string]interface{}
    mu       sync.RWMutex
}

type IoTPlatform struct {
    Devices map[string]*Device
    Data    []SensorData
    mu      sync.RWMutex
}

func NewIoTPlatform() *IoTPlatform {
    return &IoTPlatform{
        Devices: make(map[string]*Device),
        Data:    make([]SensorData, 0),
    }
}

func (p *IoTPlatform) RegisterDevice(device *Device) {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    p.Devices[device.ID] = device
    fmt.Printf("Device registered: %s (%s) in %s\n", device.Name, device.Type, device.Location)
}

func (p *IoTPlatform) AddSensorData(data SensorData) {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    p.Data = append(p.Data, data)
    
    // Автоматические действия на основе данных
    p.processAutomation(data)
}

func (p *IoTPlatform) processAutomation(data SensorData) {
    // Пример автоматизации: включение света при движении в темное время
    if data.Motion && data.LightLevel < 30 {
        if light, exists := p.Devices["light_living_room"]; exists {
            light.mu.Lock()
            light.State["on"] = true
            light.State["brightness"] = 80
            light.mu.Unlock()
            fmt.Printf("Automation: Living room light turned on (motion detected)\n")
        }
    }
    
    // Автоматическое включение кондиционера при высокой температуре
    if data.Temperature > 25 {
        if ac, exists := p.Devices["ac_living_room"]; exists {
            ac.mu.Lock()
            ac.State["on"] = true
            ac.State["temperature"] = 22.0
            ac.mu.Unlock()
            fmt.Printf("Automation: AC turned on (temperature: %.1f°C)\n", data.Temperature)
        }
    }
}

func (p *IoTPlatform) GetDeviceStatus(deviceID string) map[string]interface{} {
    p.mu.RLock()
    defer p.mu.RUnlock()
    
    if device, exists := p.Devices[deviceID]; exists {
        device.mu.RLock()
        defer device.mu.RUnlock()
        
        status := make(map[string]interface{})
        for k, v := range device.State {
            status[k] = v
        }
        status["name"] = device.Name
        status["type"] = device.Type
        status["location"] = device.Location
        
        return status
    }
    return nil
}

func (p *IoTPlatform) ControlDevice(deviceID string, command map[string]interface{}) bool {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    if device, exists := p.Devices[deviceID]; exists {
        device.mu.Lock()
        defer device.mu.Unlock()
        
        for key, value := range command {
            device.State[key] = value
        }
        
        fmt.Printf("Device %s controlled: %v\n", deviceID, command)
        return true
    }
    return false
}

func (p *IoTPlatform) GetStats() map[string]interface{} {
    p.mu.RLock()
    defer p.mu.RUnlock()
    
    stats := make(map[string]interface{})
    stats["total_devices"] = len(p.Devices)
    stats["total_data_points"] = len(p.Data)
    
    // Статистика по типам устройств
    typeCount := make(map[string]int)
    for _, device := range p.Devices {
        typeCount[device.Type]++
    }
    stats["devices_by_type"] = typeCount
    
    return stats
}

func SimulateSensor(platform *IoTPlatform, deviceID string) {
    for {
        data := SensorData{
            DeviceID:   deviceID,
            Temperature: 20 + rand.Float64()*10, // 20-30°C
            Humidity:   40 + rand.Float64()*30,  // 40-70%
            LightLevel: rand.Float64() * 100,    // 0-100%
            Motion:     rand.Float64() > 0.8,    // 20% chance
            Timestamp:  time.Now().Unix(),
        }
        
        platform.AddSensorData(data)
        time.Sleep(5 * time.Second)
    }
}

func StartWebServer(platform *IoTPlatform, port string) {
    http.HandleFunc("/api/devices", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        platform.mu.RLock()
        defer platform.mu.RUnlock()
        
        devices := make([]map[string]interface{}, 0)
        for id, device := range platform.Devices {
            device.mu.RLock()
            info := map[string]interface{}{
                "id":       id,
                "name":     device.Name,
                "type":     device.Type,
                "location": device.Location,
                "state":    device.State,
            }
            device.mu.RUnlock()
            devices = append(devices, info)
        }
        
        json.NewEncoder(w).Encode(devices)
    })
    
    http.HandleFunc("/api/device/", func(w http.ResponseWriter, r *http.Request) {
        deviceID := r.URL.Path[len("/api/device/"):]
        
        switch r.Method {
        case "GET":
            status := platform.GetDeviceStatus(deviceID)
            if status == nil {
                http.NotFound(w, r)
                return
            }
            json.NewEncoder(w).Encode(status)
            
        case "POST":
            var command map[string]interface{}
            if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
                http.Error(w, "Invalid command", http.StatusBadRequest)
                return
            }
            
            if platform.ControlDevice(deviceID, command) {
                w.WriteHeader(http.StatusOK)
                fmt.Fprintf(w, `{"status": "success"}`)
            } else {
                http.NotFound(w, r)
            }
        }
    })
    
    http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
        stats := platform.GetStats()
        json.NewEncoder(w).Encode(stats)
    })
    
    http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        platform.mu.RLock()
        defer platform.mu.RUnlock()
        
        // Возвращаем последние 10 записей
        start := len(platform.Data) - 10
        if start < 0 {
            start = 0
        }
        json.NewEncoder(w).Encode(platform.Data[start:])
    })
    
    fmt.Printf("IoT Platform API server started on :%s\n", port)
    http.ListenAndServe(":"+port, nil)
}

func main() {
    platform := NewIoTPlatform()
    
    // Регистрация устройств умного дома
    devices := []*Device{
        {
            ID:       "sensor_living_room",
            Name:     "Living Room Sensor",
            Type:     "sensor",
            Location: "Living Room",
            State:    map[string]interface{}{},
        },
        {
            ID:       "light_living_room",
            Name:     "Living Room Light",
            Type:     "light",
            Location: "Living Room",
            State:    map[string]interface{}{
                "on": false,
                "brightness": 0,
                "color": "white",
            },
        },
        {
            ID:       "ac_living_room",
            Name:     "Living Room AC",
            Type:     "ac",
            Location: "Living Room",
            State:    map[string]interface{}{
                "on": false,
                "temperature": 22.0,
                "mode": "cool",
            },
        },
        {
            ID:       "thermostat_bedroom",
            Name:     "Bedroom Thermostat",
            Type:     "thermostat",
            Location: "Bedroom",
            State:    map[string]interface{}{
                "temperature": 21.0,
                "target": 22.0,
                "mode": "heat",
            },
        },
    }
    
    for _, device := range devices {
        platform.RegisterDevice(device)
    }
    
    // Запуск симуляции сенсоров
    go SimulateSensor(platform, "sensor_living_room")
    
    // Запуск веб-сервера
    go StartWebServer(platform, "8080")
    
    // Демонстрация работы системы
    fmt.Println("\n=== Smart Home IoT Platform Demo ===")
    fmt.Println("API endpoints:")
    fmt.Println("  GET  /api/devices - List all devices")
    fmt.Println("  GET  /api/device/{id} - Get device status")
    fmt.Println("  POST /api/device/{id} - Control device")
    fmt.Println("  GET  /api/stats - Platform statistics")
    fmt.Println("  GET  /api/data - Sensor data")
    
    // Демонстрация управления устройствами
    time.Sleep(2 * time.Second)
    
    fmt.Println("\n=== Device Control Demo ===")
    
    // Включение света
    platform.ControlDevice("light_living_room", map[string]interface{}{
        "on": true,
        "brightness": 75,
    })
    
    // Установка температуры на AC
    platform.ControlDevice("ac_living_room", map[string]interface{}{
        "on": true,
        "temperature": 23.0,
        "mode": "cool",
    })
    
    // Бесконечный цикл для отображения обновлений
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        stats := platform.GetStats()
        fmt.Printf("\nSystem Stats: %d devices, %d data points\n", 
            stats["total_devices"], stats["total_data_points"])
        
        // Показ статуса основных устройств
        lightStatus := platform.GetDeviceStatus("light_living_room")
        acStatus := platform.GetDeviceStatus("ac_living_room")
        
        fmt.Printf("Light: %v, AC: %v\n", 
            lightStatus["on"], acStatus["on"])
    }
}
//Задание: IoT платформа для умного дома с автоматизацией и REST API
```
161. Квантовые вычисления - симуляция кубитов
```go
package main
import (
    "fmt"
    "math"
    "math/cmplx"
    "math/rand"
)

type Qubit struct {
    Alpha complex128 // Амплитуда для |0⟩
    Beta  complex128 // Амплитуда для |1⟩
}

type QuantumGate struct {
    Matrix [2][2]complex128
}

type QuantumCircuit struct {
    Qubits []*Qubit
    Gates  []GateOperation
}

type GateOperation struct {
    Gate    *QuantumGate
    Targets []int
}

func NewQubit() *Qubit {
    return &Qubit{
        Alpha: complex(1, 0), // |0⟩ состояние
        Beta:  complex(0, 0),
    }
}

func (q *Qubit) Measure() int {
    prob0 := real(q.Alpha * complex(real(q.Alpha), -imag(q.Alpha))) // |α|²
    // prob1 := real(q.Beta * complex(real(q.Beta), -imag(q.Beta))) // |β|²
    
    if rand.Float64() < prob0 {
        return 0
    } else {
        return 1
    }
}

func (q *Qubit) ApplyGate(gate *QuantumGate) {
    newAlpha := gate.Matrix[0][0]*q.Alpha + gate.Matrix[0][1]*q.Beta
    newBeta := gate.Matrix[1][0]*q.Alpha + gate.Matrix[1][1]*q.Beta
    
    q.Alpha = newAlpha
    q.Beta = newBeta
}

func (q *Qubit) String() string {
    return fmt.Sprintf("α|0⟩ + β|1⟩ = (%.3f%+.3fi)|0⟩ + (%.3f%+.3fi)|1⟩", 
        real(q.Alpha), imag(q.Alpha), real(q.Beta), imag(q.Beta))
}

// Квантовые гейты
var (
    // Гейт Адамара (создание суперпозиции)
    H = &QuantumGate{
        Matrix: [2][2]complex128{
            {complex(1/math.Sqrt2, 0), complex(1/math.Sqrt2, 0)},
            {complex(1/math.Sqrt2, 0), complex(-1/math.Sqrt2, 0)},
        },
    }
    
    // Гейт Паули-X (NOT)
    X = &QuantumGate{
        Matrix: [2][2]complex128{
            {0, 1},
            {1, 0},
        },
    }
    
    // Гейт Паули-Y
    Y = &QuantumGate{
        Matrix: [2][2]complex128{
            {0, complex(0, -1)},
            {complex(0, 1), 0},
        },
    }
    
    // Гейт Паули-Z
    Z = &QuantumGate{
        Matrix: [2][2]complex128{
            {1, 0},
            {0, -1},
        },
    }
)

func NewQuantumCircuit(numQubits int) *QuantumCircuit {
    qubits := make([]*Qubit, numQubits)
    for i := range qubits {
        qubits[i] = NewQubit()
    }
    
    return &QuantumCircuit{
        Qubits: qubits,
        Gates:  make([]GateOperation, 0),
    }
}

func (qc *QuantumCircuit) ApplyGate(gate *QuantumGate, targets ...int) {
    for _, target := range targets {
        if target < len(qc.Qubits) {
            qc.Qubits[target].ApplyGate(gate)
        }
    }
    qc.Gates = append(qc.Gates, GateOperation{Gate: gate, Targets: targets})
}

func (qc *QuantumCircuit) Measure() []int {
    results := make([]int, len(qc.Qubits))
    for i, qubit := range qc.Qubits {
        results[i] = qubit.Measure()
    }
    return results
}

func (qc *QuantumCircuit) Run(shots int) map[string]int {
    results := make(map[string]int)
    
    for i := 0; i < shots; i++ {
        // Сброс кубитов в |0⟩ состояние
        for _, qubit := range qc.Qubits {
            qubit.Alpha = complex(1, 0)
            qubit.Beta = complex(0, 0)
        }
        
        // Применение всех гейтов
        for _, gateOp := range qc.Gates {
            for _, target := range gateOp.Targets {
                qc.Qubits[target].ApplyGate(gateOp.Gate)
            }
        }
        
        // Измерение
        measurement := qc.Measure()
        
        // Конвертация в бинарную строку
        var binStr string
        for _, bit := range measurement {
            if bit == 1 {
                binStr += "1"
            } else {
                binStr += "0"
            }
        }
        
        results[binStr]++
    }
    
    return results
}

// Квантовая телепортация
func QuantumTeleportation() {
    fmt.Println("=== Quantum Teleportation Demo ===")
    
    // Создание трех кубитов:
    // q0 - кубит для телепортации
    // q1, q2 - запутанная пара
    circuit := NewQuantumCircuit(3)
    
    // Шаг 1: Создание запутанной пары (q1 и q2)
    circuit.ApplyGate(H, 1)        // Применение Адамара к q1
    circuit.ApplyGate(X, 2)        // CNOT эквивалент
    // Здесь должен быть CNOT гейт, но для простоты используем X
    
    // Шаг 2: Подготовка кубита для телепортации (q0)
    // Создаем произвольное состояние
    circuit.ApplyGate(H, 0)
    circuit.ApplyGate(Z, 0)
    
    fmt.Println("Initial states:")
    for i, qubit := range circuit.Qubits {
        fmt.Printf("Qubit %d: %s\n", i, qubit)
    }
    
    // Шаг 3: Процесс телепортации
    circuit.ApplyGate(X, 1) // CNOT(q0, q1)
    circuit.ApplyGate(H, 0) // Адамара на q0
    
    // Измерение q0 и q1
    m0 := circuit.Qubits[0].Measure()
    m1 := circuit.Qubits[1].Measure()
    
    fmt.Printf("\nMeasurement results: q0=%d, q1=%d\n", m0, m1)
    
    // Коррекция на q2 на основе измерений
    if m1 == 1 {
        circuit.ApplyGate(X, 2)
    }
    if m0 == 1 {
        circuit.ApplyGate(Z, 2)
    }
    
    fmt.Printf("Qubit 2 after correction: %s\n", circuit.Qubits[2])
    fmt.Println("Teleportation completed!")
}

// Алгоритм Дойча-Йожи
func DeutschJozsa(n int) {
    fmt.Printf("\n=== Deutsch-Jozsa Algorithm (n=%d) ===\n", n)
    
    circuit := NewQuantumCircuit(n + 1) // n входных кубитов + 1 вспомогательный
    
    // Инициализация вспомогательного кубита в |1⟩
    circuit.ApplyGate(X, n)
    
    // Применение Адамара ко всем кубитам
    for i := 0; i <= n; i++ {
        circuit.ApplyGate(H, i)
    }
    
    // Здесь должна быть оракул-функция
    // Для демонстрации используем сбалансированную функцию
    circuit.ApplyGate(X, n) // Простой оракул
    
    // Снова применяем Адамара к входным кубитам
    for i := 0; i < n; i++ {
        circuit.ApplyGate(H, i)
    }
    
    // Измерение входных кубитов
    results := circuit.Run(1000)
    
    fmt.Println("Measurement results:")
    for state, count := range results {
        fmt.Printf("|%s⟩: %d\n", state, count)
    }
    
    // Анализ результатов
    allZero := true
    for state := range results {
        if state[:n] != "000" { // Проверяем только входные кубиты
            allZero = false
            break
        }
    }
    
    if allZero {
        fmt.Println("Function is CONSTANT")
    } else {
        fmt.Println("Function is BALANCED")
    }
}

func main() {
    rand.Seed(42)
    
    fmt.Println("Quantum Computing Simulator")
    
    // Демонстрация базовых операций с кубитами
    fmt.Println("\n=== Basic Qubit Operations ===")
    
    qubit := NewQubit()
    fmt.Printf("Initial state: %s\n", qubit)
    
    qubit.ApplyGate(H)
    fmt.Printf("After Hadamard: %s\n", qubit)
    
    measurement := qubit.Measure()
    fmt.Printf("Measurement result: %d\n", measurement)
    
    // Демонстрация квантовой схемы
    fmt.Println("\n=== Quantum Circuit Demo ===")
    
    circuit := NewQuantumCircuit(2)
    circuit.ApplyGate(H, 0) // Суперпозиция первого кубита
    circuit.ApplyGate(X, 1) // NOT второго кубита
    
    results := circuit.Run(1000)
    fmt.Println("Measurement results (1000 shots):")
    for state, count := range results {
        fmt.Printf("|%s⟩: %d (%.1f%%)\n", state, count, float64(count)/10)
    }
    
    // Квантовая телепортация
    QuantumTeleportation()
    
    // Алгоритм Дойча-Йожи
    DeutschJozsa(3)
    
    // Квантовый случайный генератор чисел
    fmt.Println("\n=== Quantum Random Number Generator ===")
    qrngCircuit := NewQuantumCircuit(8)
    for i := 0; i < 8; i++ {
        qrngCircuit.ApplyGate(H, i)
    }
    
    qrngResults := qrngCircuit.Run(1)
    for state := range qrngResults {
        fmt.Printf("Random 8-bit number: %s (decimal: ", state)
        // Конвертация в decimal
        var decimal int
        for i, bit := range state {
            if bit == '1' {
                decimal += 1 << (7 - i)
            }
        }
        fmt.Printf("%d)\n", decimal)
    }
    
    // Демонстрация квантовой суперпозиции
    fmt.Println("\n=== Quantum Superposition Demo ===")
    superposCircuit := NewQuantumCircuit(1)
    superposCircuit.ApplyGate(H, 0)
    
    superposResults := superposCircuit.Run(10000)
    fmt.Println("Superposition measurement (10000 shots):")
    for state, count := range superposResults {
        fmt.Printf("|%s⟩: %d (%.1f%%)\n", state, count, float64(count)/100)
    }
}
//Задание: Симулятор квантовых вычислений с кубитами и гейтами
```
Компьютерная графика и рендеринг 
162. 3D рендеринг - программный растеризатор
```go
package main
import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "math"
    "os"
)

type Vector3 struct {
    X, Y, Z float64
}

type Matrix4 struct {
    Data [4][4]float64
}

type Triangle struct {
    Vertices [3]Vector3
    Color    color.RGBA
}

type Mesh struct {
    Triangles []Triangle
}

type Camera struct {
    Position Vector3
    Target   Vector3
    Up       Vector3
    FOV      float64
}

type Renderer struct {
    Width   int
    Height  int
    ZBuffer []float64
}

func NewVector3(x, y, z float64) Vector3 {
    return Vector3{X: x, Y: y, Z: z}
}

func (v Vector3) Add(other Vector3) Vector3 {
    return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v Vector3) Subtract(other Vector3) Vector3 {
    return Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v Vector3) Multiply(scalar float64) Vector3 {
    return Vector3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v Vector3) Dot(other Vector3) float64 {
    return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vector3) Cross(other Vector3) Vector3 {
    return Vector3{
        X: v.Y*other.Z - v.Z*other.Y,
        Y: v.Z*other.X - v.X*other.Z,
        Z: v.X*other.Y - v.Y*other.X,
    }
}

func (v Vector3) Length() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3) Normalize() Vector3 {
    length := v.Length()
    if length == 0 {
        return v
    }
    return Vector3{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}

func NewMatrix4() Matrix4 {
    return Matrix4{
        Data: [4][4]float64{
            {1, 0, 0, 0},
            {0, 1, 0, 0},
            {0, 0, 1, 0},
            {0, 0, 0, 1},
        },
    }
}

func IdentityMatrix() Matrix4 {
    return NewMatrix4()
}

func LookAtMatrix(eye, target, up Vector3) Matrix4 {
    forward := target.Subtract(eye).Normalize()
    right := forward.Cross(up).Normalize()
    up = right.Cross(forward).Normalize()

    return Matrix4{
        Data: [4][4]float64{
            {right.X, right.Y, right.Z, -right.Dot(eye)},
            {up.X, up.Y, up.Z, -up.Dot(eye)},
            {-forward.X, -forward.Y, -forward.Z, forward.Dot(eye)},
            {0, 0, 0, 1},
        },
    }
}

func PerspectiveMatrix(fov, aspect, near, far float64) Matrix4 {
    tanHalfFov := math.Tan(fov * 0.5)
    
    return Matrix4{
        Data: [4][4]float64{
            {1.0 / (aspect * tanHalfFov), 0, 0, 0},
            {0, 1.0 / tanHalfFov, 0, 0},
            {0, 0, -(far + near) / (far - near), -(2.0 * far * near) / (far - near)},
            {0, 0, -1, 0},
        },
    }
}

func (m Matrix4) MultiplyVector(v Vector3) Vector3 {
    x := m.Data[0][0]*v.X + m.Data[0][1]*v.Y + m.Data[0][2]*v.Z + m.Data[0][3]
    y := m.Data[1][0]*v.X + m.Data[1][1]*v.Y + m.Data[1][2]*v.Z + m.Data[1][3]
    z := m.Data[2][0]*v.X + m.Data[2][1]*v.Y + m.Data[2][2]*v.Z + m.Data[2][3]
    w := m.Data[3][0]*v.X + m.Data[3][1]*v.Y + m.Data[3][2]*v.Z + m.Data[3][3]
    
    if w != 0 {
        x /= w
        y /= w
        z /= w
    }
    
    return Vector3{X: x, Y: y, Z: z}
}

func NewRenderer(width, height int) *Renderer {
    zBuffer := make([]float64, width*height)
    for i := range zBuffer {
        zBuffer[i] = math.MaxFloat64
    }
    
    return &Renderer{
        Width:   width,
        Height:  height,
        ZBuffer: zBuffer,
    }
}

func (r *Renderer) ClearZBuffer() {
    for i := range r.ZBuffer {
        r.ZBuffer[i] = math.MaxFloat64
    }
}

func (r *Renderer) WorldToScreen(v Vector3) (int, int) {
    x := int((v.X + 1.0) * 0.5 * float64(r.Width))
    y := int((1.0 - (v.Y + 1.0) * 0.5) * float64(r.Height))
    return x, y
}

func (r *Renderer) DrawTriangle(img *image.RGBA, tri Triangle, viewProj Matrix4) {
    // Transform vertices
    var screenVerts [3]Vector3
    for i, vert := range tri.Vertices {
        transformed := viewProj.MultiplyVector(vert)
        screenVerts[i] = transformed
    }
    
    // Convert to screen coordinates
    var screenCoords [3]struct{ X, Y int }
    for i, vert := range screenVerts {
        x, y := r.WorldToScreen(vert)
        screenCoords[i].X = x
        screenCoords[i].Y = y
    }
    
    // Find bounding box
    minX := min(screenCoords[0].X, screenCoords[1].X, screenCoords[2].X)
    maxX := max(screenCoords[0].X, screenCoords[1].X, screenCoords[2].X)
    minY := min(screenCoords[0].Y, screenCoords[1].Y, screenCoords[2].Y)
    maxY := max(screenCoords[0].Y, screenCoords[1].Y, screenCoords[2].Y)
    
    // Clamp to screen bounds
    minX = max(0, minX)
    maxX = min(r.Width-1, maxX)
    minY = max(0, minY)
    maxY = min(r.Height-1, maxY)
    
    // Rasterize
    for y := minY; y <= maxY; y++ {
        for x := minX; x <= maxX; x++ {
            if r.IsPointInTriangle(x, y, screenCoords) {
                // Simple z-buffering
                z := r.CalculateZ(x, y, screenVerts)
                idx := y*r.Width + x
                
                if z < r.ZBuffer[idx] {
                    r.ZBuffer[idx] = z
                    img.Set(x, y, tri.Color)
                }
            }
        }
    }
}

func (r *Renderer) IsPointInTriangle(x, y int, coords [3]struct{ X, Y int }) bool {
    // Barycentric coordinate check
    v0 := [2]float64{float64(coords[2].X - coords[0].X), float64(coords[2].Y - coords[0].Y)}
    v1 := [2]float64{float64(coords[1].X - coords[0].X), float64(coords[1].Y - coords[0].Y)}
    v2 := [2]float64{float64(x - coords[0].X), float64(y - coords[0].Y)}
    
    dot00 := v0[0]*v0[0] + v0[1]*v0[1]
    dot01 := v0[0]*v1[0] + v0[1]*v1[1]
    dot02 := v0[0]*v2[0] + v0[1]*v2[1]
    dot11 := v1[0]*v1[0] + v1[1]*v1[1]
    dot12 := v1[0]*v2[0] + v1[1]*v2[1]
    
    invDenom := 1.0 / (dot00*dot11 - dot01*dot01)
    u := (dot11*dot02 - dot01*dot12) * invDenom
    v := (dot00*dot12 - dot01*dot02) * invDenom
    
    return (u >= 0) && (v >= 0) && (u + v < 1)
}

func (r *Renderer) CalculateZ(x, y int, verts [3]Vector3) float64 {
    // Simple interpolation (в реальном рендерере нужно использовать барицентрические координаты)
    return (verts[0].Z + verts[1].Z + verts[2].Z) / 3.0
}

func CreateCubeMesh() *Mesh {
    // Вершины куба
    vertices := []Vector3{
        {-1, -1, -1}, {1, -1, -1}, {1, 1, -1}, {-1, 1, -1}, // задняя грань
        {-1, -1, 1}, {1, -1, 1}, {1, 1, 1}, {-1, 1, 1},     // передняя грань
    }
    
    // Треугольники куба (по 2 на грань)
    triangles := []Triangle{
        // Задняя грань
        {Vertices: [3]Vector3{vertices[0], vertices[1], vertices[2]}, Color: color.RGBA{255, 0, 0, 255}},
        {Vertices: [3]Vector3{vertices[0], vertices[2], vertices[3]}, Color: color.RGBA{255, 0, 0, 255}},
        // Передняя грань
        {Vertices: [3]Vector3{vertices[4], vertices[6], vertices[5]}, Color: color.RGBA{0, 255, 0, 255}},
        {Vertices: [3]Vector3{vertices[4], vertices[7], vertices[6]}, Color: color.RGBA{0, 255, 0, 255}},
        // Левая грань
        {Vertices: [3]Vector3{vertices[0], vertices[3], vertices[7]}, Color: color.RGBA{0, 0, 255, 255}},
        {Vertices: [3]Vector3{vertices[0], vertices[7], vertices[4]}, Color: color.RGBA{0, 0, 255, 255}},
        // Правая грань
        {Vertices: [3]Vector3{vertices[1], vertices[5], vertices[6]}, Color: color.RGBA{255, 255, 0, 255}},
        {Vertices: [3]Vector3{vertices[1], vertices[6], vertices[2]}, Color: color.RGBA{255, 255, 0, 255}},
        // Верхняя грань
        {Vertices: [3]Vector3{vertices[3], vertices[2], vertices[6]}, Color: color.RGBA{255, 0, 255, 255}},
        {Vertices: [3]Vector3{vertices[3], vertices[6], vertices[7]}, Color: color.RGBA{255, 0, 255, 255}},
        // Нижняя грань
        {Vertices: [3]Vector3{vertices[0], vertices[4], vertices[5]}, Color: color.RGBA{0, 255, 255, 255}},
        {Vertices: [3]Vector3{vertices[0], vertices[5], vertices[1]}, Color: color.RGBA{0, 255, 255, 255}},
    }
    
    return &Mesh{Triangles: triangles}
}

func main() {
    width, height := 800, 600
    
    // Создание рендерера
    renderer := NewRenderer(width, height)
    
    // Создание изображения
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    
    // Заливка фона
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            img.Set(x, y, color.RGBA{50, 50, 50, 255})
        }
    }
    
    // Создание меша куба
    cube := CreateCubeMesh()
    
    // Настройка камеры
    camera := Camera{
        Position: NewVector3(0, 0, -5),
        Target:   NewVector3(0, 0, 0),
        Up:       NewVector3(0, 1, 0),
        FOV:      math.Pi / 3, // 60 градусов
    }
    
    // Матрицы вида и проекции
    viewMatrix := LookAtMatrix(camera.Position, camera.Target, camera.Up)
    aspect := float64(width) / float64(height)
    projMatrix := PerspectiveMatrix(camera.FOV, aspect, 0.1, 100.0)
    
    // Комбинированная матрица
    viewProj := projMatrix // В реальном рендерере нужно умножить viewMatrix * projMatrix
    
    // Рендеринг
    fmt.Println("Rendering 3D cube...")
    for _, tri := range cube.Triangles {
        renderer.DrawTriangle(img, tri, viewProj)
    }
    
    // Сохранение изображения
    file, err := os.Create("cube.png")
    if err != nil {
        fmt.Printf("Error creating image: %v\n", err)
        return
    }
    defer file.Close()
    
    png.Encode(file, img)
    fmt.Println("3D rendering saved: cube.png")
    
    // Демонстрация векторных операций
    fmt.Println("\n=== Vector Operations Demo ===")
    v1 := NewVector3(1, 2, 3)
    v2 := NewVector3(4, 5, 6)
    
    fmt.Printf("v1: (%.1f, %.1f, %.1f)\n", v1.X, v1.Y, v1.Z)
    fmt.Printf("v2: (%.1f, %.1f, %.1f)\n", v2.X, v2.Y, v2.Z)
    fmt.Printf("Dot product: %.1f\n", v1.Dot(v2))
    fmt.Printf("Cross product: (%.1f, %.1f, %.1f)\n", 
        v1.Cross(v2).X, v1.Cross(v2).Y, v1.Cross(v2).Z)
    fmt.Printf("Length of v1: %.1f\n", v1.Length())
}

func min(a, b, c int) int {
    return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

func max(a, b, c int) int {
    return int(math.Max(float64(a), math.Max(float64(b), float64(c))))
}
//Задание: Программный 3D рендерер с растеризацией треугольников и z-буферингом
```
163. Робототехника - симулятор движения робота
```go
package main
import (
    "fmt"
    "math"
    "time"
)

type Vector2D struct {
    X, Y float64
}

type Pose2D struct {
    Position Vector2D
    Theta    float64 // Orientation in radians
}

type DifferentialDriveRobot struct {
    Pose          Pose2D
    WheelRadius   float64
    WheelBase     float64
    LeftVelocity  float64
    RightVelocity float64
    MaxVelocity   float64
}

type LidarSensor struct {
    MaxRange     float64
    MinRange     float64
    NumRays      int
    FieldOfView  float64
}

type Environment struct {
    Obstacles []Obstacle
    Width     float64
    Height    float64
}

type Obstacle struct {
    Position Vector2D
    Radius   float64
}

func NewDifferentialDriveRobot(wheelRadius, wheelBase, maxVelocity float64) *DifferentialDriveRobot {
    return &DifferentialDriveRobot{
        Pose:        Pose2D{Position: Vector2D{X: 0, Y: 0}, Theta: 0},
        WheelRadius: wheelRadius,
        WheelBase:   wheelBase,
        MaxVelocity: maxVelocity,
    }
}

func (r *DifferentialDriveRobot) SetVelocities(leftVel, rightVel float64) {
    // Ограничение скоростей
    r.LeftVelocity = math.Max(-r.MaxVelocity, math.Min(r.MaxVelocity, leftVel))
    r.RightVelocity = math.Max(-r.MaxVelocity, math.Min(r.MaxVelocity, rightVel))
}

func (r *DifferentialDriveRobot) Update(dt float64) {
    // Кинематика дифференциального привода
    v := (r.WheelRadius / 2) * (r.RightVelocity + r.LeftVelocity)
    w := (r.WheelRadius / r.WheelBase) * (r.RightVelocity - r.LeftVelocity)
    
    // Обновление позы
    if math.Abs(w) < 1e-6 {
        // Движение по прямой
        r.Pose.Position.X += v * math.Cos(r.Pose.Theta) * dt
        r.Pose.Position.Y += v * math.Sin(r.Pose.Theta) * dt
    } else {
        // Движение по дуге
        radius := v / w
        iccX := r.Pose.Position.X - radius * math.Sin(r.Pose.Theta)
        iccY := r.Pose.Position.Y + radius * math.Cos(r.Pose.Theta)
        
        r.Pose.Theta += w * dt
        r.Pose.Position.X = math.Cos(w*dt)*(r.Pose.Position.X-iccX) - 
                           math.Sin(w*dt)*(r.Pose.Position.Y-iccY) + iccX
        r.Pose.Position.Y = math.Sin(w*dt)*(r.Pose.Position.X-iccX) + 
                           math.Cos(w*dt)*(r.Pose.Position.Y-iccY) + iccY
    }
    
    // Нормализация угла
    r.Pose.Theta = math.Mod(r.Pose.Theta, 2*math.Pi)
    if r.Pose.Theta < 0 {
        r.Pose.Theta += 2 * math.Pi
    }
}

func (r *DifferentialDriveRobot) MoveTo(target Pose2D, kpLinear, kpAngular float64) {
    // Расчет ошибки
    dx := target.Position.X - r.Pose.Position.X
    dy := target.Position.Y - r.Pose.Position.Y
    distance := math.Sqrt(dx*dx + dy*dy)
    
    // Желаемый угол
    desiredTheta := math.Atan2(dy, dx)
    angleError := math.Mod(desiredTheta - r.Pose.Theta + math.Pi, 2*math.Pi) - math.Pi
    
    // Пропорциональное управление
    linearVel := kpLinear * distance
    angularVel := kpAngular * angleError
    
    // Конвертация в скорости колес
    leftVel := (2*linearVel - angularVel*r.WheelBase) / (2 * r.WheelRadius)
    rightVel := (2*linearVel + angularVel*r.WheelBase) / (2 * r.WheelRadius)
    
    r.SetVelocities(leftVel, rightVel)
}

func NewLidarSensor(maxRange, minRange, fov float64, numRays int) *LidarSensor {
    return &LidarSensor{
        MaxRange:    maxRange,
        MinRange:    minRange,
        NumRays:     numRays,
        FieldOfView: fov,
    }
}

func (l *LidarSensor) Scan(robotPose Pose2D, env *Environment) []float64 {
    ranges := make([]float64, l.NumRays)
    angleStep := l.FieldOfView / float64(l.NumRays-1)
    
    for i := 0; i < l.NumRays; i++ {
        angle := robotPose.Theta - l.FieldOfView/2 + float64(i)*angleStep
        ranges[i] = l.castRay(robotPose.Position, angle, env)
    }
    
    return ranges
}

func (l *LidarSensor) castRay(origin Vector2D, angle float64, env *Environment) float64 {
    direction := Vector2D{
        X: math.Cos(angle),
        Y: math.Sin(angle),
    }
    
    minDistance := l.MaxRange
    
    for _, obs := range env.Obstacles {
        // Проверка пересечения луча с окружностью
        oc := Vector2D{
            X: origin.X - obs.Position.X,
            Y: origin.Y - obs.Position.Y,
        }
        
        a := direction.X*direction.X + direction.Y*direction.Y
        b := 2 * (oc.X*direction.X + oc.Y*direction.Y)
        c := oc.X*oc.X + oc.Y*oc.Y - obs.Radius*obs.Radius
        
        discriminant := b*b - 4*a*c
        
        if discriminant >= 0 {
            t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
            t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
            
            for _, t := range []float64{t1, t2} {
                if t >= 0 && t < minDistance {
                    minDistance = t
                }
            }
        }
    }
    
    // Проверка границ среды
    if direction.X != 0 {
        t := (env.Width/2 - origin.X) / direction.X
        if t > 0 && t < minDistance {
            minDistance = t
        }
        
        t = (-env.Width/2 - origin.X) / direction.X
        if t > 0 && t < minDistance {
            minDistance = t
        }
    }
    
    if direction.Y != 0 {
        t := (env.Height/2 - origin.Y) / direction.Y
        if t > 0 && t < minDistance {
            minDistance = t
        }
        
        t = (-env.Height/2 - origin.Y) / direction.Y
        if t > 0 && t < minDistance {
            minDistance = t
        }
    }
    
    return math.Max(l.MinRange, math.Min(l.MaxRange, minDistance))
}

func NewEnvironment(width, height float64) *Environment {
    obstacles := []Obstacle{
        {Position: Vector2D{X: 2, Y: 2}, Radius: 0.5},
        {Position: Vector2D{X: -1, Y: 3}, Radius: 0.7},
        {Position: Vector2D{X: 3, Y: -2}, Radius: 0.3},
        {Position: Vector2D{X: -2, Y: -1}, Radius: 0.6},
    }
    
    return &Environment{
        Obstacles: obstacles,
        Width:     width,
        Height:    height,
    }
}

type PIDController struct {
    Kp, Ki, Kd float64
    Integral   float64
    PreviousError float64
}

func NewPIDController(kp, ki, kd float64) *PIDController {
    return &PIDController{
        Kp: kp,
        Ki: ki,
        Kd: kd,
    }
}

func (pid *PIDController) Compute(error, dt float64) float64 {
    pid.Integral += error * dt
    derivative := (error - pid.PreviousError) / dt
    
    output := pid.Kp*error + pid.Ki*pid.Integral + pid.Kd*derivative
    
    pid.PreviousError = error
    return output
}

func main() {
    fmt.Println("=== Robotics Simulation: Differential Drive Robot ===")
    
    // Создание робота
    robot := NewDifferentialDriveRobot(0.1, 0.5, 5.0)
    
    // Создание среды
    env := NewEnvironment(20.0, 20.0)
    
    // Создание лидара
    lidar := NewLidarSensor(10.0, 0.1, math.Pi*2, 360)
    
    // Создание ПИД-контроллера
    pid := NewPIDController(1.0, 0.1, 0.05)
    
    // Целевые точки для следования
    waypoints := []Pose2D{
        {Position: Vector2D{X: 5, Y: 0}, Theta: 0},
        {Position: Vector2D{X: 5, Y: 5}, Theta: math.Pi / 2},
        {Position: Vector2D{X: 0, Y: 5}, Theta: math.Pi},
        {Position: Vector2D{X: 0, Y: 0}, Theta: 3 * math.Pi / 2},
    }
    
    currentWaypoint := 0
    waypointThreshold := 0.1
    
    // Симуляция
    dt := 0.1 // 100 ms
    for step := 0; step < 1000; step++ {
        // Проверка достижения целевой точки
        target := waypoints[currentWaypoint]
        dx := target.Position.X - robot.Pose.Position.X
        dy := target.Position.Y - robot.Pose.Position.Y
        distance := math.Sqrt(dx*dx + dy*dy)
        
        if distance < waypointThreshold {
            currentWaypoint = (currentWaypoint + 1) % len(waypoints)
            fmt.Printf("Reached waypoint %d, moving to next\n", currentWaypoint)
        }
        
        // Управление движением к целевой точке
        robot.MoveTo(waypoints[currentWaypoint], 0.5, 2.0)
        
        // Сканирование окружения
        ranges := lidar.Scan(robot.Pose, env)
        
        // Обнаружение препятствий и избегание
        minRange := lidar.MaxRange
        for _, r := range ranges {
            if r < minRange {
                minRange = r
            }
        }
        
        // Избегание препятствий с помощью ПИД-регулятора
        if minRange < 2.0 {
            avoidanceForce := pid.Compute(2.0-minRange, dt)
            // Корректировка скоростей для избегания
            robot.LeftVelocity += avoidanceForce
            robot.RightVelocity -= avoidanceForce
        }
        
        // Обновление состояния робота
        robot.Update(dt)
        
        // Вывод информации
        if step%10 == 0 {
            fmt.Printf("Step %d: Pos(%.2f, %.2f), Theta: %.2f, Target: %d\n",
                step, robot.Pose.Position.X, robot.Pose.Position.Y,
                robot.Pose.Theta, currentWaypoint)
            
            // Визуализация лидара (первые 8 лучей)
            fmt.Printf("Lidar: ")
            for i := 0; i < 8 && i < len(ranges); i++ {
                fmt.Printf("%.1f ", ranges[i])
            }
            fmt.Printf("...\n")
        }
        
        time.Sleep(time.Duration(dt * 1000) * time.Millisecond)
    }
    
    // Демонстрация кинематики
    fmt.Println("\n=== Kinematics Demo ===")
    testRobot := NewDifferentialDriveRobot(0.1, 0.5, 5.0)
    
    // Движение по прямой
    fmt.Println("Moving straight...")
    testRobot.SetVelocities(2.0, 2.0)
    for i := 0; i < 10; i++ {
        testRobot.Update(0.1)
        fmt.Printf("Position: (%.2f, %.2f)\n", testRobot.Pose.Position.X, testRobot.Pose.Position.Y)
    }
    
    // Поворот на месте
    fmt.Println("\nTurning in place...")
    testRobot.Pose = Pose2D{Position: Vector2D{X: 0, Y: 0}, Theta: 0}
    testRobot.SetVelocities(-2.0, 2.0)
    for i := 0; i < 10; i++ {
        testRobot.Update(0.1)
        fmt.Printf("Theta: %.2f rad\n", testRobot.Pose.Theta)
    }
    
    // Движение по дуге
    fmt.Println("\nMoving in arc...")
    testRobot.Pose = Pose2D{Position: Vector2D{X: 0, Y: 0}, Theta: 0}
    testRobot.SetVelocities(1.0, 2.0)
    for i := 0; i < 20; i++ {
        testRobot.Update(0.1)
        fmt.Printf("Position: (%.2f, %.2f), Theta: %.2f\n",
            testRobot.Pose.Position.X, testRobot.Pose.Position.Y, testRobot.Pose.Theta)
    }
}
//Задание: Симулятор робота с дифференциальным приводом, лидаром и ПИД-регулятором
<<<<<<< HEAD
```
=======
>>>>>>> 42172ba081a2664a093f85f690fc0c553121547b
