package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("clear resources")
	}()

	fmt.Println("Start")

	// panic有调用栈，有错误信息，会执行defer函数
	panic(errors.New("Something wrong"))

	// Exit没有调用栈，没有错误信息，不会执行defer函数
	//os.Exit(-1)
}
