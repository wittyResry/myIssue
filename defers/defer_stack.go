package defers

import "fmt"

func MainDeferStack() {
	fmt.Println("begin")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer %d ,", i)
	}
	fmt.Println("done")
}
