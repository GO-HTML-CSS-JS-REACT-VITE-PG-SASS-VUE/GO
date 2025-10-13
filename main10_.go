package main
import "fmt"

func main() {
    for i := 1; i <= 5; i++ {      // Классический цикл for
        fmt.Println("Итерация:", i)
    }
     for i := 1; i <= 10; i++ {      // Классический цикл for
         for j := 1; j <= 10; j++ {      // Классический цикл for
        fmt.Println(j,"*",i,"=",j*i)
    }
    }
}