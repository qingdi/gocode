package sword

import "fmt"

func Repeatnum() {
	a := [10]int{1, 2, 6, 4, 5, 3, 2, 8, 1, 7}
	l := len(a)
	d := []int{}
	for k, _ := range a {
		if a[k] < 0 || a[k] > l-1 {
			return
		}
	}
	for i := 0; i < l; i++ {
		for a[i] != i {
			if a[i] == a[a[i]] {
				d = append(d, a[i])

			}
			temp := a[i]
			a[i] = a[temp]
			a[temp] = temp
		}
	}
	fmt.Println(d)
}
