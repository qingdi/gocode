package ch5

import (
	"fmt"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func Sq() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
