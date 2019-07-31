package ch4

//
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

//go允许我们定义不带名称的结构体成员，只需要指定类型即可。这种结构体成员称为匿名成员

type Point2 struct {
	X, Y int
}
type Circle2 struct {
	Point2
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func Cir() {
	var w Wheel2
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}
