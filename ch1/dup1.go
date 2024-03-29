package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
