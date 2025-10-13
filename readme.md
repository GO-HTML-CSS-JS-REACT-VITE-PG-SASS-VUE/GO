GO напиши 200 простых примеров кодов на использование GO с нуля, чтобы я понял с объяснением в коде, что и зачем это в каждой строке
 200+ примеров GO с нуля, разбитых на логические разделы. Начнем с основ и постепенно перейдем к более сложным концепциям.
ещё ко всему, надо сделать так, чтобы после копирования кодов сами коды при вставке в ворд оказывались бы в каждый в своей ячейке таблицы 1-я колонка нумерация вторая код и так до 200 кодов? 3-я задание с объяснением цели в кратце  - 
 200 примеров GO в формате таблицы для удобного копирования в Word. Каждый пример будет в отдельной строке таблицы с нумерацией и задание с объяснением цели в кратце.
Я создам для вас 200+ примеров Go с нуля в формате таблицы. Вот первые 50 примеров:

Основы Go
1. Простая программа
go
package main  // Объявление основного пакета

import "fmt"  // Импорт пакета для ввода-вывода

func main() {  // Главная функция программы
    fmt.Println("Hello, World!")  // Вывод текста в консоль
}
Задание: Базовая структура программы Go

2. Объявление переменных
go
package main
import "fmt"

func main() {
    var name string = "Alice"  // Явное объявление переменной
    age := 25                  // Короткое объявление с выводом типа
    fmt.Println(name, age)     // Вывод значений
}
Задание: Демонстрация способов объявления переменных

3. Константы
go
package main
import "fmt"

func main() {
    const Pi = 3.14159        // Объявление константы
    const AppName = "MyApp"   // Строковая константа
    fmt.Println(Pi, AppName)  // Вывод констант
}
Задание: Работа с константами в Go

4. Основные типы данных
go
package main
import "fmt"

func main() {
    var integer int = 42          // Целое число
    var floating float64 = 3.14   // Число с плавающей точкой
    var text string = "Go lang"   // Строка
    var boolean bool = true       // Логический тип
    fmt.Println(integer, floating, text, boolean)
}
Задание: Демонстрация основных типов данных

5. Арифметические операции
go
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
Задание: Базовые арифметические операции

6. Строковые операции
go
package main
import "fmt"

func main() {
    firstName := "John"
    lastName := "Doe"
    fullName := firstName + " " + lastName  // Конкатенация строк
    length := len(fullName)                 // Длина строки
    fmt.Println(fullName, "Длина:", length)
}
Задание: Операции со строками

7. Условные операторы
go
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
Задание: Базовые условные конструкции

8. Множественные условия
go
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
Задание: Использование else if для множественных условий

9. Оператор switch
go
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
Задание: Использование switch для выбора

10. Цикл for
go
package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {      // Классический цикл for
        fmt.Println("Итерация:", i)
    }
}
Задание: Базовый цикл for

11. While-подобный цикл
go
package main
import "fmt"

func main() {
    count := 1
    for count <= 5 {               // Цикл while-style
        fmt.Println("Счетчик:", count)
        count++
    }
}
Задание: Цикл с условием (аналог while)

12. Бесконечный цикл
go
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
Задание: Бесконечный цикл с break

13. Массивы
go
package main
import "fmt"

func main() {
    var numbers [3]int           // Объявление массива
    numbers[0] = 10              // Присвоение значения
    numbers[1] = 20
    numbers[2] = 30
    fmt.Println(numbers)         // Вывод всего массива
}
Задание: Работа с массивами

14. Срезы (Slices)
go
package main
import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}  // Создание среза
    slice = append(slice, 6)       // Добавление элемента
    fmt.Println("Срез:", slice)
    fmt.Println("Длина:", len(slice))  // Длина среза
}
Задание: Базовые операции со срезами

15. Карты (Maps)
go
package main
import "fmt"

func main() {
    ages := make(map[string]int)  // Создание карты
    ages["Alice"] = 25            // Добавление элемента
    ages["Bob"] = 30
    fmt.Println("Возраст Alice:", ages["Alice"])  // Доступ к элементу
}
Задание: Работа с ассоциативными массивами (map)

16. Функции
go
package main
import "fmt"

func add(a int, b int) int {     // Объявление функции
    return a + b                 // Возврат значения
}

func main() {
    result := add(5, 3)          // Вызов функции
    fmt.Println("Результат:", result)
}
Задание: Создание и вызов функций

17. Функция с несколькими возвращаемыми значениями
go
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
Задание: Функции с множественными возвращаемыми значениями

18. Анонимные функции
go
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
Задание: Использование анонимных функций

19. Замыкания
go
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
Задание: Демонстрация замыканий

20. Указатели
go
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
Задание: Основы работы с указателями

21. Структуры
go
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
Задание: Создание и использование структур

22. Методы структур
go
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
Задание: Добавление методов к структурам

23. Интерфейсы
go
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
Задание: Базовое использование интерфейсов

24. Горутины
go
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
Задание: Базовое использование горутин

25. Каналы
go
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
Задание: Базовое использование каналов

