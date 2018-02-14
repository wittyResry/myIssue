package arrays

import "fmt"

func MainArray() {
    var a[2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
    fmt.Println(primes[1:4])
    names := [4]string{
        "Jo",
        "Pa",
        "Ge",
        "Ri",
    }
    a1 := names[1:3]
    a1[1] = "XX"
    fmt.Println(names)

    var snil []string
    fmt.Println(snil, len(snil), cap(snil))
    if snil == nil {
        fmt.Println("nil!")
    }
    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
        fmt.Printf("2^%d = %d,", i, v)
    }
}
