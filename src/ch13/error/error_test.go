package error

import (
	"errors"
	"testing"
)

func fib(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n should be greater than or equal to 0")
	}

	ret := []int {1, 1}
	for i := 2; i < n; i ++ {
		ret = append(ret, ret[i - 1] + ret[i - 2])
	}

	return ret, nil
}

func TestFib(t *testing.T) {
	var (
		list []int
		err error
	)

	// 提早失败，避免嵌套
    if list, err = fib(5); err != nil {
    	t.Error(err)
		return
	}

	t.Log(list)
}
