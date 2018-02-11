package functions

import "fmt"

func add(x int, y int) int {
    return x + y
}

func add_(x, y int) int {
    return x + y
}

func Main() {
    fmt.Println(add_(11, 12))
}
