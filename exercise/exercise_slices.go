package exercise

import (
//"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	ret := [][]uint8{}
	var s []int
	s = append(s, 0)
	for i := 0; i < dx; i++ {
		tmp := []uint8{}
		for j := 0; j < dy; j++ {
			tmp = append(tmp, uint8(i*j))
		}
		ret = append(ret, tmp)
	}
	return ret
}

func MainPic() {
	//pic.Show(Pic)
}
