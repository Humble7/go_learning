package type_test

import "testing"

// go 不支持隐形类型转换
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64

	// 不支持
	//b = a
	b = int64(a)

	t.Log(a, b)
}

// 即使是别名，在go中也不支持隐式类型转换
type MyInt int64
func TestTypeAlias(t *testing.T) {
	var a int64 = 1
	var b MyInt

	// 不支持
	//b = a
    b = MyInt(a)
	t.Log(a, b)
}
