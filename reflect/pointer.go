package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := new(int)
    *a = 1

    fmt.Printf("reflect.TypeOf(a)\t\t%v\n", reflect.TypeOf(a))
    fmt.Printf("reflect.ValueOf(a)\t\t%v\n", reflect.ValueOf(a))
    fmt.Printf("reflect.ValueOf(a).Elem()\t%v\n", reflect.ValueOf(a).Elem())
    fmt.Printf("reflect.ValueOf(a).Elem().Int()\t%v\n", reflect.ValueOf(a).Elem().Int())
    
    var b interface{}
    b = 1
    
    fmt.Printf("reflect.TypeOf(b)\t\t%v\n", reflect.TypeOf(b))
    fmt.Printf("reflect.ValueOf(b)\t\t%v\n", reflect.ValueOf(b)) // "<int Value>" 책에 오류가 있네
    fmt.Printf("reflect.ValueOf(b).Int()\t%v\n", reflect.ValueOf(b).Int())
}