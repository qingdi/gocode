package ch6

import (
	"fmt"
	"math"
)

//关于封装和组合

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.Y, q.Y-p.Y)
}

//p为方法接收者
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.Y, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func Geometry() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func Geometry2() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{4, 1},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}

//指针接收者
//习惯上，如果point的任何一个方法使用了指针接收者，那么所有的point方法都应该使用指针接收者
