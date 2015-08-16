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
}

func (p *Person) Hello() {
    if p == nil {
        fmt.Println("No Person")
    } else {
        fmt.Println("Hello", p.name)
    }
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

    // 포인터 타입도 함수 확장이 가능하다!
    // 포인터 타입도 덕타이핑이 가능하다!
    // 포인터 타입의 함수를 확장하여 "nil 검사를 포함" 할 수 있다
    // 포인터를 써야만 가능하다
    // 강타입과 포인터가 결합하여 재밌는 상황을 만드네
    s = getSun()
    t = getNilPerson()

    s.Hello()
    t.Hello()

    fmt.Println("nil type", reflect.TypeOf(nil))
    fmt.Println("s type", reflect.TypeOf(s))
}