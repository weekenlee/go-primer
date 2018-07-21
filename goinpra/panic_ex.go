package main

import (
    "errors"
    "fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {

    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("Trap panic: %s (%T)\n", err, err)
            goto con
        }
    }()

    fmt.Println("Divide 1 by 0")

    _, err := precheckDivide(1, 0)

    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }
    fmt.Println("Divide 2 by 0")
    divide(2, 0)

con:
    fmt.Println("end")
}

func precheckDivide(a, b int) (int ,error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return divide(a, b), nil
}

func divide(a, b int) int {
    return a/b
}
