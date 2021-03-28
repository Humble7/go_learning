package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	s = "中"
	// 字符串的len返回的是字节数
	t.Log(len(s))

	// string是不可变的byte slice
	//s[1] = '2'

	// 将字符串转成unicode的slice
	c := []rune(s)
	t.Log(len(c))
	// Unicode是字符集，值编码规则
	// UTF8是编码规则的具体实现，即以什么形式存储在磁盘上
    t.Logf("中 Unicode 为：%x", c[0])
	t.Logf("中 UTF8 为：%x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _,c := range s {
		t.Logf("%[1]c, %[1]x", c)
	}
}