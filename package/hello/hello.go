package main

// 실패함
// GOPATH 쫌 귀찮네?
// 상대 경로로 참조할 수 있구나 결국 비슷하네
// 책에서는 글로벌 패키지 참조만 설명한거네
import (
    "../calc"
    "../test/memo"
    "fmt"
)

func main() {
    fmt.Println(calc.Sum(2, 3))
    fmt.Println(memo.Memo("Sun"))
}