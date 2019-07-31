package ch3

import "fmt"

func Rem(slice []int, i int) []int {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}
