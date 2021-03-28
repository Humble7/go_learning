package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1,2}
	b := [...]int{2,1}
	//c := [...]int{1,2,3}
	d := [...]int{1,2}
	t.Log(a == b)
	// 长度不同不可比较
	//t.Log(a == c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	// 取消可读权限
	a = a &^ Readable
	t.Log(Readable & a == Readable, Writable & a == Writable, Executable & a == Executable)
}