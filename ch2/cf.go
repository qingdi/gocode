package ch2

import (
	"fmt"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
)

var a = b + c
var b = f()
var c = 3

func f() int {
	return c + 1
}

func init() {

}

func Cf() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
