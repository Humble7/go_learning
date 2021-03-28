package string_test

import "testing"

// 字符串默认值值空串而不是nil
func TestString(t *testing.T) {
	var str string
	t.Log("*" + str + "*")
}