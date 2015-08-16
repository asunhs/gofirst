package main

import (
    "fmt"
    "reflect"
)

func main() {

    f := 1.3
    t := reflect.TypeOf(f)
    v := reflect.ValueOf(f)

    fmt.Printf("TypeOf.Name()\t%v\n", t.Name())
    fmt.Printf("TypeOf.Size()\t%v\n", t.Size())
    fmt.Printf("TypeOf.Kind()\t%v\n", t.Kind())
    fmt.Printf("ValueOf.Type()\t%v\n", v.Type())
    fmt.Printf("ValueOf.Kind()\t%v\n", v.Kind())
    // fmt.Printf("ValueOf.Int()\t%v\n", v.Int()) // 캐스팅을 해주는 줄 알았더니 panic 나네
    fmt.Printf("ValueOf.Float()\t%v\n", v.Float())
}