package main

func fib1(n int) int{
    if n <= 1 {
        return n
    }
    return fib1(n-1) + fib1(n-2)
}


func fib2(n int) int {
    if n <= 1 {
        return n
    }
    first := 0
    second := 1
    sum := 0
    for i := 0; i < n - 1; i++ {
        sum = first + second
        first = second
        second = sum
    }
    return sum
}


