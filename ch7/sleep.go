package ch7

import (
	"flag"
	"fmt"
	"github.com/qingdi/gocode/ch2"
	"io"
	"os"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func Sl() {
	flag.Parse()
	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

type celsiusFlag struct {
	ch2.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = ch2.Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = ch2.FToC(ch2.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value ch2.Celsius, usage string) *ch2.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func Sl2() {
	flag.Parse()
	fmt.Println(*temp)
}

func Sl3() {
	var w io.Writer
	fmt.Println(w)
	w = os.Stdout
	w.Write([]byte("hello"))
}
