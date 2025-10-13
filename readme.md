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
Задание: Система миграций базы данных

80. REST API сервер
go
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
Задание: Создание REST API сервера

81. JWT аутентификация
go
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
Задание: JWT аутентификация в REST API

82. Graceful shutdown
go
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
Задание: Graceful shutdown сервера

83. Пул соединений
go
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
Задание: Реализация пула соединений

84. Кэширование
go
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
Задание: Реализация TTL кэша

85. Rate limiting
go
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
Задание: Реализация rate limiting

86. Конфигурационные файлы
go
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
Задание: Работа с конфигурационными файлами

87. Health check
go
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
Задание: Система health check для мониторинга

88. Метрики Prometheus
go
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
Задание: Интеграция с Prometheus для сбора метрик

89. Трассировка (Tracing)
go
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
Задание: Инструментирование приложения с трассировкой

90. Feature flags
go
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
Задание: Система feature flags для управления функциональностью

91. Работа с большими файлами
go
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
Задание: Эффективная обработка больших файлов

92. Парсинг аргументов с cobra
go
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
Задание: Создание CLI приложения с cobra

93. Генерация документации
go
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
Задание: Документирование кода с примерами

94. Benchmark тесты
go
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
Задание: Написание benchmark тестов для измерения производительности

95. Table-driven тесты
go
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
Задание: Table-driven тестирование

96. Mock тестирование
go
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
Задание: Mock тестирование с зависимостями

97. Интеграционные тесты
go
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
Задание: Интеграционное тестирование с базой данных

98. Property-based тестирование
go
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
Задание: Property-based тестирование

99. Fuzz тестирование
go
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
Задание: Fuzz тестирование для поиска краевых случаев

100. E2E тестирование
go
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
Задание: End-to-end тестирование API

Это первые 100 примеров из 200. Хотите, чтобы я продолжил со следующими 100 примерами? Они будут охватывать еще более продвинутые темы и реальные сценарии использования.

да
Продолжаю со следующими 100 примерами Go, охватывающими продвинутые темы и реальные сценарии:

Продвинутые техники и реальные сценарии
101. Graceful shutdown HTTP сервера
go
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
Задание: Полная реализация graceful shutdown для HTTP сервера

102. Конфигурация с Viper
go
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
Задание: Управление конфигурацией с помощью Viper

103. Миграции с Goose
go
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
Задание: Управление миграциями базы данных с Goose

104. gRPC сервер
go
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
Задание: Создание gRPC сервера

105. gRPC клиент
go
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
Задание: Создание gRPC клиента

106. GraphQL сервер с gqlgen
go
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
Задание: Создание GraphQL сервера с gqlgen

107. WebSocket чат
go
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
Задание: Реализация многопользовательского WebSocket чата

108. OAuth2 аутентификация
go
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
Задание: Реализация OAuth2 аутентификации с Google

109. JWT с refresh токенами
go
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
Задание: JWT аутентификация с refresh токенами

110. Rate limiting с Redis
go
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
Задание: Rate limiting с использованием Redis

111. Кэширование с Redis
go
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
Задание: Кэширование данных с Redis

112. Message Queue с RabbitMQ
go
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
Задание: Асинхронная обработка задач с RabbitMQ

113. Микросервисная архитектура
go
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
Задание: Простая микросервисная архитектура с API Gateway

114. Docker-совместимое приложение
go
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
Задание: Создание Docker-совместимого приложения

115. Kubernetes-ready приложение
go
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
Задание: Создание Kubernetes-ready приложения

116. Приложение с мониторингом
go
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
Задание: Приложение с комплексным мониторингом Prometheus

117. Distributed tracing с Jaeger
go
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
Задание: Distributed tracing с Jaeger

118. Feature flags с запуском в продакшн
go
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
Задание: Продвинутая система feature flags

119. A/B тестирование
go
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
Задание: Система A/B тестирования

120. Real-time дашборд с SSE
go
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
Задание: Real-time дашборд с Server-Sent Events

Это 120 примеров из 200. Продолжить с оставшимися 80 примерами? Они будут охватывать еще более специализированные темы и продвинутые сценарии использования Go.

lf
Продолжаю с оставшимися 80 примерами, охватывающими специализированные темы Go:

Специализированные техники и продвинутые сценарии
121. Работа с бинарными данными
go
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
Задание: Работа с бинарными данными и сериализация структур

122. Bit manipulation
go
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
Задание: Работа с битовыми операциями и битовыми множествами

123. Memory pool для объектов
go
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
Задание: Реализация пула объектов для уменьшения нагрузки на GC

124. Lock-free структуры данных
go
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
Задание: Lock-free очередь на основе атомарных операций

125. SIMD-подобные операции
go
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
Задание: Векторные операции и оптимизации циклов

126. Genetic algorithm
go
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
Задание: Реализация генетического алгоритма для оптимизации

127. Neural network (простая реализация)
go
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
Задание: Простая реализация нейронной сети с обратным распространением ошибки

128. Blockchain simulation
go
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
Задание: Симуляция простого блокчейна с proof-of-work

129. Game server (WebSocket)
go
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
Задание: Многопользовательский игровой сервер на WebSocket

130. Image processing
go
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
Задание: Обработка изображений: фильтры, преобразования, рисование

Это 130 примеров из 200. Хотите, чтобы я продолжил с оставшимися 70 примерами? Они будут охватывать такие темы как: криптография, системное программирование, сетевые протоколы, оптимизация производительности и многое другое.

да
Продолжаю с оставшимися 70 примерами, охватывающими криптографию, системное программирование и продвинутые темы:

Криптография и безопасность
131. AES шифрование
go
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
Задание: AES шифрование с использованием GCM режима

132. RSA шифрование и подписи
go
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
Задание: RSA шифрование, цифровые подписи и работа с PEM файлами

133. Password hashing с argon2
go
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
Задание: Безопасное хеширование паролей с Argon2

134. TLS сервер и клиент
go
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
Задание: TLS сервер и клиент с взаимной аутентификацией

Системное программирование
135. Работа с процессами
go
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
Задание: Управление процессами: запуск, мониторинг, сигналы