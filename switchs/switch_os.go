package switchs

import (
    "fmt"
    "runtime"
)

func MainSwitchOS() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Linux.")
        default:
            fmt.Printf("%s.\n", os)
    }
}
