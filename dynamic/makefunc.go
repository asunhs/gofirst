package main

import (
    "fmt"
    "reflect"
)

func h(args []reflect.Value) []reflect.Value{
    fmt.Println("Hello, world")
    return nil
}

// 파라미터는 사용 안해도 상관없나?
func simple(s string) *string {
    return nil
}

func main() {
    var hello func()

    fn := reflect.ValueOf(&hello).Elem()
    v := reflect.MakeFunc(fn.Type(), h)

    fn.Set(v)

    hello()


    var sayHello func(string)

    fns := reflect.ValueOf(&sayHello).Elem()
    s := reflect.MakeFunc(fns.Type(), h)

    fns.Set(s)

    // sayHello() // 에러남
    sayHello("hello")
    simple("hello")

    fmt.Println("========================================================")
    fmt.Printf("reflect.TypeOf(simple)\t%v\n", reflect.TypeOf(simple))
    fmt.Printf("reflect.ValueOf(simple)\t%v\n", reflect.ValueOf(simple))
}