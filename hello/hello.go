package main

import . "fmt"

func main() {
    var x, y int = 30, 50
    var age, name = 10, "Maria"

    _ = name

    var t, u int
    var r int

    t, u, r = x, y, age

    Println(t + u + r);
}