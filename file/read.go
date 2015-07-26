package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    if b, err := ioutil.ReadFile("./hello1.txt"); err == nil {
        fmt.Printf("%s", b)
    } else {
        fmt.Println(err)
    }
}