package ch1

var global *int

func Test() {
	var x int
	x = 1
	global = &x
}

func Fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
