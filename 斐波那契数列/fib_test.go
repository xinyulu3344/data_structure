package main

import (
    "fmt"
    "testing"
)

func TestFib1(t *testing.T) {
    fmt.Println(fib1(30))
}

func TestFib2(t *testing.T) {
    fmt.Println(fib2(100))
}
