package main

import (
    "fmt"
    "os"
)

func main() {
    var n1 int
    var n2 float32
    var s string

    file1, _ := os.Open("hello1.txt")
    defer file1.Close()
    if n, err := fmt.Fscan(file1, &n1, &n2, &s); err == nil {
        fmt.Println("입력 개수:", n)
        fmt.Println(n1, n2, s)
    } else {
        fmt.Println("1.에러남:", err)
        fmt.Println(n1, n2, s)
    }

    file2, _ := os.Open("hello2.txt")
    defer file2.Close()
    if n, err := fmt.Fscanln(file2, &n1, &n2, &s); err == nil {
        fmt.Println("입력 개수:", n)
        fmt.Println(n1, n2, s)
    } else {
        fmt.Println("2.에러남:", err)
        // 뭐지? 에러가 났는데 왜.. 값은 출력돼지?
        fmt.Println(n1, n2, s)
    }

    file3, _ := os.Open("hello3.txt")
    defer file3.Close()
    if n, err := fmt.Fscanf(file3, "%d,%f,%s", &n1, &n2, &s); err == nil {
        fmt.Println("입력 개수:", n)
        fmt.Println(n1, n2, s)
    } else {
        fmt.Println("3.에러남:", err)
        fmt.Println(n1, n2, s)
    }
}