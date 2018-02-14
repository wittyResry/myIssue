package multipleresults

import "fmt"

type point struct {
	x, y int
}

func swap(p point) point {
	return point{p.y, p.x}
}

func swap2(x, y string) (string, string) {
	return y, x
}

func Main() {
	p := point{1, 2}
	p_swap := swap(p)
	fmt.Printf("%v\n", p_swap)
	c, d := swap2("hello", "world")
	fmt.Printf("c=%s, d=%s\n", c, d)
}
