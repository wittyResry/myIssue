package pointers

import "fmt"

func MainPointer() {
    i := 42
    p := &i
    fmt.Println("*p=", *p)
    *p = 21
    fmt.Println("i=", i)
}
