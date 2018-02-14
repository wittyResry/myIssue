package maps

import "fmt"

type Vertex struct {
	X, Y int
}

func MainMap() {
	m := make(map[string]Vertex)
	m["s1"] = Vertex{5, 6}
	m["s2"] = Vertex{7, 8}
	fmt.Println("")
	fmt.Println(m)
	var m1 = map[string]Vertex{
		"Bell Labs": {403, -747},
		"Google":    {3742202, -12208408},
	}
	fmt.Println(m1)
	delete(m1, "Bell Labs")
	v, ok := m1["Bell Labs"]
	fmt.Println("Value:", v, "Exist:", ok)
	v, ok = m1["Google"]
	fmt.Println("Value:", v, "Exist:", ok)
}
