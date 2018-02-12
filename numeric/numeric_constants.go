package numeric

import "fmt"

const (
    Big = 1 << 100
    Small = Big >> 99
)

func MainNumericConstants() {
    //fmt.Print(Big) overflows
    //fmt.Print("%T\n", Small)
    fmt.Println(Small)
}
