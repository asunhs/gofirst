package main

import "fmt"

func main() {
    var s1, s2 string

    // Scanln 을 쓰지 않으면 한줄에서 다 입력받으려 하네
    if n, err := fmt.Scanln(&s1, &s2); err == nil {
        fmt.Println("입력 개수:", n)
        fmt.Println(s1)
        fmt.Println(s2)
    }

    var n1, n2 int

    if n, err := fmt.Scanf("%d,%d", &n1, &n2); err == nil {
        fmt.Println("숫자 개수:", n)
        fmt.Println(n1)
        fmt.Println(n2)
    } else {
        fmt.Println("Error")
    }
}