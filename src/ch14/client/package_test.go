package client

import (
    "ch14/series"
    "testing"
)

func TestPackage(t *testing.T) {
    t.Log(series.Fib(3))

    // 小写字母开头的方法，不能被包外访问
    //series.square(5)
}
