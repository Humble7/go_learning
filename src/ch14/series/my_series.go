package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func Fib(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n should be greater than or equal to 0")
	}

	ret := []int {1, 1}
	for i := 2; i < n; i ++ {
		ret = append(ret, ret[i - 1] + ret[i - 2])
	}

	return ret, nil
}

// 小写字母开头的方法，不能被包外访问
func square(n int) int {
	return n * n
}
