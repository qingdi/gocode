package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { //使用空标识符
		s += sep + arg //追加旧的字符串，分隔符，和下一个参数，生成新字符串，然后将新的字符串赋值给s。旧的内容丢弃，被例行垃圾回收，如果字符比较大，这种方式需要优化
		sep = " "
	}
	fmt.Println(s)
}
