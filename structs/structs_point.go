package structs

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func MainStructPoint() {
	v := Vertex2{1, 2}
	p := &v
	(*p).X = 1e9
	p.Y = 1e8
	fmt.Println(v)
}
