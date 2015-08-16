package main

import (
    "fmt"
    "reflect"
)

type Say struct {
    name string
}

func (s Say) Hello() {
    fmt.Println("Hello", s.name)
}

func (s *Say) Bye() {
    if s == nil {
        fmt.Println("Nil")
    } else {
        fmt.Println("Bye", s.name)
    }
}

type Isay interface {
    Hello()
    Bye()
}

func NilSay() *Say {
    return nil
}

func main() {
    s := Say{"Sun"}
    r := &s
    t := *r
    var p *Say
    var n Say

    //n = nil // 에러

    //var i Isay = Say{} // 에러
    // Isay 타입으로는 실행타입을 결정할 수 없다
    var i Isay = NilSay() // nil이어도 타입이 결정되네


    fmt.Printf("reflect.TypeOf(s)\t%v\n", reflect.TypeOf(s))
    s.Hello()
    s.Bye()
    fmt.Printf("reflect.TypeOf(r)\t%v\n", reflect.TypeOf(r))
    r.Hello()
    r.Bye()
    fmt.Printf("reflect.TypeOf(t)\t%v\n", reflect.TypeOf(t))
    t.Hello()
    t.Bye()
    fmt.Printf("reflect.TypeOf(p)\t%v\n", reflect.TypeOf(p))
    //p.Hello() // 에러
    p.Bye()
    fmt.Printf("reflect.TypeOf(n)\t%v\n", reflect.TypeOf(n))
    n.Hello()
    n.Bye()
    fmt.Printf("reflect.TypeOf(i)\t%v\n", reflect.TypeOf(i))
    //i.Hello() // 에러
    i.Bye()
}