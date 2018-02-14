package exercise

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	tokens := strings.Fields(s)
	for _, v := range tokens {
		res[v]++
	}
	return res
}

func MainMaps() {
	wc.Test(WordCount)
}
