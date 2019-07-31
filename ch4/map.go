package ch4

import (
	"fmt"
	"sort"
)

func Maptest() {
	ages := make(map[string]int)
	fmt.Println(ages)
	ages2 := map[string]int{
		"alice":   32,
		"charlie": 34,
	}
	delete(ages2, "alice")
	fmt.Println(ages2)
}

func Maptest2() {
	ages := map[string]int{
		"alice":   32,
		"charlie": 34,
	}
	//var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
