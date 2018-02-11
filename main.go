package main

import (
    "fmt"
    "my-go-test/importmath"
    "my-go-test/packages"
    "my-go-test/functions"
    "my-go-test/multipleresults"
    "my-go-test/namedresults"
    "my-go-test/variables"
)

func main() {
    fmt.Println("Welcome to the playgroud!")
    importmath.Main()
    packages.Main()
    functions.Main()
    multipleresults.Main()
    namedresults.Main()
    variables.Main()
    variables.MainInitializers()
}
