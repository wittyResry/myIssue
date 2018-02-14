package forcontinued

import "fmt"

func MainForContinued() {
	sum := 1
	for sum < 1024 {
		sum += sum
	}
	fmt.Println(sum)
	s := 0.0
	for i := 1; i < 100; i++ {
		s += float64(i)
	}
	fmt.Println(s)
}
