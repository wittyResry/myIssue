package structs

import "fmt"

type Vertex struct {
	X int
	Y int
}

type MyStruct struct {
	X int
	Y string
}

func MainStruct() {
	fmt.Println(Vertex{1, 2})
	fmt.Println("MyStruct Init:", MyStruct{X: 1, Y: "a"})
}
