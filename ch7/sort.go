package ch7

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func St() {
	names := []string{"hou", "zhou", "wang", "lian"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}
