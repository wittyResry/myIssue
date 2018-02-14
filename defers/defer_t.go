package defers

import "fmt"

func MainDefer() {
	defer fmt.Println("defer")
	fmt.Printf("hello ")
}
