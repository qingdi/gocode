package ch6

import (
	"fmt"
	"image/color"
)

//结构体内嵌
type ColoredPoint struct {
	Point //可以理解为继承
	Color color.RGBA
}

func Cp1() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{2, 12}, blue}

	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}

//使用指针，我们可以共享通用的结构
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func Cp2() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint2{&Point{1, 1}, red}
	q := ColoredPoint2{&Point{2, 2}, blue}
}
