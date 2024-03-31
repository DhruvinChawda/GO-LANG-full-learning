package main

import (
    "fmt"
)

func main() {
    var integerVar int = 42
    var floatVar float64

    // Typecasting int to float64
    floatVar = float64(integerVar)

    fmt.Printf("Integer value: %d\n", integerVar)
    fmt.Printf("Float value: %.2f\n", floatVar)

    // Typecasting float64 to int
    var newIntegerVar int = int(floatVar)
    fmt.Printf("New Integer value from Float: %d\n", newIntegerVar)
}
