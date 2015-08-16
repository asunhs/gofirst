package main

import (
    "fmt"
    "reflect"
)

type Data struct {
    a, b int
}

func main() {
    num := 1
    fmt.Println(reflect.TypeOf(num))

    s := "Hello, world"
    fmt.Println(reflect.TypeOf(s))

    f := 1.3
    fmt.Println(reflect.TypeOf(f))

    data := Data{1, 2}
    fmt.Println(reflect.TypeOf(data))

    t := reflect.TypeOf(f)

    fmt.Println("===============================================================")
    fmt.Printf("reflect.TypeOf(t)\t\t%v\n", reflect.TypeOf(t))
    fmt.Printf("reflect.TypeOf(t.Kind())\t%v\n", reflect.TypeOf(t.Kind()))
    fmt.Printf("reflect.TypeOf(reflect.Float64)\t%v\n", reflect.TypeOf(reflect.Float64))
    fmt.Printf("t.Kind() == reflect.Float64\t%v\n", t.Kind() == reflect.Float64)
}