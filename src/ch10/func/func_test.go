package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestComplexFunc(t *testing.T) {
	ret := timeSpent(slowFunc)
	t.Log(ret(10))
}

// 可变长参数
func sum(op ...int) int {
	var sum = 0
	for _, v := range op {
		sum += v
	}

	return sum
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
}

func clear() {
	fmt.Println("clear resources.")
}

func TestDef(t *testing.T) {
	defer clear()
	fmt.Println("Start")
	panic("Error")
}
