package main

// 실패함
// GOPATH 쫌 귀찮네?
// 상대 경로로 참조할 수 있구나 결국 비슷하네
// 책에서는 글로벌 패키지 참조만 설명한거네
import (
    "../calc"
    "../test/memo"
    "../test/diff"
    "fmt"
    "reflect"
)

func main() {
    fmt.Println(calc.Sum(2, 3))
    fmt.Println(memo.Memo("Sun"))

    // 접근 변수가 패키지명을 따르네 : diff\<bye package>
    fmt.Println(bye.Bye("Sun"))

    // fmt.Println(reflect.TypeOf(bye)) // 에러남 : use of package bye without selector
}