26. Буферизованные каналы
go
package main
import "fmt"

func main() {
    ch := make(chan int, 2)     // Буферизованный канал
    ch <- 1                     // Отправка без блокировки
    ch <- 2
    fmt.Println(<-ch)           // Получение данных
    fmt.Println(<-ch)
}
Задание: Использование буферизованных каналов

27. Select с каналами
go
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
Задание: Использование select для работы с каналами

28. Обработка ошибок
go
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
Задание: Базовая обработка ошибок

29. Panic и Recover
go
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
Задание: Обработка паник с recover

30. Defer
go
package main
import "fmt"

func main() {
    defer fmt.Println("Это выполнится последним")  // Отложенный вызов
    defer fmt.Println("Это выполнится предпоследним")
    
    fmt.Println("Обычный вывод")
    fmt.Println("Еще один обычный вывод")
}
Задание: Использование defer для отложенных вызовов

31. Работа с файлами
go
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
Задание: Базовая работа с файлами

32. JSON маршалинг
go
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
Задание: Работа с JSON

33. HTTP сервер
go
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
Задание: Простой HTTP сервер

34. HTTP клиент
go
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
Задание: Создание HTTP клиента

35. Тестирование
go
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
Задание: Структура тестирования в Go

36. Пакеты и импорты
go
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
Задание: Использование стандартных пакетов

37. Строки и руны
go
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
Задание: Работа с Unicode строками

38. Ввод пользователя
go
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
Задание: Чтение ввода от пользователя

39. Время и дата
go
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
Задание: Работа с временем

40. Сортировка
go
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
Задание: Сортировка данных

41. Работа с директориями
go
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
Задание: Работа с файловой системой

42. Регулярные выражения
go
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
Задание: Использование регулярных выражений

43. Шаблоны (Templates)
go
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
Задание: Использование текстовых шаблонов

44. Генерация случайных данных
go
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
Задание: Генерация случайных данных

45. Работа с байтами
go
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
Задание: Работа с байтовыми слайсами

46. Интерфейс io.Reader
go
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
Задание: Использование интерфейса Reader

47. Интерфейс io.Writer
go
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
Задание: Использование интерфейса Writer

48. Кастомные типы
go
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
Задание: Создание кастомных типов с методами

49. Композиция структур
go
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
Задание: Композиция структур через встраивание

50. Рефлексия
go
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
Задание: Основы рефлексии в Go

Продолжение следует... Хотите, чтобы я продолжил со следующими 50 примерами?

да
Продолжаю со следующими 50 примерами Go:

Продвинутые концепции Go
51. Контекст (Context)
go
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
Задание: Использование context для управления горутинами

52. Context с таймаутом
go
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
Задание: Context с автоматическим таймаутом

53. Sync.WaitGroup
go
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
Задание: Ожидание завершения группы горутин

54. Sync.Mutex
go
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
Задание: Синхронизация доступа к общим данным

55. Sync.RWMutex
go
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
        go func(id int) {
            fmt.Printf("Чтец %d: %s\n", id, store.Get("name"))
        }(i)
    }
    
    time.Sleep(time.Millisecond * 100)
}
Задание: Оптимизированная блокировка для частого чтения

56. Sync.Once
go
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
Задание: Гарантия однократного выполнения кода

57. Sync.Pool
go
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
Задание: Использование пула объектов для оптимизации

58. Каналы с контекстом
go
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
Задание: Управление каналами через контекст

59. Fan-out, Fan-in
go
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
Задание: Паттерны Fan-out и Fan-in для параллельной обработки

60. Атомарные операции
go
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
Задание: Использование атомарных операций для счетчиков

61. Пользовательские ошибки
go
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
Задание: Создание пользовательских типов ошибок

62. errors.Is и errors.As
go
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
Задание: Современные методы проверки ошибок

63. Работа с временными зонами
go
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
Задание: Работа с временными зонами

64. Таймеры и Ticker'ы
go
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
Задание: Использование таймеров и тикеров

65. Шаблоны HTML
go
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
Задание: Генерация HTML с шаблонами

66. HTTP middleware
go
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
Задание: Создание HTTP middleware

67. WebSocket сервер
go
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
Задание: Создание WebSocket сервера

68. Работа с CSV
go
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
Задание: Работа с CSV файлами

69. Base64 кодирование
go
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
Задание: Кодирование и декодирование Base64

70. Хеширование
go
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
Задание: Вычисление хешей данных

71. Генерация UUID
go
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
Задание: Генерация и работа с UUID

72. Флаги командной строки
go
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
Задание: Обработка аргументов командной строки

73. Environment variables
go
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
Задание: Работа с переменными окружения

74. Logging с уровнями
go
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
Задание: Создание логгера с уровнями логирования

75. Работа с архивами
go
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
Задание: Создание ZIP архивов

76. Парсинг XML
go
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
Задание: Работа с XML данными

77. Парсинг YAML
go
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
Задание: Работа с YAML конфигурациями

78. SQL база данных
go
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
Задание: Работа с SQL базой данных

79. Миграции базы данных
go
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