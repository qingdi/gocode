package ch1

import (
	"fmt"
	"os"
)

func Echo() {
	var s, sep string //var定义变量，并隐式初始化
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ { //:=用于短变量声明，不支持 ++i
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
