package main

import "fmt"

func main() {
    a := [5]int {
        1,
        2,
        4,
        5,
        7,
    }

    for index, value := range a {
        fmt.Printf("%d\t%d\n", index, value)
    }

    b := a

    // true 임, 무슨의미이지?
    fmt.Println(b == a)

    b[2] = 2;

    fmt.Println(a, b)
    fmt.Println(b == a)
}