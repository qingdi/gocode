package ch5

import (
	"fmt"
	"os"
)

//变长函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//fmt.Println(sum(1,2,2,34))

func Sumtest() {
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}

	fmt.Println(sum(values...))
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
}
