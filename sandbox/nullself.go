package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    name string
}

type Self interface {
    Hello()
    Bye()
}

func (p *Person) Hello() {
    if p == nil {
        fmt.Println("No Person")
    } else {
        fmt.Println("Hello", p.name)
    }
}

func (p Person) Bye() {
    // p 는 nil 과 비교가 불가능
    fmt.Println("Bye", p.name)
}

func getNilPerson() *Person {
    return nil
}

func getSun() *Person {
    return &Person{"Sun"}
}

func main() {
    var s Self
    var t Self
    var p Person
    k := Person{"Kim"}
    l := &Person{"Lee"}

    // 포인터 타입도 함수 확장이 가능하다!
    // 포인터 타입도 덕타이핑이 가능하다!
    // 포인터 타입의 함수를 확장하여 "nil 검사를 포함" 할 수 있다
    // 포인터를 써야만 가능하다
    // 강타입과 포인터가 결합하여 재밌는 상황을 만드네
    s = getSun()
    t = getNilPerson()

    s.Hello()
    t.Hello()
    p.Hello()
    k.Hello()
    l.Hello()

    fmt.Println("========================================")

    s.Bye()
    // t.Bye() // 에러남
    p.Bye()
    k.Bye()
    l.Bye()

    fmt.Println("========================================")

    fmt.Println("nil type", reflect.TypeOf(nil))
    fmt.Println("s type", reflect.TypeOf(s))
    fmt.Println("t type", reflect.TypeOf(t))
    fmt.Println("p type", reflect.TypeOf(p))
    fmt.Println("k type", reflect.TypeOf(k))
    fmt.Println("l type", reflect.TypeOf(l))
}