package main

import "fmt"

func main() {
    s := []int{
        1,
        4,
        9,
        16,
        25,
    }

    t := s

    t[4] = 36

    //invalid operation: s == t (slice can only be compared to nil)
    //fmt.Println(s == t)
    fmt.Println(s, t)
}