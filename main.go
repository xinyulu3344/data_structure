package main

import "fmt"

func main() {
    a1 := 1
    b1 := 1
    var x1 any
    var x2 any
    x1 = &a1
    x2 = &b1
    fmt.Printf("%x\n", x1)
    fmt.Printf("%x\n", x2)
}