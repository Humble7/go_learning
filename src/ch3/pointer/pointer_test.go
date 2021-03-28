package pointer_test

import "testing"

// Go 不支持指针运算
func TestFetchPointer(t *testing.T) {
	a := 1
	aPtr := &a
	// 不支持
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}
