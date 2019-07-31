package ch5

import (
	"github.com/pkg/errors"
)

var EOF = errors.New("EOF")

func Ir() {
	//in := bufio.NewReader(os.Stdin)
	//for {
	//	r, _, err := in.ReadRune()
	//	if err == io.EOF {
	//		break
	//	}
	//
	//	if err != nil {
	//		return fmt.Errorf("read failed: %v", err)
	//	}
	//}
}
