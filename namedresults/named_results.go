package namedresults

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 10
    y = sum / 10
    return
}

func Main() {
    fmt.Println(split(17))
}
